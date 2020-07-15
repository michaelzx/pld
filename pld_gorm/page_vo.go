package pld_gorm

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/michaelzx/pld/pld_reflect"
	"github.com/michaelzx/pld/pld_sql"
	"strings"
)

type IPage interface {
	GetPageNum() int64
	GetPageSize() int64
}

type PageVO struct {
	PageNum     int64 // 第几页
	PageSize    int64 // 每页几条
	PageTotal   int64 // 总共几页
	Total       int64 // 总共几条
	IsFirstPage bool  // 是否是第一页
	IsLastPage  bool  // 是否是最后一页
	List        interface{}
}

func NewPageVO(db *gorm.DB, list interface{}, sqlTpl string, params IPage) (*PageVO, error) {
	// return &PageVO{PageNum: pageNum, PageSize: pageSize, List: list}
	p := &PageVO{
		PageNum:  params.GetPageNum(),
		PageSize: params.GetPageSize(),
		List:     list,
	}
	err := p.get(db, sqlTpl, params)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PageVO) get(db *gorm.DB, sqlTpl string, params interface{}) error {
	// 先用text/template对sql解析一波
	resolver, err := pld_sql.NewResolver(sqlTpl, params)
	if err != nil {
		return err
	}
	countSql := fmt.Sprintf(`select count(*) from (%s) as t`, resolver.GetSql())
	countSql = strings.Replace(countSql, "\n", " ", -1)
	countSql = strings.Replace(countSql, "  ", " ", -1)
	result := db.Raw(countSql, resolver.GetValues()...).Count(&p.Total)
	if result.Error != nil {
		return result.Error
	}
	p.PageTotal = p.Total / p.PageSize
	if p.Total%p.PageSize > 0 {
		p.PageTotal = p.PageTotal + 1
	}
	if p.PageNum > p.PageTotal {
		p.PageNum = p.PageTotal
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	if p.PageTotal == 0 {

		p.IsFirstPage = true
		p.IsLastPage = true
	} else {
		switch p.PageNum {
		case 1:
			p.IsFirstPage = true
		case p.PageTotal:
			p.IsLastPage = true
		}
	}
	skipRow := p.PageSize * (p.PageNum - 1)
	pageSql := fmt.Sprintf(`%s limit %d,%d`, resolver.GetSql(), skipRow, p.PageSize)
	pageSql = strings.Replace(pageSql, "\n", " ", -1)
	result = db.Raw(pageSql, resolver.GetValues()...)
	if result.Error != nil {
		if result.RecordNotFound() {
			return nil
		} else {
			panic(result.Error)
		}
	}
	if !pld_reflect.IsPtr(p.List) {
		return errors.New("字段 List 必须是指针类型")
	}
	result.Scan(p.List)
	// fmt.Printf("p.Total->%#v\n", p.Total)
	// fmt.Printf("p.PageNum->%#v\n", p.PageNum)
	// fmt.Printf("p.PageSize->%#v\n", p.PageSize)
	// fmt.Printf("p.PageTotal->%#v\n", p.PageTotal)
	// fmt.Printf("p.IsFirstPage->%#v\n", p.IsFirstPage)
	// fmt.Printf("p.IsLastPage->%#v\n", p.IsLastPage)
	// fmt.Printf("countSql->%#v\n", countSql)
	// fmt.Printf("pageSql->%#v\n", pageSql)
	return nil
}

package pld_sql

import (
	"fmt"
	"github.com/michaelzx/pld/pld_config"
	"github.com/michaelzx/pld/pld_logger"
	"regexp"
	"strings"
	"testing"
)

type SdsDocParams struct {
	SDSNo   string   `valid:"" cn:"SDSNo"`
	Title   string   `valid:"" cn:"Title"`
	Cas     string   `valid:"" cn:"Cas"`
	CasList []string `json:"-"`
}

func TestNewResolver(t *testing.T) {
	pld_logger.Init(pld_config.LoggerConfig{
		Mode: "dev",
	})
	sqlSdsDocPage := `select d.no
,d.sds_no
,d.model_version_no
,d.title
,d.version
,d.revision_date
,d.effective_date
,d.type_tag
,m.title as model_title
,v.version as model_version
from sds_doc d 
left join sds_model_version v on v.no=d.model_version_no
left join sds_model m on m.no=v.model_no
where d.deleted_at is null
{{if .SDSNo}} and d.sds_no like concat('%',#{SDSNo},'%') {{end}}
{{if .Title}} and d.title like concat('%',#{Title},'%') {{end}}
{{if .CasList}} 
	{{range $casI, $casV := .CasList}} 
	and JSON_CONTAINS(d.doc_data->'$.a4s1zpfm0c8w[*][1]', '#{CasList[{{$casI}}]}') 
	{{end}}
{{end}}
order by d.id desc
`
	params := SdsDocParams{
		SDSNo: "xxxx",
		CasList: []string{
			"111-76-2",
			"7732-18-5",
		},
	}
	resolver, err := NewResolver(sqlSdsDocPage, &params)
	if err != nil {
		panic(err)
	}
	fmt.Println(resolver.GetSql())
	fmt.Println(resolver.GetValues())
}

func TestParams(t *testing.T) {
	// 兼容切片和数组
	tplParamsRegexp := regexp.MustCompile(`#{(?P<param>(?P<name>\w*)(\[(?P<idx>\d+)])*)}`)
	tplParams := tplParamsRegexp.FindAllStringSubmatch("asdfasdfasdf#{xxxxxxx}1231312312312331#{yyyyyyyyyyy[1]}", -1)
	groupNames := tplParamsRegexp.SubexpNames()
	// 循环所有行
	for _, user := range tplParams {
		m := make(map[string]string)
		fmt.Println(strings.Join(user, ","))
		// 对每一行生成一个map
		for j, name := range groupNames {
			if j != 0 && name != "" {
				m[name] = strings.TrimSpace(user[j])
				fmt.Println(name, user[j], fmt.Sprintf("%T", user[j]))
			}
		}
		fmt.Println("-----")
	}
}
package pld_gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/michaelzx/pld/pld_sql"
)

func RawPlus(db *gorm.DB, sqlTpl string, params interface{}) *gorm.DB {
	resolver, err := pld_sql.NewResolver(sqlTpl, params)
	if err != nil {
		clone := db.New()
		_ = clone.AddError(err)
		return clone
	}
	return db.Raw(resolver.GetSql(), resolver.GetValues()...)
}

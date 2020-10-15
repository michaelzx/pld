package db_meta

import (
	"database/sql"
	"strings"
)

// Constants for return types of golang
const (
	goString        = "string"
	goTime          = "pld_types.Time"
	goInt8          = "int8"
	goInt16         = "int16"
	goInt32         = "int32"
	goInt64         = "int64"
	goDecimal       = "decimal.Decimal"
	goBool          = "pld_types.Bool"
	goJson          = "pld_types.JsonString"
	goByteArray     = "[]byte"
	goNullString    = "*string"
	goNullTime      = "*pld_types.Time"
	goNullInt8      = "*int8"
	goNullInt16     = "*int16"
	goNullInt32     = "*int32"
	goNullInt64     = "*int64"
	goNullDecimal   = "*decimal.Decimal"
	goNullBool      = "pld_types.NullBool"
	goNullJson      = "*pld_types.JsonString"
	goNullByteArray = "*[]byte"
)

func SqlType2GoType(c *sql.ColumnType) string {
	mysqlType := strings.ToLower(c.DatabaseTypeName())
	nullable, _ := c.Nullable()
	switch mysqlType {
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":
		if nullable {
			return goNullString
		}
		return goString

	case "tinyint":
		if nullable {
			return goNullInt8
		}
		return goInt8
	case "smallint":
		if nullable {
			return goNullInt16
		}
		return goInt16
	case "mediumint", "int":
		if nullable {
			return goNullInt32
		}
		return goInt32
	case "bigint":
		if nullable {
			return goNullInt64
		}
		return goInt64
	case "decimal", "float", "double":
		if nullable {
			return goNullDecimal
		}
		return goDecimal
	case "date", "datetime", "time", "timestamp":
		if nullable {
			return goNullTime
		}
		return goTime
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		if nullable {
			return goNullByteArray
		}
		return goByteArray
	case "bit":
		if nullable {
			return goNullBool
		}
		return goBool
	case "json":
		if nullable {
			return goNullJson
		}
		return goJson
	}

	return ""
}

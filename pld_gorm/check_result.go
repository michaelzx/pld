package pld_gorm

import "github.com/jinzhu/gorm"

// CheckResult 检查result
// 1.有记录被找到; 2.没错误,如果有错误直接panic
func CheckResult(result *gorm.DB) bool {
	switch {
	case result.RecordNotFound():
		return false
	case result.Error != nil:
		panic(result.Error)
	default:
		return true
	}
}

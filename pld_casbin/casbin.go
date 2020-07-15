package pld_casbin

import (
	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
	"github.com/michaelzx/pld/pld_logger"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

var casbinEnforcer *casbin.Enforcer

func Init(db *gorm.DB) *casbin.Enforcer {
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	a, err := gormAdapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatal(errors.Wrap(err, "gormAdapter初始化失败"))
	}
	casbinEnforcer, err = casbin.NewEnforcer("resource/casbin_rbac.ini", a)
	if err != nil {
		log.Fatal(errors.Wrap(err, "casbin初始化失败"))
	}
	// casbinEnforcer.EnableAutoSave(true)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	// err = CasbinEnforcer.LoadPolicy()
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "casbin加载policy失败"))
	// }
	// casbinEnforcer.Enforce("alice", "data1", "read")
	// casbinEnforcer.AddPolicy(...)
	// casbinEnforcer.RemovePolicy(...)
	// casbinEnforcer.SavePolicy()
	// casbinEnforcer.AddGroupingPolicy()
	return casbinEnforcer
}

func Casbin() *casbin.Enforcer {
	if casbinEnforcer == nil {
		pld_logger.Error("casbin 未初始化")
	}
	return casbinEnforcer
}

func CasbinUserKey(id int64) string {
	idStr := strconv.Itoa(int(id))
	return "user=" + idStr
}
func CasbinUserID(str string) (int64, error) {
	str = strings.TrimPrefix(str, "user=")
	return strconv.ParseInt(str, 10, 64)
}
func CasbinUserIDList(keyList []string) []int64 {
	list := make([]int64, 0, 0)
	for _, k := range keyList {
		id, err := CasbinUserID(k)
		if err == nil {
			list = append(list, id)
		}
	}
	return list
}

func CasbinRoleKey(id int64) string {
	idStr := strconv.Itoa(int(id))
	return "role=" + idStr
}

func CasbinRoleID(str string) (int64, error) {
	str = strings.TrimPrefix(str, "role=")
	return strconv.ParseInt(str, 10, 64)
}
func CasbinRoleIDList(keyList []string) []int64 {
	list := make([]int64, 0, 0)
	for _, k := range keyList {
		id, err := CasbinRoleID(k)
		if err == nil {
			list = append(list, id)
		}
	}
	return list
}

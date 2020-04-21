package auth

import (
	"reflect"

	"github.com/fatih/structs"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"
)

var RBAC = struct {
	AllRBAC       []string `role:"ALL"`
	SuperRBAC     []string `role:"SUPER"`
	AdminRBAC     []string `role:"ADMIN"`
	UserRBAC      []string `role:"USER"`
	DeveloperRBAC []string `role:"DEVELOPER"`
}{
	AllRBAC:       []string{"SUPER", "ADMIN", "DEVELOPER", "USER"},
	SuperRBAC:     []string{"SUPER", "ADMIN", "DEVELOPER", "USER"},
	AdminRBAC:     []string{"ADMIN", "USER"},
	UserRBAC:      []string{"USER"},
	DeveloperRBAC: []string{"DEVELOPER"},
}

func GetRBACByRole(role string) ([]string, bool) {
	s := structs.New(RBAC)
	t := reflect.TypeOf(RBAC)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("role") == role {
			fn := t.Field(i).Name
			name := s.Field(fn)
			value := name.Value().([]string)
			return value, true
		}
	}
	return []string{}, false
}

func CheckPermission(role, object string) bool {
	rbac, _ := GetRBACByRole(role)
	if _, ok := lib.Find(rbac, object); !ok {
		return false
	}
	return true
}

package pld_fiber

import (
	"github.com/michaelzx/pld/pld_validator"
)

var validate *pld_validator.DefaultValidator

func init() {
	validate = &pld_validator.DefaultValidator{}
}
func Validator() *pld_validator.DefaultValidator {
	return validate
}

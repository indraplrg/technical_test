package validator

import (
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// register custom validators on the default binding engine used by Gin.
func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("date", dateValidation)
	}
}

// dateValidation checks that the value is a date in YYYY-MM-DD format.
func dateValidation(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	_, err := time.Parse("2006-01-02", value)
	return err == nil
}

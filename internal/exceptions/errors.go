package exceptions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

type ErrorResponse struct {
	Message  string `json:"message,omitempty"`
	ErrorRes string `json:"error,omitempty"`
}

type PredestinationError struct {
	Field   string
	Message string
}

func (m PredestinationError) Error() string {
	return fmt.Sprintf("{\"%s\": \"%s\"}", m.Field, m.Message)
}

func (m PredestinationError) JSON() map[string]string {
	return map[string]string{
		m.Field: m.Message,
	}
}

// ErrorToMap converts a validator error to a map[string]string
func ErrorToMap(dbError error) map[string]interface{} {
	var jsonMap map[string]interface{} = map[string]interface{}{}

	if validationErrors, ok := dbError.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			jsonMap[ToSnakeCase(err.Field())] = strings.TrimSpace(fmt.Sprintf("%s %s", err.Tag(), err.Param()))
		}
	}

	if err, ok := dbError.(*PredestinationError); ok {
		jsonMap[ToSnakeCase(err.Field)] = err.Message
	}

	if err, ok := dbError.(*mysql.MySQLError); ok {
		jsonMap["error"] = "Invalid request"
		fmt.Println(err.Error())
	}
	return jsonMap
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

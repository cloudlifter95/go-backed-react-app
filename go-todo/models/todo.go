package models

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json: "title"`
	Completed bool   `json:"completed`
}

func (p Todo) String() string {
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	result := "Todo{"
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Name
		value := v.Field(i).Interface()
		result += fmt.Sprintf("%s: %v, ", field, value)
	}
	result = result[:len(result)-2] + "}" // Remove trailing comma and space

	return result
}

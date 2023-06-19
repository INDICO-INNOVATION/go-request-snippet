package requestsnippet

import (
	"fmt"
	"net/url"
	"reflect"
)

func StructToURLData(data interface{}) *url.Values {
	values := url.Values{}

	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		fieldName := fieldType.Tag.Get("form")
		fieldValue := fmt.Sprintf("%v", field.Interface())

		values.Set(fieldName, fieldValue)
	}

	return &values
}

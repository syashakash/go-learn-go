package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	//fmt.Println(val)
	//fmt.Println(x)
	field := val.Field(0)
	//fmt.Println(field)
	fn(field.String())
}

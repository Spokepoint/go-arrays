package arrays

import (
	"errors"
	"reflect"
)

func StringArray(data []interface{}) (strs []string, err error) {
	strs = make([]string, 0, len(data))
	for _, v := range data {
		if v == nil {
			strs = append(strs, "")
		} else if str, ok := v.(string); ok {
			strs = append(strs, str)
		} else {
			return strs, errors.New("Error converting " +
				reflect.TypeOf(v).String() + " to string")
		}
	}
	return
}

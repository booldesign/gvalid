package gvalid

import (
	"fmt"
	"reflect"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/1/1 12:34
 * @Desc:
 */

// Funcs 验证 func map
type Funcs map[string]reflect.Value

// Call call func
func (f Funcs) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	if _, ok := f[name]; !ok {
		err = fmt.Errorf("tag: %s 格式错误", toLowerCamel(name))
		return
	}

	if len(params) != f[name].Type().NumIn() {
		err = fmt.Errorf("%s 参数不匹配", name)
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		if param == nil {
			continue
		}
		in[k] = reflect.ValueOf(param)
	}
	result = f[name].Call(in)

	return
}

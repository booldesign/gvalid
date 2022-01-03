package gvalid

import (
	"fmt"
	"reflect"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/1/1 11:33
 * @Desc:
 */

// ValidCustom 自定义验证
type ValidCustom interface {
	Valid(*Validation)
}

type Validation struct {
	Errors    []*Error
	ErrorsMap map[string][]*Error
}

// HasErrors 是否有 Errors 信息
func (valid *Validation) HasErrors() bool {
	return len(valid.Errors) > 0
}

// setError 设置 Error
func (valid *Validation) setError(err *Error) {
	valid.Errors = append(valid.Errors, err)
	if valid.ErrorsMap == nil {
		valid.ErrorsMap = make(map[string][]*Error)
	}
	if _, ok := valid.ErrorsMap[err.Field]; !ok {
		valid.ErrorsMap[err.Field] = []*Error{}
	}
	valid.ErrorsMap[err.Field] = append(valid.ErrorsMap[err.Field], err)
}

// SetError 设置 Error
func (valid *Validation) SetError(fieldName string, name string, msg string) {
	valid.setError(&Error{Field: fieldName, Name: name, Message: msg})
}

// Valid 验证
func (valid *Validation) Valid(obj interface{}) (b bool, err error) {
	var vOf reflect.Value
	var tOf reflect.Type
	if _, ok := obj.(reflect.Value); ok {
		vOf = obj.(reflect.Value)
		tOf = vOf.Type()
	} else {
		tOf, vOf = reflect.TypeOf(obj), reflect.ValueOf(obj)
	}

	switch {
	case isStruct(tOf):
	case isStructPtr(tOf):
		tOf, vOf = tOf.Elem(), vOf.Elem()
	default:
		err = fmt.Errorf("%v 必须是 结构体 或者 结构体指针", obj)
		return
	}

	for i := 0; i < tOf.NumField(); i++ {
		var vfs []ValidFunc
		if vfs, err = matchValidFunc(tOf.Field(i)); err != nil {
			return
		}
		for _, vf := range vfs {
			if _, err = validFuncMap.Call(vf.Name, valid, tOf.Field(i), vOf.Field(i), vf.Params); err != nil {
				return
			}
		}
	}

	if form, ok := obj.(ValidCustom); ok {
		form.Valid(valid)
	}

	return !valid.HasErrors(), nil
}

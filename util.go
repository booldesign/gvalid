package gvalid

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/1/1 11:39
 * @Desc:
 */

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func isStructOrStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Struct || t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

// matchValidFunc 匹配验证 func
func matchValidFunc(f reflect.StructField) (vfs []ValidFunc, err error) {
	tag := f.Tag.Get(ValidTag)
	if tag == "" || tag == "-" {
		return
	}

	if vfs, tag, err = parseRegexFunc(tag, f.Name); err != nil {
		return
	}

	for _, rule := range strings.Split(tag, TagSep) {
		if rule == "" {
			continue
		}
		var vf ValidFunc
		if vf, err = parseFunc(rule); err != nil {
			return
		}
		vfs = append(vfs, vf)
	}
	return
}

// ValidFunc 验证函数
type ValidFunc struct {
	Name   string
	Params interface{}
}

var (
	InvalidExpr = errors.New("invalid expr")
)

// parseRegexFunc 匹配正则 func
func parseRegexFunc(tag, key string) (vfs []ValidFunc, str string, err error) {
	tag = strings.TrimSpace(tag)
	index := strings.Index(tag, RegexTagStart)
	if index == -1 {
		str = tag
		return
	}
	end := strings.LastIndex(tag, RegexTagEnd)
	if end < index {
		err = InvalidExpr
		return
	}
	reg, err := regexp.Compile(tag[index+len(RegexTagStart) : end])
	if err != nil {
		return
	}
	vfs = []ValidFunc{{ValidFuncPrefix + RegexFunc, reg.String()}}
	str = strings.TrimSpace(tag[:index]) + strings.TrimSpace(tag[end+len(RegexTagEnd):])
	return
}

// parseFunc 匹配要验证的 func
func parseFunc(rule string) (v ValidFunc, err error) {
	ruleSlice := strings.Split(strings.TrimSpace(rule), AssignSep)
	var params string
	if len(ruleSlice) == 2 {
		params = ruleSlice[1]
	}
	v = ValidFunc{ValidFuncPrefix + toUpperCamel(ruleSlice[0]), params}
	return
}

// toUpperCamel 首字母大写
func toUpperCamel(s string) string {
	if s == "" {
		return s
	}
	if r := rune(s[0]); r >= 97 && r <= 122 {
		s = strings.ToUpper(string(r)) + s[1:]
	}
	return s
}

// toLowerCamel 首字母小写
func toLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if r := rune(s[0]); r >= 65 && r <= 90 {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return s
}

// ValidIdCardCode 验证身份证号码
func ValidIdCardCode(val string) bool {
	if len(val) != 18 {
		return false
	}

	// 加权因子
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	//校验码范围
	checkCode := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	sum := 0
	for i := 0; i < 17; i++ {
		v, _ := strconv.Atoi(string(val[i]))
		sum = sum + v*weight[i]
	}
	return string(val[17]) == checkCode[sum%11]
}

// ValidUsername 检查用户名格式
func ValidUsername() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			flag := true
			isIncludeLetter := false
			if len(val) >= 5 && len(val) <= 25 {
				for _, v := range val {
					if !((v >= 65 && v <= 90) || (v >= 97 && v <= 122) || (v >= 48 && v <= 57) || v == 95) {
						flag = false
						break
					}

					if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
						isIncludeLetter = true
					}
				}
				// 不能以下划线开头和结尾
				if val[0] == 95 || val[len(val)-1] == 95 {
					flag = false
				}
			} else {
				flag = false
			}

			if flag && isIncludeLetter == false {
				flag = false
			}

			return flag
		},
		"5~25位数字字母下划线组合，必须包含字母，不能以下划线开头和结尾",
	}
}

// ValidationFuncRule 自定义函数
type ValidationFuncRule struct {
	Func func(val string) bool
	Msg  string
}

// ValidPassword 验证密码格式
func ValidPassword() ValidationFuncRule {
	return ValidationFuncRule{
		func(val string) bool {
			m := 0
			if len(val) >= 8 && len(val) <= 32 {
				for _, v := range val {
					// [0-9A-Za-z!"#$%&'()*+,-./:;<=>?@ [\]^_`{|}~
					if v >= 48 && v <= 57 {
						m = m | 1
					} else if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
						m = m | 2
					} else if (v >= 33 && v <= 47) || (v >= 58 && v <= 64) || (v >= 91 && v <= 96) || (v >= 123 && v <= 126) {
						m = m | 4
					} else {
						m = 0
						break
					}
				}
				if m != 0 && m != 1 && m != 2 && m != 4 {
					return true
				}
			}
			return false
		},
		"8~32位字母,数字,特殊符号的组合，且包含2种以上组合",
	}
}

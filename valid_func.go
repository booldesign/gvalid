package gvalid

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/1/1 13:29
 * @Desc:
 */

const (
	DefaultLocal    = "Asia/Shanghai"
	DefaultDate     = "2006-01-02"
	DefaultDatetime = "2006-01-02 15:04:05"
)

var (
	validFuncMap = make(Funcs)
)

var (
	loc, _ = time.LoadLocation(DefaultLocal)
)

func SetLocal(local string) (err error) {
	loc, err = time.LoadLocation(local)
	return
}

func init() {
	t := reflect.TypeOf(&Validation{})
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if strings.HasPrefix(m.Name, validFuncPrefix) {
			validFuncMap[m.Name] = m.Func
		}
	}
}

// RuleRequired 必填
func (valid *Validation) RuleRequired(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValCanNotEmpty)
	}
	return
}

// RuleGt 大于
// 支持: int8, int32, int, int64,
// float32, float64,
// string, slice, map, array
func (valid *Validation) RuleGt(tOf reflect.StructField, vOf reflect.Value, size string) {
	if vOf.IsZero() {
		return
	}

	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}

	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Kind() {
	case reflect.Int8, reflect.Int32, reflect.Int, reflect.Int64:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s < int(vOf.Int()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotGtInt, s))
	case reflect.Float32, reflect.Float64:
		s, err := strconv.ParseFloat(size, 64)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s < vOf.Float() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotGtFloat, s))
	case reflect.String:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s < utf8.RuneCountInString(vOf.String()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotGtString, s))
	case reflect.Slice, reflect.Map, reflect.Array:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s < vOf.Len() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotGtSlice, s))
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "gt", vOf.String()))
	}

	return
}

// RuleGte 大于等于
// 支持: int8, int32, int, int64,
// float32, float64,
// string, slice, map, array
func (valid *Validation) RuleGte(tOf reflect.StructField, vOf reflect.Value, size string) {
	if vOf.IsZero() {
		return
	}

	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}

	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Kind() {
	case reflect.Int8, reflect.Int32, reflect.Int, reflect.Int64:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s <= int(vOf.Int()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotGteInt, s))
	case reflect.Float32, reflect.Float64:
		s, err := strconv.ParseFloat(size, 64)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s <= vOf.Float() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotGteFloat, s))
	case reflect.String:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s <= utf8.RuneCountInString(vOf.String()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotGteString, s))
	case reflect.Slice, reflect.Map, reflect.Array:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s <= vOf.Len() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotGteSlice, s))
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "gte", vOf.String()))
	}

	return
}

// RuleLt 小于
// 支持: int8, int32, int, int64,
// float32, float64,
// string, slice, map, array
func (valid *Validation) RuleLt(tOf reflect.StructField, vOf reflect.Value, size string) {
	if vOf.IsZero() {
		return
	}
	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}

	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Kind() {
	case reflect.Int8, reflect.Int32, reflect.Int, reflect.Int64:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s > int(vOf.Int()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLtInt, s))
	case reflect.Float32, reflect.Float64:
		s, err := strconv.ParseFloat(size, 64)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s > vOf.Float() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLtFloat, s))
	case reflect.String:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s > utf8.RuneCountInString(vOf.String()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLtString, s))
	case reflect.Slice, reflect.Map, reflect.Array:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s > vOf.Len() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLtSlice, s))
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "lt", vOf.String()))
	}

	return
}

// RuleLte 小于等于
// 支持: int8, int32, int, int64,
// float32, float64,
// string, slice, map, array
func (valid *Validation) RuleLte(tOf reflect.StructField, vOf reflect.Value, size string) {
	if vOf.IsZero() {
		return
	}
	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}

	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Kind() {
	case reflect.Int8, reflect.Int32, reflect.Int, reflect.Int64:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s >= int(vOf.Int()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLteInt, s))
	case reflect.Float32, reflect.Float64:
		s, err := strconv.ParseFloat(size, 64)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s >= vOf.Float() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLteFloat, s))
	case reflect.String:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s >= utf8.RuneCountInString(vOf.String()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLteString, s))
	case reflect.Slice, reflect.Map, reflect.Array:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s >= vOf.Len() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLteSlice, s))
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "lte", vOf.String()))
	}

	return
}

// RuleLen 字符串长度或数值等于期望值
// 支持: string, slice, map, array
func (valid *Validation) RuleLen(tOf reflect.StructField, vOf reflect.Value, size string) {
	if vOf.IsZero() {
		return
	}

	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}

	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Kind() {
	case reflect.String:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s == utf8.RuneCountInString(vOf.String()) {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLenString, s))
	case reflect.Slice, reflect.Map, reflect.Array:
		s, err := strconv.Atoi(size)
		if err != nil {
			valid.SetError(name, tag, ValidateValTypeErr)
			return
		}
		if s == vOf.Len() {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotLenSlice, s))
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "len", vOf.String()))
	}
	return
}

// RuleDate 日期格式
// 支持: string
func (valid *Validation) RuleDate(tOf reflect.StructField, vOf reflect.Value, format string) {
	if vOf.IsZero() {
		return
	}

	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}

	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Kind() {
	case reflect.String:
		if _, err := time.ParseInLocation(format, vOf.String(), loc); err == nil {
			return
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValDateFormatErr, format))
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "date", vOf.String()))
	}

	return
}

// RuleIn in
// 支持: int8, int32, int, int64,
// string
func (valid *Validation) RuleIn(tOf reflect.StructField, vOf reflect.Value, size string) {

	if vOf.IsZero() {
		return
	}
	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}

	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Kind() {
	case reflect.Int8, reflect.Int32, reflect.Int, reflect.Int64:
		for _, v := range strings.Split(size, " ") {
			s, err := strconv.Atoi(v)
			if err != nil {
				valid.SetError(name, tag, ValidateValTypeErr)
				return
			}
			if int64(s) == vOf.Int() {
				return
			}
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotExists, size))
	case reflect.String:
		for _, v := range strings.Split(size, " ") {
			if v == vOf.String() {
				return
			}
		}
		valid.SetError(name, tag, fmt.Sprintf(ValidateValNotExists, size))
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "in", vOf.String()))
	}

	return
}

// RuleSin sliceInSlice
// 支持: []int, []int64,
// []string
func (valid *Validation) RuleSin(tOf reflect.StructField, vOf reflect.Value, size string) {
	if vOf.IsZero() {
		return
	}

	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}

	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch tOf.Type.String() {
	case "[]int":
		i := map[int]struct{}{}
		for _, v := range strings.Split(size, " ") {
			val, err := strconv.Atoi(v)
			if err != nil {
				valid.SetError(name, tag, ValidateValTypeErr)
				return
			}
			i[val] = struct{}{}
		}
		for _, v := range vOf.Interface().([]int) {
			if _, ok := i[v]; !ok {
				valid.SetError(name, tag, fmt.Sprintf(ValidateValNotExistsSlice, size))
				return
			}
		}
	case "[]int64":
		i := map[int]struct{}{}
		for _, v := range strings.Split(size, " ") {
			val, err := strconv.Atoi(v)
			if err != nil {
				valid.SetError(name, tag, ValidateValTypeErr)
				return
			}
			i[val] = struct{}{}
		}
		for _, v := range vOf.Interface().([]int64) {
			if _, ok := i[int(v)]; !ok {
				valid.SetError(name, tag, fmt.Sprintf(ValidateValNotExistsSlice, size))
				return
			}
		}
	case "[]string":
		i := map[string]struct{}{}
		for _, v := range strings.Split(size, " ") {
			i[v] = struct{}{}
		}
		for _, v := range vOf.Interface().([]string) {
			if _, ok := i[v]; !ok {
				valid.SetError(name, tag, fmt.Sprintf(ValidateValNotExistsSlice, size))
				return
			}
		}
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "sin", vOf.String()))
	}
}

// RuleDive 嵌套验证
func (valid *Validation) RuleDive(tOf reflect.StructField, vOf reflect.Value, _ string) {

	if vOf.Type().Kind() == reflect.Slice {
		l := vOf.Len()
		// 仅支持 slice 类型的 struct
		if l > 0 && !isStructOrStructPtr(vOf.Index(0).Type()) {
			return
		}
		for i := 0; i < l; i++ {
			_, _ = valid.Valid(vOf.Index(i))
		}
	} else if isStruct(tOf.Type) {
		_, _ = valid.Valid(vOf)
	} else if isStructPtr(tOf.Type) {
		if vOf.IsZero() {
			vOf.Set(reflect.New(tOf.Type.Elem()))
		}
		_, _ = valid.Valid(vOf)
	}

	return
}

// RuleRegex 正则
// 支持: string
// regex pattern string must in "(//)"
// 注意反斜杠需转译 如 \\d
func (valid *Validation) RuleRegex(tOf reflect.StructField, vOf reflect.Value, pattern string) {
	if vOf.IsZero() {
		return
	}
	if m, _ := regexp.MatchString(pattern, vOf.String()); !m {
		valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValNotFormatErr)
	}
	return
}

// RuleEmail 邮箱验证
func (valid *Validation) RuleEmail(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}
	if b := emailPattern.MatchString(vOf.String()); !b {
		valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValNotFormatErr)
	}
	return
}

// RuleMobile 手机验证
func (valid *Validation) RuleMobile(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}
	if b := mobilePattern.MatchString(vOf.String()); !b {
		valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValNotFormatErr)
	}
	return
}

// RuleBase64 base64 验证
func (valid *Validation) RuleBase64(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}
	if b := base64Pattern.MatchString(vOf.String()); !b {
		valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValNotFormatErr)
	}
	return
}

// RuleIp ip 验证
func (valid *Validation) RuleIp(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}
	if b := ipPattern.MatchString(vOf.String()); !b {
		valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValNotFormatErr)
	}
	return
}

// RuleUrl url验证
func (valid *Validation) RuleUrl(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}
	if b := urlPattern.MatchString(vOf.String()); !b {
		valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValNotFormatErr)
	}
	return
}

// RuleIdCard 身份证验证
func (valid *Validation) RuleIdCard(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}
	if b := ValidIdCardCode(vOf.String()); !b {
		valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValNotFormatErr)
	}
	return
}

// RuleNumeric 纯数字字符
func (valid *Validation) RuleNumeric(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}
	for _, v := range vOf.String() {
		if v < 48 || v > 57 {
			valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), ValidateValNotNumericErr)
			break
		}
	}
	return
}

// RuleDefault 默认值
// 支持: int8, int32, int, int64,
// string
func (valid *Validation) RuleDefault(tOf reflect.StructField, vOf reflect.Value, def string) {
	if vOf.IsZero() {
		if vOf.Type().Kind() == reflect.Ptr {
			vOf = vOf.Elem()
		}
		name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
		switch vOf.Type().Kind() {
		case reflect.Int8, reflect.Int32, reflect.Int, reflect.Int64:
			val, err := strconv.Atoi(def)
			if err != nil {
				valid.SetError(name, tag, ValidateValTypeErr)
				return
			}
			vOf.SetInt(int64(val))
		case reflect.String:
			vOf.SetString(def)
		}
		return
	}
	return
}

// RuleDistinct 重复值验证
// 支持: []int, []int64,
// []string
func (valid *Validation) RuleDistinct(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}
	if vOf.Kind() != reflect.Slice {
		valid.SetError(tOf.Name, "distinct", ValidateValTypeErr)
	}
	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Type().String() {
	case "[]int":
		i := map[int]struct{}{}
		for _, v := range vOf.Interface().([]int) {
			if _, ok := i[v]; ok {
				valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), fmt.Sprintf(ValidateValMustDistinct, vOf.Interface()))
				return
			}
			i[v] = struct{}{}
		}
	case "[]int64":
		i := map[int64]struct{}{}
		for _, v := range vOf.Interface().([]int64) {
			if _, ok := i[v]; ok {
				valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), fmt.Sprintf(ValidateValMustDistinct, vOf.Interface()))
				return
			}
			i[v] = struct{}{}
		}
	case "[]string":
		i := map[string]struct{}{}
		for _, v := range vOf.Interface().([]string) {
			if _, ok := i[v]; ok {
				valid.SetError(tOf.Name, tOf.Tag.Get(defaultNameTag), fmt.Sprintf(ValidateValMustDistinct, vOf.Interface()))
				return
			}
			i[v] = struct{}{}
		}

	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "distinct", vOf.Interface()))
	}
	return
}

// RuleTrimSpace 字符串去除空格
// 支持: string
func (valid *Validation) RuleTrimSpace(tOf reflect.StructField, vOf reflect.Value, _ string) {
	if vOf.IsZero() {
		return
	}

	if vOf.Kind() == reflect.Ptr {
		vOf = vOf.Elem()
	}
	name, tag := tOf.Name, tOf.Tag.Get(defaultNameTag)
	switch vOf.Kind() {
	case reflect.String:
		vOf.SetString(strings.TrimSpace(vOf.String()))
	default:
		valid.SetError(name, tag, fmt.Sprintf(ValidateMethodNotAllowSth, "trimSpace", vOf.String()))
	}
	return
}

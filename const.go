package gvalid

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/1/1 12:46
 * @Desc:
 */

const (
	defaultTagName    = "valid"
	defaultNameTag    = "name"
	tagSep            = ","
	tagKeySep         = "="
	skipValidationTag = "-"
	validFuncPrefix   = "Rule"
)

const (
	RegexFunc     = "Regex"
	RegexTagStart = "regex=(/"
	RegexTagEnd   = "/)"
)

const (
	ValidateMethodNotAllowSth = "验证方法 %s 不允许 %v"
	ValidateValTypeErr        = "验证规则写法有误"
	ValidateValCanNotEmpty    = "不能为空或零值"
	ValidateValNotGtString    = "长度必须是大于 %d"
	ValidateValNotGtSlice     = "长度必须是大于 %d"
	ValidateValNotGtInt       = "必须是大于 %d"
	ValidateValNotGtFloat     = "必须是大于 %.2f"
	ValidateValNotGteString   = "长度必须是大于等于 %d"
	ValidateValNotGteSlice    = "长度必须是大于等于 %d"
	ValidateValNotGteInt      = "必须是大于等于 %d"
	ValidateValNotGteFloat    = "必须是大于等于 %.2f"
	ValidateValNotLtString    = "长度必须是小于 %d"
	ValidateValNotLtSlice     = "长度必须是小于 %d"
	ValidateValNotLtInt       = "必须是小于 %d"
	ValidateValNotLtFloat     = "必须是小于 %.2f"
	ValidateValNotLteString   = "长度必须是小于等于 %d"
	ValidateValNotLteSlice    = "长度必须是小于等于 %d"
	ValidateValNotLteInt      = "必须是小于等于 %d"
	ValidateValNotLteFloat    = "必须是小于等于 %.2f"
	ValidateValNotLenString   = "长度必须是等于 %d"
	ValidateValNotLenSlice    = "长度必须是等于 %d"
	ValidateValDateFormatErr  = "时间格式错误 %s"
	ValidateValNotExists      = "必须是 %s 其中一个"
	ValidateValNotExistsSlice = "必须是 %s 其中一个或多个"
	ValidateValNotFormatErr   = "格式错误"
	ValidateValNotNumericErr  = "必须是有效的数字字符"
	ValidateValMustDistinct   = "含有重复的值 %+v"
)

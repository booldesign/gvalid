package gvalid

import "fmt"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/1/1 12:53
 * @Desc:
 */

// Error ...
type Error struct {
	Field, Name, Message string
}

// String Return Message
func (e *Error) String() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s %s", e.Name, e.Message)
}

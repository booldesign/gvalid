# Package gValid

gValid 验证器，可以用于结构体字段值的验证，也可以用来表单验证。

它具有以下独特的特点:
* 能够深入映射键和值进行验证
* 内置常用验证方法
* 支持自定义校验，需实现 ValidCustom 接口

# Installation

```
go get github.com/booldesign/gvalid
import "github.com/booldesign/gvalid"
```

# 常用tag使用说明

| tag           | 说明                                          | 使用示例                               |
| ------------- | ----------------------------------------     | -------------------------------------- |
| -             | 不校验                                         | valid:"-"                            |                                     
| required      | 必填字段,且不能为零值                            | valid:"required"                    |
| default       | 默认值,不和required共用,可用于非指针的基础类型 int/int64/string  | valid:"default"               |
|               |                                              |                                        |
| gt            | int/int64/string/float64/float32/slice 大于   | valid:"gt=0"                        |
| gte           | 同上 大于等于                                   | valid:"gte=0"                       |
| lt            | int/int64/string/float64/float32/slice 小于   | valid:"lt=10"                       |
| lte           | 同上 小于等于                                   | valid:"lte=10"                      |
| len           | string/slice 指定长度                          | valid:"len=1"                       |
|               |                                              |                                        |
| in            | 其中之一                                       | valid:"in=5 7 9"                     |
| sin           | slice 都在可选范围 仅支持:[]string/[]int/[]int64 | valid:"sin=5 7 9"                    |
| distinct      | 不能重复                                       | valid:"distinct"                    |
|               |                                               |                                        |
| date          | 校验日期                                        | valid:"date=2006-01-02" 格式可自定义  |
| numeric       | 纯数字字符                                      | valid:"numeric"                       |
|               |                                               |                                        |
| regex         | 正则                                           | valid:"regex=(//)"                      |
| mobile        | 手机号码                                       | valid:"mobile"                          |
| email         | 校验Email地址                                  | valid:"email"                       |
| idCard        | 校验身份证                                      | valid:idCard"                          |
| base64        | 校验base64值                                   | valid:"base64"                      |
| ip            | 校验IP地址                                     | valid:"ip"                          |
|               |                                              |                                        |
| dive          | 向下延伸验证   | valid:"required,dive"`         |



# 扩展：使用自定义函数示例

```
// 自定义校验函数
实现 ValidCustom 接口 即可

演示：

type Account struct {
	Username   string `valid:"required" name:"用户名"`
	Password   string `valid:"required" name:"密码"`
	RePassword string `valid:"required" name:"确认密码"`
}

func (a *Account) Valid(v *gvalid.Validation) {
	passFunc := gvalid.ValidUsername()
	if a.Username != "" && !passFunc.Func(a.Username) {
		v.SetError("Username", "用户名", passFunc.Msg)
	}

	passFunc = gvalid.ValidPassword()
	if a.Password != "" && !passFunc.Func(a.Password) {
		v.SetError("Password", "密码", passFunc.Msg)
	}

	if a.Password != a.RePassword {
		v.SetError("RePassword", "确认密码", "必须 和 密码一致")
	}
}
```


# Error Return Value

```
v = &gvalid.Validation{}
b, err = v.Valid(u)
if err != nil {
    t.Fatal("result err:", err)
}
if !b {
    t.Fatal("result valid err:", v.ErrorsMap)
}

```

# Question 字段必传，用指针可以解决零值问题

```
type RequestForm struct {
    Status *int `json:"status" valid:"required" name:"状态"`
}
```

# Benchmarks
Run on MacBook Pro (15-inch, 2018) go version go1.16.6 darwin/amd64

```
goos: darwin
goarch: amd64
pkg: github.com/booldesign/gvalid
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkRequest-12     	 1142132	      1044 ns/op
BenchmarkGt-12          	 1000000	      1151 ns/op
BenchmarkLt-12          	 1000000	      1159 ns/op
BenchmarkLen-12         	 1000000	      1185 ns/op
BenchmarkDate-12        	  800492	      1458 ns/op
BenchmarkIn-12          	  958705	      1240 ns/op
BenchmarkSin-12         	  837912	      1377 ns/op
BenchmarkRegex-12       	  173752	      6548 ns/op
BenchmarkEmail-12       	  798022	      1480 ns/op
BenchmarkMobile-12      	  875113	      1254 ns/op
BenchmarkUrl-12         	  981883	      1261 ns/op
BenchmarkIdCard-12      	  838698	      1425 ns/op
BenchmarkCustom-12      	  594772	      2034 ns/op
BenchmarkNumeric-12     	 1000000	      1050 ns/op
BenchmarkDefault-12     	 1000000	      1175 ns/op
BenchmarkDistinct-12    	  942357	      1211 ns/op

```

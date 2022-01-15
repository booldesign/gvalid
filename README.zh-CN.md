# Package gValid
[English](README.md) | 简体中文

![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

gValid 验证器，可以用于结构体字段值的验证，也可以用来表单验证。

它具有以下独特的特点:
  * 能够深入嵌套结构体的映射键和值进行验证
  * 内置常用验证方法，满足日常使用
  * 使用别名验证标签，允许将多个验证方法映射到单个Field，以便更轻松地定义结构上的验证
  * 支持自定义校验，仅需实现 ValidCustom 接口
  * 扩展快捷方便

## 安装

获取包
```
go get github.com/booldesign/gvalid
```

导入包到我们的项目
```
import "github.com/booldesign/gvalid"
```


## 常用tag使用说明

| 标签           | 说明                                          | 使用示例                               |
| ------------- | ----------------------------------------     | -------------------------------------- |
| -             | 不校验                                         | valid:"-"                            |                                     
| required      | 必填字段,且不能为零值                            | valid:"required"                    |
| default       | 默认值,不和required共用,可用于非指针的基础类型 int/int64/string  | valid:"default"               |
| trimSpace     | 去除空格                                       | valid:"trimSpace"               |
|               |                                              |                                        |
| gt            | 大于, 支持:int/int64/string/float64/float32/slice/map/array    | valid:"gt=0"                        |
| gte           | 同上 大于等于                                   | valid:"gte=0"                       |
| lt            | 小于, 支持:int/int64/string/float64/float32/slice/map/array    | valid:"lt=10"                       |
| lte           | 同上 小于等于                                   | valid:"lte=10"                      |
| len           | 指定长度, 支持:string/slice/map/array                           | valid:"len=1"                       |
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
| dive          | 向下延伸验证，匿名结构体默认自带                      | valid:"required,dive"`         |


## 快速开始
```
type WUser struct {
    Name string `valid:"required" name:"姓名"`
}

u := &WUser{
    Name: "BoolDesign",
}

v := &Validation{}
b, err := v.Valid(u)
if err != nil {
    // TODO: handle error
    panic(err)
}
if !b {
    // TODO: 捕获验证错误信息
    fmt.Println(v.ErrorsMap)
}
```

### 扩展：使用自定义函数示例

```
// 自定义校验函数
// 实现 ValidCustom 接口 即可

// 演示：

type Account struct {
	Username   string `valid:"required,gte=6,lte=20" name:"用户名"`
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

## 常见问题(FAQ)

#### 问题 1: 字段必传，用指针可以解决零值问题
```
type RequestForm struct {
    Status *int `json:"status" valid:"required" name:"状态"`
}
```

## Benchmarks
Run on MacBook Pro (15-inch, 2018) go version go1.16.6 darwin/amd64

```
goos: darwin
goarch: amd64
pkg: github.com/booldesign/gvalid
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkRequest-12      	 1142695	      1038 ns/op
BenchmarkGt-12           	 1000000	      1170 ns/op
BenchmarkLt-12           	 1000000	      1161 ns/op
BenchmarkLen-12          	 1000000	      1205 ns/op
BenchmarkDate-12         	  822451	      1473 ns/op
BenchmarkIn-12           	  950832	      1252 ns/op
BenchmarkSin-12          	  848228	      1385 ns/op
BenchmarkRegex-12        	  173545	      6673 ns/op
BenchmarkEmail-12        	  815535	      1473 ns/op
BenchmarkMobile-12       	  999700	      1240 ns/op
BenchmarkUrl-12          	  973112	      1225 ns/op
BenchmarkIdCard-12       	  815989	      1411 ns/op
BenchmarkCustom-12       	  586150	      1995 ns/op
BenchmarkNumeric-12      	 1000000	      1079 ns/op
BenchmarkDefault-12      	 1000000	      1192 ns/op
BenchmarkDistinct-12     	  963518	      1264 ns/op
BenchmarkTrimSpace-12    	 1031680	      1135 ns/op
```

## License
根据 MIT 许可证分发，请参阅代码中的许可证文件以获取更多详细信息。

## 问题反馈

如果您发现 bug 请及时提 issue，我会尽快确认并修改。

如果觉得对您有用的话，有劳点一下 star⭐，谢谢！
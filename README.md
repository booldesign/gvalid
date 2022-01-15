# Package gValid
English | [简体中文](README.zh-CN.md)

![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The gValid validator can be used for the validation of struct field values as well as for form validation.

It has the following unique features:
* Ability to dive into both map keys and values for validation
* Built-in common authentication methods
* Alias validation tags, which allows for mapping of several validations to a single tag for easier defining of validations on structs
* Support custom validation, need to implement ValidCustom interface
* Expansion is quick and easy

## Installation

Use go get.
```
go get github.com/booldesign/gvalid
```

Then import the gValid package into your own code.
```
import "github.com/booldesign/gvalid"
```


## Baked-in Validations

| Tag           | Description                                          | Example of use                               |
| ------------- | ----------------------------------------     | -------------------------------------- |
| -             | Do not check                                         | valid:"-"                            |                                     
| required      | Required                            | valid:"required"                    |
| default       | Default, not shared with Required, supported:int/int64/string  | valid:"default"               |
| trimSpace     | Trim Space                                       | valid:"trimSpace"               |
|               |                                              |                                        |
| gt            | int/int64/string/float64/float32/slice/map/array Greater than   | valid:"gt=0"                        |
| gte           | Greater than or equal                                   | valid:"gte=0"                       |
| lt            | int/int64/string/float64/float32/slice/map/array Less Than   | valid:"lt=10"                       |
| lte           | Less Than or Equal                                   | valid:"lte=10"                      |
| len           | Length, supported:string/slice/map/array                           | valid:"len=1"                       |
|               |                                              |                                        |
| in            | In                                       | valid:"in=5 7 9"                     |
| sin           | slice In, supported:[]string/[]int/[]int64 | valid:"sin=5 7 9"                    |
| distinct      | Distinct                                       | valid:"distinct"                    |
|               |                                               |                                        |
| date          | Date                                        | valid:"date=2006-01-02"  |
| numeric       | Numeric                                      | valid:"numeric"                       |
|               |                                               |                                        |
| regex         | Regex                                           | valid:"regex=(//)"                      |
| mobile        | Mobile String                                        | valid:"mobile"                          |
| email         | E-mail String                                  | valid:"email"                       |
| idCard        | Chinese ID card                                      | valid:idCard"                          |
| base64        | Base64 String                                   | valid:"base64"                      |
| ip            | Internet Protocol Address IP                                     | valid:"ip"                          |
|               |                                              |                                        |
| dive          | Dive                      | valid:"required,dive"`         |


## Quick Start
```
type WUser struct {
    Name string `valid:"required" name:"name"`
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
    // TODO: validation error messages
    fmt.Println(v.ErrorsMap)
}
```

### Use custom validation rules

```
// Custom check function
// Just implement the ValidCustom interface

// Demo:

type Account struct {
	Username   string `valid:"required,gte=6,lte=20" name:"username"`
	Password   string `valid:"required" name:"password"`
	RePassword string `valid:"required" name:"rePassword"`
}

func (a *Account) Valid(v *gvalid.Validation) {
	passFunc := gvalid.ValidUsername()
	if a.Username != "" && !passFunc.Func(a.Username) {
		v.SetError("Username", "username", passFunc.Msg)
	}

	passFunc = gvalid.ValidPassword()
	if a.Password != "" && !passFunc.Func(a.Password) {
		v.SetError("Password", "password", passFunc.Msg)
	}

	if a.Password != a.RePassword {
		v.SetError("RePassword", "rePassword", "must be the same as the password")
	}
}
```

## FAQ

#### Question 1: Fields must be passed, and pointers can be used to solve the zero-value problem
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
Distributed under MIT License, please see license file within the code for more details.

## Give a Star! ⭐

If you like or are using this project to learn or start your solution, please give it a star. Thanks!
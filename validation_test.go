package gvalid

import (
	"flag"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/1/1 16:56
 * @Desc:
 */

var mode string

func init() {
	testing.Init()

	flag.StringVar(&mode, "mode", "PASS", "模式")
	flag.Parse()
}

func TestRequest(t *testing.T) {
	type Address struct {
		Province string `valid:"required" name:"省"`
		City     string `valid:"required" name:"市"`
	}
	type WUser struct {
		Name    string     `valid:"required" name:"姓名"`
		Age     int        `valid:"required" name:"年龄"`
		Num     *int       `valid:"required" name:"数量"`
		Height  float64    `valid:"required" name:"身高"`
		Hobby   []string   `valid:"required" name:"爱好"`
		Status  []int      `valid:"required" name:"状态"`
		Address []*Address `valid:"required,dive" name:"地址"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{Address: []*Address{
			{},
		}}
	default:
		var n = 1
		u = &WUser{
			Name:   "222",
			Num:    &n,
			Age:    180,
			Height: 250.01,
			Hobby:  []string{"篮球", "音乐", "健身"},
			Status: []int{1, 2, 3},
			Address: []*Address{
				{Province: "shanghai", City: "shanghai2"},
			},
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkRequest(b *testing.B) {
	type WUser struct {
		Name string `valid:"required" name:"姓名"`
	}

	var u *WUser

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Name: "222",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestGt(t *testing.T) {
	type Address struct {
		Province string `valid:"gt=2" name:"省"`
		City     string `valid:"gt=2" name:"市"`
	}
	type WUser struct {
		Name    string     `valid:"gt=2" name:"姓名"`
		Age     int        `valid:"gt=2" name:"年龄"`
		Num     *int       `valid:"gt=2" name:"数量"`
		Height  float64    `valid:"gt=50.1" name:"身高"`
		Hobby   []string   `valid:"gt=1" name:"爱好"`
		Status  []int      `valid:"gt=1" name:"状态"`
		Address []*Address `valid:"gt=1,dive" name:"地址"`
	}

	var u *WUser
	var n int

	switch mode {
	case "FAIL":
		n = 1
		u = &WUser{
			Name:   "we",
			Num:    &n,
			Age:    2,
			Height: 50.01,
			Hobby:  []string{"篮球"},
			Status: []int{1},
			Address: []*Address{
				{"sh", "sh"},
			},
		}
	default:
		n = 300
		u = &WUser{
			Name:   "wei",
			Num:    &n,
			Age:    180,
			Height: 250.01,
			Hobby:  []string{"篮球", "音乐", "健身"},
			Status: []int{1, 2, 3},
			Address: []*Address{
				{"shanghai", "shanghai"},
				{"江苏省", "常州市"},
			},
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkGt(b *testing.B) {
	type WUser struct {
		Age int `valid:"gt=2" name:"年龄"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Age: 180,
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestLt(t *testing.T) {
	type Address struct {
		Province string `valid:"lt=4" name:"省"`
		City     string `valid:"lt=4" name:"市"`
	}
	type WUser struct {
		Name    string     `valid:"lt=10" name:"姓名"`
		Age     int        `valid:"lt=110" name:"年龄"`
		Num     *int       `valid:"lt=30" name:"数量"`
		Height  float64    `valid:"lt=250" name:"身高"`
		Hobby   []string   `valid:"lt=3" name:"爱好"`
		Status  []int      `valid:"lt=3" name:"状态"`
		Address []*Address `valid:"lt=3,dive" name:"地址"`
	}

	var u *WUser
	var n int

	switch mode {
	case "FAIL":
		n = 31
		u = &WUser{
			Name:   "wei1234567",
			Num:    &n,
			Age:    120,
			Height: 259.01,
			Hobby:  []string{"篮球", "音乐", "健身"},
			Status: []int{1, 2, 3},
			Address: []*Address{
				{"上海市1", "上海市1"},
				{"江苏省1", "常州市1"},
				{"山东省1", "临沂市1"},
			},
		}
	default:
		n = 29
		u = &WUser{
			Name:   "wei",
			Num:    &n,
			Age:    100,
			Height: 180.01,
			Hobby:  []string{"篮球", "音乐"},
			Status: []int{1, 2},
			Address: []*Address{
				{"上海市", "上海市"},
				{"江苏省", "常州市"},
			},
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkLt(b *testing.B) {
	type WUser struct {
		Age int `valid:"lt=110" name:"年龄"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Age: 100,
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestLen(t *testing.T) {
	type Address struct {
		Province string `valid:"len=3" name:"省"`
		City     string `valid:"len=3" name:"市"`
	}
	type WUser struct {
		Name    string     `valid:"len=2" name:"姓名"`
		Hobby   []string   `valid:"len=2" name:"爱好"`
		Status  []int      `valid:"len=2" name:"状态"`
		Address []*Address `valid:"len=2,dive" name:"地址"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			Name:   "wei1234567",
			Hobby:  []string{"篮球", "音乐", "健身"},
			Status: []int{1, 2, 3},
			Address: []*Address{
				{"上海市1", "上海市1"},
				{"江苏省1", "常州市1"},
				{"山东省1", "临沂市1"},
			},
		}
	default:
		u = &WUser{
			Name:   "we",
			Hobby:  []string{"篮球", "音乐"},
			Status: []int{1, 2},
			Address: []*Address{
				{"上海市", "上海市"},
				{"江苏省", "常州市"},
			},
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkLen(b *testing.B) {
	type WUser struct {
		Name string `valid:"len=2" name:"姓名"`
	}
	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Name: "we",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestDate(t *testing.T) {
	type WUser struct {
		StartAt  string `valid:"date=2006-01-02 15:04:05" name:"结束时间"`
		EndAt    string `valid:"date=2006-01-02 15:04:05" name:"开始时间"`
		Birthday string `valid:"date=2006-01-02" name:"生日"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			StartAt:  "2021",
			EndAt:    "2021-12-16",
			Birthday: "2021-12-16 15",
		}
	default:
		u = &WUser{
			StartAt:  "2021-12-16 15:15:40",
			EndAt:    "2021-12-17 15:15:40",
			Birthday: "2021-12-16",
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkDate(b *testing.B) {
	type WUser struct {
		Birthday string `valid:"date=2006-01-02" name:"生日"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Birthday: "2021-12-16",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestIn(t *testing.T) {
	type WUser struct {
		Type    int    `valid:"in=1 2 3" name:"类型"`
		TypeStr string `valid:"in=ENABLE DISABLE DELETE" name:"类型"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			Type:    4,
			TypeStr: "HELLO",
		}
	default:
		u = &WUser{
			Type:    1,
			TypeStr: "ENABLE",
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkIn(b *testing.B) {
	type WUser struct {
		Type int `valid:"in=1 2" name:"类型"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Type: 1,
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestSin(t *testing.T) {
	type WUser struct {
		Type    []int    `valid:"sin=1 2 3" name:"类型"`
		TypeStr []string `valid:"sin=ENABLE DISABLE DELETE" name:"类型"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			Type:    []int{4},
			TypeStr: []string{"HELLO"},
		}
	default:
		u = &WUser{
			Type:    []int{1, 2},
			TypeStr: []string{"ENABLE", "DISABLE"},
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkSin(b *testing.B) {
	type WUser struct {
		Type []int `valid:"sin=1 2" name:"类型"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Type: []int{1},
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestRegex(t *testing.T) {
	type WUser struct {
		Mobile string `valid:"regex=(/^((\\+86)|(86))?1[3456789]\\d{9}$/)" name:"手机"`
		Email  string `valid:"regex=(/^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\\.){1,4}[a-z]{2,4}$/)" name:"邮箱"`
		Num    string `valid:"regex=(/^\\d{3}$/)" name:"编号"`
	}
	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			Mobile: "135016914365",
			Email:  "wei_#123@vip.qq.com",
			Num:    "001a",
		}
	default:
		u = &WUser{
			Mobile: "13501691436",
			Email:  "wei_jianwen@vip.qq.com",
			Num:    "001",
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkRegex(b *testing.B) {
	type WUser struct {
		Num string `valid:"regex=(/^\\d{3}$/)" name:"编号"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Num: "001",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestEmail(t *testing.T) {
	type WUser struct {
		Email  string `valid:"email" name:"邮箱"`
		Email1 string `valid:"email" name:"邮箱1"`
		Email2 string `valid:"email" name:"邮箱2"`
		Email3 string `valid:"email" name:"邮箱3"`
		Email4 string `valid:"email" name:"邮箱4"`
		Email5 string `valid:"email" name:"邮箱5"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			Email:  "#1234@qq.com",
			Email2: "we`i@163.com",
			Email3: "wei_jian.wen@vip.@qq.com",
			Email4: "wei_ji&an_wen@outlook.com",
			Email5: "bool.de(sign@icloud.com",
		}
	default:
		u = &WUser{
			Email:  "1234@qq.com",
			Email2: "wei@163.com",
			Email3: "wei_jianwen@vip.qq.com",
			Email4: "wei_jianwen@outlook.com",
			Email5: "booldesign@icloud.com",
		}
	}
	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)

	}
}

func BenchmarkEmail(b *testing.B) {
	type WUser struct {
		Email string `valid:"email" name:"邮箱"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Email: "wei@163.com",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestValidateMobile(t *testing.T) {
	type WUser struct {
		Mobile  string `valid:"mobile" name:"手机"`
		Mobile1 string `valid:"mobile" name:"手机1"`
		Mobile2 string `valid:"mobile" name:"手机2"`
		Mobile3 string `valid:"mobile" name:"手机3"`
	}
	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			Mobile:  "123344",
			Mobile1: "s3456678990",
			Mobile2: "147000000010",
		}
	default:
		u = &WUser{
			Mobile:  "13501691436",
			Mobile1: "15839098799",
			Mobile2: "8615839098799",
			Mobile3: "+8615839098799",
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkMobile(b *testing.B) {
	type WUser struct {
		Mobile string `valid:"mobile" name:"手机"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Mobile: "13501691436",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestUrl(t *testing.T) {
	type WUser struct {
		Url  string `valid:"url" name:"地址"`
		Url1 string `valid:"url" name:"地址1"`
		Url2 string `valid:"url" name:"地址2"`
		Url3 string `valid:"url" name:"地址3"`
		Url4 string `valid:"url" name:"地址4"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			Url:  "ftp:://www.github.com/",
			Url1: "https:://www.github.com/",
			Url2: "http:/github.com/",
			Url3: "httpss://kubernetes.io/zh/docs/tasks/tools/install-kubectl-macos/",
			Url4: "https:///www.baidu.com/news/86328",
		}
	default:
		u = &WUser{
			Url:  "https://github.com/",
			Url1: "https://www.github.com/",
			Url2: "http://github.com/",
			Url3: "https://kubernetes.io/zh/docs/tasks/tools/install-kubectl-macos/",
			Url4: "https://code.2.baidu.com/news/86328",
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkUrl(b *testing.B) {
	type WUser struct {
		Mobile string `valid:"mobile" name:"手机"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Mobile: "13501691436",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}
func TestIdCard(t *testing.T) {
	type WUser struct {
		IdCard string `valid:"idCard" name:"身份证"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			IdCard: "1989898989898989",
		}
	default:
		u = &WUser{
			IdCard: "310104200312166537",
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkIdCard(b *testing.B) {
	type WUser struct {
		IdCard string `valid:"idCard" name:"身份证"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			IdCard: "310104200312166537",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

type Account struct {
	Username   string `valid:"required" name:"用户名"`
	Password   string `valid:"required" name:"密码"`
	RePassword string `valid:"required" name:"确认密码"`
}

func (a *Account) Valid(v *Validation) {
	passFunc := ValidUsername()
	if a.Username != "" && !passFunc.Func(a.Username) {
		v.SetError("Username", "用户名", passFunc.Msg)
	}

	passFunc = ValidPassword()
	if a.Password != "" && !passFunc.Func(a.Password) {
		v.SetError("Password", "密码", passFunc.Msg)
	}

	if a.Password != a.RePassword {
		v.SetError("RePassword", "确认密码", "必须 和 密码一致")
	}
}

func TestCustom(t *testing.T) {
	var u *Account

	switch mode {
	case "FAIL":
		u = &Account{
			Username:   "12345",
			Password:   "12244",
			RePassword: "1234567a",
		}
	default:
		u = &Account{
			Username:   "user123",
			Password:   "1234567a",
			RePassword: "1234567a",
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

type AccountBench struct {
	Password   string `valid:"required" name:"密码"`
	RePassword string `valid:"required" name:"确认密码"`
}

func (a *AccountBench) Valid(v *Validation) {
	if a.Password != a.RePassword {
		v.SetError("RePassword", "确认密码", "必须 和 密码一致")
	}
}

func BenchmarkCustom(b *testing.B) {

	var u *AccountBench

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &AccountBench{
			Password:   "1234567a",
			RePassword: "1234567a",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestNumeric(t *testing.T) {
	type WUser struct {
		SmsCode string `valid:"required,len=6,numeric" name:"验证码"`
	}
	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			SmsCode: "12345a",
		}
	default:
		u = &WUser{
			SmsCode: "123456",
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkNumeric(b *testing.B) {
	type WUser struct {
		SmsCode string `valid:"numeric" name:"验证码"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			SmsCode: "123456",
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestDefault(t *testing.T) {
	type WUser struct {
		DefInt   int    `valid:"default=1" name:"默认"`
		DefInt64 int64  `valid:"default=1" name:"默认"`
		DefStr   string `valid:"default=我是默认值" name:"默认"`
	}
	var u = &WUser{}
	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}

	Convey("test default", t, func() {
		So(u.DefInt, ShouldEqual, 1)
		So(u.DefInt64, ShouldEqual, 1)
		So(u.DefStr, ShouldEqual, "我是默认值")
	})
}

func BenchmarkDefault(b *testing.B) {
	type WUser struct {
		DefInt int `valid:"default=1" name:"默认"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

func TestDistinct(t *testing.T) {
	type WUser struct {
		Mode  []int    `valid:"distinct,sin=0 1 2" name:"模式"`
		Hobby []string `valid:"required,distinct" name:"爱好"`
	}

	var u *WUser

	switch mode {
	case "FAIL":
		u = &WUser{
			Mode:  []int{0, 1, 2, 2, 3},
			Hobby: []string{"篮球", "足球", "篮球"},
		}
	default:
		u = &WUser{
			Mode:  []int{0, 1, 2},
			Hobby: []string{"篮球", "足球", "乒乓球"},
		}
	}

	v := &Validation{}
	b, err := v.Valid(u)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

func BenchmarkDistinct(b *testing.B) {
	type WUser struct {
		Mode []int `valid:"distinct" name:"模式"`
	}

	var u *WUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u = &WUser{
			Mode: []int{0},
		}

		v := &Validation{}
		ret, err := v.Valid(u)
		if err != nil {
			b.Fatal("result err:", err)
		}
		if !ret {
			b.Fatal("result valid err:", v.ErrorsMap)
		}
	}
	b.StopTimer()
}

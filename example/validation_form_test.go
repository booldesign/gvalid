package example

import (
	"encoding/json"
	"testing"

	"github.com/booldesign/gvalid"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/1/2 17:02
 * @Desc:
 */

// Img 图片通用
type Img struct {
	ImgUrl string `json:"imgUrl" valid:"required"`
}

type GoodsSpuBase struct {
	Cate      int64  `json:"cate" valid:"required,gt=0" name:"商品分类"`
	Name      string `json:"name" valid:"required,lte=10" name:"商品名称"`
	Code      string `json:"code" valid:"required,len=10" name:"商品编号"`
	Summary   string `json:"summary" valid:"lte=255" name:"商品简介"`
	Gallery   *Img   `json:"gallery" valid:"required,dive" name:"图片"`
	SellBegin string `json:"sellBegin" valid:"required,date=2006-01-02 15:04:05" name:"销售起始日期"`
	SellEnd   string `json:"sellEnd" valid:"required,date=2006-01-02 15:04:05" name:"销售结束日期"`
	Status    int    `json:"status" valid:"required,in=1 2" name:"上架状态"`
	Mode      []int  `json:"mode"  valid:"required,distinct,sin=0 1" name:"配送方式"`
}

type GoodsSkuBase struct {
	SellingPrice     float64 `json:"sellingPrice" valid:"required,gt=0" name:"销售价格"`
	IsMemberDiscount *int    `json:"isMemberDiscount" valid:"required,in=0 1 2" name:"会员折扣"`
	Stock            int     `json:"stock" valid:"required,gte=1" name:"库存"`
}

type GoodsBase struct {
	GoodsSpuBase `valid:"dive"`
	GoodsSkuBase `valid:"dive"`
}

func TestGoodsAdd(t *testing.T) {
	data := `{"cate":1,"name":"衣服12345678","code":"1234567890","summary":"商品简介","gallery":{"imgUrl":"https://www.baidu.com/"},"sellingPrice":10.2,"isMemberDiscount":0,"stock":10,"sellBegin":"2022-01-01 10:00:00","sellEnd":"2022-01-03 23:59:59","status":1,"mode":[0,1]}`

	input := &GoodsBase{}
	err := json.Unmarshal([]byte(data), input)
	if err != nil {
		t.Fatal(err)
	}

	v := &gvalid.Validation{}
	b, err := v.Valid(input)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

type User struct {
	Username    string `json:"username" valid:"required" name:"用户名"`
	Password    string `json:"password" valid:"required" name:"密码"`
	RePassword  string `json:"rePassword" valid:"required" name:"确认密码"`
	Mobile      string `json:"mobile" valid:"required,mobile" name:"手机号"`
	MobileRegex string `json:"mobileRegex" valid:"required,regex=(/^((\\+86)|(86))?1[3456789]\\d{9}$/)" name:"手机号正则验证"`
	SmsCode     string `json:"smsCode" valid:"required,len=6,numeric" name:"验证码"`
	IdCard      string `json:"idCard" valid:"required,idCard" name:"身份证"`
	Birthday    string `json:"birthday" valid:"required,date=2006-01-02" name:"生日"`
	Email       string `json:"email" valid:"required,email" name:"邮箱"`
	LoginIp     string `json:"loginIp" valid:"required,ip" name:"用户IP"`
}

func (u *User) Valid(v *gvalid.Validation) {
	passFunc := gvalid.ValidUsername()
	if u.Username != "" && !passFunc.Func(u.Username) {
		v.SetError("Username", "用户名", passFunc.Msg)
	}

	passFunc = gvalid.ValidPassword()
	if u.Password != "" && !passFunc.Func(u.Password) {
		v.SetError("Password", "密码", passFunc.Msg)
	}

	if u.Password != u.RePassword {
		v.SetError("RePassword", "确认密码", "必须 和 密码一致")
	}
}

func TestUserAdd(t *testing.T) {

	data := `{"username":"adgte","password":"1234567a","rePassword":"1234567a","mobile":"13501691436","mobileRegex":"13501691436","smsCode":"095081","idCard":"310104200312166537","birthday":"2001-03-03","email":"booldesign@163.com","loginIp":"127.0.0.1"}`

	input := &User{}
	err := json.Unmarshal([]byte(data), input)
	if err != nil {
		t.Fatal(err)
	}

	v := &gvalid.Validation{}
	b, err := v.Valid(input)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

type Gift struct {
	Base *GiftBase     `json:"base" valid:"required,dive" name:"基本信息"`
	List []*GiftExtend `json:"list" valid:"required,dive" name:"权益包"`
}

type GiftBase struct {
	Name         string  `json:"name" valid:"required,gte=1,lte=10" name:"礼包名称"`
	SellingPrice float64 `json:"sellingPrice" valid:"gte=0" name:"销售价"`
	MarkingPrice float64 `json:"markingPrice" valid:"required,gte=0" name:"划线价"`
	IsTwitter    *int    `json:"isTwitter" valid:"in=0 1" name:"身份权益"`
	CardBg       Img     `json:"cardBg" valid:"-" name:"卡面背景图"`
}

type GiftExtend struct {
	Title  string `json:"title" valid:"required,gte=1,lte=10" name:"权益包标题"`
	Sort   int    `json:"sort" valid:"default=99" name:"排序"`
	DefStr string `json:"defStr" valid:"default=我是默认值" name:"默认"`
}

func TestGift(t *testing.T) {

	data := `{"base":{"name":"test sort","sellingPrice":129.00,"markingPrice":1.00,"isTwitter":0,"cardBg":{"imgUrl":"https://www.baidu.com/"}},"list":[{"title":"111","sort":0}]}`

	input := &Gift{}
	err := json.Unmarshal([]byte(data), input)
	if err != nil {
		t.Fatal(err)
	}

	v := &gvalid.Validation{}
	b, err := v.Valid(input)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

type Pager struct {
	Page     int `json:"pageNum" valid:"default=1" name:"页码"`
	PageSize int `json:"pageSize" valid:"default=10" name:"页条数"`
}

// PageListDto
type PageListDto struct {
	Pager
	Keywords string `json:"keywords" valid:"-"`
	Status   string `json:"status" valid:"default=ENABLED" name:"状态"`
}

func TestPageList(t *testing.T) {

	data := `{}`

	input := &PageListDto{}
	err := json.Unmarshal([]byte(data), input)
	if err != nil {
		t.Fatal(err)
	}
	v := &gvalid.Validation{}
	b, err := v.Valid(input)
	if err != nil {
		t.Fatal("result err:", err)
	}
	if !b {
		t.Fatal("result valid err:", v.ErrorsMap)
	}
}

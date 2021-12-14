// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// HuaweiPeoplePrice is the golang structure for table huawei_people_price.
type HuaweiPeoplePrice struct {
	Id            uint   `orm:"id,primary"     json:"id"`            //
	Model         string `orm:"model"          json:"model"`         // 型号
	Category      string `orm:"category"       json:"category"`      // 分类  大类/小类
	ServiceMethod string `orm:"service_method" json:"serviceMethod"` // 服务方式
	ServiceCate   string `orm:"service_cate"   json:"serviceCate"`   // 服务分类
	Price         int    `orm:"price"          json:"price"`         // 人工费
}

// Item is the golang structure for table item.
type Item struct {
	Id    uint `orm:"id,primary" json:"id"`    //
	Power int  `orm:"power"      json:"power"` //
	Price int  `orm:"price"      json:"price"` //
}

// User is the golang structure for table user.
type User struct {
	Id       uint        `orm:"id,primary" json:"id"`       // 用户ID
	Passport string      `orm:"passport"   json:"passport"` // 用户账号
	Password string      `orm:"password"   json:"password"` // 用户密码
	Nickname string      `orm:"nickname"   json:"nickname"` // 用户昵称
	CreateAt *gtime.Time `orm:"create_at"  json:"createAt"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"updateAt"` // 更新时间
}
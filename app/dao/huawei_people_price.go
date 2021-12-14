// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-demo/app/dao/internal"
)

// huaweiPeoplePriceDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type huaweiPeoplePriceDao struct {
	*internal.HuaweiPeoplePriceDao
}

var (
	// HuaweiPeoplePrice is globally public accessible object for table huawei_people_price operations.
	HuaweiPeoplePrice = huaweiPeoplePriceDao{
		internal.NewHuaweiPeoplePriceDao(),
	}
)

// Fill with you ideas below.

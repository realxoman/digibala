package models

import "time"

// Voucher model comes here

/*

Voucher Types:
	- Percentage
	- Fixed Price

Voucher Conditions:
	- For All
	- Per Product
	- Per User
	- Specific Time Limitation (E.g. Time)

*/

type ProductVoucher struct {
	IsActive   bool  `json:"is_active"`
	ProductsId []int `json:"products_id"`
}

type UserVoucher struct {
	IsActive bool  `json:is_active`
	UsersId  []int `json:user_id`
}

type Voucher struct {
	IsActive       bool           `json:is_active`
	VoucherType    string         `json:voucher_type`
	VoucherNum     int            `json:voucher_num`
	ProductVoucher ProductVoucher `json:products`
	UserVoucher    UserVoucher    `json:users`
	ExpiredTime    time.Time      `json:exp_time`
}

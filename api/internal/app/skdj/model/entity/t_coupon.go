// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TCoupon is the golang structure for table t_coupon.
type TCoupon struct {
	Id            int    `json:"id"            description:"唯一自增ID"`
	CouponId      string `json:"couponId"      description:"优惠券id"`
	FullMoney     int    `json:"fullMoney"     description:"满多少可以使用"`
	SubtractMoney int    `json:"subtractMoney" description:"减多少钱"`
	Scene         string `json:"scene"         description:"应用场景，注册时register,rg_invite邀请注册时,first_xd首次下单"`
	Enabled       int    `json:"enabled"       description:"是否有效 0无效"`
}

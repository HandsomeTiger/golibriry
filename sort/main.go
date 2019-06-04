package sort

import "git.funxdata.com/funxcloud/funxcloud/types"

type feeList []*types.Order

func (f feeList) Len() int      { return len(f) }
func (f feeList) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f feeList) Less(i, j int) bool {
	if f[i].Type == types.OrderTypeRoom {
		if f[j].Type == types.OrderTypeRoom {
			return f[i].Money > f[j].Money
		}
		return true
	}
	return f[i].Money > f[j].Money
}

var (
	// 优惠券类型的优先级
	couponsModePriority = map[types.CouponMode]int{
		types.CouponModeRemit:    1,
		types.CouponModeDiscount: 2,
		types.CouponModeCash:     3,
	}

	// 优惠券作用域的优先级
	couponsLimitPriority = map[types.CouponLimit]int{
		types.CouponLimitRoom:        1,
		types.CouponLimitManagement:  2,
		types.CouponLimitWater:       3,
		types.CouponLimitGas:         4,
		types.CouponLimitElectricity: 5,
		types.CouponLimitNone:        10,
	}
)

// TODO 完善这两个匹配的优先级
type couponsList []*types.Coupon

func (f couponsList) Len() int      { return len(f) }
func (f couponsList) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f couponsList) Less(i, j int) bool {
	a, b := f[i].CouponType, f[j].CouponType
	if a.Mode != b.Mode {
		return couponsModePriority[a.Mode] < couponsModePriority[b.Mode]
	}

	if a.Limit != b.Limit {
		return couponsLimitPriority[a.Limit] < couponsLimitPriority[b.Limit]
	}

	if a.Mode == types.CouponModeCash {
		return a.Discount > b.Discount
	} else if a.Mode == types.CouponModeRemit {
		return false
	}
	return a.Discount < b.Discount
}

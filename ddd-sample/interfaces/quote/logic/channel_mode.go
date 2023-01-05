/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_mode
	@Date    2022/5/30 11:39
	@Desc
*/

package logic

// UnitPrice 单价模式 eg: 0-10kg，5元/kg,不足1kg按1kg计算.
func UnitPrice(unitWeight, chargeWeight int, unitWeightFee float64) (shippingCost float64) {
	if unitWeight > 0 {
		shippingCost = float64(chargeWeight) / float64(unitWeight) * unitWeightFee
	}

	return shippingCost
}

// ContinuedUnitPrice 续单价模式 eg: 0-10kg，首重2kg 15元, 续重3元/0.5kg.
func ContinuedUnitPrice(unitWeight, chargeWeight, firstWeight int,
	unitWeightFee, firstWeightFee float64,
) (shippingCost float64) {
	if firstWeight > 0 {
		shippingCost = firstWeightFee
	}

	if firstWeight > 0 && chargeWeight > firstWeight && unitWeight > 0 {
		shippingCost += float64(chargeWeight-firstWeight) / float64(unitWeight) * unitWeightFee
	}

	return shippingCost
}

// TotalOrUnitPrice 总价或单价模式 eg <=2kg 按10元计算， >2kg 按8元/KG  2.5=3*8=24元.
func TotalOrUnitPrice(unitWeight, chargeWeight, firstWeight int,
	unitWeightFee, firstWeightFee float64,
) (shippingCost float64) {
	if firstWeight > 0 {
		shippingCost = firstWeightFee
	}

	if firstWeight > 0 && chargeWeight > firstWeight && unitWeight > 0 {
		shippingCost = float64(chargeWeight/unitWeight) * unitWeightFee
	}

	return shippingCost
}

// UnitPriceNoCeil 单价模式 eg: 0-10kg，5元/kg.
func UnitPriceNoCeil(unitWeight, chargeWeight int,
	unitWeightFee float64,
) (shippingCost float64) {
	if unitWeight > 0 {
		shippingCost = float64(chargeWeight/unitWeight) * unitWeightFee
	}

	return shippingCost
}

// ContinuedUnitPriceNoCeil  续单价模式.
func ContinuedUnitPriceNoCeil(unitWeight, chargeWeight, firstWeight int,
	unitWeightFee, firstWeightFee float64,
) (shippingCost float64) {
	if firstWeight > 0 {
		shippingCost = firstWeightFee
	}

	if firstWeight > 0 && chargeWeight > firstWeight && unitWeight > 0 {
		shippingCost += float64((chargeWeight-firstWeight)/unitWeight) * unitWeightFee
	}

	return shippingCost
}

// TotalOrUnitPriceNoCeil 总价或单价模式.
func TotalOrUnitPriceNoCeil(unitWeight, chargeWeight, firstWeight int,
	unitWeightFee, firstWeightFee float64,
) (shippingCost float64) {
	if firstWeight > 0 {
		shippingCost = firstWeightFee
	}

	if firstWeight > 0 && chargeWeight > firstWeight && unitWeight > 0 {
		shippingCost = float64((chargeWeight-firstWeight)/unitWeight) * unitWeightFee
	}

	return shippingCost
}

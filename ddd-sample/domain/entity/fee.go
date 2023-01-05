/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    fee
	@Date    2022/5/13 15:22
	@Desc
*/

package entity

type ShippingFee struct {
	totalFee        float64
	fuelFee         float64
	processingFee   float64
	registrationFee float64
	miscFee         float64
	shippingFee     float64
}

func (shippingFee ShippingFee) TotalFee() float64 {
	return shippingFee.totalFee
}

func (shippingFee ShippingFee) FuelFee() float64 {
	return shippingFee.fuelFee
}

func (shippingFee ShippingFee) ProcessingFee() float64 {
	return shippingFee.processingFee
}

func (shippingFee ShippingFee) RegistrationFee() float64 {
	return shippingFee.registrationFee
}

func (shippingFee ShippingFee) MiscFee() float64 {
	return shippingFee.miscFee
}

func (shippingFee ShippingFee) ShippingFee() float64 {
	return shippingFee.shippingFee
}

func NewShippingFee(
	totalFee float64,
	fuelFee float64,
	processingFee float64,
	registrationFee float64,
	miscFee float64,
	shippingFee float64,
) *ShippingFee {
	return &ShippingFee{
		totalFee:        totalFee,
		fuelFee:         fuelFee,
		processingFee:   processingFee,
		registrationFee: registrationFee,
		miscFee:         miscFee,
		shippingFee:     shippingFee,
	}
}

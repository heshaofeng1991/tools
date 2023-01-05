/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    shipping_option
	@Date    2022/5/12 11:53
	@Desc
*/

package entity

type ShippingFeeForFilterCost struct {
	fee           ShippingFee
	mode          int8
	channelCostID int64
	channelID     int64
	channelName   string
	currency      string
	chargeWeight  int64
	actualWeight  int64
	volumeWeight  int64
	logisticsType int8
	maxNormalDays int
	minNormalDays int
	isRecommended bool
	channelType   int8
	deliverDuty   string
	description   string
	tag           string
	averageDays   int
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) Fee() ShippingFee {
	return shippingFeeForFilterCost.fee
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) Mode() int8 {
	return shippingFeeForFilterCost.mode
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) ChannelCostID() int64 {
	return shippingFeeForFilterCost.channelCostID
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) ChannelID() int64 {
	return shippingFeeForFilterCost.channelID
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) ChannelName() string {
	return shippingFeeForFilterCost.channelName
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) ChannelType() int8 {
	return shippingFeeForFilterCost.channelType
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) Currency() string {
	return shippingFeeForFilterCost.currency
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) ChargeWeight() int64 {
	return shippingFeeForFilterCost.chargeWeight
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) ActualWeight() int64 {
	return shippingFeeForFilterCost.actualWeight
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) VolumeWeight() int64 {
	return shippingFeeForFilterCost.volumeWeight
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) MaxNormalDays() int {
	return shippingFeeForFilterCost.maxNormalDays
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) MinNormalDays() int {
	return shippingFeeForFilterCost.minNormalDays
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) IsRecommended() bool {
	return shippingFeeForFilterCost.isRecommended
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) DeliverDuty() string {
	return shippingFeeForFilterCost.deliverDuty
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) LogisticsType() int8 {
	return shippingFeeForFilterCost.logisticsType
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) Description() *string {
	return &shippingFeeForFilterCost.description
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) Tag() *string {
	return &shippingFeeForFilterCost.tag
}

func (shippingFeeForFilterCost ShippingFeeForFilterCost) AverageDays() int {
	return shippingFeeForFilterCost.averageDays
}

type Option func(u *ShippingFeeForFilterCost)

func WithDescription(description *string) Option {
	return func(fee *ShippingFeeForFilterCost) {
		if description != nil {
			fee.description = *description
		}
	}
}

func WithTag(tag *string) Option {
	return func(fee *ShippingFeeForFilterCost) {
		if tag != nil {
			fee.description = *tag
		}
	}
}

func NewShippingFeeForFilterCost(
	fee ShippingFee,
	mode int8,
	channelCostID int64,
	channelID int64,
	channelName string,
	currency string,
	chargeWeight int64,
	actualWeight int64,
	volumeWeight int64,
	logisticsType int8,
	maxNormalDays int,
	minNormalDays int,
	isRecommended bool,
	channelType int8,
	deliverDuty string,
	description string,
	tag string,
	averageDays int,
) *ShippingFeeForFilterCost {
	return &ShippingFeeForFilterCost{
		fee:           fee,
		mode:          mode,
		channelCostID: channelCostID,
		channelID:     channelID,
		channelName:   channelName,
		currency:      currency,
		chargeWeight:  chargeWeight,
		actualWeight:  actualWeight,
		volumeWeight:  volumeWeight,
		logisticsType: logisticsType,
		maxNormalDays: maxNormalDays,
		minNormalDays: minNormalDays,
		isRecommended: isRecommended,
		channelType:   channelType,
		deliverDuty:   deliverDuty,
		description:   description,
		tag:           tag,
		averageDays:   averageDays,
	}
}

/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel
	@Date    2022/5/12 14:32
	@Desc
*/

package entity

import (
	"time"

	ent "github.com/heshaofeng1991/entgo/ent/gen"
)

type Channel struct {
	id                  int64
	warehouseID         int64
	courierPlatform     string
	name                string
	code                string
	channelLogisticType int8
	quotationCurrency   string
	volumeFactor        int32
	enName              string
	displayName         string
	hasTrackingNumber   int8
	maxWeight           int
	maxLength           int
	minLength           int
	maxThreeSideSum     int
	description         string
	sortingPort         int32
	prepayTariff        bool
	status              int8
	test                bool
	excludeAttributes   string
	options             string
	createdAt           time.Time
	updatedAt           time.Time
	channelType         int8
	virtual             int8
	deliverDuty         string
	special             int8
}

func (channel *Channel) ID() int64 {
	return channel.id
}

func (channel *Channel) WarehouseID() int64 {
	return channel.warehouseID
}

func (channel *Channel) CourierPlatform() string {
	return channel.courierPlatform
}

func (channel *Channel) Name() string {
	return channel.name
}

func (channel *Channel) Code() string {
	return channel.code
}

func (channel *Channel) ChannelLogisticType() int8 {
	return channel.channelLogisticType
}

func (channel *Channel) QuotationCurrency() string {
	return channel.quotationCurrency
}

func (channel *Channel) VolumeFactor() int32 {
	return channel.volumeFactor
}

func (channel *Channel) EnName() string {
	return channel.enName
}

func (channel *Channel) DisplayName() string {
	return channel.displayName
}

func (channel *Channel) HasTrackingNumber() int8 {
	return channel.hasTrackingNumber
}

func (channel *Channel) MaxWeight() int {
	return channel.maxWeight
}

func (channel *Channel) MaxLength() int {
	return channel.maxLength
}

func (channel *Channel) MinLength() int {
	return channel.minLength
}

func (channel *Channel) MaxThreeSideSum() int {
	return channel.maxThreeSideSum
}

func (channel *Channel) Description() string {
	return channel.description
}

func (channel *Channel) SortingPort() int32 {
	return channel.sortingPort
}

func (channel *Channel) PrepayTariff() bool {
	return channel.prepayTariff
}

func (channel *Channel) Status() int8 {
	return channel.status
}

func (channel *Channel) Test() bool {
	return channel.test
}

func (channel *Channel) ExcludeAttributes() string {
	return channel.excludeAttributes
}

func (channel *Channel) Options() string {
	return channel.options
}

func (channel *Channel) CreatedAt() time.Time {
	return channel.createdAt
}

func (channel *Channel) UpdatedAt() time.Time {
	return channel.updatedAt
}

func (channel *Channel) DeliverDuty() string {
	return channel.deliverDuty
}

func (channel *Channel) Virtual() int8 {
	return channel.virtual
}

func (channel *Channel) ChannelType() int8 {
	return channel.channelType
}

func (channel *Channel) Special() int8 {
	return channel.special
}

func NewChannel(
	id int64,
	warehouseID int64,
	courierPlatform string,
	name string,
	code string,
	channelLogisticType int8,
	quotationCurrency string,
	volumeFactor int32,
	enName string,
	displayName string,
	hasTrackingNumber int8,
	maxWeight int,
	maxLength int,
	minLength int,
	maxThreeSideSum int,
	description string,
	sortingPort int32,
	prepayTariff bool,
	status int8,
	test bool,
	excludeAttributes string,
	options string,
	createdAt time.Time,
	updatedAt time.Time,
	channelType int8,
	virtual int8,
	deliverDuty string,
	special int8,
) *Channel {
	return &Channel{
		id:                  id,
		warehouseID:         warehouseID,
		courierPlatform:     courierPlatform,
		name:                name,
		code:                code,
		channelType:         channelType,
		quotationCurrency:   quotationCurrency,
		volumeFactor:        volumeFactor,
		enName:              enName,
		displayName:         displayName,
		hasTrackingNumber:   hasTrackingNumber,
		maxWeight:           maxWeight,
		maxLength:           maxLength,
		minLength:           minLength,
		maxThreeSideSum:     maxThreeSideSum,
		description:         description,
		sortingPort:         sortingPort,
		prepayTariff:        prepayTariff,
		status:              status,
		test:                test,
		excludeAttributes:   excludeAttributes,
		options:             options,
		createdAt:           createdAt,
		updatedAt:           updatedAt,
		channelLogisticType: channelLogisticType,
		virtual:             virtual,
		deliverDuty:         deliverDuty,
		special:             special,
	}
}

func UnmarshalChannelFromDB(
	id int64,
	warehouseID int64,
	courierPlatform string,
	name string,
	code string,
	channelLogisticType int8,
	quotationCurrency string,
	volumeFactor int32,
	enName string,
	displayName string,
	hasTrackingNumber int8,
	maxWeight int,
	maxLength int,
	minLength int,
	maxThreeSideSum int,
	description string,
	sortingPort int32,
	prepayTariff bool,
	status int8,
	test bool,
	excludeAttributes string,
	options string,
	createdAt time.Time,
	updatedAt time.Time,
	channelType int8,
	virtual int8,
	deliverDuty string,
	special int8,
) *Channel {
	return NewChannel(
		id,
		warehouseID,
		courierPlatform,
		name,
		code,
		channelLogisticType,
		quotationCurrency,
		volumeFactor,
		enName,
		displayName,
		hasTrackingNumber,
		maxWeight,
		maxLength,
		minLength,
		maxThreeSideSum,
		description,
		sortingPort,
		prepayTariff,
		status,
		test,
		excludeAttributes,
		options,
		createdAt,
		updatedAt,
		channelType,
		virtual,
		deliverDuty,
		special,
	)
}

func CovertDBToChannel(channels []*ent.Channel) []*Channel {
	results := make([]*Channel, 0)

	for _, channel := range channels {
		chl := NewChannel(
			channel.ID,
			channel.WarehouseID,
			channel.CourierPlatform,
			channel.Name,
			channel.Code,
			channel.Type,
			channel.QuotationCurrency,
			channel.VolumeFactor,
			channel.EnName,
			channel.DisplayName,
			channel.HasTrackingNumber,
			channel.MaxWeight,
			channel.MaxLength,
			channel.MinLength,
			channel.MaxThreeSideSum,
			channel.Description,
			channel.SortingPort,
			channel.PrepayTariff,
			channel.Status,
			channel.Test,
			channel.ExcludeAttributes,
			channel.Options,
			channel.CreatedAt,
			channel.UpdatedAt,
			channel.ChannelType,
			channel.Virtual,
			channel.DeliverDuty,
			channel.Special,
		)
		results = append(results, chl)
	}

	return results
}

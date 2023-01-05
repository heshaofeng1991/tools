/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost
	@Date    2022/5/12 21:12
	@Desc
*/

package entity

import (
	"time"

	ent "github.com/heshaofeng1991/entgo/ent/gen"
)

type ChannelCost struct {
	id                 int64
	channelCostBatchID int64
	channelID          int64
	mode               int8
	countryCode        string
	zone               string
	startWeight        int
	endWeight          int
	firstWeight        int
	firstWeightFee     float64
	unitWeight         int
	unitWeightFee      float64
	fuelFee            float64
	processingFee      float64
	registrationFee    float64
	miscFee            float64
	createdAt          time.Time
	updatedAt          time.Time
	channel            *Channel
	minNormalDays      int
	maxNormalDays      int
	status             int8
	averageDays        int
}

func (channelCost *ChannelCost) ID() int64 {
	return channelCost.id
}

func (channelCost *ChannelCost) ChannelCostBatchID() int64 {
	return channelCost.channelCostBatchID
}

func (channelCost *ChannelCost) ChannelID() int64 {
	return channelCost.channelID
}

func (channelCost *ChannelCost) Mode() int8 {
	return channelCost.mode
}

func (channelCost *ChannelCost) CountryCode() string {
	return channelCost.countryCode
}

func (channelCost *ChannelCost) Zone() string {
	return channelCost.zone
}

func (channelCost *ChannelCost) StartWeight() int {
	return channelCost.startWeight
}

func (channelCost *ChannelCost) EndWeight() int {
	return channelCost.endWeight
}

func (channelCost *ChannelCost) FirstWeight() int {
	return channelCost.firstWeight
}

func (channelCost *ChannelCost) FirstWeightFee() float64 {
	return channelCost.firstWeightFee
}

func (channelCost *ChannelCost) UnitWeightFee() float64 {
	return channelCost.unitWeightFee
}

func (channelCost *ChannelCost) UnitWeight() int {
	return channelCost.unitWeight
}

func (channelCost *ChannelCost) FuelFee() float64 {
	return channelCost.fuelFee
}

func (channelCost *ChannelCost) ProcessingFee() float64 {
	return channelCost.processingFee
}

func (channelCost *ChannelCost) RegistrationFee() float64 {
	return channelCost.registrationFee
}

func (channelCost *ChannelCost) MiscFee() float64 {
	return channelCost.miscFee
}

func (channelCost *ChannelCost) CreatedAt() time.Time {
	return channelCost.createdAt
}

func (channelCost *ChannelCost) UpdatedAt() time.Time {
	return channelCost.updatedAt
}

func (channelCost *ChannelCost) Channel() *Channel {
	return channelCost.channel
}

func (channelCost *ChannelCost) MaxNormalDays() int {
	return channelCost.maxNormalDays
}

func (channelCost *ChannelCost) MinNormalDays() int {
	return channelCost.minNormalDays
}

func (channelCost *ChannelCost) Status() int8 {
	return channelCost.status
}

func (channelCost *ChannelCost) AverageDays() int {
	return channelCost.averageDays
}

func NewChannelCost(
	id int64,
	channelCostBatchID int64,
	channelID int64,
	mode int8,
	countryCode string,
	zone string,
	startWeight int,
	endWeight int,
	firstWeight int,
	firstWeightFee float64,
	unitWeight int,
	unitWeightFee float64,
	fuelFee float64,
	processingFee float64,
	registrationFee float64,
	miscFee float64,
	createdAt time.Time,
	updatedAt time.Time,
	channel *Channel,
	minNormalDays int,
	maxNormalDays int,
	status int8,
	averageDays int,
) *ChannelCost {
	return &ChannelCost{
		id:                 id,
		channelCostBatchID: channelCostBatchID,
		channelID:          channelID,
		mode:               mode,
		countryCode:        countryCode,
		zone:               zone,
		startWeight:        startWeight,
		endWeight:          endWeight,
		firstWeight:        firstWeight,
		firstWeightFee:     firstWeightFee,
		unitWeight:         unitWeight,
		unitWeightFee:      unitWeightFee,
		fuelFee:            fuelFee,
		processingFee:      processingFee,
		registrationFee:    registrationFee,
		miscFee:            miscFee,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
		channel:            channel,
		minNormalDays:      minNormalDays,
		maxNormalDays:      maxNormalDays,
		status:             status,
		averageDays:        averageDays,
	}
}

func UnmarshalChannelCostFromDB(
	id int64,
	channelCostBatchID int64,
	channelID int64,
	mode int8,
	countryCode string,
	zone string,
	startWeight int,
	endWeight int,
	firstWeight int,
	firstWeightFee float64,
	unitWeight int,
	unitWeightFee float64,
	fuelFee float64,
	processingFee float64,
	registrationFee float64,
	miscFee float64,
	createdAt time.Time,
	updatedAt time.Time,
	channel *Channel,
	minNormalDays int,
	maxNormalDays int,
	status int8,
	averageDays int,
) *ChannelCost {
	return NewChannelCost(
		id,
		channelCostBatchID,
		channelID,
		mode,
		countryCode,
		zone,
		startWeight,
		endWeight,
		firstWeight,
		firstWeightFee,
		unitWeight,
		unitWeightFee,
		fuelFee,
		processingFee,
		registrationFee,
		miscFee,
		createdAt,
		updatedAt,
		channel,
		minNormalDays,
		maxNormalDays,
		status,
		averageDays,
	)
}

func CovertChannel(channel *ent.Channel) *Channel {
	if channel == nil {
		return nil
	}

	return NewChannel(
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
}

func CovertChannelCost(channelCosts []*ent.ChannelCost) []*ChannelCost {
	results := make([]*ChannelCost, 0)

	for _, channelCost := range channelCosts {
		cost := NewChannelCost(
			channelCost.ID,
			channelCost.ChannelCostBatchID,
			channelCost.ChannelID,
			channelCost.Mode,
			channelCost.CountryCode,
			channelCost.Zone,
			channelCost.StartWeight,
			channelCost.EndWeight,
			channelCost.FirstWeight,
			channelCost.FirstWeightFee,
			channelCost.UnitWeight,
			channelCost.UnitWeightFee,
			channelCost.FuelFee,
			channelCost.ProcessingFee,
			channelCost.RegistrationFee,
			channelCost.MiscFee,
			channelCost.CreatedAt,
			channelCost.UpdatedAt,
			CovertChannel(channelCost.Edges.Channels),
			channelCost.MinNormalDays,
			channelCost.MaxNormalDays,
			channelCost.Status,
			channelCost.AverageDays,
		)

		results = append(results, cost)
	}

	return results
}

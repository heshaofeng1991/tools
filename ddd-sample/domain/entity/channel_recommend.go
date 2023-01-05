/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_recommend
	@Date    2022/5/25 14:29
	@Desc
*/

package entity

import (
	"time"

	ent "github.com/heshaofeng1991/entgo/ent/gen"
)

type ChannelRecommend struct {
	id            int64
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     time.Time
	countryCode   string
	channelID     int64
	isRecommended int8
	status        int8
	value         string
}

func (recommend *ChannelRecommend) ID() int64 {
	return recommend.id
}

func (recommend *ChannelRecommend) CreatedAt() time.Time {
	return recommend.createdAt
}

func (recommend *ChannelRecommend) UpdatedAt() time.Time {
	return recommend.updatedAt
}

func (recommend *ChannelRecommend) DeleteAt() time.Time {
	return recommend.deletedAt
}

func (recommend *ChannelRecommend) ChannelID() int64 {
	return recommend.channelID
}

func (recommend *ChannelRecommend) CountryCode() string {
	return recommend.countryCode
}

func (recommend *ChannelRecommend) Status() int8 {
	return recommend.status
}

func (recommend *ChannelRecommend) IsRecommended() int8 {
	return recommend.isRecommended
}

func (recommend *ChannelRecommend) Value() string {
	return recommend.value
}

func NewChannelRecommend(
	id int64,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
	countryCode string,
	channelID int64,
	isRecommended int8,
	status int8,
	value string,
) *ChannelRecommend {
	return &ChannelRecommend{
		id:            id,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		deletedAt:     deletedAt,
		countryCode:   countryCode,
		channelID:     channelID,
		isRecommended: isRecommended,
		status:        status,
		value:         value,
	}
}

func CovertDBToChannelRecommend(recommends []*ent.ChannelRecommend) []*ChannelRecommend {
	result := make([]*ChannelRecommend, 0)

	for _, val := range recommends {
		result = append(result,
			NewChannelRecommend(
				val.ID,
				val.CreatedAt,
				val.UpdatedAt,
				val.DeletedAt,
				val.CountryCode,
				val.ChannelID,
				val.IsRecommended,
				val.Status,
				val.Value,
			))
	}

	return result
}

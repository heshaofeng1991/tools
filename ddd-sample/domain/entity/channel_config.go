/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_config
	@Date    2022/5/25 14:28
	@Desc
*/

package entity

import (
	"time"

	ent "github.com/heshaofeng1991/entgo/ent/gen"
)

type ChannelConfig struct {
	id                 int64
	createdAt          time.Time
	updatedAt          time.Time
	deletedAt          time.Time
	channelID          int64
	excludeCountryCode string
	status             int8
}

func (cfg *ChannelConfig) ID() int64 {
	return cfg.id
}

func (cfg *ChannelConfig) CreatedAt() time.Time {
	return cfg.createdAt
}

func (cfg *ChannelConfig) UpdatedAt() time.Time {
	return cfg.updatedAt
}

func (cfg *ChannelConfig) DeleteAt() time.Time {
	return cfg.deletedAt
}

func (cfg *ChannelConfig) ChannelID() int64 {
	return cfg.channelID
}

func (cfg *ChannelConfig) ExcludeCountryCode() string {
	return cfg.excludeCountryCode
}

func (cfg *ChannelConfig) Status() int8 {
	return cfg.status
}

func NewChannelConfig(
	id int64,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
	channelID int64,
	excludeCountryCode string,
	status int8,
) *ChannelConfig {
	return &ChannelConfig{
		id:                 id,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
		deletedAt:          deletedAt,
		channelID:          channelID,
		excludeCountryCode: excludeCountryCode,
		status:             status,
	}
}

func CovertDBToChannelConfig(configs []*ent.CustomerConfig) []*ChannelConfig {
	result := make([]*ChannelConfig, 0)

	for _, val := range configs {
		result = append(result,
			NewChannelConfig(
				val.ID,
				val.CreatedAt,
				val.UpdatedAt,
				val.DeletedAt,
				val.ChannelID,
				val.ExcludeCountryCode,
				val.Status,
			))
	}

	return result
}

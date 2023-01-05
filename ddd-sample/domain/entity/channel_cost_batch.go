/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost_batch
	@Date    2022/5/12 18:01
	@Desc
*/

package entity

import (
	"time"

	ent "github.com/heshaofeng1991/entgo/ent/gen"
)

type ChannelCostBatch struct {
	id            int64
	channelID     int64
	effectiveDate time.Time
	expiryDate    time.Time
	status        bool
	createdAt     time.Time
	updatedAt     time.Time
}

func (channelCostBatch *ChannelCostBatch) ID() int64 {
	return channelCostBatch.id
}

func (channelCostBatch *ChannelCostBatch) ChannelID() int64 {
	return channelCostBatch.channelID
}

func (channelCostBatch *ChannelCostBatch) EffectiveDate() time.Time {
	return channelCostBatch.effectiveDate
}

func (channelCostBatch *ChannelCostBatch) ExpiryDate() time.Time {
	return channelCostBatch.expiryDate
}

func (channelCostBatch *ChannelCostBatch) Status() bool {
	return channelCostBatch.status
}

func (channelCostBatch *ChannelCostBatch) CreatedAt() time.Time {
	return channelCostBatch.createdAt
}

func (channelCostBatch *ChannelCostBatch) UpdatedAt() time.Time {
	return channelCostBatch.updatedAt
}

func NewChannelCostBatch(
	id int64,
	channelID int64,
	effectiveDate time.Time,
	expiryDate time.Time,
	status bool,
	createdAt time.Time,
	updatedAt time.Time,
) *ChannelCostBatch {
	return &ChannelCostBatch{
		id:            id,
		channelID:     channelID,
		effectiveDate: effectiveDate,
		expiryDate:    expiryDate,
		status:        status,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
}

func UnmarshalChannelCostBatchFromDB(
	id int64,
	channelID int64,
	effectiveDate time.Time,
	expiryDate time.Time,
	status bool,
	createdAt time.Time,
	updatedAt time.Time,
) *ChannelCostBatch {
	return NewChannelCostBatch(
		id,
		channelID,
		effectiveDate,
		expiryDate,
		status,
		createdAt,
		updatedAt,
	)
}

func CovertDBToChannelCostBatch(batches []*ent.ChannelCostBatche) []*ChannelCostBatch {
	results := make([]*ChannelCostBatch, 0)

	for _, batch := range batches {
		btc := NewChannelCostBatch(
			batch.ID,
			batch.ChannelID,
			batch.EffectiveDate,
			batch.ExpiryDate,
			batch.Status,
			batch.CreatedAt,
			batch.UpdatedAt,
		)
		results = append(results, btc)
	}

	return results
}

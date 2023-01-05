/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    common
	@Date    2022/5/13 15:21
	@Desc
*/

package entity

type ChannelBatch struct {
	channelData      *Channel
	channelCostBatch *ChannelCostBatch
	chargeWeight     int64
	actualWeight     int64
	volumeWeight     int64
}

func (channelBatch *ChannelBatch) ChannelData() *Channel {
	return channelBatch.channelData
}

func (channelBatch *ChannelBatch) ChannelCostBatch() *ChannelCostBatch {
	return channelBatch.channelCostBatch
}

func (channelBatch *ChannelBatch) ChargeWeight() int64 {
	return channelBatch.chargeWeight
}

func (channelBatch *ChannelBatch) ActualWeight() int64 {
	return channelBatch.actualWeight
}

func (channelBatch *ChannelBatch) VolumeWeight() int64 {
	return channelBatch.volumeWeight
}

func NewChannelBatch(
	channelData *Channel,
	channelCostBatch *ChannelCostBatch,
	chargeWeight int64,
	actualWeight int64,
	volumeWeight int64,
) *ChannelBatch {
	return &ChannelBatch{
		channelData:      channelData,
		channelCostBatch: channelCostBatch,
		chargeWeight:     chargeWeight,
		actualWeight:     actualWeight,
		volumeWeight:     volumeWeight,
	}
}

type ChannelCostRsp struct {
	channelCost  *ChannelCost
	chargeWeight int64
	actualWeight int64
	volumeWeight int64
}

func (channelCostRsp *ChannelCostRsp) ChannelCost() *ChannelCost {
	return channelCostRsp.channelCost
}

func (channelCostRsp *ChannelCostRsp) ChargeWeight() int64 {
	return channelCostRsp.chargeWeight
}

func (channelCostRsp *ChannelCostRsp) ActualWeight() int64 {
	return channelCostRsp.actualWeight
}

func (channelCostRsp *ChannelCostRsp) VolumeWeight() int64 {
	return channelCostRsp.volumeWeight
}

func NewChannelCostRsp(
	channelCost *ChannelCost,
	chargeWeight int64,
	actualWeight int64,
	volumeWeight int64,
) *ChannelCostRsp {
	return &ChannelCostRsp{
		channelCost:  channelCost,
		chargeWeight: chargeWeight,
		actualWeight: actualWeight,
		volumeWeight: volumeWeight,
	}
}

type Vols struct {
	VolumeWeight int64
	BatchID      int64
}

type MapVols struct {
	VolumeWeight int64
	BatchID      []int64
}

type UserKey string

const (
	UserID UserKey = "UserID"
)

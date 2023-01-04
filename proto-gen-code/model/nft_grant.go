package model

import (
	"time"
)

type NftGrant struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	SendUserId     int64     `gorm:"column:send_user_id;type:bigint(20);comment:发起转赠用户ID;NOT NULL" json:"send_user_id"`
	ReceiverUserId int64     `gorm:"column:receiver_user_id;type:bigint(20);comment:接受转赠用户ID;NOT NULL" json:"receiver_user_id"`
	Status         int       `gorm:"column:status;type:tinyint(1);comment:赠与状态;NOT NULL" json:"status"`
	CollectionId   int64     `gorm:"column:collection_id;type:bigint(20);comment:赠送藏品ID;NOT NULL" json:"collection_id"`
	AssetId        int64     `gorm:"column:asset_id;type:bigint(20);comment:赠送资产ID;NOT NULL" json:"asset_id"`
	ReceiverTime   time.Time `gorm:"column:receiver_time;type:datetime;comment:领取时间" json:"receiver_time"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	DeleteFlag     int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftGrant) TableName() string {
	return "nft_grant"
}

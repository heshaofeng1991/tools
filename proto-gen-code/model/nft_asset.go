package model

import (
	"time"
)

type NftAsset struct {
	Id                  int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	CollectionId        int64     `gorm:"column:collection_id;type:bigint(20);comment:藏品ID;NOT NULL" json:"collection_id"`
	Status              int       `gorm:"column:status;type:tinyint(1);comment:状态;NOT NULL" json:"status"`
	ChainStatus         int       `gorm:"column:chain_status;type:tinyint(1);comment:上链状态;NOT NULL" json:"chain_status"`
	IsBlindBox          int       `gorm:"column:is_blind_box;type:tinyint(1);comment:是否属于盲盒;NOT NULL" json:"is_blind_box"`
	IsOpen              int       `gorm:"column:is_open;type:tinyint(1);comment:是否开盒;NOT NULL" json:"is_open"`
	AssetIndex          string    `gorm:"column:asset_index;type:varchar(255);comment:资产编号;NOT NULL" json:"asset_index"`
	AssetIndexNo        int       `gorm:"column:asset_index_no;type:int(11);comment:资产编号(数字);NOT NULL" json:"asset_index_no"`
	OwnerUserId         int64     `gorm:"column:owner_user_id;type:bigint(20);comment:持有人用户ID" json:"owner_user_id"`
	OwnerTime           time.Time `gorm:"column:owner_time;type:datetime;comment:持有时间" json:"owner_time"`
	FirstOwnerUserId    int64     `gorm:"column:first_owner_user_id;type:bigint(20);comment:首持用户ID" json:"first_owner_user_id"`
	FirstOwnerStartTime time.Time `gorm:"column:first_owner_start_time;type:datetime;comment:首持用户持有开始时间" json:"first_owner_start_time"`
	FirstOwnerEndTime   time.Time `gorm:"column:first_owner_end_time;type:datetime;comment:首持用户持有结束时间" json:"first_owner_end_time"`
	Source              int       `gorm:"column:source;type:tinyint(1);comment:来源;NOT NULL" json:"source"`
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	DeleteFlag          int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftAsset) TableName() string {
	return "nft_asset"
}

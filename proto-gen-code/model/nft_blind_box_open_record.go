package model

import (
	"time"
)

// NFT盲盒藏品表
type NftBlindBoxOpenRecord struct {
	Id           int       `gorm:"column:id;type:int(11);primary_key;comment:ID" json:"id"`
	UserId       int64     `gorm:"column:user_id;type:bigint(20);comment:用户ID;NOT NULL" json:"user_id"`
	BlindBoxId   int64     `gorm:"column:blind_box_id;type:bigint(20);comment:盲盒ID;NOT NULL" json:"blind_box_id"`
	CollectionId int64     `gorm:"column:collection_id;type:bigint(20);comment:开盒藏品ID;NOT NULL" json:"collection_id"`
	AssetId      int64     `gorm:"column:asset_id;type:bigint(20);comment:开盒资产ID;NOT NULL" json:"asset_id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	DeleteFlag   int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftBlindBoxOpenRecord) TableName() string {
	return "nft_blind_box_open_record"
}

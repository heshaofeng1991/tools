package model

import (
	"time"
)

type NftOrderAsset struct {
	Id           int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	OrderId      int64     `gorm:"column:order_id;type:bigint(20);comment:订单ID;NOT NULL" json:"order_id"`
	CollectionId int64     `gorm:"column:collection_id;type:bigint(20);comment:藏品ID;NOT NULL" json:"collection_id"`
	BlindBoxId   int64     `gorm:"column:blind_box_id;type:bigint(20);comment:盲盒ID" json:"blind_box_id"`
	AssetId      int64     `gorm:"column:asset_id;type:bigint(20);comment:资产ID;NOT NULL" json:"asset_id"`
	AssetIndex   string    `gorm:"column:asset_index;type:varchar(255);comment:资产编号;NOT NULL" json:"asset_index"`
	Price        int64     `gorm:"column:price;type:bigint(20);comment:藏品价格;NOT NULL" json:"price"`
	Quantity     string    `gorm:"column:quantity;type:varchar(255);comment:藏品数量" json:"quantity"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	DeleteFlag   int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftOrderAsset) TableName() string {
	return "nft_order_asset"
}

package model

import (
	"time"
)

type NftSyntheticRecord struct {
	Id                  int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	UserId              int64     `gorm:"column:user_id;type:bigint(20);comment:用户ID;NOT NULL" json:"user_id"`
	SyntheticId         int64     `gorm:"column:synthetic_id;type:bigint(20);comment:合成ID;NOT NULL" json:"synthetic_id"`
	CollectionId        int64     `gorm:"column:collection_id;type:bigint(20);comment:合成藏品ID;NOT NULL" json:"collection_id"`
	AssetId             int64     `gorm:"column:asset_id;type:bigint(20);comment:合成资产ID;NOT NULL" json:"asset_id"`
	AssetNo             int       `gorm:"column:asset_no;type:int(11);comment:合成资产编号;NOT NULL" json:"asset_no"`
	Status              int       `gorm:"column:status;type:int(11);comment:状态;NOT NULL" json:"status"`
	MaterialCount       int       `gorm:"column:material_count;type:int(11);comment:消耗材料总数;NOT NULL" json:"material_count"`
	MaterialTotalAmount int64     `gorm:"column:material_total_amount;type:bigint(20);comment:材料总价(分);NOT NULL" json:"material_total_amount"`
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	DeleteFlag          int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftSyntheticRecord) TableName() string {
	return "nft_synthetic_record"
}

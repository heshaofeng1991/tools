package model

import (
	"time"
)

type NftSyntheticRecordDetail struct {
	Id                int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	UserId            int64     `gorm:"column:user_id;type:bigint(20);comment:用户ID;NOT NULL" json:"user_id"`
	SyntheticRecordId int64     `gorm:"column:synthetic_record_id;type:bigint(20);comment:合成记录ID;NOT NULL" json:"synthetic_record_id"`
	CollectionId      int64     `gorm:"column:collection_id;type:bigint(20);comment:消耗材料藏品ID;NOT NULL" json:"collection_id"`
	AssetId           int64     `gorm:"column:asset_id;type:bigint(20);comment:消耗材料资产ID;NOT NULL" json:"asset_id"`
	Amount            int64     `gorm:"column:amount;type:bigint(20);comment:消耗材料金额;NOT NULL" json:"amount"`
	CreatedAt         time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	DeleteFlag        int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftSyntheticRecordDetail) TableName() string {
	return "nft_synthetic_record_detail"
}

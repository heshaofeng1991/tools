package model

import (
	"time"
)

type NftAirdropRecord struct {
	Id                 int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	AirdropId          int64     `gorm:"column:airdrop_id;type:bigint(20);comment:空投ID;NOT NULL" json:"airdrop_id"`
	AirdropOperationId int64     `gorm:"column:airdrop_operation_id;type:bigint(20);comment:空投操作ID;NOT NULL" json:"airdrop_operation_id"`
	AssetId            int64     `gorm:"column:asset_id;type:bigint(20);comment:资产ID;NOT NULL" json:"asset_id"`
	UserId             int64     `gorm:"column:user_id;type:bigint(20);comment:用户ID;NOT NULL" json:"user_id"`
	CreatedAt          time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy          int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt          time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy          int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag         int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftAirdropRecord) TableName() string {
	return "nft_airdrop_record"
}

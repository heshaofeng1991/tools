package model

import (
	"time"
)

type NftAirdrop struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	Type           int       `gorm:"column:type;type:tinyint(1);comment:空投类型;NOT NULL" json:"type"`
	Status         int       `gorm:"column:status;type:tinyint(1);comment:发放状态;NOT NULL" json:"status"`
	Name           string    `gorm:"column:name;type:varchar(255);comment:空投名称;NOT NULL" json:"name"`
	CoverUrl       string    `gorm:"column:cover_url;type:varchar(255);comment:空投图片;NOT NULL" json:"cover_url"`
	CollectionId   int64     `gorm:"column:collection_id;type:bigint(20);comment:藏品ID;NOT NULL" json:"collection_id"`
	CollectionName string    `gorm:"column:collection_name;type:varchar(255);comment:藏品名称;NOT NULL" json:"collection_name"`
	CollectionType int       `gorm:"column:collection_type;type:tinyint(1);comment:藏品类型;NOT NULL" json:"collection_type"`
	Stock          int       `gorm:"column:stock;type:int(11);comment:总库存;NOT NULL" json:"stock"`
	SurplusStock   int       `gorm:"column:surplus_stock;type:int(11);comment:剩余库存;NOT NULL" json:"surplus_stock"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy      int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy      int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag     int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftAirdrop) TableName() string {
	return "nft_airdrop"
}

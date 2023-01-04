package model

import (
	"time"
)

type NftSyntheticMaterial struct {
	Id           int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	SyntheticId  int64     `gorm:"column:synthetic_id;type:bigint(20);comment:合成活动ID;NOT NULL" json:"synthetic_id"`
	CollectionId int64     `gorm:"column:collection_id;type:bigint(20);comment:材料藏品ID;NOT NULL" json:"collection_id"`
	Quantity     int       `gorm:"column:quantity;type:int(11);comment:材料数量;NOT NULL" json:"quantity"`
	SortOrder    int       `gorm:"column:sort_order;type:int(11);comment:优先级;NOT NULL" json:"sort_order"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy    int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy    int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag   int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftSyntheticMaterial) TableName() string {
	return "nft_synthetic_material"
}

package model

import (
	"time"
)

type NftBindBoxItem struct {
	Id           int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	BlindBoxId   int64     `gorm:"column:blind_box_id;type:bigint(20);comment:NFT盲盒ID;NOT NULL" json:"blind_box_id"`
	CollectionId int64     `gorm:"column:collection_id;type:bigint(20);comment:藏品ID;NOT NULL" json:"collection_id"`
	AssignCount  int       `gorm:"column:assign_count;type:int(11);comment:分配数量;NOT NULL" json:"assign_count"`
	SurplusCount int       `gorm:"column:surplus_count;type:int(11);comment:剩余数量;NOT NULL" json:"surplus_count"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy    int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy    int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag   int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftBindBoxItem) TableName() string {
	return "nft_bind_box_item"
}

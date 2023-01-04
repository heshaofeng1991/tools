package model

import (
	"time"
)

type NftFirstRule struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	FirstId        int64     `gorm:"column:first_id;type:bigint(20);comment:NFT优先购ID;NOT NULL" json:"first_id"`
	CollectionId   int64     `gorm:"column:collection_id;type:bigint(20);comment:NFT藏品ID;NOT NULL" json:"collection_id"`
	CollectionName string    `gorm:"column:collection_name;type:varchar(255);comment:NFT藏品名称;NOT NULL" json:"collection_name"`
	MaxAssertNo    int       `gorm:"column:max_assert_no;type:int(11);comment:最大资产编号;NOT NULL" json:"max_assert_no"`
	AssignValue    int       `gorm:"column:assign_value;type:int(11);comment:指定的数量/指定的编号;NOT NULL" json:"assign_value"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy      int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy      int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag     int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftFirstRule) TableName() string {
	return "nft_first_rule"
}

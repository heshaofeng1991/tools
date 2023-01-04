package model

import (
	"time"
)

type NftCollectionCertification struct {
	Id              int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	CollectionId    int64     `gorm:"column:collection_id;type:bigint(20);comment:藏品ID;NOT NULL" json:"collection_id"`
	ContractAddress string    `gorm:"column:contract_address;type:varchar(255);comment:合约地址" json:"contract_address"`
	Standard        string    `gorm:"column:standard;type:varchar(255);comment:认证标准" json:"standard"`
	Network         string    `gorm:"column:network;type:varchar(255);comment:认证网络" json:"network"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy       int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	CreatedAtTime   time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy       int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag      int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftCollectionCertification) TableName() string {
	return "nft_collection_certification"
}

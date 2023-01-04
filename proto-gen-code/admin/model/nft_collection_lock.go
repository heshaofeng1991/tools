package model

import (
	"time"
)

type NftCollectionLock struct {
	Id           int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	CollectionId int64     `gorm:"column:collection_id;type:bigint(20);comment:藏品ID;NOT NULL" json:"collection_id"`
	Status       int       `gorm:"column:status;type:tinyint(1);comment:锁定状态【0->已释放；1->已锁定】;NOT NULL" json:"status"`
	Stock        int       `gorm:"column:stock;type:int(11);default:0;comment:锁定库存;NOT NULL" json:"stock"`
	SurplusStock int       `gorm:"column:surplus_stock;type:int(11);default:0;comment:剩余库存;NOT NULL" json:"surplus_stock"`
	LockSource   int       `gorm:"column:lock_source;type:tinyint(1);comment:锁定来源;NOT NULL" json:"lock_source"`
	SourceId     string    `gorm:"column:source_id;type:varchar(24);comment:锁定来源活动ID;NOT NULL" json:"source_id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy    int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy    int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag   int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftCollectionLock) TableName() string {
	return "nft_collection_lock"
}

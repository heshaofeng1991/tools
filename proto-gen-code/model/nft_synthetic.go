package model

import (
	"time"
)

type NftSynthetic struct {
	Id           uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;comment:NFT合成活动ID" json:"id"`
	CollectionId int64     `gorm:"column:collection_id;type:bigint(20);comment:NFT藏品ID;NOT NULL" json:"collection_id"`
	CoverUrl     string    `gorm:"column:cover_url;type:varchar(255);comment:图片链接;NOT NULL" json:"cover_url"`
	Stock        int       `gorm:"column:stock;type:int(11);comment:合成库存;NOT NULL" json:"stock"`
	SurplusStock int       `gorm:"column:surplus_stock;type:int(11);comment:剩余库存;NOT NULL" json:"surplus_stock"`
	ShowTime     time.Time `gorm:"column:show_time;type:datetime(6);comment:展示时间;NOT NULL" json:"show_time"`
	StartTime    time.Time `gorm:"column:start_time;type:datetime(6);comment:开始时间;NOT NULL" json:"start_time"`
	EndTime      time.Time `gorm:"column:end_time;type:datetime(6);comment:结束时间;NOT NULL" json:"end_time"`
	ComposeLimit int       `gorm:"column:compose_limit;type:int(11);comment:参与次数;NOT NULL" json:"compose_limit"`
	SortOrder    int       `gorm:"column:sort_order;type:int(11);comment:优先级" json:"sort_order"`
	Intro        string    `gorm:"column:intro;type:varchar(255);comment:合成说明" json:"intro"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy    int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy    int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag   int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftSynthetic) TableName() string {
	return "nft_synthetic"
}

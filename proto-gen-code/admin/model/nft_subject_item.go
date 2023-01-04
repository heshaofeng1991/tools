package model

import (
	"time"
)

type NftSubjectItem struct {
	Id                  int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	SubjectId           int64     `gorm:"column:subject_id;type:bigint(20);comment:NFT专题ID;NOT NULL" json:"subject_id"`
	CollectionId        int64     `gorm:"column:collection_id;type:bigint(20);comment:藏品ID;NOT NULL" json:"collection_id"`
	BlindBoxId          int64     `gorm:"column:blind_box_id;type:bigint(20);comment:盲盒ID;NOT NULL" json:"blind_box_id"`
	Stock               int       `gorm:"column:stock;type:int(11);comment:库存;NOT NULL" json:"stock"`
	ShowStock           int       `gorm:"column:show_stock;type:int(11);comment:展示库存;NOT NULL" json:"show_stock"`
	BuyLimit            int       `gorm:"column:buy_limit;type:int(11);comment:限购数;NOT NULL" json:"buy_limit"`
	WhiteList           string    `gorm:"column:white_list;type:json;comment:限购白名单" json:"white_list"`
	WhiteListLimit      int       `gorm:"column:white_list_limit;type:int(11);comment:限购白名单限购数" json:"white_list_limit"`
	FirstStock          int       `gorm:"column:first_stock;type:int(11);comment:优先购库存" json:"first_stock"`
	FirstWhiteList      string    `gorm:"column:first_white_list;type:json;comment:优先购白名单" json:"first_white_list"`
	FirstWhiteListLimit int       `gorm:"column:first_white_list_limit;type:int(11);comment:优先购白名单限购数" json:"first_white_list_limit"`
	FirstList           string    `gorm:"column:first_list;type:json;comment:管理的优先购权益ID" json:"first_list"`
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy           int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy           int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag          int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftSubjectItem) TableName() string {
	return "nft_subject_item"
}

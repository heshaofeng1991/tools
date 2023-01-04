package model

import (
	"time"
)

type NftBlindBox struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	Status     int       `gorm:"column:status;type:tinyint(1);comment:状态;NOT NULL" json:"status"`
	Name       string    `gorm:"column:name;type:varchar(255);comment:名称;NOT NULL" json:"name"`
	CoverUrl   string    `gorm:"column:cover_url;type:varchar(255);comment:图片链接;NOT NULL" json:"cover_url"`
	Price      int64     `gorm:"column:price;type:bigint(20);default:0;comment:售价(分);NOT NULL" json:"price"`
	OpenWay    int       `gorm:"column:open_way;type:tinyint(1);comment:开盒方式【1->定时开盒；2->自主开盒】;NOT NULL" json:"open_way"`
	Intro      string    `gorm:"column:intro;type:varchar(255);comment:详情;NOT NULL" json:"intro"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy  int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy  int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftBlindBox) TableName() string {
	return "nft_blind_box"
}

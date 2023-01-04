package model

import (
	"time"
)

type NftSubject struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;comment:NFT专题ID" json:"id"`
	ColumnId       int64     `gorm:"column:column_id;type:bigint(20);comment:NFT栏目ID;NOT NULL" json:"column_id"`
	Name           string    `gorm:"column:name;type:varchar(255);comment:名称;NOT NULL" json:"name"`
	CoverUrl       string    `gorm:"column:cover_url;type:varchar(255);comment:图片链接;NOT NULL" json:"cover_url"`
	ShowTime       time.Time `gorm:"column:show_time;type:datetime(6);comment:展示时间;NOT NULL" json:"show_time"`
	StartTime      time.Time `gorm:"column:start_time;type:datetime(6);comment:出售开始时间;NOT NULL" json:"start_time"`
	EndTime        time.Time `gorm:"column:end_time;type:datetime(6);comment:出售结束时间;NOT NULL" json:"end_time"`
	IsOpenFirst    int       `gorm:"column:is_open_first;type:tinyint(1);default:0;comment:是否开启优先购【0->禁用；1->启用】;NOT NULL" json:"is_open_first"`
	FirstStartTime time.Time `gorm:"column:first_start_time;type:datetime(6);comment:优先购开始时间" json:"first_start_time"`
	SortOrder      int       `gorm:"column:sort_order;type:int(11);default:0;comment:优先级;NOT NULL" json:"sort_order"`
	Intro          string    `gorm:"column:intro;type:varchar(255);comment:专题说明;NOT NULL" json:"intro"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy      int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy      int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag     int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftSubject) TableName() string {
	return "nft_subject"
}

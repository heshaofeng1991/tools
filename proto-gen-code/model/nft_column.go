package model

import (
	"time"
)

type NftColumn struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(10);comment:名称;NOT NULL" json:"name"`
	Type       int       `gorm:"column:type;type:tinyint(1);comment:类型【1->常驻；1->自定义】;NOT NULL" json:"type"`
	SortOrder  int       `gorm:"column:sort_order;type:int(11);default:0;comment:优先级;NOT NULL" json:"sort_order"`
	Remark     string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy  int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy  int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftColumn) TableName() string {
	return "nft_column"
}

package model

import (
	"time"
)

// NFT藏品表
type NftCollection struct {
	Id                  int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	Status              int       `gorm:"column:status;type:tinyint(1);comment:状态【0->已下架；1->已上架】;NOT NULL" json:"status"`
	Name                string    `gorm:"column:name;type:varchar(255);comment:名称;NOT NULL" json:"name"`
	CoverUrl            string    `gorm:"column:cover_url;type:varchar(255);comment:图片链接;NOT NULL" json:"cover_url"`
	Supplier            string    `gorm:"column:supplier;type:varchar(255);comment:供应商;NOT NULL" json:"supplier"`
	Issuer              string    `gorm:"column:issuer;type:varchar(255);comment:发行方;NOT NULL" json:"issuer"`
	Stock               int       `gorm:"column:stock;type:int(11);default:0;comment:库存;NOT NULL" json:"stock"`
	Price               int64     `gorm:"column:price;type:bigint(20);default:0;comment:售价(分);NOT NULL" json:"price"`
	IsShowCertification int       `gorm:"column:is_show_certification;type:tinyint(1);default:1;comment:是否展示授权信息;NOT NULL" json:"is_show_certification"`
	IsGrant             int       `gorm:"column:is_grant;type:tinyint(1);default:1;comment:是否允许转赠;NOT NULL" json:"is_grant"`
	IsShowThreeModel    int       `gorm:"column:is_show_three_model;type:tinyint(1);default:1;comment:是否展示3D模型;NOT NULL" json:"is_show_three_model"`
	Model3DUrl          string    `gorm:"column:model_3d_url;type:varchar(255);comment:3D模型链接;NOT NULL" json:"model_3d_url"`
	Model3DGifUr        string    `gorm:"column:model_3d_gif_ur;type:varchar(255);comment:3D模型Gif图链接;NOT NULL" json:"model_3d_gif_ur"`
	Intro               string    `gorm:"column:intro;type:varchar(255);comment:详情;NOT NULL" json:"intro"`
	PurchaseNote        string    `gorm:"column:purchase_note;type:varchar(255);comment:购买须知;NOT NULL" json:"purchase_note"`
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	CreatedBy           int64     `gorm:"column:created_by;type:bigint(20);comment:创建人;NOT NULL" json:"created_by"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	UpdatedBy           int64     `gorm:"column:updated_by;type:bigint(20);comment:更新人;NOT NULL" json:"updated_by"`
	DeleteFlag          int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftCollection) TableName() string {
	return "nft_collection"
}

package model

import (
	"time"
)

type NftOrder struct {
	Id            int64     `gorm:"column:id;type:bigint(20);primary_key;comment:ID" json:"id"`
	UserId        int64     `gorm:"column:user_id;type:bigint(20);comment:用户ID;NOT NULL" json:"user_id"`
	Status        int       `gorm:"column:status;type:tinyint(4);comment:状态;NOT NULL" json:"status"`
	TotalAmount   int64     `gorm:"column:total_amount;type:bigint(20);comment:订单总金额(分);NOT NULL" json:"total_amount"`
	TotalQuantity int64     `gorm:"column:total_quantity;type:bigint(20);comment:订单总数量;NOT NULL" json:"total_quantity"`
	PayAmount     int64     `gorm:"column:pay_amount;type:bigint(20);comment:应付金额(分);NOT NULL" json:"pay_amount"`
	SubjectId     int64     `gorm:"column:subject_id;type:bigint(20);comment:专题ID" json:"subject_id"`
	IsFirst       int       `gorm:"column:is_first;type:tinyint(1);comment:是否优先购" json:"is_first"`
	Remark        string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	OrderType     int       `gorm:"column:order_type;type:tinyint(1);comment:购买类型;NOT NULL" json:"order_type"`
	PayTime       time.Time `gorm:"column:pay_time;type:datetime;comment:支付时间" json:"pay_time"`
	PayType       int       `gorm:"column:pay_type;type:tinyint(1);comment:支付类型" json:"pay_type"`
	PayOrderId    int64     `gorm:"column:pay_order_id;type:bigint(20);comment:支付充值订单ID" json:"pay_order_id"`
	OutTradeNo    string    `gorm:"column:out_trade_no;type:varchar(255);comment:支付交易号" json:"out_trade_no"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(6);default:CURRENT_TIMESTAMP;comment:更新时间;NOT NULL" json:"updated_at"`
	DeleteFlag    int       `gorm:"column:delete_flag;type:tinyint(1);default:0;comment:逻辑删除【0->正常；1->已删除】;NOT NULL" json:"delete_flag"`
}

func (m *NftOrder) TableName() string {
	return "nft_order"
}

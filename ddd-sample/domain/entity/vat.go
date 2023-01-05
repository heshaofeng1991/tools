/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    vat
	@Date    2022/5/13 15:21
	@Desc
*/

package entity

type Vat struct {
	id          int64
	countryCode string
	// 标准税率 12 => 12%.
	standardRate float64
	// 没有 IOSS 的税率 14 => 14%.
	withoutIossRate float64
	// 免征额,超出才要计算税费.
	exemptionInUsd float64
}

func (vat Vat) ID() int64 {
	return vat.id
}

func (vat Vat) CountryCode() string {
	return vat.countryCode
}

func (vat Vat) StandardRate() float64 {
	return vat.standardRate
}

func (vat Vat) WithoutIossRate() float64 {
	return vat.withoutIossRate
}

func (vat Vat) ExemptionInUsd() float64 {
	return vat.exemptionInUsd
}

func NewVat(
	id int64,
	countryCode string,
	standardRate float64,
	withoutIossRate float64,
	exemptionInUsd float64,
) *Vat {
	return &Vat{
		id:              id,
		countryCode:     countryCode,
		standardRate:    standardRate,
		withoutIossRate: withoutIossRate,
		exemptionInUsd:  exemptionInUsd,
	}
}

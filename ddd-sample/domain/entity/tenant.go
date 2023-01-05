/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    tenant
	@Date    2022/5/13 17:37
	@Desc
*/

package entity

type Tenant struct {
	tenantID         int64
	prepayTariff     bool
	presetChannelIds string
	testChannelIds   string
}

func (t Tenant) TenantID() int64 {
	return t.tenantID
}

func (t Tenant) PrepayTariff() bool {
	return t.prepayTariff
}

func (t Tenant) PresetChannelIds() string {
	return t.presetChannelIds
}

func (t Tenant) TestChannelIds() string {
	return t.testChannelIds
}

func NewTenant(
	tenantID int64,
	prepayTariff bool,
	presetChannelIds string,
	testChannelIds string,
) *Tenant {
	return &Tenant{
		tenantID:         tenantID,
		prepayTariff:     prepayTariff,
		presetChannelIds: presetChannelIds,
		testChannelIds:   testChannelIds,
	}
}

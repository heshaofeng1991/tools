/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    country_zone
	@Date    2022/5/30 11:30
	@Desc
*/

package logic

import (
	"context"
	"strings"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	"github.com/pkg/errors"
)

// 通过 zone 过滤渠道报价.
func (h HTTPServer) filterCostsByCountryCodeAndZipCode(ctx context.Context, channelCosts []*domainEntity.ChannelCostRsp,
	countryCode, zipCode string,
) (rsp []*domainEntity.ChannelCostRsp, err error) {
	var countryZones []*domainEntity.CountryZone

	countryZones, err = h.application.Queries.GetCountryZones.Handle(ctx, zipCode, countryCode)

	if err != nil {
		return rsp, errors.Wrap(err, "")
	}

	for _, channelCost := range channelCosts {
		chlCost := domainEntity.NewChannelCostRsp(
			channelCost.ChannelCost(),
			channelCost.ChargeWeight(),
			channelCost.ActualWeight(),
			channelCost.VolumeWeight(),
		)

		if chlCost.ChannelCost().Zone() == "" {
			rsp = append(rsp, channelCost)

			continue
		}

		for _, countryZone := range countryZones {
			if countryZone.ChannelID() == chlCost.ChannelCost().ChannelID() &&
				strings.EqualFold(strings.ToUpper(countryZone.Zone()),
					strings.ToUpper(chlCost.ChannelCost().Zone())) {
				rsp = append(rsp, channelCost)
			}
		}
	}

	return rsp, nil
}

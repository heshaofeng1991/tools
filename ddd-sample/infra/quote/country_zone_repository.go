/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    country_zone
	@Date    2022/5/12 10:09
	@Desc
*/

package quote

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/countryzone"
	"github.com/pkg/errors"
)

const (
	ZtoSzFedexUsPrefix = "ZTO_SZ_FEDEX_US_PREFIX_"
)

func (c ShippingOptionRepository) GetCountryZone(ctx context.Context,
	zipCode, countryCode string,
) ([]*domainEntity.CountryZone, error) {
	query := c.entClient.CountryZone.Query().Where(countryzone.CountryCode(countryCode))

	var (
		err          error
		countryZones []*ent.CountryZone
		zipCodes     []string
	)

	if zipCode == "000000" {
		zipCode = ""
	}

	if zipCode == "" {
		query = query.Where(
			countryzone.Or(
				countryzone.And(
					countryzone.StartZipCode(""),
					countryzone.EndZipCode(""),
					countryzone.ZipCode(""),
				)))
	} else {
		if countryCode == "US" {
			// 美国的邮编是以zipCode的首位判断所在的地区
			zipCodes = append(zipCodes, ZtoSzFedexUsPrefix+zipCode[0:1])
		}

		zipCodes = append(zipCodes, zipCode)

		query = query.Where(
			countryzone.Or(
				countryzone.And(
					countryzone.StartZipCode(""),
					countryzone.EndZipCode(""),
					countryzone.ZipCode(""),
				),
				countryzone.ZipCodeIn(zipCodes...),
				countryzone.And(
					countryzone.StartZipCodeLTE(zipCode),
					countryzone.EndZipCodeGTE(zipCode),
				)))
	}

	countryZones, err = query.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return CovertCountryZones(countryZones), nil
}

func CovertCountryZones(countryZones []*ent.CountryZone) []*domainEntity.CountryZone {
	results := make([]*domainEntity.CountryZone, 0)

	for _, countryZone := range countryZones {
		zone := domainEntity.NewCountryZone(
			countryZone.ID,
			countryZone.ChannelID,
			countryZone.CountryCode,
			countryZone.ZipCode,
			countryZone.StartZipCode,
			countryZone.EndZipCode,
			countryZone.Zone,
		)

		results = append(results, zone)
	}

	return results
}

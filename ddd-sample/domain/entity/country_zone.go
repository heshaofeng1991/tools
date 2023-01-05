/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    country_zone
	@Date    2022/5/13 10:29
	@Desc
*/

package entity

type CountryZone struct {
	id           int
	channelID    int64
	countryCode  string
	zipCode      string
	startZipCode string
	endZipCode   string
	zone         string
}

func (countryZone *CountryZone) ID() int {
	return countryZone.id
}

func (countryZone *CountryZone) ChannelID() int64 {
	return countryZone.channelID
}

func (countryZone *CountryZone) CountryCode() string {
	return countryZone.countryCode
}

func (countryZone *CountryZone) ZipCode() string {
	return countryZone.zipCode
}

func (countryZone *CountryZone) StartZipCode() string {
	return countryZone.startZipCode
}

func (countryZone *CountryZone) EndZipCode() string {
	return countryZone.endZipCode
}

func (countryZone *CountryZone) Zone() string {
	return countryZone.zone
}

func NewCountryZone(
	id int,
	channelID int64,
	countryCode string,
	zipCode string,
	startZipCode string,
	endZipCode string,
	zone string,
) *CountryZone {
	return &CountryZone{
		id:           id,
		channelID:    channelID,
		countryCode:  countryCode,
		zipCode:      zipCode,
		startZipCode: startZipCode,
		endZipCode:   endZipCode,
		zone:         zone,
	}
}

func UnmarshalCountryZoneFromDB(
	id int,
	channelID int64,
	countryCode string,
	zipCode string,
	startZipCode string,
	endZipCode string,
	zone string,
) *CountryZone {
	return NewCountryZone(
		id,
		channelID,
		countryCode,
		zipCode,
		startZipCode,
		endZipCode,
		zone,
	)
}

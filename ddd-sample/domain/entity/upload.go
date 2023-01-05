/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    upload
	@Date    2022/5/17 11:20
	@Desc
*/

package entity

type UploadChannel struct {
	countryCode     string
	countryName     string
	startWeight     int
	endWeight       int
	firstWeight     int
	firstWeightFee  float64
	unitWeight      int
	unitWeightFee   float64
	fuelFee         float64
	processingFee   float64
	registrationFee float64
	miscFee         float64
	mode            int8
	channelCode     string
	zone            string
	minNormalDays   int
	maxNormalDays   int
	averageDays     int
}

func (channel *UploadChannel) CountryCode() string {
	return channel.countryCode
}

func (channel *UploadChannel) CountryName() string {
	return channel.countryName
}

func (channel *UploadChannel) StartWeight() int {
	return channel.startWeight
}

func (channel *UploadChannel) EndWeight() int {
	return channel.endWeight
}

func (channel *UploadChannel) FirstWeight() int {
	return channel.firstWeight
}

func (channel *UploadChannel) FirstWeightFee() float64 {
	return channel.firstWeightFee
}

func (channel *UploadChannel) UnitWeightFee() float64 {
	return channel.unitWeightFee
}

func (channel *UploadChannel) UnitWeight() int {
	return channel.unitWeight
}

func (channel *UploadChannel) FuelFee() float64 {
	return channel.fuelFee
}

func (channel *UploadChannel) ProcessingFee() float64 {
	return channel.processingFee
}

func (channel *UploadChannel) RegistrationFee() float64 {
	return channel.registrationFee
}

func (channel *UploadChannel) MiscFee() float64 {
	return channel.miscFee
}

func (channel *UploadChannel) ChannelCode() string {
	return channel.channelCode
}

func (channel *UploadChannel) Zone() string {
	return channel.zone
}

func (channel *UploadChannel) MaxNormalDays() int {
	return channel.maxNormalDays
}

func (channel *UploadChannel) MinNormalDays() int {
	return channel.minNormalDays
}

func (channel *UploadChannel) AverageDays() int {
	return channel.averageDays
}

func (channel *UploadChannel) Mode() int8 {
	return channel.mode
}

func NewUploadChannel(
	countryCode string,
	countryName string,
	startWeight int,
	endWeight int,
	firstWeight int,
	firstWeightFee float64,
	unitWeight int,
	unitWeightFee float64,
	fuelFee float64,
	processingFee float64,
	registrationFee float64,
	miscFee float64,
	mode int8,
	channelCode string,
	zone string,
	minNormalDays int,
	maxNormalDays int,
	averageDays int,
) *UploadChannel {
	return &UploadChannel{
		countryCode:     countryCode,
		countryName:     countryName,
		startWeight:     startWeight,
		endWeight:       endWeight,
		firstWeight:     firstWeight,
		firstWeightFee:  firstWeightFee,
		unitWeight:      unitWeight,
		unitWeightFee:   unitWeightFee,
		fuelFee:         fuelFee,
		processingFee:   processingFee,
		registrationFee: registrationFee,
		miscFee:         miscFee,
		mode:            mode,
		channelCode:     channelCode,
		zone:            zone,
		minNormalDays:   minNormalDays,
		maxNormalDays:   maxNormalDays,
		averageDays:     averageDays,
	}
}

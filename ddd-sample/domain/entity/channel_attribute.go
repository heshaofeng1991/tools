/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_attribute
	@Date    2022/5/20 12:03
	@Desc
*/

package entity

type ChannelAttribute struct {
	attribute string
}

func (channelAttribute *ChannelAttribute) Attribute() string {
	return channelAttribute.attribute
}

func NewChannelAttribute(attribute string) *ChannelAttribute {
	return &ChannelAttribute{
		attribute: attribute,
	}
}

type UpdateChannelAttribute struct {
	channelID int64
	status    int32
	reason    string
}

func (ucp *UpdateChannelAttribute) ChannelID() int64 {
	return ucp.channelID
}

func (ucp *UpdateChannelAttribute) Status() int32 {
	return ucp.status
}

func (ucp *UpdateChannelAttribute) Reason() string {
	return ucp.reason
}

func NewUpdateChannelAttribute(
	channelID int64,
	status int32,
	reason string,
) *UpdateChannelAttribute {
	return &UpdateChannelAttribute{
		channelID: channelID,
		status:    status,
		reason:    reason,
	}
}

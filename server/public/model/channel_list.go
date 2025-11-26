// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

type ChannelList []*Channel

func (o *ChannelList) Etag() string {
	id := "0"
	var t int64
	var delta int64

	for _, v := range *o {
		if v.LastPostAt > t {
			t = v.LastPostAt
			id = v.ID
		}

		if v.UpdateAt > t {
			t = v.UpdateAt
			id = v.ID
		}
	}

	return Etag(id, t, delta, len(*o))
}

type ChannelListWithTeamData []*ChannelWithTeamData

func (o *ChannelListWithTeamData) Etag() string {
	id := "0"
	var t int64
	var delta int64

	for _, v := range *o {
		if v.LastPostAt > t {
			t = v.LastPostAt
			id = v.ID
		}

		if v.UpdateAt > t {
			t = v.UpdateAt
			id = v.ID
		}

		if v.TeamUpdateAt > t {
			t = v.TeamUpdateAt
			id = v.ID
		}
	}

	return Etag(id, t, delta, len(*o))
}

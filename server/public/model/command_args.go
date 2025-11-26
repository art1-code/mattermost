// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"github.com/mattermost/mattermost/server/public/shared/i18n"
)

type CommandArgs struct {
	UserID          string             `json:"user_id"`
	ChannelID       string             `json:"channel_id"`
	TeamID          string             `json:"team_id"`
	RootID          string             `json:"root_id"`
	ParentID        string             `json:"parent_id"`
	TriggerID       string             `json:"trigger_id,omitempty"`
	Command         string             `json:"command"`
	SiteURL         string             `json:"-"`
	T               i18n.TranslateFunc `json:"-"`
	UserMentions    UserMentionMap     `json:"-"`
	ChannelMentions ChannelMentionMap  `json:"-"`
}

func (o *CommandArgs) Auditable() map[string]any {
	return map[string]any{
		"user_id":    o.UserID,
		"channel_id": o.ChannelID,
		"team_id":    o.TeamID,
		"root_id":    o.RootID,
		"parent_id":  o.ParentID,
		"trigger_id": o.TriggerID,
		"command":    o.Command,
		"site_url":   o.SiteURL,
	}
}

// AddUserMention adds or overrides an entry in UserMentions with name username
// and identifier userID
func (o *CommandArgs) AddUserMention(username, userID string) {
	if o.UserMentions == nil {
		o.UserMentions = make(UserMentionMap)
	}

	o.UserMentions[username] = userID
}

// AddChannelMention adds or overrides an entry in ChannelMentions with name
// channelName and identifier channelID
func (o *CommandArgs) AddChannelMention(channelName, channelID string) {
	if o.ChannelMentions == nil {
		o.ChannelMentions = make(ChannelMentionMap)
	}

	o.ChannelMentions[channelName] = channelID
}

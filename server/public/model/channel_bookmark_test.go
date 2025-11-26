// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChannelBookmarkIsValid(t *testing.T) {
	testCases := []struct {
		Description     string
		Bookmark        *ChannelBookmark
		ExpectedIsValid bool
	}{
		{
			"nil bookmark",
			&ChannelBookmark{},
			false,
		},
		{
			"bookmark without create at timestamp",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   "",
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    0,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark without update at timestamp",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   "",
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    0,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with missing channel id",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   "",
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with invalid channel id",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   "invalid",
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with missing owner id",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     "",
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with invalid user id",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     "invalid",
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with missing displayname",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with missing type",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        "",
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with invalid type",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        "invalid",
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type link with missing link url",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type link with invalid link url",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "invalid",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type link with valid link url",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			true,
		},
		{
			"bookmark of type link with empty image url",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			true,
		},
		{
			"bookmark of type link with invalid image url",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				ImageUrl:    "invalid",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type link with valid image url",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				ImageUrl:    "https://mattermost.com/some-image-without-extension", // we don't care if the URL is an actual image as the client should handle the error
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			true,
		},
		{
			"bookmark of type file with missing file id",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type file with invalid file id",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      "invalid",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type file with valid file id",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      NewID(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			true,
		},
		{
			"bookmark of type file with invalid original id",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      NewID(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
				OriginalID:  "invalid",
			},
			false,
		},
		{
			"bookmark of type file with invalid parent id",
			&ChannelBookmark{
				ID:          NewID(),
				ChannelID:   NewID(),
				OwnerID:     NewID(),
				FileID:      NewID(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
				ParentID:    "invalid",
			},
			false,
		},
		{
			"bookmark of type link with a file ID attached",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   NewID(),
				FileID:      NewID(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "http://somelink",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    0,
			},
			false,
		},
		{
			"bookmark of type file with a url",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   NewID(),
				FileID:      NewID(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "http://somelink",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    0,
			},
			false,
		},
		{
			"bookmark with long display name > limit",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   NewID(),
				FileID:      "",
				DisplayName: strings.Repeat("1", 65),
				SortOrder:   0,
				LinkUrl:     "http://somelink",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    3,
				UpdateAt:    3,
				DeleteAt:    0,
			},
			false,
		},
		{
			"bookmark with long display name < limit",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   NewID(),
				FileID:      "",
				DisplayName: strings.Repeat("1", 64),
				SortOrder:   0,
				LinkUrl:     "http://somelink",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    3,
				UpdateAt:    3,
				DeleteAt:    0,
			},
			true,
		},

		{
			"bookmark with link url > limit",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   NewID(),
				FileID:      "",
				DisplayName: "not last test",
				SortOrder:   0,
				LinkUrl:     "http://somelink?" + strings.Repeat("h", 1024),
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    3,
				UpdateAt:    3,
				DeleteAt:    0,
			},
			false,
		},
		{
			"bookmark with image url > limit",
			&ChannelBookmark{
				ID:          NewID(),
				OwnerID:     NewID(),
				ChannelID:   NewID(),
				FileID:      "",
				DisplayName: "last test",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "http://somelink?" + strings.Repeat("h", 1024),
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    3,
				UpdateAt:    3,
				DeleteAt:    0,
			},
			false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			if testCase.ExpectedIsValid {
				require.Nil(t, testCase.Bookmark.IsValid())
			} else {
				require.NotNil(t, testCase.Bookmark.IsValid())
			}
		})
	}
}

func TestChannelBookmarkPreSave(t *testing.T) {
	bookmark := &ChannelBookmark{
		ID:          NewID(),
		ChannelID:   NewID(),
		OwnerID:     NewID(),
		DisplayName: "display name",
		SortOrder:   0,
		LinkUrl:     "https://mattermost.com",
		Type:        ChannelBookmarkLink,
		DeleteAt:    0,
	}

	originalBookmark := &ChannelBookmark{
		ID:          bookmark.ID,
		ChannelID:   bookmark.ChannelID,
		OwnerID:     bookmark.OwnerID,
		DisplayName: bookmark.DisplayName,
		SortOrder:   bookmark.SortOrder,
		LinkUrl:     bookmark.LinkUrl,
		Type:        bookmark.Type,
		DeleteAt:    bookmark.DeleteAt,
	}

	bookmark.PreSave()
	assert.NotEqual(t, 0, bookmark.CreateAt)
	assert.NotEqual(t, 0, bookmark.UpdateAt)

	originalBookmark.CreateAt = bookmark.CreateAt
	originalBookmark.UpdateAt = bookmark.UpdateAt
	assert.Equal(t, originalBookmark, bookmark)
}

func TestChannelBookmarkPreUpdate(t *testing.T) {
	bookmark := &ChannelBookmark{
		ID:          NewID(),
		ChannelID:   NewID(),
		OwnerID:     NewID(),
		DisplayName: "display name",
		SortOrder:   0,
		LinkUrl:     "https://mattermost.com",
		Type:        ChannelBookmarkLink,
		CreateAt:    2,
		DeleteAt:    0,
	}

	originalBookmark := &ChannelBookmark{
		ID:          bookmark.ID,
		ChannelID:   bookmark.ChannelID,
		OwnerID:     bookmark.OwnerID,
		DisplayName: bookmark.DisplayName,
		SortOrder:   bookmark.SortOrder,
		LinkUrl:     bookmark.LinkUrl,
		Type:        bookmark.Type,
		DeleteAt:    bookmark.DeleteAt,
	}

	bookmark.PreSave()
	assert.NotEqual(t, 0, bookmark.UpdateAt)

	originalBookmark.CreateAt = bookmark.CreateAt
	originalBookmark.UpdateAt = bookmark.UpdateAt
	assert.Equal(t, originalBookmark, bookmark)

	bookmark.PreUpdate()
	assert.Greater(t, bookmark.UpdateAt, originalBookmark.UpdateAt)
}

func TestToBookmarkWithFileInfo(t *testing.T) {
	testCases := []struct {
		name          string
		bookmark      *ChannelBookmark
		fileInfo      *FileInfo
		expectedEmoji string
	}{
		{
			name: "emoji with colons",
			bookmark: &ChannelBookmark{
				ID:          NewID(),
				DisplayName: "test bookmark",
				Emoji:       ":smile:",
				Type:        ChannelBookmarkLink,
			},
			fileInfo:      nil,
			expectedEmoji: "smile",
		},
		{
			name: "emoji without colons",
			bookmark: &ChannelBookmark{
				ID:          NewID(),
				DisplayName: "test bookmark",
				Emoji:       "smile",
				Type:        ChannelBookmarkLink,
			},
			fileInfo:      nil,
			expectedEmoji: "smile",
		},
		{
			name: "empty emoji",
			bookmark: &ChannelBookmark{
				ID:          NewID(),
				DisplayName: "test bookmark",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
			},
			fileInfo:      nil,
			expectedEmoji: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bookmark.ToBookmarkWithFileInfo(tc.fileInfo)
			assert.Equal(t, tc.expectedEmoji, result.Emoji)
		})
	}
}

func TestChannelBookmarkPatch(t *testing.T) {
	p := &ChannelBookmarkPatch{
		DisplayName: NewPointer(NewID()),
		SortOrder:   NewPointer(int64(1)),
		LinkUrl:     NewPointer(NewID()),
	}

	b := ChannelBookmark{
		ID:          NewID(),
		DisplayName: NewID(),
		Type:        ChannelBookmarkLink, // should not update
		LinkUrl:     NewID(),
	}
	b.Patch(p)

	require.Empty(t, b.FileID)
	require.Equal(t, *p.DisplayName, b.DisplayName)
	require.Equal(t, *p.SortOrder, b.SortOrder)
	require.Equal(t, *p.LinkUrl, b.LinkUrl)
	require.Equal(t, ChannelBookmarkLink, b.Type)
}

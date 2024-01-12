// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package gcal

import (
	"github.com/firstfoundry/ff-mattermost-plugin-google-calendar/gcal/calendar/remote"
)

func (c *client) GetSuperuserToken() (string, error) {
	return "", remote.ErrNotImplemented
}

package gcal

import (
	"github.com/firstfoundry/ff-mattermost-plugin-google-calendar/gcal/calendar/remote"
)

func (c *client) GetSchedule(_ []*remote.ScheduleUserInfo, _, _ *remote.DateTime, _ int) ([]*remote.ScheduleInformation, error) {
	return nil, remote.ErrNotImplemented
}

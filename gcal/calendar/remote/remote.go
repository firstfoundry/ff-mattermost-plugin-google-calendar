// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package remote

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/firstfoundry/ff-mattermost-plugin-google-calendar/gcal/calendar/config"
	"github.com/firstfoundry/ff-mattermost-plugin-google-calendar/gcal/calendar/utils/bot"
)

var (
	ErrSuperUserClientNotSupported = errors.New("superuser client is not supported")
	ErrNotImplemented              = errors.New("not implemented")
)

type Remote interface {
	MakeClient(context.Context, *oauth2.Token) Client
	MakeSuperuserClient(ctx context.Context) (Client, error)
	NewOAuth2Config() *oauth2.Config
	HandleWebhook(http.ResponseWriter, *http.Request) []*Notification
	CheckConfiguration(configuration config.StoredConfig) error
}

var Makers = map[string]func(*config.Config, bot.Logger) Remote{}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

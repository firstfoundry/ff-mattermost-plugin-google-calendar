package main

import (
	mattermostplugin "github.com/mattermost/mattermost-server/v6/plugin"

	"github.com/firstfoundry/ff-mattermost-plugin-google-calendar/gcal"
	"github.com/firstfoundry/ff-mattermost-plugin-google-calendar/gcal/calendar/config"
	"github.com/firstfoundry/ff-mattermost-plugin-google-calendar/gcal/calendar/engine"
	"github.com/firstfoundry/ff-mattermost-plugin-google-calendar/gcal/calendar/plugin"
)

var BuildHash string
var BuildHashShort string
var BuildDate string
var CalendarProvider string

func main() {
	config.Provider = gcal.GetGcalProviderConfig()

	mattermostplugin.ClientMain(
		plugin.NewWithEnv(
			engine.Env{
				Config: &config.Config{
					PluginID:       manifest.ID,
					PluginVersion:  manifest.Version,
					BuildHash:      BuildHash,
					BuildHashShort: BuildHashShort,
					BuildDate:      BuildDate,
					Provider:       config.Provider,
				},
				Dependencies: &engine.Dependencies{},
			}))
}

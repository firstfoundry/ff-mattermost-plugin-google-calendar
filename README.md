Google Calendar for FF

This plugin is a modificaion of [this repo](https://github.com/mattermost/mattermost-plugin-google-calendar)

### Requirements

- [golang](https://golang.org/doc/install)
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [npm](https://www.npmjs.com/get-npm)

### Build the plugin

1. Clone this repository.
1. Run `make dist`.
2. When the build process finishes the plugin tarball will be available at `dist/com.mattermost.gcal-$(VERSION).tar.gz`
3. In your Mattermost Server, go to **System Console > Plugin Management** and upload the `.tar.gz` file to install the plugin.
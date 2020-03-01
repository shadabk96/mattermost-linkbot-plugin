# Mattermost Linkbot Plugin

## Getting Started

* Clone this repo
* Edit `serverUrl` in `webapp/webpack.config.js` to Mattermost URL.
* Edit `ServeHTTP` function in `server/plugin.go` to reflect teamID and db location user by linkbot.
* Follow below steps:

Build your plugin:
```
make
```

This will produce a single plugin file (with support for multiple architectures) for upload to your Mattermost server:

```
com.github.shadabk96.mattermost-linkbot-plugin-0.1.0.tar.gz
```
In production, deploy and upload your plugin via the System Console.
# miniflux-discord

This is a small "proxy" program between [Miniflux](https://miniflux.app) and [Discord](https://discord.com/).

Miniflux already supports Discord as an integration, however the "rich" embeds it provides are kinda crap, and there's no way to customize them. So instead, this program acts as a proxy between.

The format this outputs is `__**$TITLE**__\n$URL`, this way Discord can embed links itself.

## Setup

1. Download from [releases](https://github.com/p3lim/miniflux-discord/releases) or use [the container image](https://github.com/p3lim/miniflux-discord/pkgs/container/miniflux-discord).
2. Run it with the following environment variables to configure it:
	- `LISTEN_PORT` - defaults to `8080`
	- `LISTEN_ADDR` - defaults to `0.0.0.0`
	- `DISCORD_WEBHOOK_URL` - required
		- alternatively `DISCORD_WEBHOOK_URL_FILE` to better work with secrets
	- `LOG_LEVEL` - defaults to `INFO`
3. Use this as a Webhook integration (Settings > Integrations > Webhook)
	- e.g. `http://localhost:8080/`

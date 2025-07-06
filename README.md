# miniflux-discord

This is a small wrapper program between [Miniflux](https://miniflux.app) and [Discord](https://discord.com/).

Miniflux already supports Discord as an integration, however the "rich" embeds it provides are kinda crap, and there's no way to customize them. So instead, this program acts as a proxy between.

## Setup

1. Build the program (just `go build ./cmd/...`) or use the container.
2. Run it with the following environment variables to configure it:
	- `LISTEN_PORT` - defaults to `8080`
	- `LISTEN_ADDR` - defaults to `0.0.0.0`
	- `DISCORD_WEBHOOK_URL` - required
	- `LOG_LEVEL` - defaults to `INFO`
3. Use this as a Webhook integration (Settings > Integrations > Webhook)
	- e.g. `http://localhost:8080/`

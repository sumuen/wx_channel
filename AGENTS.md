# Repository Guide (AGENTS)

## Summary
- Windows-focused Go app that runs a local proxy to help download WeChat Channel videos.
- Entry point is `cmd/root.go` -> `internal/app`.

## Repo Layout
- `cmd/`: CLI entry (cobra).
- `internal/`: app core (api, handlers, router, services, storage, utils, websocket).
- `pkg/`: shared helpers (certificate, proxy).
- `web/`: web console static files (served on `port+1`).
- `assets/`, `winres/`: UI/assets and Windows resource metadata.
- `third_party/`: vendored replacements (SunnyNet, go-libutp).
- `docs/`: user/developer documentation.

## Build
- Go version: module uses `go 1.23` with `toolchain go1.24.3`.
- Build (Windows binary):
  - `go build -o wx_channel.exe`
  - Smaller build: `go build -ldflags="-s -w" -o wx_channel.exe`

## Run
- Default:
  - `.\wx_channel.exe`
- Change proxy port:
  - `.\wx_channel.exe -p 8080` or `.\wx_channel.exe --port 8080`
- Config file:
  - `.\wx_channel.exe --config path\to\config.yaml`
- Uninstall certificate:
  - `.\wx_channel.exe --uninstall`

## Configuration
- Loaded from (highest to lowest):
  1) Database overrides (if present)
  2) Environment variables `WX_CHANNEL_*`
  3) Config file `config.yaml` (repo root or `$HOME/.wx_channel/`)
  4) Defaults
- Common env vars:
  - `WX_CHANNEL_PORT` (default 2025)
  - `WX_CHANNEL_DOWNLOADS_DIR` (default `downloads`)
  - `WX_CHANNEL_LOG_FILE` (default `logs/wx_channel.log`)
  - `WX_CHANNEL_LOG_MAX_MB` (default `5`)
  - `WX_CHANNEL_DOWNLOAD_CONCURRENCY`
  - `WX_CHANNEL_TOKEN` (API auth)
  - `WX_CHANNEL_ALLOWED_ORIGINS`

## Runtime Notes
- Proxy listens on `port` (default 2025).
- WebSocket + web console listen on `port+1`.
- On first run, a root cert is installed; if it fails, the cert is written to
  `downloads/SunnyRoot.cer` for manual install.

## Tests
- Run all tests:
  - `go test ./...`
- Tests live under `internal/*_test.go`.


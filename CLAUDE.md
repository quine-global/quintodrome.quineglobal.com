# quintodrome-quineglobal-com

Website and auto-update feed server for [Quintodrome](https://github.com/quine-global/quintodrome), a native desktop app that wraps [Navidrome](https://www.navidrome.org) (a self-hosted music server) for macOS, Windows, and Linux.

## Commands

```bash
make dev        # start Go server + Tailwind watcher (Ctrl-C stops both)
make start      # Go server only
make test       # go test -shuffle on ./...
make css        # one-shot Tailwind build
make lint       # golangci-lint
make build      # Docker image (linux/amd64)
make push       # build + push to forge.quinefoundation.com
```

## Architecture

Pure Go HTTP server (`go-chi/chi`). HTML is rendered server-side with `maragu.dev/gomponents` (no templates). CSS is Tailwind, compiled to `public/styles/app.css`.

```
cmd/app/        entry point
http/           route handlers, server, release cache, update-check logic
html/           gomponents page renderers and shared types (Release, DownloadProps, etc.)
public/         static assets (images, compiled CSS)
scripts/        release refresh helper script
```

## Key concepts

**Release cache** (`http/releases.go`): fetches GitHub releases from the `quine-global/quintodrome` repo via the GitHub API, caches them in memory. `TryRefresh` is debounced (called per-request, no-ops if cache is fresh or a fetch is already running). `ForceRefresh` is used by the internal `/internal/refresh` endpoint.

**Asset classification** (`classifyAsset` in `http/releases.go`): maps GitHub asset filenames to `(os, arch)` keys used in `Release.Assets`. Supported platforms: `mac/{arm64,x64}` (`.dmg`), `windows/x64` (`.exe`), `linux/{arm64,x64}` (AppImage), `linux-rpm`, `linux-deb`, `linux-snap`, `linux-pacman`, `linux-flatpak` (each `{arm64,x64}`). Arch is detected from either `arm64` or `aarch64` in the filename. armv7l and other unrecognised ARM variants are explicitly excluded (return `ok=false`) so they don't get misclassified as x64. Quintodrome's actual releases currently ship `.dmg`/`.exe`/`.flatpak`/`.snap` only (no Linux AppImage/deb/rpm/pacman yet); `.msi` and `.app.zip` variants are intentionally left unclassified in favor of `.exe`/`.dmg`. The upstream repo also publishes `navidrome_*` server-binary assets (tar.gz/deb/rpm for the bundled Navidrome server itself) — those are unrelated to the desktop app and are not classified here.

**Auto-update feed** (`http/updatecheck.go`): implements a Hazel/Squirrel-style update protocol. Responds to `GET /api/update-check` and `GET /update` with 204 (no update) or 200 JSON `{name, notes, pub_date, url}`. Handles stable and canary tracks; platform strings from the client (`darwin`, `win32`, `deb`, etc.) are normalised via `normalizePlatformForAssets`.

**Download page** (`http/home.go`, `html/download.go`): detects platform from `User-Agent`, supports manual override via `?os=` and `?arch=` query params. Linux defaults to the Flatpak format (`?os=linux-flatpak`) since that's what's actually published today; the format toggle also offers Snap, AppImage, DEB, RPM, and Pacman for forward compatibility.

## Testing

Tests live alongside source in `package http` (internal) or `package http_test` (integration). Run with `make test`. Notable test files:

- `http/releases_test.go` — `classifyAsset` and `parseTag`
- `http/updatecheck_test.go` — `semverCmp`, `newerCanary`, `normalizePlatformForAssets`
- `http/server_test.go` — integration test that starts the real server

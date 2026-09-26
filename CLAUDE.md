# quineglobal-sites

Monorepo for Quine Global's website + auto-update feed servers: currently
[Quintodrome](https://github.com/quine-global/quintodrome) (a native desktop
app wrapping [Navidrome](https://www.navidrome.org)) and
[Quine Hyper](https://github.com/quine-global/hyper) (a fork of the Hyper
terminal). Both are the same server shape — release feed, download page,
Hazel/Squirrel-style auto-update API — with only branding, asset-naming
rules, and landing-page content differing per site.

## Commands

All commands take `SITE=quintodrome` (default) or `SITE=hyper`.

```bash
make dev                  # start Go server + Tailwind watcher for SITE (Ctrl-C stops both)
make dev SITE=hyper       # same, for the hyper site
make start                # Go server only
make test                 # go test -shuffle on ./... (covers every site)
make css                  # one-shot Tailwind build (shared across sites)
make lint                 # golangci-lint
make build                # Docker image for SITE (linux/amd64)
make push                 # build + push to forge.quinefoundation.com
```

## Architecture

Pure Go HTTP server (`go-chi/chi`). HTML is rendered server-side with `maragu.dev/gomponents` (no templates). CSS is Tailwind, compiled to `public/styles/app.css`.

```
internal/site/       shared server framework: Server, routes, release cache, update-check protocol
internal/component/  shared gomponents chrome: page/nav/footer, download page, FeatureCard, DeviceWindow
sites/quintodrome/   Quintodrome's Config, classifyAsset, landing-page content, cmd/main.go
sites/hyper/         Quine Hyper's Config, classifyAsset, landing-page content, cmd/main.go
public/              static assets shared by all sites (images, compiled CSS)
scripts/             release refresh helper script (shared, identical across sites)
```

A site is defined entirely by `sites/<site>/site.go`'s `Config` value
(`internal/site.Config`, embedding `component.Branding` and
`component.DownloadConfig`) plus that package's `classifyAsset` and
`HomeContent`. Everything else — HTTP routing, the release cache, the
update-check handler, the download page's mechanics — is shared code in
`internal/site` and `internal/component`.

## Key concepts

**Release cache** (`internal/site/releases.go`): fetches GitHub releases from `Config.RepoOwner/Config.RepoName` via the GitHub API, caches them in memory per-server. `TryRefresh` is debounced (called per-request, no-ops if cache is fresh or a fetch is already running). `ForceRefresh` is used by the internal `/internal/refresh` endpoint.

**Asset classification** (`sites/<site>/assets.go`): each site's own `classifyAsset` maps GitHub release asset filenames to `(os, arch)` keys used in `Release.Assets`, since real release artifact naming conventions differ per product (this is intentionally *not* shared logic). Supported OS keys: `mac`, `windows`, `linux` (AppImage), `linux-rpm`, `linux-deb`, `linux-snap`, `linux-pacman`, `linux-flatpak`. armv7l and other unrecognised ARM variants are explicitly excluded (return `ok=false`) rather than misclassified as x64.

**Auto-update feed** (`internal/site/updatecheck.go`): implements a Hazel/Squirrel-style update protocol, shared by all sites. Responds to `GET /api/update-check` and `GET /update` with 204 (no update) or 200 JSON `{name, notes, pub_date, url}`. Handles stable and canary tracks; platform strings from the client (`darwin`, `win32`, `deb`, etc.) are normalised via `normalizePlatformForAssets`.

**Download page** (`internal/site/home.go`, `internal/component/download.go`): detects platform from `User-Agent` (Linux default comes from `Config.DefaultLinuxOS`), supports manual override via `?os=` and `?arch=` query params. The Linux format toggle's button list, order, and per-format architecture availability are entirely driven by `Config.LinuxFormats` — e.g. Quintodrome offers Flatpak first and ships Snap for both arm64/x64, while Hyper offers AppImage first and only ships Snap for x64 (hidden on arm64 automatically since the config declares it x64-only).

## Testing

Tests live alongside source: `package site`/`site_test` in `internal/site`, `package quintodrome`/`hyper` in each site dir. Run with `make test`. Notable test files:

- `sites/<site>/assets_test.go` — per-site `classifyAsset` cases
- `internal/site/releases_test.go` — shared `parseTag`
- `internal/site/updatecheck_test.go` — shared `semverCmp`, `newerCanary`, `normalizePlatformForAssets`
- `internal/site/server_test.go` — integration test that starts a real server against a minimal test `Config`

package site

import (
	. "maragu.dev/gomponents"

	"app/internal/component"
)

// Config is the full set of per-site knobs the shared server and page
// chrome need. Each site under sites/ builds one Config value and passes
// it to [NewServer].
type Config struct {
	component.Branding
	component.DownloadConfig

	// RepoOwner/RepoName identify the GitHub repo releases are fetched from
	// and linked to, e.g. "quine-global" / "quintodrome".
	RepoOwner string
	RepoName  string

	// ClassifyAsset maps a GitHub release asset filename to an (os, arch)
	// key used in Release.Assets. Each site's real release artifacts have
	// their own naming quirks, so this is not shared logic.
	ClassifyAsset func(name string) (os, arch string, ok bool)

	// HomeContent renders the site's landing page (hero, mockup, feature
	// cards, origin story). It's expected to call component.Page itself.
	// It takes the full Config (rather than closing over the site's package-
	// level Config var) to avoid a package initialization cycle.
	HomeContent func(cfg Config, props component.PageProps) Node
}

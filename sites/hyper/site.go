// Package hyper holds the Quine Hyper site's identity: branding, asset
// classification, download-format list, and landing-page content.
// Everything else (server, routes, release cache, update-check protocol,
// page chrome) lives in the shared internal/site and internal/component
// packages.
package hyper

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/internal/component"
	"app/internal/site"
)

// Config is the Quine Hyper site configuration, passed to site.NewServer.
var Config = site.Config{
	Branding: component.Branding{
		Name:      "Quine Hyper",
		LogoPath:  "/images/hyper-logo.png",
		GitHubURL: "https://github.com/quine-global/hyper",
		FooterNote: Group{
			Text("Built on "),
			A(Href("https://hyper.is"), Class("text-zinc-400 hover:text-white underline"), Text("Hyper")),
			Text(" by Vercel."),
		},
	},
	DownloadConfig: component.DownloadConfig{
		DefaultLinuxOS: "linux",
		Formats: []component.LinuxFormat{
			{OS: "linux", Label: "AppImage", Archs: []string{"arm64", "x64"}},
			{OS: "linux-deb", Label: "DEB", Archs: []string{"arm64", "x64"}},
			{OS: "linux-rpm", Label: "RPM", Archs: []string{"arm64", "x64"}},
			{OS: "linux-pacman", Label: "Pacman", Archs: []string{"arm64", "x64"}},
			{OS: "linux-flatpak", Label: "Flatpak", Archs: []string{"arm64", "x64"}},
			// Hyper only ships Snap for x64.
			{OS: "linux-snap", Label: "Snap", Archs: []string{"x64"}},
		},
	},
	RepoOwner:     "quine-global",
	RepoName:      "hyper",
	ClassifyAsset: classifyAsset,
	HomeContent:   HomeContent,
}

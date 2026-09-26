// Package quintodrome holds the Quintodrome site's identity: branding,
// asset classification, download-format list, and landing-page content.
// Everything else (server, routes, release cache, update-check protocol,
// page chrome) lives in the shared internal/site and internal/component
// packages.
package quintodrome

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/internal/component"
	"app/internal/site"
)

// Config is the Quintodrome site configuration, passed to site.NewServer.
var Config = site.Config{
	Branding: component.Branding{
		Name:      "Quintodrome",
		LogoPath:  "/images/quintodrome-logo.png",
		GitHubURL: "https://github.com/quine-global/quintodrome",
		FooterNote: Group{
			Text("Built on "),
			A(Href("https://www.navidrome.org"), Class("text-zinc-400 hover:text-white underline"), Text("Navidrome")),
			Text("."),
		},
	},
	DownloadConfig: component.DownloadConfig{
		// Quintodrome currently ships Linux builds as Flatpak/Snap only (no AppImage).
		DefaultLinuxOS: "linux-flatpak",
		Formats: []component.LinuxFormat{
			{OS: "linux-flatpak", Label: "Flatpak", Archs: []string{"arm64", "x64"}},
			{OS: "linux-snap", Label: "Snap", Archs: []string{"arm64", "x64"}},
			{OS: "linux", Label: "AppImage", Archs: []string{"arm64", "x64"}},
			{OS: "linux-deb", Label: "DEB", Archs: []string{"arm64", "x64"}},
			{OS: "linux-rpm", Label: "RPM", Archs: []string{"arm64", "x64"}},
			{OS: "linux-pacman", Label: "Pacman", Archs: []string{"arm64", "x64"}},
		},
	},
	RepoOwner:     "quine-global",
	RepoName:      "quintodrome",
	ClassifyAsset: classifyAsset,
	HomeContent:   HomeContent,
}

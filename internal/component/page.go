// Package component holds all the common HTML components and utilities
// shared across sites. Site-specific branding is threaded through via
// [Branding] and [DownloadConfig] rather than hardcoded.
package component

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

var hashOnce sync.Once
var appCSSPath string

// PageProps are properties for the [Page] component.
type PageProps struct {
	Title       string
	Description string
}

// Branding holds the per-site identity used by the shared page chrome.
type Branding struct {
	Name      string // e.g. "Quintodrome"
	LogoPath  string // e.g. "/images/quintodrome-logo.png"
	GitHubURL string // e.g. "https://github.com/quine-global/quintodrome"

	// FooterNote is the "Built on <link>." fragment rendered after the
	// shared "© 2026 Quine Global. " copyright line. Each site supplies
	// its own link and trailing text (e.g. "Navidrome." vs "Hyper by Vercel.").
	FooterNote Node
}

// Page layout with nav, main content, and footer.
func Page(b Branding, props PageProps, children ...Node) Node {
	hashOnce.Do(func() {
		appCSSPath = getHashedPath("public/styles/app.css")
	})

	return HTML5(HTML5Props{
		Title:       props.Title,
		Description: props.Description,
		Language:    "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href(appCSSPath)),
			Link(Rel("preconnect"), Href("https://fonts.googleapis.com")),
			Link(Rel("preconnect"), Href("https://fonts.gstatic.com"), Attr("crossorigin", "")),
			Link(Rel("stylesheet"), Href("https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600;700&display=swap")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
		},
		Body: Group{
			siteNav(b),
			Main(Class("mx-auto max-w-5xl px-6 py-12"),
				Group(children),
			),
			siteFooter(b),
		},
	})
}

func siteNav(b Branding) Node {
	return Header(Class("border-b border-zinc-800 bg-zinc-950"),
		Div(Class("mx-auto flex max-w-5xl items-center justify-between px-6 py-4"),
			A(Href("/"), Class("flex items-center gap-3 no-underline"),
				Img(Src(b.LogoPath), Alt(b.Name), Class("h-8 w-8 rounded")),
				Span(Class("font-mono text-lg font-bold tracking-tight text-white"),
					Text(b.Name),
				),
			),
			Div(Class("flex items-center gap-4"),
				A(Href(b.GitHubURL),
					Class("font-mono text-sm text-zinc-400 hover:text-white transition-colors"),
					Text("GitHub"),
				),
				A(Href("/download"),
					Class("rounded-md bg-primary-500 px-4 py-2 font-mono text-sm font-semibold text-black hover:bg-primary-400 transition-colors"),
					Text("Download"),
				),
			),
		),
	)
}

func siteFooter(b Branding) Node {
	return Footer(Class("mt-24 border-t border-zinc-800 py-8"),
		Div(Class("mx-auto flex max-w-5xl flex-col items-center justify-between gap-4 px-6 sm:flex-row"),
			Div(Class("flex items-center gap-3"),
				Img(Src("/images/quineglobal-logo.png"), Alt("Quine Global"), Class("h-6 w-6 rounded-full")),
				Span(Class("font-mono text-sm text-zinc-500"),
					Text("© 2026 Quine Global. "),
					b.FooterNote,
				),
			),
			Div(Class("flex gap-6 font-mono text-sm"),
				A(Href("/"), Class("text-zinc-500 hover:text-white transition-colors"), Text("Home")),
				A(Href(b.GitHubURL), Class("text-zinc-500 hover:text-white transition-colors"), Text("GitHub")),
				A(Href("/download"), Class("text-zinc-500 hover:text-white transition-colors"), Text("Download")),
			),
		),
	)
}

func getHashedPath(path string) string {
	externalPath := strings.TrimPrefix(path, "public")
	ext := filepath.Ext(path)
	if ext == "" {
		panic("no extension found")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("%v.x%v", strings.TrimSuffix(externalPath, ext), ext)
	}

	return fmt.Sprintf("%v.%x%v", strings.TrimSuffix(externalPath, ext), sha256.Sum256(data), ext)
}

package quintodrome

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/internal/component"
	"app/internal/site"
)

// HomeContent is the Quintodrome landing page.
func HomeContent(cfg site.Config, props component.PageProps) Node {
	props.Title = "Quintodrome — Your music, self-hosted"
	props.Description = "Quintodrome wraps Navidrome, your self-hosted music server, into a native desktop app for macOS, Windows, and Linux."

	return component.Page(cfg.Branding, props,
		// Hero
		Div(Class("py-20 text-center"),
			Div(Class("mb-6 inline-flex items-center gap-2 rounded-full border border-zinc-700 bg-zinc-900 px-4 py-1.5"),
				Span(Class("h-2 w-2 rounded-full bg-primary-400")),
				Span(Class("font-mono text-xs text-zinc-400"), Text("Powered by Navidrome")),
			),
			H1(Class("font-mono text-5xl font-bold tracking-tight text-white sm:text-6xl"),
				Text("Your music. Your server."),
			),
			P(Class("mx-auto mt-6 max-w-2xl font-mono text-lg text-zinc-400"),
				Text("Quintodrome wraps your self-hosted music library into a native desktop app. No setup, no subscription, no one else's server."),
			),
			Div(Class("mt-10 flex flex-wrap items-center justify-center gap-4"),
				A(Href("/download"),
					Class("inline-flex items-center gap-2 rounded-lg bg-primary-500 px-8 py-3 font-mono text-base font-semibold text-black hover:bg-primary-400 transition-colors"),
					Text("Download for free"),
				),
				A(Href(cfg.GitHubURL),
					Class("inline-flex items-center gap-2 rounded-lg border border-zinc-700 bg-zinc-900 px-8 py-3 font-mono text-base font-semibold text-white hover:border-zinc-500 transition-colors"),
					Text("View on GitHub"),
				),
			),
		),

		// Player mockup
		musicPlayerMockup(),

		// Features
		Div(Class("pb-20"),
			H2(Class("text-center font-mono text-xs font-semibold uppercase tracking-widest text-zinc-500"),
				Text("Why Quintodrome"),
			),
			Div(Class("mt-12 grid grid-cols-1 gap-6 md:grid-cols-2"),
				component.FeatureCard("Zero Setup",
					Text("Launch the app and it provisions and starts its own music server for you — no admin account to create, no config file to write."),
					"double-click → listening",
				),
				component.FeatureCard("Cross-Platform",
					Text("Native apps for macOS, Windows, and Linux, all wrapping the same self-hosted server underneath."),
					"Works everywhere you do.",
				),
				component.FeatureCard("Your Library, Your Rules",
					Group{
						Text("Point it at your music folder and it's yours: no subscription, no catalog gaps, no algorithm. Built on "),
						A(Href("https://www.navidrome.org"),
							Class("text-primary-400 hover:text-primary-300 underline"),
							Text("Navidrome"),
						),
						Text("."),
					},
					"~/Music → your library",
				),
				component.FeatureCard("Subsonic Compatible",
					Text("Speaks the Subsonic/OpenSubsonic API, so your library also works with the mobile and desktop clients you already use."),
					"stream anywhere, from one server",
				),
			),
		),

		// Origin story
		Div(Class("mb-20 rounded-xl border border-zinc-800 bg-zinc-900 p-10 text-center"),
			Img(Src("/images/quineglobal-logo.png"), Alt("Quine Global"), Class("mx-auto mb-6 h-12 w-12 rounded-full")),
			H2(Class("font-mono text-2xl font-bold text-white"), Text("Built on Navidrome")),
			P(Class("mx-auto mt-4 max-w-xl font-mono text-base text-zinc-400"),
				Text("Quintodrome is a fork of "),
				A(Href("https://github.com/navidrome/navidrome"), Class("text-primary-400 hover:text-primary-300 underline"), Text("Navidrome")),
				Text(", the open-source self-hosted music server. We wrap it in a native desktop shell so there's nothing to install or configure separately."),
			),
			Div(Class("mt-8"),
				A(Href(cfg.GitHubURL),
					Class("inline-block rounded-lg border border-zinc-700 bg-zinc-800 px-6 py-3 font-mono text-sm font-semibold text-white hover:border-zinc-500 transition-colors"),
					Text("Star us on GitHub"),
				),
			),
		),
	)
}

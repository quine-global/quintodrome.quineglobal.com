package hyper

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/internal/component"
	"app/internal/site"
)

// HomeContent is the Quine Hyper landing page.
func HomeContent(cfg site.Config, props component.PageProps) Node {
	props.Title = "Quine Hyper — The thoughtful terminal"
	props.Description = "Quine Hyper is a beautiful, extensible terminal emulator forked from Vercel's Hyper. Built on web technology with a thriving plugin ecosystem."

	return component.Page(cfg.Branding, props,
		// Hero
		Div(Class("py-20 text-center"),
			Div(Class("mb-6 inline-flex items-center gap-2 rounded-full border border-zinc-700 bg-zinc-900 px-4 py-1.5"),
				Span(Class("h-2 w-2 rounded-full bg-primary-400")),
				Span(Class("font-mono text-xs text-zinc-400"), Text("Forked from Vercel's Hyper")),
			),
			H1(Class("font-mono text-5xl font-bold tracking-tight text-white sm:text-6xl"),
				Text("The thoughtful terminal."),
			),
			P(Class("mx-auto mt-6 max-w-2xl font-mono text-lg text-zinc-400"),
				Text("A beautiful, extensible terminal built with care."),
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

		// Terminal window mockup
		terminalMockup(),

		// Features
		Div(Class("pb-20"),
			H2(Class("text-center font-mono text-xs font-semibold uppercase tracking-widest text-zinc-500"),
				Text("Why Quine Hyper"),
			),
			Div(Class("mt-12 grid grid-cols-1 gap-6 md:grid-cols-2"),
				component.FeatureCard("Plugin Ecosystem",
					Group{
						Text("Install plugins and themes from npm. Browse "),
						A(Href("https://hyper.is/plugins"),
							Class("text-primary-400 hover:text-primary-300 underline"),
							Text("featured plugins"),
						),
						Text(" and "),
						A(Href("https://hyper.is/themes"),
							Class("text-primary-400 hover:text-primary-300 underline"),
							Text("featured themes"),
						),
						Text(". Build your own with HTML, CSS, and JavaScript."),
					},
					"npm install hyper-snazzy",
				),
				component.FeatureCard("Cross-Platform",
					Text("Runs on macOS, Windows, and Linux. One consistent experience across every machine you use."),
					"Works everywhere you do.",
				),
				component.FeatureCard("Web-Native",
					Text("Built on Electron, React, and xterm.js. Hack on the terminal itself using the same tools you use every day."),
					"HTML + CSS + JS = your terminal",
				),
				component.FeatureCard("Fully Customizable",
					Text("Tweak every detail with CSS. Change fonts, colors, animations, and layouts — no config file limits."),
					"hyper.json → your rules",
				),
			),
		),

		// Origin story
		Div(Class("mb-20 rounded-xl border border-zinc-800 bg-zinc-900 p-10 text-center"),
			Img(Src("/images/quineglobal-logo.png"), Alt("Quine Global"), Class("mx-auto mb-6 h-12 w-12 rounded-full")),
			H2(Class("font-mono text-2xl font-bold text-white"), Text("Built on Hyper")),
			P(Class("mx-auto mt-4 max-w-xl font-mono text-base text-zinc-400"),
				Text("Quine Hyper is a fork of "),
				A(Href("https://hyper.is"), Class("text-primary-400 hover:text-primary-300 underline"), Text("Hyper")),
				Text(", the open-source terminal by Vercel. We're building on that foundation with new capabilities and active development."),
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

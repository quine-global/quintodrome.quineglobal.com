package component

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// FeatureCard is one card in the "Why <product>" features grid.
func FeatureCard(title string, body Node, code string) Node {
	return Div(Class("rounded-xl border border-zinc-800 bg-zinc-900 p-8"),
		H3(Class("font-mono text-base font-semibold text-white"), Text(title)),
		P(Class("mt-3 font-mono text-sm text-zinc-400 leading-relaxed"), body),
		Div(Class("mt-5 inline-block rounded-md bg-zinc-800 px-3 py-1.5"),
			Span(Class("font-mono text-xs text-primary-400"), Text(code)),
		),
	)
}

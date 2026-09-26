package component

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// DeviceWindow is the shared "device window" chrome (rounded border, shadow,
// traffic-light title bar) used by each site's home-page mockup. Sites
// supply their own body content (a music player, a terminal, etc.).
func DeviceWindow(titleBarLabel string, body Node) Node {
	return Div(Class("mx-auto mb-20 max-w-3xl overflow-hidden rounded-xl border border-zinc-800 bg-zinc-900 shadow-2xl"),
		// Title bar
		Div(Class("flex items-center gap-2 border-b border-zinc-800 bg-zinc-900 px-4 py-3"),
			Div(Class("h-3 w-3 rounded-full bg-red-500 opacity-80")),
			Div(Class("h-3 w-3 rounded-full bg-yellow-500 opacity-80")),
			Div(Class("h-3 w-3 rounded-full bg-primary-500 opacity-80")),
			Span(Class("ml-4 font-mono text-xs text-zinc-500"), Text(titleBarLabel)),
		),
		body,
	)
}

package quintodrome

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/internal/component"
)

// musicPlayerMockup is the "now playing" device-window body shown on the
// landing page.
func musicPlayerMockup() Node {
	return component.DeviceWindow("quintodrome — now playing",
		Div(Class("p-8"),
			Div(Class("flex items-center gap-5"),
				Div(Class("h-20 w-20 shrink-0 rounded-lg bg-gradient-to-br from-primary-400 to-primary-500")),
				Div(Class("min-w-0 text-left"),
					P(Class("truncate font-mono text-base font-semibold text-white"), Text("Brandenburg Concerto No. 3")),
					P(Class("mt-1 truncate font-mono text-sm text-zinc-400"), Text("J.S. Bach — Library/Classical")),
				),
			),
			Div(Class("mt-6"),
				Div(Class("h-1.5 w-full overflow-hidden rounded-full bg-zinc-800"),
					Div(Class("h-full w-1/3 rounded-full bg-primary-400")),
				),
				Div(Class("mt-2 flex justify-between font-mono text-xs text-zinc-500"),
					Span(Text("1:42")),
					Span(Text("5:07")),
				),
			),
			Div(Class("mt-6 flex items-center justify-center gap-6 font-mono text-lg text-zinc-400"),
				Span(Text("⏮")),
				Span(Class("text-2xl text-white"), Text("⏸")),
				Span(Text("⏭")),
			),
		),
	)
}

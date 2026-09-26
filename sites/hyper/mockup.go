package hyper

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/internal/component"
)

// terminalMockup is the terminal device-window body shown on the landing page.
func terminalMockup() Node {
	return component.DeviceWindow("quine-hyper — bash — 80×24",
		Div(Class("p-6 font-mono text-sm leading-relaxed"),
			termPrompt("deep-thought --answer"),
			termOutput("42"),
			termPrompt("hyper install infinite-improbability"),
			termOutput("✔ Engaging drive..."),
			termOutput("✔ Probability: 2:1 against."),
			termOutput("✔ Don't Panic."),
			termPrompt("ls ~/projects"),
			termOutput("drwxr-xr-x  earth/"),
			termOutput("drwxr-xr-x  magrathea/"),
			termOutput("drwxr-xr-x  slartibartfast/"),
			Div(Class("flex items-center gap-0 text-primary-400"),
				Span(Class("text-zinc-500"), Text("❯ ")),
				Span(Class("animate-pulse"), Text("█")),
			),
		),
	)
}

func termPrompt(cmd string) Node {
	return Div(Class("flex items-center gap-2"),
		Span(Class("text-primary-400"), Text("❯")),
		Span(Class("text-white"), Text(cmd)),
	)
}

func termOutput(line string) Node {
	return Div(Class("pl-5 text-zinc-400"), Text(line))
}

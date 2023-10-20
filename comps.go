package main

import (
	"time"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	s "github.com/maragudk/gomponents/svg"
)

func NavbarDaisy(links []PageLink) g.Node {
	return h.Nav(h.Class("navbar bg-base-100"),
		h.Div(h.Class("navbar-start"),
			h.A(h.Href("/"), h.Class("btn btn-ghost normal-case text-xl"), g.Text("Bourgomponents")),
		),
		h.Div(h.Class("navbar-end"),
			h.Div(h.Class("dropdown dropdown-end"),
				h.Label(h.TabIndex("0"), h.Class("btn btn-ghost btn-circle"),
					h.SVG(h.Class("h-5 w-5"), s.Fill("none"), s.ViewBox("0 0 24 24"), s.Stroke("currentColor"),
						s.Path(s.StrokeWidth("2"), s.D("M4 6h16M4 18h7")),
					),
				),
				h.Ul(h.TabIndex("0"), h.Class("menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"),
					g.Group(g.Map(links, func(l PageLink) g.Node {
						return h.Li(h.A(h.Href(l.Path), g.Text(l.Name)))
					})),
				),
			),
		),
	)
}

func Container(children ...g.Node) g.Node {
	return h.Main(h.Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}

func Prose(children ...g.Node) g.Node {
	return h.Section(h.Class("prose"), g.Group(children))
}

func PageFooter() g.Node {
	return h.Footer(h.Class("prose prose-sm prose-indigo"),
		h.P(
			// We can use string interpolation directly, like fmt.Sprintf.
			g.Textf("Rendered %v. ", time.Now().Format(time.RFC3339)),

			// Conditional inclusion
			g.If(time.Now().Second()%2 == 0, g.Text("It's an even second.")),
			g.If(time.Now().Second()%2 == 1, g.Text("It's an odd second.")),
		),

		h.P(h.A(h.Href("https://www.gomponents.com"), g.Text("gomponents"))),
	)
}

func Button(class, text string) g.Node {
	return h.Button(h.Class(class), g.Text(text))
}

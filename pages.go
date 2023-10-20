package main

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func indexPage() (string, g.Node) {
	return "HTML2GO", Container(h.Class("hero min-h-screen"),
		h.Div(h.Class("hero-overlay bg-opacity-60")),
		h.Div(h.Class("hero-content text-center text-neutral-content"),
			h.Div(h.Class("max-w-md"),
				h.H1(h.Class("mb-5 text-xl font-bold"), g.Text("Hello there")),
				h.P(h.Class("mb-5"), g.Text("Hey there")),
				Button("btn btn-primary btn-outline", "Let's Go!"),
			),
		),
	)
}

func contactPage() (string, g.Node) {
	return "Contact", h.Article(
		h.H1(g.Text("Contact us")),
		h.P(g.Text("Just do it.")),
	)
}

func aboutPage() (string, g.Node) {
	return "About", h.Article(
		h.H1(g.Text("About this site")),
		h.P(g.Text("This is a site showing off gomponents.")),
	)
}

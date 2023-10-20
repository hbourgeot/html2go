package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

func main() {
	app := fiber.New()

	app.Get("/", adaptor.HTTPHandler(createHandler(indexPage())))
	app.Get("/contact", adaptor.HTTPHandler(createHandler(contactPage())))
	app.Get("/about", adaptor.HTTPHandler(createHandler(aboutPage())))

	app.Listen(":5050")
}

func createHandler(title string, body g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = Page(title, r.URL.Path, body).Render(w)
	}
}

func Page(title, path string, body g.Node) g.Node {
	links := []PageLink{
		{Path: "/contact", Name: "Contact"},
		{Path: "/about", Name: "About"},
	}

	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			h.Link(h.Rel("stylesheet"), h.Type("text/css"), h.Href("https://cdn.jsdelivr.net/npm/daisyui@3.9.2/dist/full.css")),
			h.Script(h.Src("https://cdn.tailwindcss.com?plugins=typography")),
		},
		Body: []g.Node{
			NavbarDaisy(links), Container(Prose(body), PageFooter()),
		},
	})
}

package article

import (
	"net/http"
	"github.com/gowiki/data"
	"html/template"
)
var Templates = template.Must(template.ParseGlob("html/template/*"))
func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := data.LoadArticle(title)

	if err != nil {
		p = &data.Article{
			Title:title,
		}
	}

	Templates.ExecuteTemplate(w, "edit.html", p)

}

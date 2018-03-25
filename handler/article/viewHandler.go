package article

import (
	"net/http"
	"github.com/gowiki/data"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := data.LoadArticle(title)
	Templates.ExecuteTemplate(w, "view.html", p)

}

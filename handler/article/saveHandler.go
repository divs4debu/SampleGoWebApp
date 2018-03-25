package article

import (
	"net/http"
	"github.com/gowiki/data"
	"fmt"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	article := &data.Article{
		Title:title,
		Body:[]byte(body),
	}
	fmt.Println("moving on")
	val := article.Save()
	if val != nil {
		fmt.Print("error")
	}
	http.Redirect(w,r,"/view/"+title, http.StatusFound)
}
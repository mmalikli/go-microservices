package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	// "os"
	// "path/filepath"
)

var tpl *template.Template

func init() {
  wd, _ := os.Getwd()
	//folder, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//fmt.Println(os.Getwd())
	tpl = template.Must(template.ParseGlob(wd + "/frontend/cmd/web/templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	// partials := []string{
	// 	"./cmd/web/templates/base.layout.gohtml",
	// 	"./cmd/web/templates/header.partial.gohtml",
	// 	"./cmd/web/templates/footer.partial.gohtml",
	// }

	// var templateSlice []string
	// templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	// for _, x := range partials {
	// 	templateSlice = append(templateSlice, x)
	// }

	//tmpl, err := template.ParseFiles(templateSlice...)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	if err := tpl.ExecuteTemplate(w, t, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
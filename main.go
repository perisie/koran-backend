package main

import (
	"github.com/arikama/koran-backend/src/handler"
	"github.com/arikama/koran-backend/src/quran"
	"html/template"
	"net/http"
)

func main() {
	var tmpl *template.Template
	var mngr_quran quran.Mngr

	tmpl = template.Must(template.ParseGlob("src/template/*.html"))
	fs := http.FileServer(http.Dir("static"))
	mngr_quran, _ = quran.Mngr_impl_new("qurancsv")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home(tmpl, mngr_quran))
	mux.HandleFunc("/static/", handler.Static(fs))

	_ = http.ListenAndServe(":8001", mux)
}

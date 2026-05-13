package routes

import "net/http"

func RegisterTemplateRoutes(mux *http.ServeMux) {
	AdminfileServer := http.FileServer(http.Dir("./templates/admin"))
	ClientfileServer := http.FileServer(http.Dir("./templates/client/dist"))

	mux.Handle("GET /", ClientfileServer)
	mux.Handle("GET /admin/", http.StripPrefix("/admin/", AdminfileServer))
}

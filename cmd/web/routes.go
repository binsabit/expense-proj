package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", app.MainPage).Methods("GET")
	r.HandleFunc("/expenses", app.GetExpenses).Methods("GET")
	r.HandleFunc("/newexpense", app.PostExpense).Methods("POST")
	r.HandleFunc("/categories", app.GetCategories).Methods("GET")
	r.HandleFunc("/newcategory", app.PostCategory).Methods("POST")

	//serving static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r
}

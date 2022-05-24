package main

import (
	"encoding/json"
	"fmt"
	models "money-tracker/model"
	"net/http"
	"path"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

type Test struct {
	cats     []string `json:"cats,omitempty"`
	dateFrom string   `json:"dataForm,omitempty"`
	dateTo   string   `json:"dataTo,omitempty"`
}

func (app *application) MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("main page")

	fp := path.Join("static", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, ""); err != nil {
		app.serverError(w, err)
		return
	}
}
func (app *application) GetExpenses(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("You are on GET Expense\n")

	results, err := app.expenses.FindAllExpenses()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// app.infoLog.Println("Get all Expenses:")
	// spew.Dump(results)
	json.NewEncoder(w).Encode(results)

}
func (app *application) PostExpense(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("You are in Post Expense\n")
	//in order to hander the post request from the form you shoulf par it first!!!
	//Parsing
	if err := r.ParseForm(); err != nil {
		app.clientError(w, http.StatusNoContent)
		return
	}

	//expense params
	name := r.FormValue("expenseName")
	category := r.FormValue("expenseCat")
	amount := r.FormValue("expenseAmount")

	id, err := app.expenses.InsertExpense(name, category, amount)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.infoLog.Printf("ID of the new Expense is %v", id)
	http.Redirect(w, r, "/", http.StatusFound)

}

func (app *application) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("You are in Delete Expense")

	params := mux.Vars(r)
	expenseId := params["id"]

	err := app.expenses.DeleteExpenseByID(expenseId)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.infoLog.Println(expenseId)
}

func (app *application) GetExpensesFiltered(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("You are in Get filtered category")
	var t Test
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		app.serverError(w, err)
		return
	}
	spew.Dump(t)

}
func (app *application) PostCategory(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("You are in Post Category")
	if err := r.ParseForm(); err != nil {
		app.serverError(w, err)
		return
	}

	name := r.FormValue("catName")

	id, err := app.expenses.InsertCategory(name)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("ID of the new category is %v \n ", id)

	http.Redirect(w, r, "/", http.StatusFound)

}

func (app *application) GetCategories(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("You are in GET Category")

	var results []*models.Category
	//finding multiple elements return the a cursor
	results, err := app.expenses.FindAllCategories()
	if err != nil {
		app.serverError(w, err)
	}

	// app.infoLog.Println("Get all Categories:")
	// spew.Dump(results)

	json.NewEncoder(w).Encode(results)

}

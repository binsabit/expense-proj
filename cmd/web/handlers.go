package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"money-tracker/model"
	"net/http"
	"path"
	"text/template"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (app *application) MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("main page")

	fp := path.Join("static", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// list := client.Database("expenses-db").Collection("lists")

}
func (app *application) GetExpenses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get expense page")
	collection := client.Database("expenseDB").Collection("expenses")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findOptions := options.Find()
	// findOptions.SetLimit(5)

	var results []*model.Expense
	//finding multiple elements return the a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		fmt.Println("Error Finding Categories")
		return
	}

	//iterate thorgh the curser and add them to array.

	if err = cur.All(ctx, &results); err != nil {
		log.Print("Error GET items 2")
		return
	}
	fmt.Println(results)

	json.NewEncoder(w).Encode(results)

}
func (app *application) PostExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post expense page")
	//in order to hander the post request from the form you shoulf par it first!!!
	//Parsing
	if err := r.ParseForm(); err != nil {
		fmt.Println("Parsing Error Expense", err)
		return
	}

	//make new struct
	newExpense := model.Expense{
		ExpenseName: r.FormValue("expenseName"),
		ExpenseCat:  r.FormValue("expenseCat"),
	}

	//select a collection
	collection := client.Database("expenseDB").Collection("expenses")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//insert new expense to the database
	res, err := collection.InsertOne(ctx, newExpense)
	if err != nil {
		fmt.Print("Inserting Error")
		return
	}

	id := res.InsertedID

	fmt.Println(newExpense.ExpenseName, newExpense.ExpenseCat, id)
	http.Redirect(w, r, "/", http.StatusFound)

}

func (app *application) PostCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post category page")
	if err := r.ParseForm(); err != nil {
		fmt.Println("Parsing Error Expense", err)
		return
	}

	newCategory := model.Category{
		CatName: r.FormValue("catName"),
	}

	//select a collection
	collection := client.Database("expenseDB").Collection("expenses")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//insert new expense to the database
	res, err := collection.InsertOne(ctx, newCategory)
	if err != nil {
		fmt.Print("Inserting Error")
		return
	}

	id := res.InsertedID

	fmt.Println(newCategory.CatName, id)
	http.Redirect(w, r, "/", http.StatusFound)

}

func (app *application) GetCategories(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("expenseDB").Collection("categories")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findOptions := options.Find()

	var results []*model.Category
	//finding multiple elements return the a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		fmt.Println("Error Finding Categories")
		return
	}

	//iterate thorgh the curser and add them to array.

	if err = cur.All(ctx, &results); err != nil {
		log.Print("Error GET items 2")
		return
	}
	fmt.Println(results)

	json.NewEncoder(w).Encode(results)

}

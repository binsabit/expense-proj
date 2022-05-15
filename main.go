package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Expense struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	ExpenseName string `json:"expenseName,omitempty" bson:"expenseName, omitempty"`
	ExpenseCat  string `json:"expenseCat,omitempty" bson:"expenseCat, omitempty"`
}

type Category struct {
	ID      string `json:"_id,omitempty" bson:"_id,omitempty"`
	CatName string `json:"catName,omitempty" bson:"catName,omitempty"`
}

func MainPage(w http.ResponseWriter, r *http.Request) {
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
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get expense page")
	collection := client.Database("expenseDB").Collection("expenses")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findOptions := options.Find()
	// findOptions.SetLimit(5)

	var results []Expense
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
func PostExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post expense page")
	//in order to hander the post request from the form you shoulf par it first!!!
	//Parsing
	if err := r.ParseForm(); err != nil {
		fmt.Println("Parsing Error Expense", err)
		return
	}

	//make new struct
	newExpense := Expense{
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

func PostCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post category page")
	if err := r.ParseForm(); err != nil {
		fmt.Println("Parsing Error Expense", err)
		return
	}

	newCategory := Category{
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

func GetCategories(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("expenseDB").Collection("categories")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findOptions := options.Find()

	var results []Category
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

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", MainPage).Methods("GET")
	r.HandleFunc("/expenses", GetExpenses).Methods("GET")
	r.HandleFunc("/newexpense", PostExpense).Methods("POST")
	r.HandleFunc("/categories", GetCategories).Methods("GET")
	r.HandleFunc("/newcategory", PostCategory).Methods("POST")

	//serving static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	srv := &http.Server{
		Addr:    *addr,
		Handler: r,
	}

	err := srv.ListenAndServe()
	log.Fatal(err)

}

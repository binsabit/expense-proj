package main

import (
	"context"
	"flag"
	"log"
	"money-tracker/model/mongoDB"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	expenses *mongoDB.ModelDB
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	flag.Parse()

	db, err := InitDB()
	if err != nil {
		log.Panic("DB error", err)
	}

	app := application{
		infoLog:  infoLog,
		errorLog: errorLog,
		expenses: &mongoDB.ModelDB{DB: db},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

func InitDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Panic("Db connection Error", err)
		return nil, err
	}
	return client.Database("expenseDB"), nil

}

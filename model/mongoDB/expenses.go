package mongoDB

import (
	"context"
	"fmt"
	models "money-tracker/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModelDB struct {
	DB *mongo.Database
}

func (m *ModelDB) FindAllExpenses() ([]*models.Expense, error) {
	collection := m.DB.Collection("expenses")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findOptions := options.Find()
	//set a limit
	// findOptions.SetLimit(5)

	var results []*models.Expense

	//finding multiple elements return the a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	//iterate thorgh the curser and add them to array.

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *ModelDB) InsertExpense(name, category, amount string) (primitive.ObjectID, error) {
	collection := m.DB.Collection("expenses")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	currentTime := time.Now().Format(time.UnixDate)
	newExpense := models.Expense{
		ExpenseName:   name,
		ExpenseCat:    category,
		ExpenseAmount: amount,
		ExpenseDate:   currentTime,
	}

	result, err := collection.InsertOne(ctx, newExpense)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)

	return id, nil
}

func (m *ModelDB) DeleteExpenseByID(id string) error {
	collection := m.DB.Collection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := collection.DeleteOne(ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return err
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return nil
}

// func (m *ModelDB) FindExpensesByCategory(categoryName []string) {
// 	collection := m.DB.Collection("expenses")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// }
func (m *ModelDB) InsertCategory(name string) (primitive.ObjectID, error) {
	collection := m.DB.Collection("categories")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newCategory := models.Category{
		CatName: name,
	}

	result, err := collection.InsertOne(ctx, newCategory)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)

	return id, nil
}

func (m *ModelDB) FindAllCategories() ([]*models.Category, error) {
	collection := m.DB.Collection("categories")
	//set a context(time to finish the go routine)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findOptions := options.Find()
	//set a limit
	// findOptions.SetLimit(5)

	var results []*models.Category

	//finding multiple elements return the a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	//iterate thorgh the curser and add them to array.

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

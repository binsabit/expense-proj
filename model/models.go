package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ExpenseName   string             `json:"expenseName,omitempty" bson:"expenseName, omitempty"`
	ExpenseAmount string             `json:"expenseAmount, omitempty" bson:"expenseAmount, omitempty"`
	ExpenseCat    string             `json:"expenseCat,omitempty" bson:"expenseCat, omitempty"`
}

type Category struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CatName string             `json:"catName,omitempty" bson:"catName,omitempty"`
}

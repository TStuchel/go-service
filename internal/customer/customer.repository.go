package customer

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// ----------------------------------------------------- INTERFACE -----------------------------------------------------

//go:generate counterfeiter . Service
type Repository interface {
	GetCustomer(string) (*Customer, error)
	CreateCustomer(customer Customer) (*Customer, error)
}

// -------------------------------------------------- IMPLEMENTATION ---------------------------------------------------

type repositoryImpl struct {
	db *mongo.Database
}

// --------------------------------------------------- CONSTRUCTORS ----------------------------------------------------

// NewCustomerRepository creates and returns a new Repository.
func NewCustomerRepository(db *mongo.Database) Repository {
	return &repositoryImpl{
		db: db,
	}
}

// ------------------------------------------------------ METHODS ------------------------------------------------------

// GetCustomer returns a Customer given a customer ID. If it simply doesn't exist, then both the customer and error will
// be nil.
func (impl *repositoryImpl) GetCustomer(customerId string) (*Customer, error) {

	// Build the query
	objectId, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return nil, err
	}

	// Find the customer
	customer := &Customer{}
	filter := bson.D{{"_id", objectId}}
	err = impl.db.Collection("customers").FindOne(context.TODO(), filter).Decode(customer)

	// Not found
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	// Everything else (including errors)
	customer.ID = customerId
	return customer, err
}

// CreateCustomer persists a new customer with the given data.
func (impl *repositoryImpl) CreateCustomer(customer Customer) (*Customer, error) {

	// Insert
	result, err := impl.db.Collection("customers").InsertOne(context.TODO(), customer)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// Set the ID
	customer.ID = result.InsertedID.(primitive.ObjectID).Hex()

	// Done
	return &customer, nil
}

package repository

import (
	"context"
	"log"
	"time"

	"github.com/vtesin/StonkTendency/config"
	"github.com/vtesin/StonkTendency/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TickerMongo mongo client
type TickerMongo struct {
	db *mongo.Client
}

// NewTickerMongo NoSql repositry of Tickers
func NewTickerMongo(db *mongo.Client) *TickerMongo {
	return &TickerMongo{
		db: db,
	}
}

// Create Create a new ticker
func (r *TickerMongo) Create(t *entity.Ticker) error {
	collection := r.db.Database(config.DbName).Collection(config.DbCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := bson.Marshal(t)
	_, err = collection.InsertOne(ctx, val)

	return err
}

// Update Update a ticker
func (r *TickerMongo) Update(t *entity.Ticker) error {
	collection := r.db.Database(config.DbName).Collection(config.DbCollection)

	// find the document for which the _id field matches id
	// specify the Sort option to sort the documents by age
	// the first document in the sorted order will be returned
	opts := options.Update().SetUpsert(true)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{"Symbol", t.Symbol}}
	update := bson.D{{"$set", bson.D{{"Symbol", t.Symbol}, {"Sentiment", t.Sentiment}, {"Timestamp", t.Timestamp}}}}

	result, err := collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	if result.MatchedCount != 0 {
		log.Printf("matched and replaced an existing document")
		return nil
	}
	if result.UpsertedCount != 0 {
		log.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}

	return nil
}

// Delete Delete a ticker
func (r *TickerMongo) Delete(symbol string) error {
	return nil
}

// Get a disctinct ticker
func (r *TickerMongo) Get(symbol string) (*entity.Ticker, error) {
	collection := r.db.Database(config.DbName).Collection(config.DbCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{"Symbol", symbol}}

	res := collection.FindOne(ctx, filter)

	if res.Err() != nil {
		log.Printf("%v", res.Err())
		return nil, res.Err()
	}

	var t entity.Ticker
	err := res.Decode(&t)

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return &t, nil
}

// Search Find a particular ticker satisfying a query
func (r *TickerMongo) Search(query string) ([]*entity.Ticker, error) {
	return nil, nil
}

// List Lists all tickers in system
func (r *TickerMongo) List() ([]*entity.Ticker, error) {
	return nil, nil
}

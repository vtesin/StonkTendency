package repository

import (
	"context"
	"time"

	"github.com/vtesin/StonkTendency/config"
	"github.com/vtesin/StonkTendency/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	return nil
}

// Delete Delete a ticker
func (r *TickerMongo) Delete(symbol string) error {
	return nil
}

// Get a disctinct ticker
func (r *TickerMongo) Get(symbol string) (*entity.Ticker, error) {
	return nil, nil
}

// Search Find a particular ticker satisfying a query
func (r *TickerMongo) Search(query string) ([]*entity.Ticker, error) {
	return nil, nil
}

// List Lists all tickers in system
func (r *TickerMongo) List() ([]*entity.Ticker, error) {
	return nil, nil
}

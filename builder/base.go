package builder

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QueryBuilder struct {
	Collection string
	Fields     []string
	Group      bson.M
	Sort       bson.M
	HavingCond bson.M
	LimitVal   int64
	OffsetVal  int64 // Tambahkan OffsetVal untuk OFFSET
	Pipeline   []bson.D
}

// NewQueryBuilder initializes a new QueryBuilder.
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		Fields:    []string{},
		Pipeline:  []bson.D{},
		LimitVal:  0,
		OffsetVal: 0,
	}
}

// Limit sets the maximum number of documents to return.
func (qb *QueryBuilder) Limit(limit int64) *QueryBuilder {
	qb.LimitVal = limit
	return qb
}

// Offset sets the number of documents to skip.
func (qb *QueryBuilder) Offset(offset int64) *QueryBuilder {
	qb.OffsetVal = offset
	return qb
}

// Execute executes the query pipeline.
func (qb *QueryBuilder) Execute(db *mongo.Database) ([]map[string]interface{}, error) {
	if qb.Collection == "" {
		return nil, errors.New("collection is not specified")
	}

	collection := db.Collection(qb.Collection)

	// Build the pipeline
	if qb.OffsetVal > 0 {
		qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$skip", Value: qb.OffsetVal}})
	}
	if qb.LimitVal > 0 {
		qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$limit", Value: qb.LimitVal}})
	}

	// Execute the pipeline
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Aggregate(ctx, qb.Pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	for cursor.Next(ctx) {
		var result map[string]interface{}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

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

// From specifies the collection to query.
func (qb *QueryBuilder) From(collection string) *QueryBuilder {
	qb.Collection = collection
	return qb
}

// Join adds a $lookup stage to the aggregation pipeline for joining collections.
func (qb *QueryBuilder) Join(localField, fromCollection, foreignField, as string) *QueryBuilder {
	qb.Pipeline = append(qb.Pipeline, bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         fromCollection,
			"localField":   localField,
			"foreignField": foreignField,
			"as":           as,
		}},
	})
	return qb
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

// Select specifies the fields to include in the query result.
func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
	qb.Fields = append(qb.Fields, fields...)
	qb.Pipeline = append(qb.Pipeline, bson.D{
		{Key: "$project", Value: qb.buildProjection()},
	})
	return qb
}

// buildProjection builds a MongoDB $project stage from the Fields.
func (qb *QueryBuilder) buildProjection() bson.M {
	projection := bson.M{}
	for _, field := range qb.Fields {
		projection[field] = 1
	}
	return projection
}

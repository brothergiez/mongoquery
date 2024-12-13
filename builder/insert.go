package builder

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// InsertBuilder helps in inserting documents into a MongoDB collection.
type InsertBuilder struct {
	Collection string
	Fields     []string
	ValuesList [][]interface{}
}

// NewInsertBuilder initializes a new InsertBuilder for a specific collection.
func NewInsertBuilder() *InsertBuilder {
	return &InsertBuilder{}
}

// InsertInto specifies the collection and the fields for the insert.
func (ib *InsertBuilder) InsertInto(collection string, fields []string) *InsertBuilder {
	ib.Collection = collection
	ib.Fields = fields
	return ib
}

// Values adds a row of values corresponding to the specified fields.
func (ib *InsertBuilder) Values(values []interface{}) *InsertBuilder {
	if len(values) != len(ib.Fields) {
		panic("number of values must match the number of fields")
	}
	ib.ValuesList = append(ib.ValuesList, values)
	return ib
}

// Execute performs the insert operation.
func (ib *InsertBuilder) Execute(db *mongo.Database) (interface{}, error) {
	if ib.Collection == "" {
		return nil, errors.New("collection name is not specified")
	}

	collection := db.Collection(ib.Collection)
	documents := []interface{}{}

	// Convert rows into MongoDB-compatible documents
	for _, row := range ib.ValuesList {
		document := map[string]interface{}{}
		for i, field := range ib.Fields {
			document[field] = row[i]
		}
		documents = append(documents, document)
	}

	// Perform the insert
	if len(documents) == 1 {
		res, err := collection.InsertOne(context.TODO(), documents[0])
		if err != nil {
			return nil, fmt.Errorf("failed to insert document: %v", err)
		}
		return res.InsertedID, nil
	} else if len(documents) > 1 {
		res, err := collection.InsertMany(context.TODO(), documents)
		if err != nil {
			return nil, fmt.Errorf("failed to insert documents: %v", err)
		}
		return res.InsertedIDs, nil
	}

	return nil, errors.New("no documents to insert")
}

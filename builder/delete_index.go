package builder

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteIndexBuilder helps in deleting indexes from a MongoDB collection.
type DeleteIndexBuilder struct {
	Collection string
	Indexes    []string
}

// NewDeleteIndexBuilder initializes a new DeleteIndexBuilder for a specific collection.
func NewDeleteIndexBuilder(collection string) *DeleteIndexBuilder {
	return &DeleteIndexBuilder{
		Collection: collection,
		Indexes:    []string{},
	}
}

// Index adds an index name to the list of indexes to delete.
func (dib *DeleteIndexBuilder) Index(name string) *DeleteIndexBuilder {
	dib.Indexes = append(dib.Indexes, name)
	return dib
}

// Execute deletes all specified indexes from the collection.
func (dib *DeleteIndexBuilder) Execute(db *mongo.Database) error {
	if dib.Collection == "" {
		return fmt.Errorf("collection name is not specified")
	}

	collection := db.Collection(dib.Collection)
	for _, index := range dib.Indexes {
		_, err := collection.Indexes().DropOne(context.TODO(), index)
		if err != nil {
			return fmt.Errorf("failed to delete index %s: %v", index, err)
		}
	}

	return nil
}

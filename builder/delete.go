package builder

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteBuilder helps in deleting documents from a MongoDB collection.
type DeleteBuilder struct {
	Collection string
	Filter     map[string]interface{}
	Multi      bool // If true, deletes multiple documents
}

// NewDeleteBuilder initializes a new DeleteBuilder for a specific collection.
func NewDeleteBuilder(collection string) *DeleteBuilder {
	return &DeleteBuilder{
		Collection: collection,
		Filter:     map[string]interface{}{},
		Multi:      false,
	}
}

// Where specifies the filter condition for the delete operation.
func (db *DeleteBuilder) Where(condition string) *DeleteBuilder {
	qb := QueryBuilder{}
	db.Filter = qb.parseConditions(condition) // Reuse parseConditions from QueryBuilder
	return db
}

// SetMulti enables or disables deleting multiple documents.
func (db *DeleteBuilder) SetMulti(multi bool) *DeleteBuilder {
	db.Multi = multi
	return db
}

// Execute performs the delete operation.
func (db *DeleteBuilder) Execute(dbInstance *mongo.Database) (int64, error) {
	if db.Collection == "" {
		return 0, errors.New("collection name is not specified")
	}

	collection := dbInstance.Collection(db.Collection)

	// DeleteOne or DeleteMany
	var result *mongo.DeleteResult
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if db.Multi {
		result, err = collection.DeleteMany(ctx, db.Filter)
	} else {
		result, err = collection.DeleteOne(ctx, db.Filter)
	}

	if err != nil {
		return 0, fmt.Errorf("failed to delete documents: %v", err)
	}

	return result.DeletedCount, nil
}

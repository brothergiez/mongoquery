package builder

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateBuilder helps in updating documents in a MongoDB collection.
type UpdateBuilder struct {
	Collection string
	UpdateData bson.M
	Filter     bson.M
	Multi      bool // If true, updates multiple documents
}

// NewUpdateBuilder initializes a new UpdateBuilder for a specific collection.
func NewUpdateBuilder(collection string) *UpdateBuilder {
	return &UpdateBuilder{
		Collection: collection,
		UpdateData: bson.M{},
		Filter:     bson.M{},
		Multi:      false,
	}
}

// Set specifies the fields and their new values to be updated.
func (ub *UpdateBuilder) Set(data map[string]interface{}) *UpdateBuilder {
	ub.UpdateData["$set"] = data
	return ub
}

// Where specifies the filter condition for the update.
func (ub *UpdateBuilder) Where(condition string) *UpdateBuilder {
	qb := QueryBuilder{}
	ub.Filter = qb.parseConditions(condition) // Reuse parseConditions from QueryBuilder
	return ub
}

// SetMulti enables or disables updating multiple documents.
func (ub *UpdateBuilder) SetMulti(multi bool) *UpdateBuilder {
	ub.Multi = multi
	return ub
}

// Execute performs the update operation.
func (ub *UpdateBuilder) Execute(db *mongo.Database) (int64, error) {
	if ub.Collection == "" {
		return 0, errors.New("collection name is not specified")
	}

	collection := db.Collection(ub.Collection)

	// UpdateOne or UpdateMany
	var result *mongo.UpdateResult
	var err error
	if ub.Multi {
		result, err = collection.UpdateMany(context.TODO(), ub.Filter, ub.UpdateData)
	} else {
		result, err = collection.UpdateOne(context.TODO(), ub.Filter, ub.UpdateData)
	}

	if err != nil {
		return 0, fmt.Errorf("failed to update documents: %v", err)
	}

	return result.ModifiedCount, nil
}

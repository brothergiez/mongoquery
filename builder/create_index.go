package builder

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateIndexBuilder helps in creating indexes for a MongoDB collection.
type CreateIndexBuilder struct {
	Collection string
	Indexes    []mongo.IndexModel
}

// NewCreateIndexBuilder initializes a new CreateIndexBuilder for a specific collection.
func NewCreateIndexBuilder(collection string) *CreateIndexBuilder {
	return &CreateIndexBuilder{
		Collection: collection,
		Indexes:    []mongo.IndexModel{},
	}
}

// Index adds a new index with the specified name and fields.
func (ib *CreateIndexBuilder) Index(name string, fields string) *CreateIndexBuilder {
	keys := bson.D{}

	// Parse fields like "status ASC, amount DESC"
	fieldParts := strings.Split(fields, ",")
	for _, part := range fieldParts {
		part = strings.TrimSpace(part)
		field := strings.Fields(part)
		if len(field) != 2 {
			continue
		}

		direction := 1
		if strings.ToUpper(field[1]) == "DESC" {
			direction = -1
		}

		keys = append(keys, bson.E{Key: field[0], Value: direction})
	}

	ib.Indexes = append(ib.Indexes, mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetName(name),
	})
	return ib
}

// Execute creates all specified indexes on the collection.
func (ib *CreateIndexBuilder) Execute(db *mongo.Database) error {
	if ib.Collection == "" {
		return fmt.Errorf("collection name is not specified")
	}

	collection := db.Collection(ib.Collection)
	for _, index := range ib.Indexes {
		_, err := collection.Indexes().CreateOne(context.TODO(), index)
		if err != nil {
			return fmt.Errorf("failed to create index %v: %v", index.Options.Name, err)
		}
	}

	return nil
}

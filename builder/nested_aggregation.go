package builder

import (
	"go.mongodb.org/mongo-driver/bson"
)

// NestedGroupBy adds a nested $group stage to the pipeline.
func (qb *QueryBuilder) NestedGroupBy(field string, aggregations ...string) *QueryBuilder {
	nestedGroup := bson.M{"_id": "$" + field}

	for _, agg := range aggregations {
		alias := qb.parseAlias(agg) // Get the alias
		aggregation, err := qb.parseAggregation(agg)
		if err != nil {
			continue // Skip unsupported aggregations
		}
		nestedGroup[alias] = aggregation
	}

	qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$group", Value: nestedGroup}})
	return qb
}

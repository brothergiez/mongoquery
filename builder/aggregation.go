package builder

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// Match adds a $match stage to the pipeline (supports expressions).
func (qb *QueryBuilder) Match(condition string) *QueryBuilder {
	filter, err := qb.parseExpression(condition)
	if err != nil {
		filter = qb.parseConditions(condition) // Fallback to simple conditions
	}
	qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$match", Value: filter}})
	return qb
}

// GroupBy adds a $group stage to the pipeline.
func (qb *QueryBuilder) GroupBy(field string) *QueryBuilder {
	qb.Group = bson.M{"_id": "$" + field}
	qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$group", Value: qb.Group}})
	return qb
}

// OrderBy adds a $sort stage to the pipeline.
func (qb *QueryBuilder) OrderBy(order string) *QueryBuilder {
	sort := bson.M{}
	parts := strings.Split(order, " ")
	if len(parts) == 2 {
		direction := 1
		if strings.ToUpper(parts[1]) == "DESC" {
			direction = -1
		}
		sort[parts[0]] = direction
	}
	qb.Sort = sort
	qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$sort", Value: qb.Sort}})
	return qb
}

// AggregationLimit adds a $limit stage to the pipeline.
func (qb *QueryBuilder) AggregationLimit(n int64) *QueryBuilder {
	if n > 0 {
		qb.LimitVal = n
		qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$limit", Value: n}})
	}
	return qb
}

// AggregationOffset adds a $offset stage to the pipeline.
func (qb *QueryBuilder) AggregationOffset(n int64) *QueryBuilder {
	if n > 0 {
		qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$skip", Value: n}})
	}
	return qb
}

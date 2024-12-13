package builder

import "go.mongodb.org/mongo-driver/bson"

// Having adds a $match stage after $group to filter aggregated results (supports expressions).
func (qb *QueryBuilder) Having(condition string) *QueryBuilder {
	filter, err := qb.parseExpression(condition)
	if err != nil {
		filter = qb.parseConditions(condition) // Fallback to simple conditions
	}
	qb.Pipeline = append(qb.Pipeline, bson.D{{Key: "$match", Value: filter}})
	return qb
}

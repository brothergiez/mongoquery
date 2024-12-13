package builder

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// parseAggregation parses aggregation functions like "SUM(amount)".
func (qb *QueryBuilder) parseAggregation(field string) (bson.M, error) {
	field = strings.TrimSpace(field)

	// Handle SUM
	if strings.HasPrefix(strings.ToUpper(field), "SUM(") {
		innerField := strings.TrimSuffix(strings.TrimPrefix(field, "SUM("), ")")
		return bson.M{"$sum": "$" + strings.TrimSpace(innerField)}, nil
	}

	// Handle COUNT
	if strings.HasPrefix(strings.ToUpper(field), "COUNT(") {
		return bson.M{"$sum": 1}, nil
	}

	// Handle other aggregation functions (if needed)
	// Example: MAX, MIN, etc.
	return nil, errors.New("unsupported aggregation function")
}

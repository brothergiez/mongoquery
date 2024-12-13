package builder

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// parseConditions parses multiple conditions like "amount > 1000 AND status = 'active'".
func (qb *QueryBuilder) parseConditions(conditions string) bson.M {
	conditions = strings.TrimSpace(conditions)

	// Split by AND/OR
	if strings.Contains(strings.ToUpper(conditions), " AND ") {
		parts := strings.Split(conditions, " AND ")
		andConditions := []bson.M{}
		for _, part := range parts {
			andConditions = append(andConditions, qb.parseCondition(strings.TrimSpace(part)))
		}
		return bson.M{"$and": andConditions}
	}

	if strings.Contains(strings.ToUpper(conditions), " OR ") {
		parts := strings.Split(conditions, " OR ")
		orConditions := []bson.M{}
		for _, part := range parts {
			orConditions = append(orConditions, qb.parseCondition(strings.TrimSpace(part)))
		}
		return bson.M{"$or": orConditions}
	}

	// Single condition
	return qb.parseCondition(conditions)
}

// parseCondition parses a single condition like "amount > 1000".
func (qb *QueryBuilder) parseCondition(condition string) bson.M {
	parts := strings.Fields(condition)
	if len(parts) != 3 {
		return bson.M{}
	}

	field, operator, value := parts[0], parts[1], strings.Trim(parts[2], "'")
	mongoOperator := mapOperatorToMongo(operator)

	return bson.M{field: bson.M{mongoOperator: qb.convertValue(value)}}
}

// convertValue converts a value string to the appropriate type (e.g., int, float, string).
func (qb *QueryBuilder) convertValue(value string) interface{} {
	// Try to convert to an integer
	if num, err := strconv.Atoi(value); err == nil {
		return num
	}

	// Try to convert to a float
	if num, err := strconv.ParseFloat(value, 64); err == nil {
		return num
	}

	// Fallback to string
	return value
}

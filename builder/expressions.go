package builder

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// parseExpression parses expressions like "SUM(amount) / COUNT(*) > 1000".
func (qb *QueryBuilder) parseExpression(expression string) (bson.M, error) {
	expression = strings.TrimSpace(expression)

	re := regexp.MustCompile(`([\w\(\)\*]+)\s*([+\-*/><=]+)\s*([\w\(\)\*]+)`)
	matches := re.FindStringSubmatch(expression)
	if len(matches) < 4 {
		return nil, errors.New("invalid expression format")
	}

	operand1, operator, operand2 := matches[1], matches[2], matches[3]
	mongoOperator := mapOperatorToMongo(operator)
	if mongoOperator == "" {
		return nil, errors.New("unsupported operator in expression")
	}

	return bson.M{
		"$expr": bson.M{
			mongoOperator: []interface{}{
				qb.parseFieldOrValue(operand1),
				qb.parseFieldOrValue(operand2),
			},
		},
	}, nil
}

// parseFieldOrValue parses a field (e.g., SUM(amount)) or a literal value.
func (qb *QueryBuilder) parseFieldOrValue(input string) interface{} {
	input = strings.TrimSpace(input)

	if num, err := strconv.ParseFloat(input, 64); err == nil {
		return num
	}

	if strings.HasPrefix(strings.ToUpper(input), "SUM(") {
		field := strings.TrimSuffix(strings.TrimPrefix(input, "SUM("), ")")
		return bson.M{"$sum": "$" + field}
	}

	if strings.HasPrefix(strings.ToUpper(input), "COUNT(") {
		return bson.M{"$sum": 1}
	}

	return "$" + input
}

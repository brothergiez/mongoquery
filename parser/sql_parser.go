package parser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/brothergiez/mongoquery/builder"
)

// SQLParser is a utility to parse SQL-like syntax into MongoDB query components.
type SQLParser struct {
	query string
}

// NewSQLParser creates a new instance of SQLParser.
func NewSQLParser(query string) *SQLParser {
	return &SQLParser{query: query}
}

// ParseSQL parses an SQL-like query into a QueryBuilder.
func (sp *SQLParser) ParseSQL() (*builder.QueryBuilder, error) {
	sp.query = strings.TrimSpace(sp.query)
	qb := builder.NewQueryBuilder()

	// Parse SELECT
	fields, rest := sp.extractFields(strings.Split(sp.query, " "))
	qb.Fields = fields

	// Parse FROM
	collection, rest := sp.extractCollection(rest)
	qb.Collection = collection

	// Parse WHERE
	if strings.Contains(strings.ToUpper(rest), "WHERE") {
		whereClause, remaining := sp.extractClause("WHERE", rest)
		qb.Match(strings.TrimSpace(whereClause))
		rest = remaining
	}

	// Parse GROUP BY
	if strings.Contains(strings.ToUpper(rest), "GROUP BY") {
		groupByClause, remaining := sp.extractClause("GROUP BY", rest)
		qb.GroupBy(strings.TrimSpace(groupByClause))
		rest = remaining
	}

	// Parse HAVING
	if strings.Contains(strings.ToUpper(rest), "HAVING") {
		havingClause, remaining := sp.extractClause("HAVING", rest)
		qb.Having(strings.TrimSpace(havingClause))
		rest = remaining
	}

	// Parse ORDER BY
	if strings.Contains(strings.ToUpper(rest), "ORDER BY") {
		orderByClause, remaining := sp.extractClause("ORDER BY", rest)
		qb.OrderBy(strings.TrimSpace(orderByClause))
		rest = remaining
	}

	// Parse LIMIT
	if strings.Contains(strings.ToUpper(rest), "LIMIT") {
		limitClause, _ := sp.extractClause("LIMIT", rest)
		limit, err := sp.parseLimit(strings.TrimSpace(limitClause))
		if err != nil {
			return nil, err
		}
		qb.Limit(limit)
	}

	return qb, nil
}

// extractFields extracts fields from the SELECT clause.
func (sp *SQLParser) extractFields(parts []string) ([]string, string) {
	if strings.ToUpper(parts[0]) != "SELECT" {
		return nil, strings.Join(parts, " ")
	}

	fields := []string{}
	rest := strings.Join(parts[1:], " ")

	// Handle field extraction until the FROM keyword
	fromIndex := strings.Index(strings.ToUpper(rest), " FROM ")
	if fromIndex != -1 {
		fields = strings.Split(rest[:fromIndex], ",")
		rest = rest[fromIndex+len(" FROM "):]
	}

	// Trim whitespace and return
	for i, field := range fields {
		fields[i] = strings.TrimSpace(field)
	}
	return fields, rest
}

// extractCollection extracts the collection name from the FROM clause.
func (sp *SQLParser) extractCollection(query string) (string, string) {
	parts := strings.SplitN(query, " ", 2)
	return parts[0], parts[1]
}

// extractClause extracts a clause and the remaining query after it.
func (sp *SQLParser) extractClause(keyword string, query string) (string, string) {
	keywordIndex := strings.Index(strings.ToUpper(query), keyword)
	if keywordIndex == -1 {
		return "", query
	}

	remaining := query[keywordIndex+len(keyword):]
	nextKeywordIndex := sp.findNextKeyword(remaining)
	if nextKeywordIndex == -1 {
		return strings.TrimSpace(remaining), ""
	}

	return strings.TrimSpace(remaining[:nextKeywordIndex]), strings.TrimSpace(remaining[nextKeywordIndex:])
}

// findNextKeyword finds the position of the next SQL keyword.
func (sp *SQLParser) findNextKeyword(query string) int {
	keywords := []string{"WHERE", "GROUP BY", "HAVING", "ORDER BY", "LIMIT"}
	for _, keyword := range keywords {
		keywordIndex := strings.Index(strings.ToUpper(query), keyword)
		if keywordIndex != -1 {
			return keywordIndex
		}
	}
	return -1
}

// parseLimit parses the LIMIT clause into an integer.
func (sp *SQLParser) parseLimit(limit string) (int64, error) {
	parsedLimit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return 0, errors.New("invalid LIMIT value")
	}
	return parsedLimit, nil
}

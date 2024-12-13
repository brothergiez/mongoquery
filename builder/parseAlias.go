package builder

import "strings"

// parseAlias extracts the alias from a field like "SUM(amount) AS totalAmount".
func (qb *QueryBuilder) parseAlias(field string) string {
	// Check if there's an alias (contains "AS")
	if strings.Contains(strings.ToUpper(field), " AS ") {
		parts := strings.SplitN(field, " AS ", 2)
		return strings.TrimSpace(parts[1])
	}
	// Default to the field itself if no alias is provided
	return field
}

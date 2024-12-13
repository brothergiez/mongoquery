package builder

// mapOperatorToMongo maps SQL-like operators to MongoDB operators.
func mapOperatorToMongo(operator string) string {
	switch operator {
	case "=":
		return "$eq"
	case ">":
		return "$gt"
	case "<":
		return "$lt"
	case ">=":
		return "$gte"
	case "<=":
		return "$lte"
	case "!=":
		return "$ne"
	case "+":
		return "$add"
	case "-":
		return "$subtract"
	case "*":
		return "$multiply"
	case "/":
		return "$divide"
	default:
		return ""
	}
}

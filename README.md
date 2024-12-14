# MongoQuery Documentation

MongoQuery is a powerful and intuitive query builder for MongoDB, designed to provide a SQL-like syntax for building and executing queries. It simplifies complex MongoDB operations such as aggregation, filtering, updates, and indexing, allowing developers to write queries in a format familiar to those accustomed to SQL.

With MongoQuery, you can seamlessly handle dynamic conditions, nested aggregations, pagination (via `LIMIT` and `OFFSET`), and advanced filtering using mathematical and logical expressions. This library is perfect for developers who prefer the simplicity of SQL while leveraging the flexibility and power of MongoDB.

---

## Features

MongoQuery provides two key features:

### 1. Query Builder
A flexible and user-friendly interface for constructing MongoDB queries programmatically. The Query Builder allows developers to chain methods to define queries, making it easy to handle complex operations such as:
- Filtering (`Where` conditions)
- Grouping (`GroupBy` and `NestedGroupBy`)
- Sorting (`OrderBy`)
- Pagination (`Limit` and `Offset`)
- Aggregation Pipelines
- CRUD Operations (Insert, Update, Delete)

### 2. SQL Parser
The SQL Parser enables developers to write SQL-like queries that are parsed into MongoDB operations. This feature is ideal for those familiar with SQL syntax and simplifies the transition to MongoDB. The parser supports:
- SELECT, INSERT, UPDATE, DELETE queries
- WHERE conditions with logical operators (`AND`, `OR`)
- Grouping and aggregation with `GROUP BY` and `HAVING`
- Sorting (`ORDER BY`)
- Pagination (`LIMIT`, `OFFSET`)

---

## Installation

To install MongoQuery, you can use `go get` to fetch it directly into your Go project:

```bash
go get github.com/brothergiez/mongoquery
```

Once installed, import the library into your Go files:

```go
import (
    "github.com/brothergiez/mongoquery/builder"
    "github.com/brothergiez/mongoquery/sqlparser"
    "github.com/brothergiez/mongoquery/client"
)
```

---

## Basic Implementation

### Connecting to MongoDB

Before running queries, you need to establish a connection to your MongoDB instance:

```go
package main

import (
    "fmt"
    "log"

    "github.com/brothergiez/mongoquery/builder"
    "github.com/brothergiez/mongoquery/sqlparser"
    "github.com/brothergiez/mongoquery/client"
)

func main() {
    // Initialize MongoDB client
    mdb, err := client.New("mongodb://localhost:27017", "test_db")
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    fmt.Println("MongoDB connection established.")
    
    // Example With QueryBuilder
	fmt.Println("Example: Basic SELECT with QueryBuilder")
	basicQuery := builder.NewQueryBuilder().
		Select("field1", "field2").
		From("orders").
		Where("status = 'active'").
		OrderBy("field1 ASC").
		Limit(10)

	results, err := basicQuery.Execute(mdb.Database)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	fmt.Printf("Results: %v\n\n", results)
	
	
	//Example with sqlparser
	sql := `
    SELECT field1, field2 
    FROM orders 
    WHERE status = 'active' 
    ORDER BY field1 ASC 
    LIMIT 10`

    parser := sqlparser.New()
    sp, err := parser.Parse(sql)
    if err != nil {
        log.Fatalf("Failed to parse SQL query: %v", err)
    }

    resultsSqlParser, err := sp.Execute(mdb.Database)
    if err != nil {
        log.Fatalf("Query execution failed: %v", err)
    }
    fmt.Println("Results:", resultsSqlParser)
}
```
---
## **QUERY BUILDER**
## 1. SELECT

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Select(fields ...string)`      | Specifies the columns to select.                                            |
| `From(collection string)`       | Specifies the collection to query.                                          |
| `Where(condition string)`       | Defines filter conditions (`AND`, `OR`, `=`, `!=`, `<`, `>`, `<=`, `>=`). Supports single and multiple conditions, logical operators, and grouping with parentheses. Converts SQL-like syntax to MongoDB filters. |
| `GroupBy(field string)`         | Groups the results by a specific field.                                     |
| `Having(condition string)`      | Filters aggregation results (`SUM`, `COUNT`, etc.).                         |
| `OrderBy(fieldOrder string)`    | Sorts the results (`ASC` / `DESC`).                                         |
| `Limit(limit int64)`            | Limits the number of query results.                                         |
| `Offset(offset int64)`          | Skips a specific number of documents before retrieving results.             |

### Example

#### Basic SELECT Query
```go
qb := builder.NewQueryBuilder().
    Select("field1", "field2").
    From("orders").
    Where("status = 'active'").
    OrderBy("field1 ASC").
    Limit(10)

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Println("Results:", results)
```

#### SELECT with GroupBy and Having
```go
qb := builder.NewQueryBuilder().
    Select("category", "SUM(amount) AS totalAmount").
    From("orders").
    GroupBy("category").
    Having("totalAmount > 5000").
    OrderBy("totalAmount DESC").
    Limit(10)

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Println("Results:", results)
```

---

## 2. INSERT

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `InsertInto(collection string, fields []string)` | Specifies the collection and columns for inserting data.                 |
| `Values(values []interface{})`        | Adds values for the specified columns.                                   |
| `BulkValues(values [][]interface{})`  | Adds multiple sets of values for the columns.                            |

### Example

#### Single Insert
```go
qb := builder.NewInsertBuilder().
    InsertInto("orders", []string{"field1", "field2"}).
    Values([]interface{}{"value1", 100}).
    Execute()
```

#### Multiple Inserts
```go
qb := builder.NewInsertBuilder().
    InsertInto("orders", []string{"field1", "field2"}).
    Values([]interface{}{"value1", 100}).
    Values([]interface{}{"value2", 200}).
    Execute()
```

---

## 3. UPDATE

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Set(data map[string]interface{})` | Specifies the columns and values to update.                                |
| `Where(condition string)`       | Defines filter conditions for the update.                                   |
| `SetMulti(multi bool)`          | Enables updating multiple documents.                                        |

### Example

#### Single Document Update
```go
qb := builder.NewUpdateBuilder("products").
    Set(map[string]interface{}{"status": "inactive"}).
    Where("product_id = 123").
    Execute()
```

#### Multiple Document Update
```go
qb := builder.NewUpdateBuilder("products").
    Set(map[string]interface{}{"status": "inactive"}).
    Where("stock < 10").
    SetMulti(true).
    Execute()
```

---

## 4. DELETE

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Where(condition string)`       | Defines filter conditions for deletion.                                     |
| `SetMulti(multi bool)`          | Enables deleting multiple documents.                                        |

### Example

#### Single Document Deletion
```go
qb := builder.NewDeleteBuilder("orders").
    Where("status = 'inactive'").
    Execute()
```

#### Multiple Document Deletion
```go
qb := builder.NewDeleteBuilder("orders").
    Where("status = 'inactive'").
    SetMulti(true).
    Execute()
```

---

## 5. JOIN

| Function                              | Description                                                                 |
|---------------------------------------|-----------------------------------------------------------------------------|
| `Join(localField, fromCollection, foreignField, as string)` | Adds a `$lookup` stage to perform joins between collections.              |

### Example

#### Basic JOIN Query
```go
qb := builder.NewQueryBuilder().
    From("orders").
    Join("customerId", "customers", "_id", "customerDetails").
    Match("status = 'Completed'").
    Limit(5)

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Printf("JOIN Results: %v\n", results)
```

---

## 6. GROUP BY

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `GroupBy(field string)`               | Groups the results by a specific field.                                   |
| `NestedGroupBy(fields ...string)`     | Groups results at multiple levels (nested grouping).                      |

### Example

#### Basic GROUP BY
```go
qb := builder.NewQueryBuilder().
    From("sales").
    GroupBy("region").
    Select("region", "SUM(amount) AS totalSales").
    Having("totalSales > 1000")

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Printf("GroupBy Results: %v\n", results)
```

---

## 7. AGGREGATION

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Match(condition string)`             | Filters documents based on conditions.                                    |
| `GroupBy(field string)`               | Groups results and performs aggregation.                                  |
| `OrderBy(fieldOrder string)`          | Sorts aggregated results.                                                 |
| `AggregationLimit(limit int64)`       | Limits the number of results in the aggregation pipeline.                 |
| `AggregationOffset(offset int64)`     | Skips a specific number of documents in the aggregation pipeline.         |

### Example

#### Basic Aggregation Pipeline
```go
qb := builder.NewQueryBuilder().
    From("orders").
    Match("status = 'Completed'").
    GroupBy("customer", "SUM(amount) AS totalSpent", "COUNT(*) AS totalOrders").
    OrderBy("totalSpent DESC").
    Limit(5)

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Printf("Aggregation Results: %v\n", results)
```

---

## 8. HAVING

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Having(condition string)`            | Filters grouped results after aggregation (e.g., `SUM`, `COUNT`).         |

### Example

#### Using HAVING Clause
```go
qb := builder.NewQueryBuilder().
    From("sales").
    GroupBy("region").
    Having("SUM(amount) > 10000").
    OrderBy("region ASC")

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Printf("Having Results: %v\n", results)
```

---

## 9. CONDITIONS

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Where(condition string)`             | Handles single and multiple conditions (`AND`, `OR`, parentheses).        |

### Example

#### Single Condition
```go
qb := builder.NewQueryBuilder().
    From("products").
    Where("price > 100").
    Limit(5)

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Printf("Single Condition Results: %v\n", results)
```

#### Multiple Conditions
```go
qb := builder.NewQueryBuilder().
    From("products").
    Where("price > 100 AND (stock > 50 OR rating >= 4.5)").
    Limit(10)

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Printf("Multiple Conditions Results: %v\n", results)
```

---

## 10. EXPRESSION PARSING

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `parseConditions(condition string)`   | Parses and converts conditions into MongoDB filters.                      |
| `parseExpression(expression string)`  | Parses mathematical and logical expressions.                              |

---

## 11. CREATE INDEX

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Index(name string, fields string)`   | Creates an index on the specified fields in a collection.                 |

### Example

#### Create Index
```go
qb := builder.NewCreateIndexBuilder("orders").
    Index("idx_status", "status ASC")

err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Create Index Query failed: %v", err)
}
fmt.Println("Create Index executed successfully")
```

---

## 12. DELETE INDEX

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Index(name string)`                  | Removes an index from the specified collection.                           |

### Example

#### Delete Index
```go
qb := builder.NewDeleteIndexBuilder("orders").
    Index("idx_status")

err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Delete Index Query failed: %v", err)
}
fmt.Println("Delete Index executed successfully")
```
---

## Notes and Examples

This documentation provides detailed explanations and examples for each feature in MongoQuery. Use these as references to build efficient and readable MongoDB queries with SQL-like syntax.

---

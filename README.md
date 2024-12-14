# MongoQuery - SQL-like Query Builder for MongoDB

MongoQuery is a powerful and intuitive query builder for MongoDB, designed to provide a SQL-like syntax for building and executing queries. It simplifies complex MongoDB operations such as aggregation, filtering, updates, and indexing, allowing developers to write queries in a format familiar to those accustomed to SQL.

With MongoQuery, you can seamlessly handle dynamic conditions, nested aggregations, pagination (via `LIMIT` and `OFFSET`), and advanced filtering using mathematical and logical expressions. This library is perfect for developers who prefer the simplicity of SQL while leveraging the flexibility and power of MongoDB.

---

## Key Features

- **SQL-like Syntax**: Write MongoDB queries with a structure similar to SQL (e.g., `SELECT`, `WHERE`, `GROUP BY`, `ORDER BY`, etc.).
- **Dynamic Filtering**: Supports complex conditions (`AND`, `OR`, `=`, `>`, `<`, etc.).
- **Aggregation Pipelines**: Simplifies the creation of pipelines with support for grouping, sorting, and nested operations.
- **Pagination**: Includes `LIMIT` and `OFFSET` for easy implementation of paginated queries.
- **CRUD Operations**: Supports `INSERT`, `UPDATE`, `DELETE`, and indexing operations.
- **Expressive Queries**: Handle advanced mathematical and logical expressions in your queries.
- **Modular and Extendable**: Designed to be modular for easy integration and extension in various projects.

---


## Query Functions

Below is the complete list of query functions available in the MongoQuery library:

---

## **1. SELECT**

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Select(fields ...string)`      | Specifies the columns to select.                                            |
| `From(collection string)`       | Specifies the collection to query.                                          |
| `Where(condition string)`       | Defines filter conditions (AND, OR, =, !=, <, >, <=, >=). Supports single and multiple conditions, logical operators, and grouping with parentheses (e.g., price > 10 AND (stock > 50 OR rating >= 4)). Converts SQL-like syntax to MongoDB filters.    |
| `GroupBy(field string)`         | Groups the results by a specific field.                                     |
| `Having(condition string)`      | Filters aggregation results (`SUM`, `COUNT`, etc.).                         |
| `OrderBy(fieldOrder string)`    | Sorts the results (`ASC` / `DESC`).                                         |
| `Limit(limit int64)`            | Limits the number of query results.                                         |
| `Offset(offset int64)`          | Skips a specific number of documents before retrieving results.             |

---

## **2. INSERT**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `InsertInto(collection string, fields []string)` | Specifies the collection and columns for inserting data.                 |
| `Values(values []interface{})`        | Adds values for the specified columns.                                   |
| `BulkValues(values [][]interface{})`  | Adds multiple sets of values for the columns.                            |

---

## **3. UPDATE**

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Set(data map[string]interface{})` | Specifies the columns and values to update.                                |
| `Where(condition string)`       | Defines filter conditions for the update.                                   |
| `SetMulti(multi bool)`          | Enables updating multiple documents.                                        |

---

## **4. DELETE**

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Where(condition string)`       | Defines filter conditions for deletion.                                     |
| `SetMulti(multi bool)`          | Enables deleting multiple documents.                                        |

---

## **5. JOIN**

| Function                              | Description                                                                 |
|---------------------------------------|-----------------------------------------------------------------------------|
| `Join(localField, fromCollection, foreignField, as string)` | Adds a `$lookup` stage to perform joins between collections.              |

---

## **6. GROUP BY**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `GroupBy(field string)`               | Groups the results by a specific field.                                   |
| `NestedGroupBy(fields ...string)`     | Groups results at multiple levels (nested grouping).                      |

---

## **7. AGGREGATION**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Match(condition string)`             | Filters documents based on conditions.                                    |
| `GroupBy(field string)`               | Groups results and performs aggregation.                                  |
| `OrderBy(fieldOrder string)`          | Sorts aggregated results.                                                 |
| `AggregationLimit(limit int64)`       | Limits the number of results in the aggregation pipeline.                 |
| `AggregationOffset(offset int64)`     | Skips a specific number of documents in the aggregation pipeline.         |

---

## **8. HAVING**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Having(condition string)`            | Filters grouped results after aggregation (e.g., `SUM`, `COUNT`).         |

---

## **9. CONDITIONS**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Where(condition string)`             | Handles single and multiple conditions (`AND`, `OR`, parentheses).        |

---

## **10. EXPRESSION PARSING**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `parseConditions(condition string)`   | Parses and converts conditions into MongoDB filters.                      |
| `parseExpression(expression string)`  | Parses mathematical and logical expressions.                              |

---

## **11. CREATE INDEX**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Index(name string, fields string)`   | Creates an index on the specified fields in a collection.                 |

---

## **12. DELETE INDEX**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `Index(name string)`                  | Removes an index from the specified collection.                           |

---


## **HOW TO USE THIS QUERY**
### **Example Implementation**
```go
package main

import (
	"fmt"
	"log"

	"github.com/brothergiez/mongoquery/builder"
	"github.com/brothergiez/mongoquery/client"
)

func main() {
	// Initialize MongoDB client
	mdb, err := client.New("mongodb://localhost:27017", "test_db")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Example 1: Basic SELECT query
	fmt.Println("Example 1: Basic SELECT query")
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

	// Example 2: SELECT with Group By and Having
	fmt.Println("Example 2: SELECT with Group By and Having")
	groupQuery := builder.NewQueryBuilder().
		Select("category", "SUM(amount) AS totalAmount").
		From("orders").
		GroupBy("category").
		Having("totalAmount > 5000").
		OrderBy("totalAmount DESC").
		Limit(5)

	groupResults, err := groupQuery.Execute(mdb.Database)
	if err != nil {
		log.Fatalf("Group query failed: %v", err)
	}
	fmt.Printf("Group Results: %v\n\n", groupResults)

	// Example 3: SELECT with Offset for Pagination
	fmt.Println("Example 3: SELECT with Offset for Pagination")
	paginationQuery := builder.NewQueryBuilder().
		Select("field1", "field2").
		From("orders").
		Where("status = 'active'").
		OrderBy("field1 ASC").
		Limit(10).
		Offset(20)

	paginatedResults, err := paginationQuery.Execute(mdb.Database)
	if err != nil {
		log.Fatalf("Pagination query failed: %v", err)
	}
	fmt.Printf("Paginated Results: %v\n", paginatedResults)
}
```

### **SELECT QUERY**
#### **Description**
The **SELECT** query in MongoQuery allows you to retrieve data from a MongoDB collection using SQL-like syntax. It supports filtering, grouping, sorting, and limiting results, making it easy to construct complex queries in a readable and intuitive manner.

#### Example Usage

**1. Basic SELECT Query**
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

**2. SELECT with Group By and Having**
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

**3. SELECT with Offset for Pagination**
```go
qb := builder.NewQueryBuilder().
    Select("field1", "field2").
    From("orders").
    Where("status = 'active'").
    OrderBy("field1 ASC").
    Limit(10).
    Offset(20)

results, err := qb.Execute(mdb.Database)
if err != nil {
    log.Fatalf("Query failed: %v", err)
}
fmt.Println("Results:", results)
```
**Output**

Example Result If the query is:
```sql
SELECT field1, field2 
FROM orders 
WHERE status = 'active' 
ORDER BY field1 ASC 
LIMIT 10 OFFSET 20;
```
The corresponding MongoDB query pipeline will look like:
```json
[
    { "$match": { "status": "active" } },
    { "$sort": { "field1": 1 } },
    { "$skip": 20 },
    { "$limit": 10 }
]
```

#### **Notes**
- The Where condition supports nested expressions and dynamic operators (AND, OR, etc.).
- GroupBy must be used with aggregation functions like SUM, COUNT, etc., for meaningful results.
- Offset is implemented using the MongoDB $skip stage in the aggregation pipeline.
- Combines seamlessly with Limit for pagination.

---

### INSERT

#### Description
The **INSERT** query in MongoQuery allows you to insert one or multiple documents into a MongoDB collection using SQL-like syntax. You can specify the fields and their corresponding values for insertion, making the process intuitive and structured.

---

#### Example Usage

**Insert a Single Row**
```go
qb := builder.NewInsertBuilder().
    InsertInto("orders", []string{"field1", "field2"}).
    Values([]interface{}{"value1", 100}).
    Execute()
```

**Insert Multiple Rows**
```go
qb := builder.NewInsertBuilder().
    InsertInto("orders", []string{"field1", "field2"}).
    Values([]interface{}{"value1", 100}).
    Values([]interface{}{"value2", 200}).
    Execute()
```

**Output**

If the query is:
```sql
INSERT INTO orders (field1, field2) VALUES ('value1', 100), ('value2', 200);
```

The corresponding MongoDB operation will insert the following documents:
```json
[
    { "field1": "value1", "field2": 100 },
    { "field1": "value2", "field2": 200 }
]
```

#### **Notes**

- The InsertInto method specifies the target collection and the fields.
- The Values method allows you to insert a single row of values.
- The BulkValues method supports inserting multiple rows at once.
- Ensure the number of fields matches the number of values provided.


#### **Error Handling**
- If the number of fields and values do not match, the function will throw an error.
- MongoDB connection or schema issues will result in runtime errors; handle them gracefully.
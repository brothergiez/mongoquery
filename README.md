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

### **1. SELECT**

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Select(fields ...string)`      | Specifies the columns to select.                                            |
| `From(collection string)`       | Specifies the collection to query.                                          |
| `Where(condition string)`       | Defines filter conditions (`AND`, `OR`, `=`).                               |
| `GroupBy(field string)`         | Groups the results by a specific field.                                     |
| `Having(condition string)`      | Filters aggregation results (`SUM`, `COUNT`, etc.).                         |
| `OrderBy(fieldOrder string)`    | Sorts the results (`ASC`/`DESC`).                                           |
| `Limit(limit int64)`            | Limits the number of query results.                                         |
| `Offset(offset int64)`          | Skips a specific number of documents before retrieving results.             |

---

### **2. INSERT**

| Function                              | Description                                                               |
|---------------------------------------|---------------------------------------------------------------------------|
| `InsertInto(collection string, fields []string)` | Specifies the collection and columns for inserting data.                 |
| `Values(values []interface{})`        | Adds values for the specified columns.                                   |
| `BulkValues(values [][]interface{})`  | Adds multiple sets of values for the columns.                            |

---

### **3. UPDATE**

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Set(data map[string]interface{})` | Specifies the columns and values to update.                                |
| `Where(condition string)`       | Defines filter conditions for the update.                                   |
| `SetMulti(multi bool)`          | Enables updating multiple documents.                                        |

---

### **4. DELETE**

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Where(condition string)`       | Defines filter conditions for deletion.                                     |
| `SetMulti(multi bool)`          | Enables deleting multiple documents.                                        |

---

### **5. CREATE INDEX**

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Index(name string, fields string)` | Defines the index name and the fields to index.                            |

---

### **6. DELETE INDEX**

| Function                        | Description                                                                 |
|---------------------------------|-----------------------------------------------------------------------------|
| `Index(name string)`            | Specifies the name of the index to delete.                                  |

---

### **7. AGGREGATION**

| Function                          | Description                                                                 |
|-----------------------------------|-----------------------------------------------------------------------------|
| `Match(condition string)`         | Adds a filter (`WHERE`) to the aggregation pipeline.                        |
| `GroupBy(field string)`           | Adds grouping to the aggregation pipeline.                                  |
| `NestedGroupBy(fields ...string)` | Supports nested aggregation grouping.                                       |
| `Having(condition string)`        | Adds filtering for aggregation results.                                     |
| `OrderBy(fieldOrder string)`      | Adds sorting to the aggregation pipeline.                                   |
| `AggregationLimit(limit int64)`   | Limits the number of documents in the aggregation pipeline.                 |
| `AggregationOffset(offset int64)` | Skips a specific number of documents in the aggregation pipeline.           |

---

### **8. EXPRESSIONS & CONDITIONS**

| Function                          | Description                                                                 |
|-----------------------------------|-----------------------------------------------------------------------------|
| `parseConditions(condition string)` | Processes dynamic conditions (`AND`, `OR`, `=`, etc.).                     |
| `parseExpression(expression string)` | Supports parsing mathematical and logical expressions.                     |
| `mapOperatorToMongo(operator string)` | Converts SQL-like operators to MongoDB operators (`>`, `<`, `=`, `!=`).    |

---
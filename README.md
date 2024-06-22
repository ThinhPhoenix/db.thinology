# Structured Query API `(sqa)`

## Introduction

The Structured Query API (sqa) allows you to execute SQL queries against various databases programmatically. This document provides an overview of how to use the API to interact with your databases.

## API Endpoint

**Endpoint:** `https://sqa.thinology.id.vn/sql`

**Method:** `POST`

## Request Body

The API expects a JSON body with the following parameters:

```json
{
  "database_type": "postgres",
  "database_string": "databasestring here",
  "query": "select * from users"
}
```

### Parameters

- **database_type**: The type of database you are querying (e.g., `postgres`, `mysql`, etc.).
- **database_string**: The connection string or configuration details for your database.
- **query**: The SQL query you want to execute.

## Example

### Example Request

```bash
curl -X POST https://sqa.thinology.id.vn/sql \
  -H "Content-Type: application/json" \
  -d '{
        "database_type": "postgres",
        "database_string": "your_postgres_connection_string_here",
        "query": "select * from users"
      }'
```

### Example Response

```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "username": "john_doe",
      "email": "john.doe@example.com"
    },
    {
      "id": 2,
      "username": "jane_smith",
      "email": "jane.smith@example.com"
    }
  ]
}
```

## Error Handling

- If there is an error with the SQL query or database connection, the API will return an appropriate error message with a non-200 HTTP status code.

## Security

- Ensure that database connection strings and queries are passed securely and not exposed in public repositories or client-side code.

## Notes

- Always sanitize and validate user inputs to prevent SQL injection attacks.
- Check the API documentation for any additional features or limitations specific to the `sqa` service.

## Support

For support or inquiries, please contact [support@thinology.id.vn](mailto:support@thinology.id.vn).

---

Adjust the specifics like `database_type`, `database_string`, and `query` based on the actual requirements and capabilities of your `sqa` API service. This template covers essential information to help users understand how to interact with the API effectively.

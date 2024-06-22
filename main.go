package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/denisenkom/go-mssqldb" // SQL Server driver
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/lib/pq"              // PostgreSQL driver
)

type CustomQuery struct {
	DatabaseType   string `json:"database_type"`   // postgres, mysql, sqlserver
	DatabaseString string `json:"database_string"` // connection string
	Query          string `json:"query"`           // SQL query
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

func getDBConnection(dbType, dbString string) (*sql.DB, error) {
	db, err := sql.Open(dbType, dbString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func executeQuery(db *sql.DB, query string) (bool, interface{}) {
	rows, err := db.Query(query)
	if err != nil {
		return false, err.Error()
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return false, err.Error()
	}

	var result []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return false, err.Error()
		}
		entry := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				entry[col] = string(b)
			} else {
				entry[col] = val
			}
		}
		result = append(result, entry)
	}

	return true, result
}

func sqlEndpoint(c *gin.Context) {
	var req CustomQuery
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := getDBConnection(req.DatabaseType, req.DatabaseString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	success, result := executeQuery(db, req.Query)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Query executed successfully", "data": result})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error: %s", result), "data": nil})
	}
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/sql", sqlEndpoint)
	log.Fatal(r.Run(":8080"))
}

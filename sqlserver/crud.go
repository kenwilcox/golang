package main

import (
    _ "github.com/denisenkom/go-mssqldb"
    "database/sql"
    "context"
    "log"
    "fmt"
)

var db *sql.DB

var server = "localhost"
var port = 1433
var user = "sa"
var password = "******"
var database = "SampleDB"

func main() {
    // Build connection string
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        server, user, password, port, database)

    var err error

    // Create connection pool
    db, err = sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal("Error creating connection pool:", err.Error())
    }
    fmt.Printf("Connected!\n")

    // Create employee
    createId, err := CreateEmployee("Jake", "United States")
    fmt.Printf("Inserted ID: %d successfully.\n", createId)

    // Read employees
    count, err := ReadEmployees()
    fmt.Printf("Read %d rows successfully.\n", count)

    // Update from database
    updateId, err := UpdateEmployee("Jake", "Poland")
    fmt.Printf("Updated row with ID: %d successfully.\n", updateId)

    // Delete from database
    rows, err := DeleteEmployee("Jake")
    fmt.Printf("Deleted %d rows successfully.\n", rows)
}

func CreateEmployee(name string, location string) (int64, error) {
    ctx := context.Background()
    var err error

    if db == nil {
        log.Fatal("What?")
    }

    // Check if database is alive.
    err = db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("INSERT INTO TestSchema.Employees (Name, Location) VALUES (@Name,@Location);")

    // Execute non-query with named parameters
    result, err := db.ExecContext(
        ctx,
        tsql,
        sql.Named("Location", location),
        sql.Named("Name", name))

	if err != nil {
		log.Fatal("Error inserting new row: " + err.Error())
		return -1, err
	}

	return result.LastInsertId()
}

func ReadEmployees() (int, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

    // Execute query
    rows, err := db.QueryContext(ctx, tsql)
    if err != nil {
        log.Fatal("Error reading rows: " + err.Error())
        return -1, err
    }

    defer rows.Close()

    var count int = 0

    // Iterate through the result set.
    for rows.Next() {
        var name, location string
        var id int

        // Get values from row.
        err := rows.Scan(&id, &name, &location)
        if err != nil {
            log.Fatal("Error reading rows: " + err.Error())
            return -1, err
        }

        fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
        count++
    }

    return count, nil
}

// Update an employee's information
func UpdateEmployee(name string, location string) (int64, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("UPDATE TestSchema.Employees SET Location = @Location WHERE Name= @Name")

    // Execute non-query with named parameters
    result, err := db.ExecContext(
        ctx,
        tsql,
        sql.Named("Location", location),
        sql.Named("Name", name))
    if err != nil {
        log.Fatal("Error updating row: " + err.Error())
        return -1, err
    }

    return result.LastInsertId()
}

// Delete an employee from database
func DeleteEmployee(name string) (int64, error) {
    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name=@Name;")

    // Execute non-query with named parameters
    result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
    if err != nil {
        fmt.Println("Error deleting row: " + err.Error())
        return -1, err
    }

    return result.RowsAffected()
}


package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/Zephyruston/sql2pb/core"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	dbType := flag.String("db_type", "mysql", "the database type (mysql or pg)")
	host := flag.String("host", "localhost", "the database host")
	port := flag.Int("port", 3306, "the database port")
	user := flag.String("user", "root", "the database user")
	password := flag.String("password", "", "the database password")
	dbName := flag.String("db", "", "the database name")
	schema := flag.String("schema", "public", "the database schema name, default is 'public' for postgresql")
	table := flag.String("table", "*", "the table schema，multiple tables ',' split. ")
	serviceName := flag.String("service_name", *dbName, "the protobuf service name , defaults to the database name.")
	packageName := flag.String("package", *dbName, "the protocol buffer package. defaults to the database name.")
	goPackageName := flag.String("go_package", "", "the protocol buffer go_package. defaults to the database name.")
	ignoreTableStr := flag.String("ignore_tables", "", "a comma spaced list of tables to ignore")
	ignoreColumnStr := flag.String("ignore_columns", "", "a comma spaced list of columns to ignore")
	fieldStyle := flag.String("field_style", "sqlPb", "gen protobuf field style, sql_pb | sqlPb")

	flag.Parse()

	if *dbName == "" {
		fmt.Println(" - please input the database name using -db flag")
		return
	}

	// Adjust default port for PostgreSQL
	if *dbType == "pg" && *port == 3306 {
		*port = 5432
	}

	var connStr string
	var driverName string
	switch *dbType {
	case "mysql":
		driverName = "mysql"
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *user, *password, *host, *port, *dbName)
	case "pg":
		driverName = "postgres"
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *host, *port, *user, *password, *dbName)
	default:
		log.Fatal("unsupported database type: ", *dbType, ". Use 'mysql' or 'pg'")
	}

	db, err := sql.Open(driverName, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	ignoreTables := strings.Split(*ignoreTableStr, ",")
	ignoreColumns := strings.Split(*ignoreColumnStr, ",")

	s, err := core.GenerateSchemaWithSchema(db, *table, ignoreTables, ignoreColumns, *serviceName, *goPackageName, *packageName, *fieldStyle, *schema)

	if nil != err {
		log.Fatal(err)
	}

	if nil != s {
		fmt.Println(s)
	}
}

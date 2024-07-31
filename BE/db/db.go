package db

import (
	"certainwager-be/models"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	host := os.Getenv("DBHOST")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("SSLMODE")
	

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s ", user, dbname, sslmode, password, host)
	// connStr := "user=postgres dbname=certainwager sslmode=disable password=0000 host=172.18.96.1"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")

	createTables()
}

func createTables() {
	createTableFromModel(models.Offer{}, "offers")
	createTableFromModel(models.Review{}, "reviews")
	createTableFromModel(models.Blog{}, "blogs")
}

func createTableFromModel(model interface{}, tableName string) {
	val := reflect.ValueOf(model)
	typ := reflect.TypeOf(model)

	var columns []string

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		columnName := strings.ToLower(field.Name)
		columnType := getSQLType(field.Type, field.Tag.Get("json"))

		if columnType != "" {
			columns = append(columns, fmt.Sprintf("%s %s", columnName, columnType))
		}
	}

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		%s
	)`, tableName, strings.Join(columns[1:], ",\n"))

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table %s: %v", tableName, err)
	}

	fmt.Printf("Table %s created or already exists.\n", tableName)
}

func getSQLType(t reflect.Type, jsonTag string) string {
	switch t.Kind() {
	case reflect.String:
		return "TEXT"
	case reflect.Int:
		return "INTEGER"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Slice:
		// Handling slices as TEXT for simplicity. In production, this should be handled better, perhaps using JSONB.
		if t.Elem().Kind() == reflect.String {
			return "TEXT"
		}
		return ""
	case reflect.Struct:
		if jsonTag == "offer" {
			return "INTEGER REFERENCES offers(id)"
		}
		return "INTEGER" // Assuming foreign key or nested structure. Adjust as necessary.
	default:
		return "TEXT"
	}
}

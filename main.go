package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

type Column struct {
	Name string
	Type string
}

type Table struct {
	Name    string
	Columns []Column
}

func mapType(postgresType string) string {
	switch postgresType {
	case "integer", "serial":
		return "int"
	case "bigint":
		return "int64"
	case "text", "varchar", "char":
		return "string"
	case "boolean":
		return "bool"
	case "timestamp", "date":
		return "time.Time"
	case "jsonb":
		return "map[string]interface{}"
	case "float8", "real":
		return "float64"
	default:
		return "string" // Default to string if type is unknown
	}
}

func loadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getTables(db *sql.DB) ([]string, error) {
	// Query for all table names
	rows, err := db.Query(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	return tables, nil
}

func generateCRUDForTable(db *sql.DB, tableName string, customFields string, onlyHandlers bool, onlyModels bool) error {
	// Query the columns of the specified table
	rows, err := db.Query(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_name = $1`, tableName)
	if err != nil {
		return err
	}
	defer rows.Close()

	var cols []Column
	for rows.Next() {
		var col Column
		if err := rows.Scan(&col.Name, &col.Type); err != nil {
			return err
		}
		col.Type = mapType(col.Type) // Map to Go type
		cols = append(cols, col)
	}

	// Add custom fields, if provided
	if customFields != "" {
		for _, field := range strings.Split(customFields, ",") {
			cols = append(cols, Column{Name: field, Type: "string"})
		}
	}

	table := Table{Name: tableName, Columns: cols}

	// Generate CRUD files based on user choices
	if !onlyHandlers {
		renderTemplate("templates/model.tmpl", "output/"+tableName+"_model.go", table)
	}
	if !onlyModels {
		renderTemplate("templates/handler.tmpl", "output/"+tableName+"_handler.go", table)
		renderTemplate("templates/routes.tmpl", "output/"+tableName+"_routes.go", table)
	}

	// If desired, generate additional files like services or DTOs here.
	return nil
}

func main() {
	// Load environment variables
	loadEnv()

	var connStr, customFields, generate string
	var onlyHandlers, onlyModels bool

	var rootCmd = &cobra.Command{
		Use:   "crudgen",
		Short: "Generates CRUD API from a PostgreSQL table",
		Run: func(cmd *cobra.Command, args []string) {
			// Connect to the PostgreSQL database
			db, err := sql.Open("postgres", connStr)
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			// Get all tables from the database
			tables, err := getTables(db)
			if err != nil {
				log.Fatal(err)
			}

			// Generate CRUD for each table
			for _, tableName := range tables {
				err := generateCRUDForTable(db, tableName, customFields, onlyHandlers, onlyModels)
				if err != nil {
					log.Printf("Error generating CRUD for table %s: %v", tableName, err)
				}
			}
		},
	}

	// Define flags
	rootCmd.Flags().StringVarP(&connStr, "conn", "c", "", "PostgreSQL connection string (required)")
	rootCmd.Flags().StringVarP(&customFields, "fields", "f", "", "Comma-separated list of custom fields to add")
	rootCmd.Flags().StringVarP(&generate, "generate", "g", "all", "Specify which components to generate: model, handler, or all")
	rootCmd.Flags().BoolVarP(&onlyHandlers, "handlers", "h", false, "Generate only handlers")
	rootCmd.Flags().BoolVarP(&onlyModels, "models", "m", false, "Generate only models")

	// Mark flags as required
	rootCmd.MarkFlagRequired("conn")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Function to render the template
func renderTemplate(templatePath, outputPath string, data Table) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		log.Fatal(err)
	}
}

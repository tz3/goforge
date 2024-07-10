// Package steps defines the steps involved in setting up a Go project.
package steps

// StepSchema represents a single step in the setup process.
// It includes the name of the step, the options available in this step,
// the headers to be displayed, and a pointer to the field where the user's
// choice will be stored.
type StepSchema struct {
	StepName string
	Options  []Option
	Headers  string
	Field    *string
}

// Option represents a single option that can be chosen in a step.
// It includes the title of the option and a description.
type Option struct {
	Flag  string
	Title string
	Desc  string
}

// Steps is a collection of StepSchema. It represents all the steps
// involved in the setup process.
type Steps struct {
	Steps map[string]StepSchema
}

// InitSteps initializes the steps of the setup process.
// The Steps struct includes all the steps involved in the setup process.
func InitSteps() *Steps {
	steps := &Steps{
		map[string]StepSchema{
			"web-framework": {
				StepName: "Web Framework",
				Options: []Option{
					{
						Flag:  "standard-library",
						Title: "Standard-library",
						Desc:  "Built in standard golang library",
					},
					{
						Flag:  "chi",
						Title: "Chi",
						Desc:  "use go-chi from: https://github.com/go-chi/chi",
					},
					{
						Flag:  "gin",
						Title: "Gin",
						Desc:  "use gin-gonic from: https://github.com/gin-gonic/gin",
					},
					{
						Flag:  "fiber",
						Title: "Fiber",
						Desc:  "use gofiber from: https://github.com/gofiber/fiber",
					},
					{
						Flag:  "gorilla/mux",
						Title: "Gorilla/mux",
						Desc:  "use gorilla/mux from: https://github.com/gorilla/mux",
					},
					{
						Flag:  "httprouter",
						Title: "Httprouter",
						Desc:  "use julienschmidt/httprouter from: https://github.com/julienschmidt/httprouter",
					},
					{
						Flag:  "echo",
						Title: "Echo",
						Desc:  "use echo from: https://github.com/labstack/echo",
					},
				},
				Headers: "What web framework do you want to use in your Go project?",
			},
			"db-driver": {
				StepName: "Database Driver",
				Options: []Option{
					{
						Flag:  "mysql",
						Title: "Mysql",
						Desc:  "Use go-mysql-driver from: https://github.com/go-sql-driver/mysql",
					},
					{
						Flag:  "postgres",
						Title: "Postgres",
						Desc:  "Use pgx, PostgreSQL driver and toolkit from: get github.com/jackc/pgx/v5",
					},
					{
						Flag:  "sqlite",
						Title: "Sqlite",
						Desc:  "Use go-sqlite3, SQLite driver for go that using database/sql from: https://github.com/mattn/go-sqlite3",
					},
					{
						Flag:  "mongo",
						Title: "Mongo",
						Desc:  "Use mongo-driver, the Go driver for MongoDB from: https://github.com/mongodb/mongo-go-driver",
					},
					{
						Flag:  "none",
						Title: "None",
						Desc:  "Project with no Database setup!",
					},
				},
				Headers: "What database driver do you want to use in your Go project?",
			},
			"advanced": {
				StepName: "Advanced Features",
				Headers:  "Which advanced features do you want?",
				Options: []Option{
					{
						Flag:  "AddHTMXTempl",
						Title: "HTMX/Templ",
						Desc:  "Add starter HTMX and Templ files.",
					},
					{
						Flag:  "Go Project Workflow",
						Title: "Go Project Workflow",
						Desc:  "Workflow templates for testing, cross-compiling and releasing Go projects",
					},
				},
			},
		},
	}

	return steps
}

// Package project provides the functionality for creating a new Go project.
package project

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	// "github.com/melkeydev/go-blueprint/cmd/utils"

	"github.com/spf13/cobra"
	tpl "github.com/tz3/goforge/internal/templates"
	"github.com/tz3/goforge/internal/templates/db"
	"github.com/tz3/goforge/internal/templates/web"
)

// ProjectConfig represents the configuration for a new Go project.
// It includes the project name, type, a map of web frameworks, a flag to indicate
// whether to exit the CLI, and the absolute path of the project.
type ProjectConfig struct {
	ProjectName       string
	ProjectType       string
	DatabaseDriver    string
	DatabaseDriverMap map[string]DatabaseDriver // can be any of the supported Db Drivers
	FrameworkMap      map[string]WebFramework   // Can be any of the supported router Packages.
	Exit              bool
	AbsolutePath      string
}

// WebFramework represents a web framework that can be used in the project.
// It includes the dependencies of the framework and a template generator.
type WebFramework struct {
	dependencies []string
	templateGen  TemplateGenerator
}

// DatabaseDriver represents a database driver that can be used in the project.
// It includes the dependencies of the driver and a template generator.
type DatabaseDriver struct {
	dependencies []string
	templateGen  DBDriverTemplateGenerator
}

// TemplateGenerator is an interface that defines the methods for generating
// templates for the main, server, and routes files.
type TemplateGenerator interface {
	Main() []byte
	Server() []byte
	Routes() []byte
	RoutesWithDB() []byte
	ServerWithDB() []byte
}

type DBDriverTemplateGenerator interface {
	Service() []byte
	Env() []byte
}

// Web framework dependencies, and supported web-frameworks
var (
	SupportedWebframeworks   = []string{"chi", "echo", "fiber", "gin", "gorilla/mux", "httprouter", "standard-library"}
	SupportedDatabaseDrivers = []string{"mysql", "postgres", "sqlite", "mongo", "none"}
	chiDependencies          = []string{"github.com/go-chi/chi/v5"}
	gorillaDependencies      = []string{"github.com/gorilla/mux"}
	routerDependencies       = []string{"github.com/julienschmidt/httprouter"}
	ginDependencies          = []string{"github.com/gin-gonic/gin"}
	fiberDependencies        = []string{"github.com/gofiber/fiber/v2"}
	echoDependencies         = []string{"github.com/labstack/echo/v4", "github.com/labstack/echo/v4/middleware"}
	mysqlDependencies        = []string{"github.com/go-sql-driver/mysql"}
	postgresDependencies     = []string{"github.com/lib/pq"}
	sqliteDependencies       = []string{"github.com/mattn/go-sqlite3"}
	mongoDependencies        = []string{"go.mongodb.org/mongo-driver"}
	godotenvDependencies     = []string{"github.com/joho/godotenv"}
)

// File paths and names.
const (
	root                 = "/"
	cmdApiPath           = "cmd/api"
	internalServerPath   = "internal/server"
	internalDatabasePath = "internal/database"
	mainFile             = "main.go"
	databaseFile         = "database.go"
	serverFile           = "server.go"
	routesFile           = "routes.go"
)

// ExitCLI releases the terminal and exits the program if the Exit flag is set.
func (p *ProjectConfig) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		err := tprogram.ReleaseTerminal()
		if err != nil {
			log.Printf("Error in ReleaseTerminal %v\n", err)
		}

		os.Exit(1)
	}
}

// createFrameworkMap initializes the FrameworkMap with the available web frameworks.
func (p *ProjectConfig) createFrameworkMap() {
	p.FrameworkMap["standard-library"] = WebFramework{
		dependencies: []string{},
		templateGen:  web.StandardLibraryTemplate{},
	}

	p.FrameworkMap["chi"] = WebFramework{
		dependencies: chiDependencies,
		templateGen:  web.ChiTemplate{},
	}

	p.FrameworkMap["gin"] = WebFramework{
		dependencies: ginDependencies,
		templateGen:  web.GinTemplate{},
	}

	p.FrameworkMap["fiber"] = WebFramework{
		dependencies: fiberDependencies,
		templateGen:  web.FiberTemplate{},
	}

	p.FrameworkMap["gorilla/mux"] = WebFramework{
		dependencies: gorillaDependencies,
		templateGen:  web.GorillaTemplate{},
	}

	p.FrameworkMap["httprouter"] = WebFramework{
		dependencies: routerDependencies,
		templateGen:  web.HttpRouterTemplate{},
	}

	p.FrameworkMap["echo"] = WebFramework{
		dependencies: echoDependencies,
		templateGen:  web.EchoTemplate{},
	}
}

// createFrameworkMap initializes the FrameworkMap with the available web frameworks.
func (p *ProjectConfig) createDatabaseDriverMap() {
	p.DatabaseDriverMap["mysql"] = DatabaseDriver{
		dependencies: mysqlDependencies,
		templateGen:  db.MysqlTemplate{},
	}
	p.DatabaseDriverMap["postgres"] = DatabaseDriver{
		dependencies: postgresDependencies,
		templateGen:  db.PostgresTemplate{},
	}
	p.DatabaseDriverMap["sqlite"] = DatabaseDriver{
		dependencies: sqliteDependencies,
		templateGen:  db.SqliteTemplate{},
	}
	p.DatabaseDriverMap["mongo"] = DatabaseDriver{
		dependencies: mongoDependencies,
		templateGen:  db.MongoTemplate{},
	}
}

// Todo: Decompose this function into smaller ones.
// CreateMainFile creates the main file for the project.
// It creates the project directory, initializes the Go module, installs the dependencies,
// creates the necessary paths and files, and formats the Go code.
func (p *ProjectConfig) CreateMainFile() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			log.Printf("Could not create directory: %v", err)
			return err
		}
	}

	p.ProjectName = strings.TrimSpace(p.ProjectName)

	// Create a new directory with the project name
	if _, err := os.Stat(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName), 0751)
		if err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	}

	projectPath := fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)

	// Create the map for our program
	p.createFrameworkMap()

	// Create go.mod
	err := initGoMod(p.ProjectName, projectPath)
	if err != nil {
		log.Printf("Could not initialize go.mod in new project %v\n", err)
		cobra.CheckErr(err)
	}

	// Install the correct package for the selected framework
	if p.ProjectType != "standard library" {
		err = goGetDependencies(projectPath, p.FrameworkMap[p.ProjectType].dependencies)
		if err != nil {
			log.Printf("Could not install go dependency for the chosen framework %v\n", err)
			cobra.CheckErr(err)
		}
	}

	// Install the correct package for the selected driver
	if p.DatabaseDriver != "none" {
		p.createDatabaseDriverMap()
		err = goGetDependencies(projectPath, p.DatabaseDriverMap[p.DatabaseDriver].dependencies)
		if err != nil {
			log.Printf("Could not install go dependency for chosen driver %v\n", err)
			cobra.CheckErr(err)
		}

		err = p.createPath(internalDatabasePath, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", internalDatabasePath)
			cobra.CheckErr(err)
			return err
		}

		err = p.createFileAndWriteTemplate(internalDatabasePath, projectPath, databaseFile, "database")
		if err != nil {
			log.Printf("Error injecting server.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	}

	// Install the godotenv package
	err = goGetDependencies(projectPath, godotenvDependencies)
	if err != nil {
		log.Printf("Could not install go dependency %v\n", err)
		cobra.CheckErr(err)
	}

	err = p.createPath(cmdApiPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", projectPath)
		cobra.CheckErr(err)
		return err
	}

	err = p.createFileAndWriteTemplate(cmdApiPath, projectPath, mainFile, "main")
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	makeFile, err := os.Create(fmt.Sprintf("%s/Makefile", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	defer makeFile.Close()

	// inject makefile template
	makeFileTemplate := template.Must(template.New("makefile").Parse(string(tpl.MakeTemplate)))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
		return err
	}

	readmeFile, err := os.Create(fmt.Sprintf("%s/README.md", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}
	defer readmeFile.Close()

	// inject readme template
	readmeFileTemplate := template.Must(template.New("readme").Parse(string(tpl.ReadmeTemplate)))
	err = readmeFileTemplate.Execute(readmeFile, p)
	if err != nil {
		return err
	}

	err = p.createPath(internalServerPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", internalServerPath)
		cobra.CheckErr(err)
		return err
	}

	if p.DatabaseDriver != "none" {
		err = p.createFileAndWriteTemplate(internalServerPath, projectPath, "routes.go", "routesWithDB")
		if err != nil {
			log.Printf("Error injecting routes.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
		err = p.createFileAndWriteTemplate(internalServerPath, projectPath, "server.go", "serverWithDB")
		if err != nil {
			log.Printf("Error injecting server.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	} else {
		err = p.createFileAndWriteTemplate(internalServerPath, projectPath, "routes.go", "routes")
		if err != nil {
			log.Printf("Error injecting routes.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
		err = p.createFileAndWriteTemplate(internalServerPath, projectPath, "server.go", "server")
		if err != nil {
			log.Printf("Error injecting server.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	}

	err = p.createFileAndWriteTemplate(root, projectPath, ".env", "env")
	if err != nil {
		log.Printf("Error injecting .env file: %v", err)
		cobra.CheckErr(err)
		return err
	}

	// Initialize git repo
	err = initGitRepo(projectPath)
	if err != nil {
		log.Printf("Error initializing git repo: %v", err)
		cobra.CheckErr(err)
		return err
	}

	// Create gitignore
	gitignoreFile, err := os.Create(fmt.Sprintf("%s/.gitignore", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}
	defer gitignoreFile.Close()

	// inject gitignore template
	gitignoreTemplate := template.Must(template.New(".gitignore").Parse(string(tpl.GitIgnoreTemplate)))
	err = gitignoreTemplate.Execute(gitignoreFile, p)
	if err != nil {
		return err
	}

	// Create .air.toml file
	airTomlFile, err := os.Create(fmt.Sprintf("%s/.air.toml", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	defer airTomlFile.Close()

	// inject air.toml template
	airTomlTemplate := template.Must(template.New("airtoml").Parse(string(tpl.AirTomlTemplate)))
	err = airTomlTemplate.Execute(airTomlFile, p)
	if err != nil {
		return err
	}

	err = goFormat(projectPath)
	if err != nil {
		log.Printf("Could not gofmt in new project %v\n", err)
		cobra.CheckErr(err)
		return err
	}

	err = goTidy(projectPath)
	if err != nil {
		log.Printf("Could not go tidy in new project %v\n", err)
		cobra.CheckErr(err)
	}

	return nil
}

// createPath creates a new directory at the given path.
func (p *ProjectConfig) createPath(pathToCreate string, projectPath string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", projectPath, pathToCreate)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", projectPath, pathToCreate), 0751)
		if err != nil {
			fmt.Printf("Error creating path directory %v\n", err)
			return err
		}
	}

	return nil
}

// createFileAndWriteTemplate creates a new file at the given path and writes a template to it.
func (p *ProjectConfig) createFileAndWriteTemplate(pathToCreate string, projectPath string, fileName string, methodName string) error {
	createdFile, err := os.Create(fmt.Sprintf("%s/%s/%s", projectPath, pathToCreate, fileName))
	if err != nil {
		return err
	}

	defer createdFile.Close()

	// inject template
	switch methodName {
	case "main":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templateGen.Main())))
		err = createdTemplate.Execute(createdFile, p)
	case "server":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templateGen.Server())))
		err = createdTemplate.Execute(createdFile, p)
	case "serverWithDB":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templateGen.ServerWithDB())))
		err = createdTemplate.Execute(createdFile, p)
	case "routes":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templateGen.Routes())))
		err = createdTemplate.Execute(createdFile, p)
	case "routesWithDB":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templateGen.RoutesWithDB())))
		err = createdTemplate.Execute(createdFile, p)
	case "database":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DatabaseDriverMap[p.DatabaseDriver].templateGen.Service())))
		err = createdTemplate.Execute(createdFile, p)
	case "env":
		if p.DatabaseDriver != "none" {
			envBytes := [][]byte{
				tpl.EnvTemplate(),
				p.DatabaseDriverMap[p.DatabaseDriver].templateGen.Env(),
			}
			createdTemplate := template.Must(template.New(fileName).Parse(string(bytes.Join(envBytes, []byte("\n")))))
			err = createdTemplate.Execute(createdFile, p)
		} else {
			createdTemplate := template.Must(template.New(fileName).Parse(string(tpl.EnvTemplate())))
			err = createdTemplate.Execute(createdFile, p)
		}
	}
	if err != nil {
		return err
	}

	return nil
}

// isValidWebFramework check if the input is supported or not
func IsValidWebFramework(input string) bool {
	for _, t := range SupportedWebframeworks {
		if input == t {
			return true
		}
	}
	return false
}

// IsValidDatabaseDriver check if the input is supported or not
func IsValidDatabaseDriver(input string) bool {
	for _, t := range SupportedDatabaseDrivers {
		if input == t {
			return true
		}
	}
	return false
}

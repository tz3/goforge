// Package project provides the functionality for creating a new Go project.
package project

import (
	"fmt"
	"html/template"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	tpl "github.com/tz3/goforge/internal/templates"
	"github.com/tz3/goforge/internal/templates/web"
)

// ProjectConfig represents the configuration for a new Go project.
// It includes the project name, type, a map of web frameworks, a flag to indicate
// whether to exit the CLI, and the absolute path of the project.
type ProjectConfig struct {
	ProjectName  string
	ProjectType  string                  // Can be an api or serverless.
	FrameworkMap map[string]WebFramework // Can be any of the suggested router Packages.
	Exit         bool
	AbsolutePath string
}

// WebFramework represents a web framework that can be used in the project.
// It includes the dependencies of the framework and a template generator.
type WebFramework struct {
	dependencies []string
	templateGen  TemplateGenerator
}

// TemplateGenerator is an interface that defines the methods for generating
// templates for the main, server, and routes files.
type TemplateGenerator interface {
	Main() []byte
	Server() []byte
	Routes() []byte
}

// Todo:- add names of framework into const variable.

// Web framework dependencies, and supported web-frameworks
var (
	SupportedWebframeworks = []string{"chi", "echo", "fiber", "gin", "gorilla/mux", "httpRouter", "standard-library"}
	chiDependencies        = []string{"github.com/go-chi/chi/v5"}
	gorillaDependencies    = []string{"github.com/gorilla/mux"}
	routerDependencies     = []string{"github.com/julienschmidt/httprouter"}
	ginDependencies        = []string{"github.com/gin-gonic/gin"}
	fiberDependencies      = []string{"github.com/gofiber/fiber/v2"}
	echoDependencies       = []string{"github.com/labstack/echo/v4", "github.com/labstack/echo/v4/middleware"}
)

// File paths and names.
const (
	cmdApiPath         = "cmd/api"
	internalServerPath = "internal/server"
	mainFile           = "main.go"
	serverFile         = "server.go"
	routesFile         = "routes.go"
)

// ExitCLI releases the terminal and exits the program if the Exit flag is set.
func (p *ProjectConfig) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		tprogram.ReleaseTerminal()
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

	p.FrameworkMap["httpRouter"] = WebFramework{
		dependencies: routerDependencies,
		templateGen:  web.HttpRouterTemplate{},
	}

	p.FrameworkMap["echo"] = WebFramework{
		dependencies: echoDependencies,
		templateGen:  web.EchoTemplate{},
	}

}

// CreateMainFile creates the main file for the project.
// It creates the project directory, initializes the Go module, installs the dependencies,
// creates the necessary paths and files, and formats the Go code.
func (p *ProjectConfig) CreateMainFile() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	}

	// First lets create a new director with the project name
	if _, err := os.Stat(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName), 0751)
		if err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	}

	projectPath := fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)

	// create the router based on user input
	p.createFrameworkMap()

	err := initGoMod(p.ProjectName, projectPath)
	if err != nil {
		log.Printf("Failed to initialize Go module: %v\n", err)
		cobra.CheckErr(err)
		return err
	}

	// we need to install the correct package
	if p.ProjectType != "standard-library" {
		err = goGetDependencies(projectPath, p.FrameworkMap[p.ProjectType].dependencies)
		if err != nil {
			log.Printf("Failed to get package for project type %s: %v\n", p.ProjectType, err)
			cobra.CheckErr(err)
			return err
		}
	}

	err = p.createPath(cmdApiPath, projectPath)
	if err != nil {
		log.Printf("Failed to create path %s: %v\n", cmdApiPath, err)
		cobra.CheckErr(err)
		return err
	}

	err = p.createFileAndWriteTemplate(cmdApiPath, projectPath, mainFile, "main")
	if err != nil {
		log.Printf("Failed to create file and template %s: %v\n", cmdApiPath, err)
		cobra.CheckErr(err)
		return err
	}

	makeFile, err := os.Create(fmt.Sprintf("%s/Makefile", projectPath))
	if err != nil {
		log.Printf("Failed to create Makefile at path %s: %v\n", projectPath, err)
		cobra.CheckErr(err)
		return err
	}

	readMeFile, err := os.Create(fmt.Sprintf("%s/README.md", projectPath))
	if err != nil {
		log.Printf("Failed to create README file at path %s: %v\n", projectPath, err)
		cobra.CheckErr(err)
		return err
	}

	defer readMeFile.Close()
	defer makeFile.Close()

	// inject makefile template
	makeFileTemplate := template.Must(template.New("makefile").Parse(string(tpl.MakeTemplate())))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
		log.Printf("Failed to execute makefile template: %v\n", err)
		return err
	}

	// inject readmeFile template
	readMeFileTemplate := template.Must(template.New("README").Parse(string(tpl.ReadmeTemplate())))
	err = readMeFileTemplate.Execute(readMeFile, p)
	if err != nil {
		log.Printf("Failed to execute readMeFile template: %v\n", err)
		return err
	}

	err = p.createPath(internalServerPath, projectPath)
	if err != nil {
		log.Printf("Failed to create path %s: %v\n", internalServerPath, err)
		return err
	}

	err = p.createFileAndWriteTemplate(internalServerPath, projectPath, serverFile, "server")
	if err != nil {
		log.Printf("Failed to create and write to server file at path %s: %v\n", internalServerPath, err)
		return err
	}

	err = p.createFileAndWriteTemplate(internalServerPath, projectPath, routesFile, "routes")
	if err != nil {
		log.Printf("Failed to create and write to routes file at path %s: %v\n", internalServerPath, err)
		return err
	}

	err = goFormat(projectPath)
	if err != nil {
		log.Printf("Failed to run 'gofmt' in the created project %v\n", err)
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
	case "routes":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templateGen.Routes())))
		err = createdTemplate.Execute(createdFile, p)
	}

	if err != nil {
		return err
	}

	return nil
}

// isValidWebFramework check if the input is supported or not
func IsValidWebFramework(input string) bool {
	fmt.Println(input)
	for _, t := range SupportedWebframeworks {
		if input == t {
			return true
		}
	}
	return false
}

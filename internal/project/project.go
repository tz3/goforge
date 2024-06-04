package project

import (
	"fmt"
	"html/template"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	tpl "github.com/tz3/goforge/internal/templates"
)

type ProjectConfig struct {
	ProjectName  string
	ProjectType  string                  // Can be an api or serverless.
	FrameworkMap map[string]WebFramework // Can be any of the suggested router Packages.
	Exit         bool
	AbsolutePath string
}

type WebFramework struct {
	packageName string
	templater   TemplateGenerator
}

type TemplateGenerator interface {
	Main() []byte
	Server() []byte
	Routes() []byte
}

// Router Packages
const (
	chiPackage         = "github.com/go-chi/chi/v5"
	gorillaPackage     = "github.com/gorilla/mux"
	routerPackage      = "github.com/julienschmidt/httprouter"
	ginPackage         = "github.com/gin-gonic/gin"
	fiberPackage       = "github.com/gofiber/fiber/v2"
	cmdApiPath         = "cmd/api"
	internalServerPath = "internal/server"
	mainFile           = "main.go"
	serverFile         = "server.go"
	routesFile         = "routes.go"
)

func (p *ProjectConfig) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		tprogram.ReleaseTerminal()
		os.Exit(1)
	}
}

func (p *ProjectConfig) createFrameworkMap() {
	p.FrameworkMap["standard lib"] = WebFramework{
		packageName: "",
		templater:   tpl.StandardLibraryRouteTemplate{},
	}

	p.FrameworkMap["chi"] = WebFramework{
		packageName: chiPackage,
		templater:   tpl.ChiRouteTemplate{},
	}

	p.FrameworkMap["gin"] = WebFramework{
		packageName: ginPackage,
		templater:   tpl.GinRouteTemplate{},
	}

	p.FrameworkMap["fiber"] = WebFramework{
		packageName: fiberPackage,
		templater:   tpl.FiberRouteTemplate{},
	}

	p.FrameworkMap["gorilla/mux"] = WebFramework{
		packageName: gorillaPackage,
		templater:   tpl.GorillaRouteTemplate{},
	}

	p.FrameworkMap["httpRouter"] = WebFramework{
		packageName: routerPackage,
		templater:   tpl.HttpRouterRouteTemplate{},
	}
}

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
			fmt.Printf("Error creating root project directory %v\n", err)
		}
	}

	projectPath := fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)

	// create the router based on user input
	p.createFrameworkMap()

	err := initGoMod(p.ProjectName, projectPath)
	if err != nil {
		cobra.CheckErr(err)
	}

	// we need to install the correct package
	if p.ProjectType != "standard lib" {
		err = goGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
		if err != nil {
			cobra.CheckErr(err)
		}
	}

	err = p.CreatePath(cmdApiPath, projectPath)
	if err != nil {
		cobra.CheckErr(err)
	}

	err = p.CreateFileAndWriteTemplate(cmdApiPath, projectPath, mainFile, "main")
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
	makeFileTemplate := template.Must(template.New("makefile").Parse(string(tpl.MakeTemplate())))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
		return err
	}

	err = p.CreatePath(internalServerPath, projectPath)
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileAndWriteTemplate(internalServerPath, projectPath, serverFile, "server")
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileAndWriteTemplate(internalServerPath, projectPath, routesFile, "routes")
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	return nil
}

// cmd/api
func (p *ProjectConfig) CreatePath(pathToCreate string, projectPath string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", projectPath, pathToCreate)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", projectPath, pathToCreate), 0751)
		if err != nil {
			fmt.Printf("Error creating path directory %v\n", err)
			return err
		}
	}

	return nil
}

func (p *ProjectConfig) CreateFileAndWriteTemplate(pathToCreate string, projectPath string, fileName string, methodName string) error {
	createdFile, err := os.Create(fmt.Sprintf("%s/%s/%s", projectPath, pathToCreate, fileName))
	if err != nil {
		return err
	}

	defer createdFile.Close()

	// inject template
	switch methodName {
	case "main":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Main())))
		err = createdTemplate.Execute(createdFile, p)
	case "server":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Server())))
		err = createdTemplate.Execute(createdFile, p)
	case "routes":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Routes())))
		err = createdTemplate.Execute(createdFile, p)
	}

	if err != nil {
		return err
	}

	return nil
}

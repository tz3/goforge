package project

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"

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
	chiPackage     = "github.com/go-chi/chi/v5"
	gorillaPackage = "github.com/gorilla/mux"
	routerPackage  = "github.com/julienschmidt/httprouter"
	ginPackage     = "github.com/gin-gonic/gin"
	fiberPackage   = "github.com/gofiber/fiber/v2"
)

func (p *ProjectConfig) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		tprogram.ReleaseTerminal()
		os.Exit(1)
	}
}

func executeCmd(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

func initGoMod(projectName string, appDir string) {
	if err := executeCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		cobra.CheckErr(err)
	}
}

func goGetPackage(appDir, packageName string) {
	fmt.Printf("Package name is: %s\n", packageName)
	if err := executeCmd("go",
		[]string{"get", "-u", packageName},
		appDir); err != nil {
		cobra.CheckErr(err)
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

	// we need to create a go mod init
	initGoMod(p.ProjectName, projectPath)

	// create the router based on user input
	p.createFrameworkMap()

	// we need to install the correct package
	if p.ProjectType != "standard lib" {
		goGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
	}

	// create /cmd/api
	if _, err := os.Stat(fmt.Sprintf("%s/cmd/api", projectPath)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/cmd/api", projectPath), 0751)
		if err != nil {
			fmt.Printf("Error creating directory %v\n", err)
		}
	}

	mainFile, err := os.Create(fmt.Sprintf("%s/cmd/api/main.go", projectPath))
	if err != nil {
		return err
	}

	defer mainFile.Close()

	// inject template
	mainTemplate := template.Must(template.New("main").Parse(string(p.FrameworkMap[p.ProjectType].templater.Main())))
	err = mainTemplate.Execute(mainFile, p)
	if err != nil {
		fmt.Printf("this is the err %v\n", err)
		return err
	}

	makeFile, err := os.Create(fmt.Sprintf("%s/Makefile", projectPath))
	if err != nil {
		return err
	}

	defer makeFile.Close()

	// inject makefile template
	makeFileTemplate := template.Must(template.New("makefile").Parse(string(tpl.MakeTemplate())))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
		return err
	}

	// create /internal/server
	if _, err := os.Stat(fmt.Sprintf("%s/internal/server", projectPath)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/internal/server", projectPath), 0751)
		if err != nil {
			fmt.Printf("Error creating directory %v\n", err)
		}
	}

	serverFile, err := os.Create(fmt.Sprintf("%s/internal/server/server.go", projectPath))
	if err != nil {
		return err
	}

	serverFileTemplate := template.Must(template.New("server").Parse(string(p.FrameworkMap[p.ProjectType].templater.Server())))
	err = serverFileTemplate.Execute(serverFile, p)
	if err != nil {
		return err
	}

	defer serverFile.Close()

	routesFile, err := os.Create(fmt.Sprintf("%s/internal/server/routes.go", projectPath))
	if err != nil {
		return err
	}

	routesFileTemplate := template.Must(template.New("routes").Parse(string(p.FrameworkMap[p.ProjectType].templater.Routes())))
	err = routesFileTemplate.Execute(routesFile, p)
	if err != nil {
		return err
	}

	defer routesFile.Close()

	return nil
}

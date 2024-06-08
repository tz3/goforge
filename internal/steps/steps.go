// Package steps defines the steps involved in setting up a Go project.
package steps

import "github.com/tz3/goforge/cmd/ui/textinput"

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
	Title string
	Desc  string
}

// Steps is a collection of StepSchema. It represents all the steps
// involved in the setup process.
type Steps struct {
	Steps []StepSchema
}

// Options represents the options for initializing the steps.
// It includes the name of the project and the type of the project.
type Options struct {
	ProjectName *textinput.Output
	ProjectType string
}

// InitSteps initializes the steps of the setup process.
// It takes an Options struct as input and returns a pointer to a Steps struct.
// The Steps struct includes all the steps involved in the setup process.
func InitSteps(options *Options) *Steps {
	steps := &Steps{
		[]StepSchema{
			{
				StepName: "Go Project Framework",
				Options: []Option{
					{
						Title: "standard lib",
						Desc:  "Built in standard golang library",
					},
					{
						Title: "chi",
						Desc:  "use go-chi from: https://github.com/go-chi/chi",
					},
					{
						Title: "gin",
						Desc:  "use gin-gonic from: https://github.com/gin-gonic/gin",
					},
					{
						Title: "fiber",
						Desc:  "use gofiber from: https://github.com/gofiber/fiber",
					},
					{
						Title: "gorilla/mux",
						Desc:  "use gorilla/mux from: https://github.com/gorilla/mux",
					},
					{
						Title: "httpRouter",
						Desc:  "use julienschmidt/httprouter from: https://github.com/julienschmidt/httprouter",
					},
					{
						Title: "echo",
						Desc:  "use echo from: https://github.com/labstack/echo",
					},
				},
				Headers: "What framework do you want to use in your Go project?",
				Field:   &options.ProjectType,
			},
		},
	}

	return steps
}

package steps

import "github.com/tz3/goforge/cmd/ui/textinput"

type StepSchema struct {
	StepName string
	Options  []Option
	Headers  string
	Field    *string
}

type Option struct {
	Title string
	Desc  string
}

type Steps struct {
	Steps []StepSchema
}

type Options struct {
	ProjectName *textinput.Output
	ProjectType string
}

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
				},
				Headers: "What framework do you want to use in your Go project?",
				Field:   &options.ProjectType,
			},
		},
	}

	return steps
}

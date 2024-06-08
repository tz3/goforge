// Package template provides a set of templates for the main function, HTTP server, README, and Makefile.
package template

func ReadmeTemplate() []byte {
	return []byte(`# Project {{.ProjectName}}

A brief description of what this project does and who it's for.

## Installation 

Instructions on how to install and get the project running on local machine.

## Usage 

Examples of how to use this project or code, including any relevant code examples or screenshots.

## Running Tests

Instructions on how to run any testing frameworks you've used.

## Deployment

Notes on how to deploy this on a live or production system.

## Built With

Any frameworks, libraries or APIs used.

## Contributing

Details on how others can contribute to your project.

## License

Any licensing for the project.

## Authors

Who has contributed to this project.

## Acknowledgments

Any acknowledgments or credits.
`)
}

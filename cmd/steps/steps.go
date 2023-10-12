package steps

import textinput "github.com/melkeydev/go-blueprint/cmd/ui/textinput"

type StepSchema struct {
	StepName string
	Options  []string
	Headers  string
	Field    *string
}

type Steps struct {
	Steps []StepSchema
}

type Options struct {
	ProjectName *textinput.Output
	GitRepo     string
	ProjectType string
}

func InitSteps(options *Options) *Steps {
	steps := &Steps{
		[]StepSchema{
			{
				StepName: "Git Repo",
				Options:  []string{"Yes", "No"},
				Headers:  "Do you want to init a git project?",
				Field:    &options.GitRepo,
			},
			{
				StepName: "Project Type",
				Options:  []string{"API Server", "Serverless Lambda"},
				Headers:  "What kind of Go project are you building?",
				Field:    &options.ProjectType,
			},
		},
	}

	return steps
}

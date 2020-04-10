package models

//Pipeline is chain of jobs
type Pipeline struct {
	Name        string
	Description string
	Tags        Tagline
	Jobs        []Job
}

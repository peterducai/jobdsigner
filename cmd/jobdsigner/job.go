package jobdsigner

//Job defines task/job
type Job struct {
	ID          int64
	Handler     func(string) error
	Title       string
	Description string
	DependsOn   []string
	Args        string
	Interaction string
}

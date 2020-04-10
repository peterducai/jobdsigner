package models

//Pool pool structure
type Pool struct {
	Pipelines []Pipeline
}

//PipelinePool is pool of pipelines
var PipelinePool Pool

package agent

/*
Tool is an interface that can be implemented to provide additional
functionality to a large language model.
*/
type Tool interface {
	Use() string
}

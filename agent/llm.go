package agent

/*
LLM is an interface to implement when you want to wrap a large language model.
*/
type LLM interface {
	Predict([]map[string]string) chan string
}

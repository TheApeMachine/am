package tool

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

type Search struct {
	input  []string
	output string
}

func NewSearch(input string) *Search {
	errnie.Trace()
	return &Search{strings.Split(input, " "), ""}
}

func (search *Search) Use() string {
	errnie.Trace()

	key := tweaker.GetString("gcp.searchKey")
	cx := tweaker.GetString("gcp.engineID")

	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
	}

	svc, err := customsearch.New(client)
	if errnie.Handles(err) != nil {
		return err.Error()
	}

	resp, err := svc.Cse.List().Cx(cx).Q(
		strings.Join(search.input, " "),
	).Do()

	if errnie.Handles(err) != nil {
		return err.Error()
	}

	for i, result := range resp.Items {
		search.output += fmt.Sprintf("#%d: %s\n", i+1, result.Title)
		search.output += fmt.Sprintf("%s\n", result.Snippet)
	}

	return search.output
}

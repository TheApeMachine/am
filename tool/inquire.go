package tool

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/wrk-grp/errnie"
)

type Inquire struct {
	input []string
}

func NewInquire(input string) *Inquire {
	return &Inquire{strings.Split(input, " ")}
}

func (inquire *Inquire) Use() string {
	var (
		input string
		err   error
	)

	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)

	if input, err = reader.ReadString('\n'); errnie.Handles(err) != nil {
		return err.Error()
	}

	return input
}

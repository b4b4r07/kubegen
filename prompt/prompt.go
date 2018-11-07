package prompt

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

// Prompt is
type Prompt struct {
	Title           string
	PreEnteredValue string
	DefaultValue    string
}

// New is
func New(title, defaultValue string, optional ...string) Prompt {
	preEnteredValue := ""
	if len(optional) > 0 {
		preEnteredValue = optional[0]
	}
	return Prompt{
		Title:           title,
		DefaultValue:    defaultValue,
		PreEnteredValue: preEnteredValue,
	}
}

// Run is
func (p Prompt) Run() (string, error) {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          p.Title,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		Stdout:          os.Stderr,
	})
	if err != nil {
		return "", err
	}
	defer l.Close()

	var line string
	for {
		if p.PreEnteredValue == "" {
			line, err = l.Readline()
		} else {
			line, err = l.ReadlineWithDefault(p.PreEnteredValue)
		}
		if err == readline.ErrInterrupt {
			if len(line) <= len(p.PreEnteredValue) {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			if p.DefaultValue == "" {
				continue
			}
			line = p.DefaultValue
		}
		return line, nil
	}
	return "", errors.New("canceled")
}

package cmd

import (
	"encoding/base64"
	"html/template"
	"os"

	"github.com/b4b4r07/kubegen/prompt"
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Generate secret manifest",
	Long:  "Generate secret manifest",
	RunE:  secretGenerator,
}

// Secret is
type Secret struct {
	Name   string
	Keys   []string
	Values []string
}

const secretTemplate = `apiVersion: v1
kind: Secret
metadata:
  name: {{.Name}}
data:
{{range $i, $v := .Keys}}  {{$v}}: {{index $.Values $i}}
{{end}}
`

func secretGenerator(cmd *cobra.Command, args []string) error {
	tmpl, err := template.New("").Parse(secretTemplate)
	if err != nil {
		return err
	}

	name, err := prompt.New("Name> ", "SECRET_NAME").Run()
	if err != nil {
		return err
	}

	var keys, values []string
	for {
		key, err := prompt.New("Key (Skip with Enter)> ", "default").Run()
		if err != nil {
			break
		}
		if key == "default" {
			break
		}
		keys = append(keys, key)
		value, err := prompt.New("Value> ", "").Run()
		if err != nil {
			break
		}
		encodedValue := base64.StdEncoding.EncodeToString([]byte(value))
		values = append(values, encodedValue)
	}

	if len(keys) == 0 {
		keys = append(keys, "KEY")
		values = append(values, "ENCORDED_VALUE")
	}

	return tmpl.Execute(os.Stdout, &Secret{
		Name:   name,
		Keys:   keys,
		Values: values,
	})
}

func init() {
	RootCmd.AddCommand(secretCmd)
}

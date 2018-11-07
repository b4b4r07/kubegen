package cmd

import (
	"html/template"
	"os"

	"github.com/b4b4r07/kubegen/prompt"
	"github.com/spf13/cobra"
)

var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "Generate Pod manifest",
	Long:  "Generate Pod manifest",
	RunE:  podGenerator,
}

// Pod is
type Pod struct {
	Name  string
	Image string
}

const podTemplate = `apiVersion: v1
kind: Pod
metadata:
  labels:
    run: {{.Name}}
  name: {{.Name}}
spec:
  containers:
  - name: {{.Name}}
    image: {{.Image}}
    imagePullPolicy: IfNotPresent
  dnsPolicy: ClusterFirst
  restartPolicy: Never
`

func podGenerator(cmd *cobra.Command, args []string) error {
	tmpl, err := template.New("").Parse(podTemplate)
	if err != nil {
		return err
	}

	name, err := prompt.New("Name> ", "POD_NAME").Run()
	if err != nil {
		return err
	}
	image, err := prompt.New("Image> ", "IMAGE_NAME").Run()
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, &Pod{
		Name:  name,
		Image: image,
	})
}

func init() {
	RootCmd.AddCommand(podCmd)
}

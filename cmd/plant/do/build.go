package do

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"automatica.team/plant/cmd/plant/tmpl"
)

func Dockerfile(project string) (func(), error) {
	path := filepath.Join(project, "Dockerfile")

	fmt.Println("[🐳] Generating Dockerfile")
	{
		var buf bytes.Buffer
		if err := tmpl.Build.ExecuteTemplate(&buf, "Dockerfile", nil); err != nil {
			return nil, err
		}
		if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
			return nil, err
		}
	}

	return func() { os.Remove(path) }, nil
}

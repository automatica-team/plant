package do

import (
	"fmt"

	"automatica.team/plant/cmd/plant/exec"
)

func Get(path string) error {

	fmt.Println("[📦] Getting " + path)

	return exec.Run("go", "get", path)
}

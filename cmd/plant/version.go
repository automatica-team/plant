package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	verPlant   = "v0.1"
	verTelebot = "v3.3.8"
)

func Version(c *cobra.Command, _ []string) error {
	fmt.Printf("%s\n\nTelebot: %s\n", Header, verTelebot)

	resp, err := http.Get("https://api.github.com/repos/automatica-team/plant/tags")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var tags []struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return err
	}

	if len(tags) > 0 {
		fmt.Printf("Latest: %s\n\n", tags[0].Name)
		fmt.Printf(`Use "plant upgrade" to upgrade to the latest version.`)
	}

	return nil
}

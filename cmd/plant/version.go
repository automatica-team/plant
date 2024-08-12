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
	fmt.Printf(Header + "\n\n")

	tagsPlant, err := githubTags("automatica-team/plant")
	if err != nil {
		return err
	}

	latestPlant := verPlant
	if len(tagsPlant) > 0 {
		latestPlant = tagsPlant[0]
	}

	tagsTelebot, err := githubTags("go-telebot/telebot")
	if err != nil {
		return err
	}

	latestTelebot := verTelebot
	if len(tagsTelebot) > 1 {
		latestTelebot = tagsTelebot[1]
	}

	fmt.Printf("Current version:\n  Plant:   %s\n  Telebot: %s\n\n", verPlant, verTelebot)
	fmt.Printf("Latest version:\n  Plant:   %s\n  Telebot: %s\n\n", latestPlant, latestTelebot)
	fmt.Println(`Use "plant upgrade" to upgrade to the latest version.`)

	return nil
}

func githubTags(repo string) ([]string, error) {
	resp, err := http.Get("https://api.github.com/repos/" + repo + "/tags")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tags []struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, err
	}

	var names []string
	for _, tag := range tags {
		names = append(names, tag.Name)
	}

	return names, nil
}

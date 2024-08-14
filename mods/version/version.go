package version

import (
	"encoding/json"
	"net/http"

	"automatica.team/plant"
	tele "gopkg.in/telebot.v3"
)

func (mod *Version) Name() string {
	return "plant/version"
}

type Version struct {
	plant.Handler
	b    *plant.Bot `plant:"bot"`
	repo string
}

type TagData struct {
	Repo string
	Tags []string
}

func New() *Version {
	return &Version{}
}

func (mod *Version) Import(v plant.V) error {
	mod.Handle("/tags", mod.onTags)
	mod.repo = v.GetString("repo")
	return nil
}

func (mod *Version) onTags(c tele.Context) error {
	repo := mod.repo
	tags, err := mod.GithubTags(repo)
	if err != nil {
		return c.Send(mod.b.Text(c, "tags"))
	}

	tagData := TagData{
		Repo: repo,
		Tags: tags,
	}

	if len(tags) == 0 {
		return c.Send(mod.b.Text(c, "tags", tagData))
	}

	return c.Send(mod.b.Text(c, "tags", tagData))
}

func (mod *Version) GithubTags(repo string) ([]string, error) {
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

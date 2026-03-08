package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	hello "github.com/g3org3/hello_goo"

	"github.com/charmbracelet/huh"
)

type githubUser struct {
	Name  string `json:"name"`
	Login string `json:"login"`
}

func fetchGitHubName(username string) (string, error) {
	resp, err := http.Get("https://api.github.com/users/" + username)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var user githubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", err
	}

	if user.Name != "" {
		return user.Name, nil
	}
	return user.Login, nil
}

func main() {
	var username string
	var style string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What's your GitHub username?").
				Placeholder("g3org3").
				Value(&username),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Greeting style").
				Options(
					huh.NewOption("Casual", "casual"),
					huh.NewOption("Formal", "formal"),
				).
				Value(&style),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Cancelled.")
		os.Exit(0)
	}

	name, err := fetchGitHubName(username)
	if err != nil {
		fmt.Printf("Could not fetch GitHub profile for %q, using username instead.\n", username)
		name = username
	}

	switch style {
	case "formal":
		fmt.Println(hello.Formal(name))
	default:
		fmt.Println(hello.Hello(name))
	}
}

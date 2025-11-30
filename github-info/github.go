package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	loginUrl := "https://api.github.com/users/vrashabh-sontakke"
	name, repos, err := githubInfo(loginUrl)
	if err != nil {
		log.Fatal("error:", err)
	} else {
		fmt.Printf("Name:%v\nRepos:%v\n", name, repos)
	}
}

func githubInfo(login string) (string, int, error) {
	resp, err := http.Get(login)
	if err != nil {
		log.Fatal("error getting response:", err)
		return "", 0, err
	}

	var r Info

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatal("error decoding resp.body:", err)
		return "", 0, err
	}

	return r.Name, r.Public_Repos, nil

}

type Info struct {
	Name         string
	Public_Repos int
}

package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"math/rand"
	"strconv"
	"encoding/json"
)

const (
	apiUrlBase string = "https://pokeapi.co/api/v2/pokemon/"
	wikiUrlBase string = "https://bulbapedia.bulbagarden.net/wiki/"
)

type pokemon struct {
    Name string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(807)

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(r))
	if err != nil {
		log.Fatalln(err)
	}
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	p := pokemon{}
	json.Unmarshal([]byte(body), &p)

	n := "Today’s Pokémon is: " + strings.Title(p.Name) + "\nLink: " + wikiUrlBase + p.Name
	fmt.Println(n)
}
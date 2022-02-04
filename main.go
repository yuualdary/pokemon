package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`

	//json format
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`

	//sub json (pokemon) format
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

//agar bisa diakses dengan custom endpoint menggunakan gin
func test(c *gin.Context) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://pokeapi.co/api/v2/pokedex/kanto/", nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	//mengambil api dari pokeapi

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	//format json
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)

	c.JSON(http.StatusOK, responseObject)
	return
}

func main() {

	router := gin.Default()

	router.GET("/api/all", test)

	//endpoint

	router.Run(":8000")

}

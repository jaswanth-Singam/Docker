/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kyokomi/emoji/v2"

	"github.com/spf13/cobra"
)

type RandomJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Random joke from the package Docker",
	Long:  `Get a Random joke from the package Docker`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

// getRandomJoke func use the method getJokeData to get JSON and umarshal into string to print in the screen
func getRandomJoke() {
	fmt.Println("Here is your Joke")
	emoji.Println(":beer::beer:Beer!!!")
	url := "https://icanhazdadjoke.com/"

	responseBytes := getJokeData(url)

	joke := RandomJoke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	fmt.Println(string(joke.Joke))
}

// getJokeData connect the external URL and get a JSON respose from the API
func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)
	if err != nil {
		log.Printf("Could not request a dadjoke. %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Ranjith KA (https://github.com/ranjith-ka/Docker)")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}
	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

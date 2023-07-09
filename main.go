package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	openai "github.com/sashabaranov/go-openai"
)

// FOR TESTING
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: homePage")

	resp, err := getResponseFromChatGpt("say you are testing this endpoint")

	if (err != nil) {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, resp)
}

func getResponseFromChatGpt(prompt string) (string, error){
	openAiApiKey := os.Getenv("OPENAI_API_KEY")
    fmt.Println(openAiApiKey)

	client := openai.NewClient(openAiApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
 

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", homePage)

    log.Fatal(http.ListenAndServe(":3259", router))
}

func main() {

	fmt.Println(os.Getenv("OPENAI_API_KEY"))

    handleRequests()
}
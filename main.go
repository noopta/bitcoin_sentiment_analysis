package main

import (
    "context"
    "fmt"
    // "net/http"
    "os"
	// "strings"
    // "reflect"
    // "encoding/json"	
    openai "github.com/sashabaranov/go-openai"
	// "time"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func getTweets() {
    scraper := twitterscraper.New()
	err := scraper.Login("Anupta2", "Spartan1")

	if err != nil {
		fmt.Println("yo")
		fmt.Println(err)
		return
	}

    for tweet := range scraper.SearchTweets(context.Background(),
        "twitter scraper data -filter:retweets", 50) {
        if tweet.Error != nil {
            panic(tweet.Error)
        }
        fmt.Println(tweet.Text)
    }
}

func callGPT() {
    fmt.Println("calling Chat GPT")
    fmt.Println()
	
	vaar strArr[10]strings

	strings[0] = "Billionaire Paul Tudor Jones:

	“[Bitcoin]’s the only thing that humans can’t adjust the supply in…I’m going to always stick with it.”"
    
	client := openai.NewClient(os.Getenv("OPEN_API_KEY"))
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model: openai.GPT3TextDavinci003,
			// Messages: []openai.ChatCompletionMessage{
			// 	{
			// 		Role:    openai.ChatMessageRoleUser,
			// 		Content: "Describe the view and quality of the seats at section 323 and row 15 for Chicago Bulls home games from 2021",
			// 	},
			// },
			Prompt: "Given three options, describe the following text as either positive, neutral, or negative: I am eager to learn and hope bitcoin price goes up!",
			MaxTokens: 50,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	
	responseText = resp.Choices[0].Text

	if (responseText == "positive") {

	} else if (responseText == "negative") {

	} else {

	}
}

func main() {
    callGPT()
	// getTweets()
}

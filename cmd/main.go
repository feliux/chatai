package main

import (
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/subosito/gotenv"
)

var (
	endpoint = "https://api.openai.com/v1"
	apiKey   = os.Getenv("OPENAI_API_KEY")
)

func init() {
	if err := gotenv.Load(); err != nil {
		log.Fatalf("error ocurred reading .env file %s", err)
	}
}

func main() {
	keyCredential := azcore.NewKeyCredential(apiKey)

	// NOTE: this constructor creates a client that connects to the public OpenAI endpoint.
	// To connect to an Azure OpenAI endpoint, use azopenai.NewClient() or azopenai.NewClientWithyKeyCredential.
	client, err := azopenai.NewClientForOpenAI(endpoint, keyCredential, nil)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Fatalf("ERROR: %s", err)
	}

	_ = client
}

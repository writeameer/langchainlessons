package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	// Create a new Ollama LLM instance

	llm, err := ollama.New(ollama.WithServerURL("http://localhost:11434"), ollama.WithModel("llama3.2:latest"))
	if err != nil {
		panic(err)
	}

	// Send a prompt
	ctx := context.Background()
	response, err := llm.Call(ctx, "Hello World")
	if err != nil {
		panic(err)
	}

	// Output the response
	fmt.Println(response)
}

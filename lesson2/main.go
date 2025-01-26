package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/ollama"
)

// DatePlugin simulates a function that returns the current date
func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func main() {
	// Initialize Ollama LLM
	llm, err := ollama.New(ollama.WithModel("llama3.2:latest"))
	if err != nil {
		log.Fatalf("failed to initialize Ollama LLM: %v", err)
	}

	// Create a LangChain agent with the LLM
	agent := agents.New(llm)

	// Define function calling behavior (Auto function selection)
	agent.RegisterFunction("GetCurrentDate", GetCurrentDate)

	// Create a context
	ctx := context.Background()

	// Invoke the prompt
	response, err := agent.Run(ctx, "What's today's date?")
	if err != nil {
		log.Fatalf("failed to generate response: %v", err)
	}

	// Output the response
	fmt.Println(response)
}

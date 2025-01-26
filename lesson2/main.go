package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// FunctionCall represents the structure of a function call from the model
type FunctionCall struct {
	Tool      string                 `json:"tool"`
	ToolInput map[string]interface{} `json:"tool_input"`
}

// getCurrentDate returns the current date in YYYY-MM-DD format
func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// handleFunctionCall processes the function call and returns the result
func handleFunctionCall(call FunctionCall) (string, error) {
	switch call.Tool {
	case "GetCurrentDate":
		return getCurrentDate(), nil
	default:
		return "", fmt.Errorf("unknown function: %s", call.Tool)
	}
}

func main() {
	// Initialize the Ollama LLM with the desired model and JSON output format
	llm, err := ollama.New(
		ollama.WithModel("llama3.2:latest"),
		ollama.WithFormat("json"),
	)
	if err != nil {
		log.Fatalf("failed to initialize Ollama LLM: %v", err)
	}

	// Define the system message with available functions
	systemMessage := `{"functions": [{"name": "GetCurrentDate", "description": "Returns the current date in YYYY-MM-DD format"}]}`

	// Create the initial messages
	var messages []llms.MessageContent

	messages = append(messages, llms.TextParts(llms.ChatMessageTypeSystem, systemMessage))
	messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, "What's today's date?"))

	// Create a context
	ctx := context.Background()

	// Generate a response from the LLM
	response, err := llm.GenerateContent(ctx, messages)
	fmt.Println("Model's Response:", response.Choices[0].Content)

	if err != nil {
		log.Fatalf("failed to generate response: %v", err)
	}

	// Parse the model's response
	var functionCall FunctionCall
	if err := json.Unmarshal([]byte(response.Choices[0].Content), &functionCall); err != nil {
		log.Fatalf("failed to parse function call: %v", err)
	}

	// Handle the function call
	result, err := handleFunctionCall(functionCall)
	if err != nil {
		log.Fatalf("failed to handle function call: %v", err)
	}

	// Output the result
	fmt.Println("Model's Response:", response.Choices[0].Content)
	fmt.Println("Function Result:", result)
}

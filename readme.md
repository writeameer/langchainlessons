# Learning LangChain with Ollama and go


This repository contains a series of lessons demonstrating how to use LangChainGo with Ollama, an AI development kit for building applications powered by large language models (LLMs).

## Install Ollama

Download [Ollama](https://ollama.com/) for your operating system. Start Ollama to host the `llama3.2` model for your applications:

```
ollama run llama3.2
```

## Lessons

Each `lessonxx` folder contains an independent Go project demonstrating a specific LangChainGo feature.

| Lesson | Description |
|--------|-------------|
| [Lesson 01](/lesson01/readme.md) | Basic LLM call with Ollama (Hello World) |
| [Lesson 02](/lesson02/readme.md) | Function calling with LangChainGo |
| [Lesson 03](/lesson03/readme.md) | Function calling with parameters |

## Notes
- The `llama3.2` model runs locally on Ollama.
- Code samples are in **Go**.
- Each lesson is self-contained and can be run independently.

Get started by trying [Lesson 01](./lesson01/readme.md)!


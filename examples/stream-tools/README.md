# Streaming Tools Agent

Real-time agent that executes tools during streaming responses. Demonstrates building an interactive agent with the Inference Gateway SDK.

## Quick Start

```bash
# Set environment (optional)
export LLM_PROVIDER="deepseek"
export LLM_MODEL="deepseek-chat"

# Run the agent
go run agent.go
```

## Features

-   **Real-time execution**: Tools execute immediately when JSON is complete
-   **Multi-tool support**: Weather lookup + calculator
-   **Iterative conversations**: Queue-based agent pattern
-   **Streaming reasoning**: See the AI's thinking process

## Available Tools

-   `get_current_weather`: Get weather for major cities
-   `calculate`: Basic arithmetic (add, subtract, multiply, divide)

## Example

```
👤 User: What's the weather in San Francisco and what's 15 multiplied by 7?

🤖 Assistant:
🔧 Executing: get_current_weather({"location": "san francisco"})
📋 Result: {"temperature":14,"unit":"celsius","description":"Foggy"}
🔧 Executing: calculate({"a": 15, "b": 7, "operation": "multiply"})
📋 Result: {"operation":"15.00 multiply 7.00","result":105}

The weather in San Francisco is foggy at 14°C.
The result of 15 × 7 is 105.
```

## Agent Pattern

1. Process user message queue
2. Stream LLM response with tools
3. Execute tools when JSON is complete
4. Add results to conversation history
5. Continue until no more tool calls

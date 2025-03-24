# Content Generation Example

This example demonstrates how to use the Inference Gateway SDK to generate content from Language Models (LLMs) using various providers.

## What This Example Shows

-   How to create a client connection to the Inference Gateway
-   How to construct a conversation with system and user messages
-   How to generate content using the SDK's `GenerateContent` method
-   How to handle and display the LLM response
-   How to access usage statistics (token counts)

## Running the Example

```sh
# Set your API URL (optional)
export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"

# Set your preferred provider (optional, default: openai)
export LLM_PROVIDER="openai"

# Set your preferred model (optional, default: gpt-4o)
export LLM_MODEL="gpt-4o"

# Run the example
go run [main.go](main.go)
```

## Example Output

```sh
Generating content using openai gpt-4o...

Model: gpt-4o
Response: Goroutines and threads are both mechanisms used for concurrent execution, but they have several key differences:

1. **Resource Usage**:

    - Goroutines are extremely lightweight, consuming only about 2KB of stack space initially.
    - Threads are more heavyweight, typically requiring 1-2MB of memory per thread.

2. **Creation Overhead**:

    - Goroutines can be created in microseconds.
    - Thread creation has higher overhead, often taking milliseconds.

3. **Scheduling**:

    - Goroutines are scheduled by Go's runtime scheduler, which implements cooperative multitasking.
    - Threads are scheduled by the operating system kernel, which uses preemptive scheduling.

4. **Communication**:

    - Goroutines are designed to communicate through channels (following the "don't communicate by sharing memory; share memory by communicating" principle).
    - Threads typically communicate through shared memory and locks.

5. **Scalability**:

    - You can easily run thousands or even millions of goroutines concurrently.
    - Most systems struggle with more than a few thousand threads.

6. **Context Switching**:

    - Context switching between goroutines is faster as it's handled by the Go runtime.
    - Context switching between threads is more expensive as it involves the OS kernel.

7. **Parallelism**:

    - Goroutines can achieve parallelism when the Go program uses multiple OS threads (controlled by GOMAXPROCS).
    - Threads can run in parallel directly on multiple CPU cores.

8. **Stack Size**:

    - Goroutines have dynamically sized stacks that can grow and shrink as needed.
    - Threads typically have a fixed stack size determined at creation time.

9. **Platform Dependency**:
    - Goroutines function the same way across all platforms supported by Go.
    - Thread implementation details can vary across operating systems.

Because of these advantages, especially regarding resource usage and scalability, Go programs can handle highly concurrent workloads efficiently using goroutines where traditional thread-based approaches might struggle.

Usage Statistics:
Prompt tokens: 84
Completion tokens: 436
Total tokens: 520
```

## How It Works

1. The example creates a client connected to the Inference Gateway API
2. It creates a context with a timeout to manage the request lifetime
3. It defines a conversation with system and user messages
4. It calls the GenerateContent method to get a response from the LLM
5. It prints the model name and response content
6. If available, it prints token usage statistics

## Customization

You can modify this example to:

-   Change the system prompt to give the LLM different instructions
-   Change the user prompt to ask different questions
-   Use different providers and models through environment variables
-   Add additional parameters like temperature or max tokens

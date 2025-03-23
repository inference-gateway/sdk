# Streaming Content Generation Example

This example demonstrates how to use the Inference Gateway SDK to stream content from LLMs. Streaming allows you to receive and display content tokens as they're generated, rather than waiting for the entire response.

## What This Example Shows

-   How to set up a streaming request with the SDK
-   How to process streaming events as they arrive
-   How to handle different event types in the stream
-   How to concatenate content parts into a complete response

## Running the Example

```sh
# Set your API URL (optional)
export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"

# Set your preferred provider (optional, default: ollama)
export LLM_PROVIDER="ollama"

# Set your preferred model (optional, default: llama2)
export LLM_MODEL="llama2"

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

1. The example creates a client and prepares a conversation with system and user messages
2. It sends the request to the specified LLM provider and model
3. Upon receiving the response, it displays the model's answer
4. It also shows usage statistics such as token counts for the prompt and completion

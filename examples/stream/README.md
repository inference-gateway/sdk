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
export LLM_PROVIDER="groq"

# Set your preferred model (optional, default: llama2)
export LLM_MODEL="qwen-2.5-coder-32b"

# Run the example
go run [main.go](main.go)
```

## Example Output

```markdown
Generating content using groq qwen-2.5-coder-32b...

Model: qwen-2.5-coder-32b
Response: Goroutines and threads are both units of execution in concurrent programming, but they operate in significantly different ways, especially in terms of their implementation and the environment in which they are used.

### Threads

1. **Resource Heavy**: Threads are supported at the system level and thus have a significant overhead in terms of memory and processing power. Each thread typically needs to allocate a stack of a certain size (often in the range of 1-8 MB), which can add up to considerable memory consumption, especially when dealing with a large number of concurrent threads.

2. **Managed by OS**: Threads are managed by the operating system. This means that the OS scheduler handles the execution and switching of threads, which can lead to varying performance depending on the OS and the specific workload.

3. **Slower Context Switch**: Context switching between threads can be relatively slow due to the resources involved; this includes saving and restoring the thread's state, and handling the delicate process of stopping and restarting a thread's execution.

4. **Concurrency Control Complexity**: Threads require careful management to avoid issues like race conditions, deadlocks, and starvation. Programming with threads can therefore become complex, as one must employ synchronization mechanisms provided by the language (e.g., mutexes, semaphores).

### Goroutines

1. **Lightweight**: Goroutines are managed by the Go runtime (also known as the Go Scheduler) and are not directly backed by any OS-level threads. They start off with a small stack (approximately 2 KB) and can grow or shrink dynamically based on the needs of the running program. This makes goroutines very lightweight and efficient, allowing for the concurrent execution of thousands if not millions of goroutines.

2. **Managed by Go Runtime**: Goroutines are scheduled to run on a pool of so-called logical processors (which, in turn, are mapped to OS-level threads). The Go runtime handles the execution and context switching between goroutines, which is more efficient and faster than OS-level context switching.

3. **Faster Context Switch**: Because goroutines are managed by the Go runtime and operate within the user space, the cost of context switching between them is much lower compared to threads. This is one of the main advantages of using goroutines in Go for high concurrency.

4. **Ease of Use**: Goroutines abstract away much of the complexity involved in concurrent programming. There are no locks or mutexes; instead, goroutines use channels to communicate and synchronize, which leads to code that is often simpler and less prone to race conditions. The error handling mechanism provided by Go, such as deferred function calls and panic/recover, further aids in managing complex concurrent programs.

### Summary

In summary, the key differences are in the overhead and control of execution:

-   **Efficiency and Scalability**: Goroutines offer better scalability and a higher degree of efficiency due to their lightweight nature.
-   **Control and Management**: Threads are heavier, and their management is more complex, often leading to performance issues with a large number of threads. Goroutines are more straightforward, with Go's runtime handling the complexities of distribution over OS threads.
-   **Concurrency Model**: Goroutines use channels for communication and synchronization, leading to cleaner and less error-prone concurrent programs. Threads rely on locks and other synchronization primitives.

Using goroutines can often result in more efficient and maintainable concurrent applications, especially in Go.

Usage Statistics:
Prompt tokens: 34
Completion tokens: 682
Total tokens: 716
```

## How It Works

1. The example creates a client and prepares a conversation with system and user messages
2. It sends the request to the specified LLM provider and model
3. Upon receiving the response, it displays the model's answer
4. It also shows usage statistics such as token counts for the prompt and completion

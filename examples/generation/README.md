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

# Set your preferred provider (optional, default: groq)
export LLM_PROVIDER="openai"

# Set your preferred model (optional, default: deepseek-r1-distill-llama-70b)
export LLM_MODEL="gpt-4o"

# Run the example
go run [main.go](main.go)
```

## Example Output

```sh
Generating content using groq deepseek-r1-distill-llama-70b...

Model: deepseek-r1-distill-llama-70b
Reasoning: Okay, so I'm trying to understand the differences between goroutines and threads. I've heard that Go uses goroutines instead of threads, but I'm not exactly sure how they're different. Let me think through this step by step.

First, I know that both goroutines and threads are ways to handle concurrency, which means doing multiple things at the same time. But I've heard that goroutines are lighter weight than threads. What does that mean? Maybe it means they use less memory or resources. So if I create a lot of goroutines, my program won't slow down as much as if I used threads. That could be useful for scaling applications.

I remember that in other languages like Java or C#, threads are managed by the operating system. Each thread has its own stack and context, which takes up more memory. So creating a thread might be more resource-intensive. Goroutines, on the other hand, are managed by the Go runtime. That probably means the Go runtime handles scheduling them, which could be more efficient.

Wait, how does scheduling work? I think threads are scheduled by the OS, which has its own algorithm. Goroutines are scheduled by Go's scheduler, which is designed to be lightweight and efficient. Maybe that's why goroutines can switch contexts faster, making them better for handling many concurrent tasks without bogging down the system.

I've heard the term M:N scheduling before. I think M stands for goroutines and N stands for OS threads. So each goroutine doesn't have to correspond one-to-one with an OS thread. Instead, multiple goroutines can run on a single OS thread. That makes sense because OS threads are heavier, so using fewer of them would save resources. But how does that work exactly? Maybe the Go scheduler multiplexes goroutines onto threads, so if one goroutine blocks, like waiting for I/O, another can take its place on the same thread without creating a new one.

Communication between goroutines is done with channels, right? I've used channels before, and they're easier than shared memory with mutexes. In threads, you'd typically use shared memory and locks to communicate, which can be error-prone and lead to race conditions. Channels in Go provide a safer way to pass data between goroutines, which is a big plus for avoiding bugs.

What about the stack size? I think threads have a fixed stack size, which can be a problem if a function requires more stack space than allocated. Goroutines, on the other hand, have stacks that can grow or shrink as needed. That would prevent stack overflow issues that might happen with threads.

I'm also trying to remember how many goroutines vs. threads a program can handle. I've heard that Go can handle tens of thousands of goroutines without a problem, while using that many threads in other languages would be impossible due to memory constraints. So for highly concurrent systems, goroutines are much more scalable.

Wait, but does that mean goroutines are better in every case? Probably not. There might be situations where you need to use threads for certain operations, especially if they're CPU-intensive. But in Go, the standard approach is to use goroutines for concurrency, and the runtime handles the underlying threads efficiently.

So, putting it all together, the main differences are:

1. **Lightweight vs. Heavyweight**: Goroutines are lighter, using less memory and resources than threads.
2. **Scheduling**: Goroutines are scheduled by the Go runtime, which is more efficient, while threads are scheduled by the OS.
3. **M:N Model**: Multiple goroutines run on fewer OS threads, improving resource usage.
4. **Communication**: Channels vs. shared memory and locks.
5. **Stack Management**: Goroutines have dynamic stack sizes, preventing stack overflow issues.
6. **Concurrency Limits**: Goroutines allow for a much higher number of concurrent tasks compared to threads.

I think I've covered the main points. Maybe I should check if I missed anything, like how goroutines handle blocking operations or how the scheduler prioritizes them. But overall, this gives a good overview of why goroutines are preferred in Go for concurrency.

Response: The differences between goroutines and threads can be summarized as follows:

1. **Lightweight vs. Heavyweight**: Goroutines are lightweight, consuming fewer resources and memory compared to threads, which are heavier and more resource-intensive.

2. **Scheduling**: Goroutines are managed by the Go runtime's scheduler, which is efficient and lightweight. Threads are scheduled by the operating system, which can be less efficient for large numbers of concurrent tasks.

3. **M:N Scheduling Model**: Goroutines follow an M:N model, where multiple goroutines (M) run on a smaller number of OS threads (N). This model efficiently multiplexes goroutines onto threads, reducing resource overhead.

4. **Communication**: Goroutines use channels for safe and efficient communication, avoiding the need for shared memory and locks that are common in thread communication, thus reducing the risk of race conditions.

5. **Stack Management**: Goroutines have dynamically adjustable stack sizes, preventing stack overflow issues. Threads typically have fixed stack sizes, which can lead to issues if exceeded.

6. **Concurrency Limits**: Goroutines allow for a much higher number of concurrent tasks compared to threads, making them suitable for highly concurrent systems.

In summary, goroutines are designed to be efficient and scalable, making them the preferred choice for concurrency in Go, while threads are more resource-intensive and managed at the OS level.

Usage Statistics:
  Prompt tokens: 24
  Completion tokens: 1124
  Total tokens: 1148
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

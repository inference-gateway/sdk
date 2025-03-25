# Listing Models Example

This example demonstrates how to use the Inference Gateway SDK to list available language models from different providers.

## What This Example Shows

-   How to list all available models across providers
-   How to list models from specific providers
-   How to handle and display model information

## Running the Example

```sh
# Set your API URL (optional)
export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"

# Run the example
go run main.go
```

## Example Output

```sh
Listing all available models...
Found 50 models
1. embed-english-light-v3.0 (owned by cohere)
2. embed-multilingual-v2.0 (owned by cohere)
3. rerank-v3.5 (owned by cohere)
4. rerank-english-v3.0 (owned by cohere)
5. command-r (owned by cohere)
6. embed-english-light-v3.0-image (owned by cohere)
7. embed-english-v3.0-image (owned by cohere)
8. command-a-03-2025 (owned by cohere)
9. command-nightly (owned by cohere)
10. command-r7b-12-2024 (owned by cohere)
11. command-r-plus (owned by cohere)
12. rerank-multilingual-v2.0 (owned by cohere)
13. c4ai-aya-vision-32b (owned by cohere)
14. command-r7b-arabic-02-2025 (owned by cohere)
15. command-light-nightly (owned by cohere)
16. embed-english-v3.0 (owned by cohere)
17. rerank-english-v2.0 (owned by cohere)
18. embed-multilingual-light-v3.0-image (owned by cohere)
19. embed-multilingual-v3.0-image (owned by cohere)
20. c4ai-aya-expanse-32b (owned by cohere)
21. command (owned by cohere)
22. claude-3-7-sonnet-20250219 (owned by anthropic)
23. claude-3-5-sonnet-20241022 (owned by anthropic)
24. claude-3-5-haiku-20241022 (owned by anthropic)
25. claude-3-5-sonnet-20240620 (owned by anthropic)
26. claude-3-haiku-20240307 (owned by anthropic)
27. claude-3-opus-20240229 (owned by anthropic)
28. llama3-8b-8192 (owned by Meta)
29. llama-3.2-11b-vision-preview (owned by Meta)
30. deepseek-r1-distill-llama-70b (owned by DeepSeek / Meta)
31. mistral-saba-24b (owned by Mistral AI)
32. qwen-qwq-32b (owned by Alibaba Cloud)
33. llama-3.2-3b-preview (owned by Meta)
34. allam-2-7b (owned by SDAIA)
35. llama-3.2-90b-vision-preview (owned by Meta)
36. llama-3.1-8b-instant (owned by Meta)
37. llama-3.3-70b-specdec (owned by Meta)
38. llama-3.3-70b-versatile (owned by Meta)
39. llama-3.2-1b-preview (owned by Meta)
40. whisper-large-v3-turbo (owned by OpenAI)
41. deepseek-r1-distill-qwen-32b (owned by DeepSeek / Alibaba Cloud)
42. llama-guard-3-8b (owned by Meta)
43. llama3-70b-8192 (owned by Meta)
44. gemma2-9b-it (owned by Google)
45. qwen-2.5-32b (owned by Alibaba Cloud)
46. distil-whisper-large-v3-en (owned by Hugging Face)
47. whisper-large-v3 (owned by OpenAI)
48. qwen-2.5-coder-32b (owned by Alibaba Cloud)
49. deepseek-chat (owned by deepseek)
50. deepseek-reasoner (owned by deepseek)

Listing models from Groq Cloud...
Provider: groq
1. llama3-8b-8192
2. llama-guard-3-8b
3. llama3-70b-8192
4. llama-3.2-90b-vision-preview
5. llama-3.1-8b-instant
6. llama-3.3-70b-versatile
7. qwen-2.5-coder-32b
8. llama-3.2-1b-preview
9. whisper-large-v3-turbo
10. mistral-saba-24b
11. qwen-2.5-32b
12. llama-3.2-3b-preview
13. allam-2-7b
14. llama-3.2-11b-vision-preview
15. deepseek-r1-distill-qwen-32b
16. qwen-qwq-32b
17. gemma2-9b-it
18. deepseek-r1-distill-llama-70b
19. whisper-large-v3
20. distil-whisper-large-v3-en
21. llama-3.3-70b-specdec

Listing models from DeepSeek...
Provider: deepseek
1. deepseek-chat
2. deepseek-reasoner

Listing models from Ollama...
Provider: ollama
1. phi3:3.8b
2. phi3:14b
3. mistral
4. llama2
```

## How It Works

-   The example creates a client connected to the Inference Gateway
-   It calls ListModels() to get all available models across providers
-   It then calls ListProviderModels() for specific providers like OpenAI and Ollama
-   The code processes and displays the model information in a readable format

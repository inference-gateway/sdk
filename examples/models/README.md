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
Found 48 models
1. claude-3-7-sonnet-20250219 (owned by anthropic)
2. claude-3-5-sonnet-20241022 (owned by anthropic)
3. claude-3-5-haiku-20241022 (owned by anthropic)
4. claude-3-5-sonnet-20240620 (owned by anthropic)
5. claude-3-haiku-20240307 (owned by anthropic)
6. claude-3-opus-20240229 (owned by anthropic)
7. embed-english-light-v3.0 (owned by cohere)
8. embed-multilingual-v2.0 (owned by cohere)
9. rerank-v3.5 (owned by cohere)
10. rerank-english-v3.0 (owned by cohere)
11. command-r (owned by cohere)
12. embed-english-light-v3.0-image (owned by cohere)
13. embed-english-v3.0-image (owned by cohere)
14. command-a-03-2025 (owned by cohere)
15. command-nightly (owned by cohere)
16. command-r7b-12-2024 (owned by cohere)
17. command-r-plus (owned by cohere)
18. rerank-multilingual-v2.0 (owned by cohere)
19. c4ai-aya-vision-32b (owned by cohere)
20. command-r7b-arabic-02-2025 (owned by cohere)
21. command-light-nightly (owned by cohere)
22. embed-english-v3.0 (owned by cohere)
23. rerank-english-v2.0 (owned by cohere)
24. embed-multilingual-light-v3.0-image (owned by cohere)
25. embed-multilingual-v3.0-image (owned by cohere)
26. c4ai-aya-expanse-32b (owned by cohere)
27. command (owned by cohere)
28. llama-3.2-3b-preview (owned by Meta)
29. qwen-2.5-32b (owned by Alibaba Cloud)
30. llama-3.1-8b-instant (owned by Meta)
31. llama-3.2-11b-vision-preview (owned by Meta)
32. deepseek-r1-distill-llama-70b (owned by DeepSeek / Meta)
33. llama-3.2-90b-vision-preview (owned by Meta)
34. mistral-saba-24b (owned by Mistral AI)
35. distil-whisper-large-v3-en (owned by Hugging Face)
36. llama-guard-3-8b (owned by Meta)
37. gemma2-9b-it (owned by Google)
38. whisper-large-v3-turbo (owned by OpenAI)
39. llama-3.3-70b-specdec (owned by Meta)
40. llama-3.3-70b-versatile (owned by Meta)
41. llama3-8b-8192 (owned by Meta)
42. llama-3.2-1b-preview (owned by Meta)
43. llama3-70b-8192 (owned by Meta)
44. deepseek-r1-distill-qwen-32b (owned by DeepSeek / Alibaba Cloud)
45. whisper-large-v3 (owned by OpenAI)
46. qwen-qwq-32b (owned by Alibaba Cloud)
47. qwen-2.5-coder-32b (owned by Alibaba Cloud)
48. allam-2-7b (owned by SDAIA)

Listing models from Groq Cloud...
Provider: groq
1. llama-3.2-1b-preview
2. deepseek-r1-distill-llama-70b
3. distil-whisper-large-v3-en
4. allam-2-7b
5. qwen-2.5-coder-32b
6. llama3-8b-8192
7. llama-3.2-90b-vision-preview
8. deepseek-r1-distill-qwen-32b
9. whisper-large-v3
10. llama-3.3-70b-specdec
11. qwen-2.5-32b
12. llama-3.3-70b-versatile
13. llama-3.2-3b-preview
14. llama3-70b-8192
15. llama-3.1-8b-instant
16. llama-guard-3-8b
17. whisper-large-v3-turbo
18. llama-3.2-11b-vision-preview
19. qwen-qwq-32b
20. mistral-saba-24b
21. gemma2-9b-it

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

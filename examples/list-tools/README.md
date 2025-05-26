# List MCP Tools Example

This example demonstrates how to list available MCP (Model Context Protocol) tools using the Inference Gateway SDK.

## Overview

The Model Context Protocol (MCP) allows the Inference Gateway to expose various tools and services that language models can use to perform tasks like reading files, making API calls, or accessing databases. This example shows how to discover what tools are available.

## Prerequisites

- An Inference Gateway instance running with `EXPOSE_MCP=true` configured
- Access to MCP tools endpoint (requires authentication in most setups)

## Running the Example

1. Set the Inference Gateway URL (optional, defaults to `http://localhost:8080/v1`):
   ```bash
   export INFERENCE_GATEWAY_URL="http://localhost:8080/v1"
   ```

2. Set the API key if authentication is required:
   ```bash
   export INFERENCE_GATEWAY_API_KEY="your-api-key"
   ```

3. Run the example:
   ```bash
   go run main.go
   ```

## Expected Output

When successful, you should see output similar to:

```
Listing available MCP tools...
Found 2 MCP tools:

Tool 1:
  Name: read_file
  Description: Read content from a file
  Server: http://mcp-filesystem-server:8083/mcp
  Input Schema: map[properties:map[file_path:map[description:Path to the file to read type:string]] required:[file_path] type:object]

Tool 2:
  Name: write_file
  Description: Write content to a file
  Server: http://mcp-filesystem-server:8083/mcp
  Input Schema: map[properties:map[content:map[description:Content to write to the file type:string] file_path:map[description:Path to the file to write type:string]] required:[file_path content] type:object]
```

If MCP tools are not exposed, you'll see:

```
No MCP tools available. Make sure EXPOSE_MCP=true is set on the server.
```

## Error Handling

The example includes proper error handling for common scenarios:

- **403 Forbidden**: MCP tools endpoint is not exposed (`EXPOSE_MCP=false`)
- **401 Unauthorized**: Missing or invalid API key
- **Network errors**: Connection issues with the Inference Gateway

## Code Structure

The example follows these steps:

1. **Configuration**: Read environment variables for URL and API key
2. **Client Creation**: Initialize the SDK client with the configuration
3. **List Tools**: Call the `ListTools` method to retrieve available tools
4. **Display Results**: Format and display the tool information
5. **Error Handling**: Handle and display any errors appropriately

## Integration

You can integrate this pattern into your applications to:

- Dynamically discover available tools
- Validate tool availability before use
- Display tool capabilities to users
- Build adaptive applications that work with different MCP configurations

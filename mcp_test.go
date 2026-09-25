package sdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func TestMCPJSONRPCToolsCall(t *testing.T) {
	var received MCPJSONRPCRequest

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/mcp", r.URL.Path, "MCP endpoint lives at the root, not under /v1")
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, MCPProtocolVersion, r.Header.Get("MCP-Protocol-Version"))
		assert.Equal(t, "tools/call", r.Header.Get("Mcp-Method"))
		assert.Equal(t, "mcp_deepwiki_ask_question", r.Header.Get("Mcp-Name"))

		require.NoError(t, json.NewDecoder(r.Body).Decode(&received))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"resultType":"complete"}}`))
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})

	request, err := NewMCPJSONRPCRequest(1, ToolsCall, map[string]any{
		"name":      "mcp_deepwiki_ask_question",
		"arguments": map[string]any{"question": "How is MCP wired up?"},
	})
	require.NoError(t, err)

	response, err := client.MCPJSONRPC(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, response)
	require.NotNil(t, response.Result)
	assert.Equal(t, "complete", (*response.Result)["resultType"])

	id, err := response.ID.AsMCPJSONRPCResponseID1()
	require.NoError(t, err)
	assert.Equal(t, 1, id)

	require.NotNil(t, received.Params)
	meta, ok := (*received.Params)["_meta"].(map[string]any)
	require.True(t, ok, "_meta should be filled in")
	assert.Equal(t, MCPProtocolVersion, meta["io.modelcontextprotocol/protocolVersion"])
	assert.Contains(t, meta, "io.modelcontextprotocol/clientInfo")
	assert.Contains(t, meta, "io.modelcontextprotocol/clientCapabilities")
}

func TestMCPJSONRPCNonASCIIToolName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "=?base64?bWNwX2FfxJs=?=", r.Header.Get("Mcp-Name"))
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{}}`))
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL})

	request, err := NewMCPJSONRPCRequest(1, ToolsCall, map[string]any{"name": "mcp_a_ě"})
	require.NoError(t, err)

	_, err = client.MCPJSONRPC(context.Background(), request)
	require.NoError(t, err)
}

func TestMCPJSONRPCToolsCallWithoutName(t *testing.T) {
	client := NewClient(&ClientOptions{BaseURL: "http://localhost:8080/v1"})

	request, err := NewMCPJSONRPCRequest(1, ToolsCall, nil)
	require.NoError(t, err)

	response, err := client.MCPJSONRPC(context.Background(), request)
	assert.Nil(t, response)
	assert.ErrorContains(t, err, "params.name")
}

func TestMCPJSONRPCErrorEnvelope(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":"abc","error":{"code":-32022,"message":"unsupported protocol version"}}`))
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})

	request, err := NewMCPJSONRPCRequest("abc", ToolsList, nil)
	require.NoError(t, err)

	response, err := client.MCPJSONRPC(context.Background(), request)
	require.NoError(t, err, "JSON-RPC errors are returned in the envelope, not as a Go error")
	require.NotNil(t, response)
	require.NotNil(t, response.Error)
	assert.Equal(t, -32022, response.Error.Code)

	id, err := response.ID.AsMCPJSONRPCResponseID0()
	require.NoError(t, err)
	assert.Equal(t, "abc", id)
}

func TestMCPJSONRPCNotExposed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"error":"MCP endpoint is not exposed. Set MCP_EXPOSE=true to enable."}`))
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL})

	request, err := NewMCPJSONRPCRequest(1, ToolsList, nil)
	require.NoError(t, err)

	response, err := client.MCPJSONRPC(context.Background(), request)
	assert.Nil(t, response)
	assert.ErrorContains(t, err, "MCP_EXPOSE=true")
}

func TestMCPJSONRPCNotification(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL})

	request, err := NewMCPJSONRPCRequest(nil, ToolsList, nil)
	require.NoError(t, err)
	assert.Nil(t, request.ID, "a notification carries no id")

	response, err := client.MCPJSONRPC(context.Background(), request)
	assert.NoError(t, err)
	assert.Nil(t, response)
}

func TestNewMCPJSONRPCRequestUnsupportedID(t *testing.T) {
	_, err := NewMCPJSONRPCRequest(1.5, ToolsList, nil)
	assert.ErrorContains(t, err, "unsupported JSON-RPC id type")
}

func TestGetMCPProtectedResourceMetadata(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/.well-known/oauth-protected-resource/mcp", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		_, _ = w.Write([]byte(`{
			"resource": "https://gateway.example.com/mcp",
			"authorization_servers": ["https://keycloak.example.com/realms/inference-gateway-realm"],
			"bearer_methods_supported": ["header"]
		}`))
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})

	metadata, err := client.GetMCPProtectedResourceMetadata(context.Background())
	require.NoError(t, err)
	require.NotNil(t, metadata)
	assert.Equal(t, "https://gateway.example.com/mcp", metadata.Resource)
	assert.Equal(t, []string{"https://keycloak.example.com/realms/inference-gateway-realm"}, metadata.AuthorizationServers)
	assert.Equal(t, []string{"header"}, metadata.BearerMethodsSupported)
}

func TestGetMCPProtectedResourceMetadataNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":"Not found"}`))
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL})

	metadata, err := client.GetMCPProtectedResourceMetadata(context.Background())
	assert.Nil(t, metadata)
	assert.ErrorContains(t, err, "Not found")
}

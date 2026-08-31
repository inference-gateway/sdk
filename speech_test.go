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

func TestCreateSpeech(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/audio/speech", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "openai", r.URL.Query().Get("provider"))

		var requestBody CreateSpeechRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		assert.NoError(t, err)
		assert.Equal(t, "gpt-4o-mini-tts", requestBody.Model)
		assert.Equal(t, "What is Go?", requestBody.Input)
		assert.Equal(t, "alloy", requestBody.Voice)

		w.Header().Set("Content-Type", "audio/mpeg")
		_, err = w.Write([]byte("fake-mp3-bytes"))
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{
		BaseURL: server.URL + "/v1",
	})

	audio, err := client.CreateSpeech(context.Background(), Openai, CreateSpeechRequest{
		Model: "gpt-4o-mini-tts",
		Input: "What is Go?",
		Voice: "alloy",
	})

	require.NoError(t, err)
	assert.Equal(t, []byte("fake-mp3-bytes"), audio)
}

func TestCreateSpeech_ProviderNotSupported(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(`{"error": "The Audio API is not supported by this provider yet."}`))
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{
		BaseURL: server.URL + "/v1",
	})

	audio, err := client.CreateSpeech(context.Background(), Openai, CreateSpeechRequest{
		Model: "gpt-4o-mini-tts",
		Input: "What is Go?",
		Voice: "alloy",
	})

	require.Error(t, err)
	assert.Nil(t, audio)
	assert.Contains(t, err.Error(), "The Audio API is not supported by this provider yet.")
}

package sdk

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func TestCreateSFX(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/audio/sfx", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "elevenlabs", r.URL.Query().Get("provider"))

		var requestBody CreateSFXRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&requestBody))
		assert.Equal(t, "sound-effects-v2", requestBody.Model)
		assert.Equal(t, "door creak", requestBody.Prompt)

		_, err := w.Write([]byte("fake-sfx"))
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})
	audio, err := client.CreateSFX(context.Background(), Provider("elevenlabs"), CreateSFXRequest{
		Model:  "sound-effects-v2",
		Prompt: "door creak",
	})

	require.NoError(t, err)
	assert.Equal(t, []byte("fake-sfx"), audio)
}

func TestCreateMusic_ProviderNotSupported(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/audio/music", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(`{"error": "Music generation is not supported by this provider yet."}`))
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})
	audio, err := client.CreateMusic(context.Background(), Openai, CreateMusicRequest{Model: "m", Prompt: "p"})

	require.Error(t, err)
	assert.Nil(t, audio)
	assert.Contains(t, err.Error(), "Music generation is not supported by this provider yet.")
}

func TestCreateVideo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/videos", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "openai", r.URL.Query().Get("provider"))
		assert.Contains(t, r.Header.Get("Content-Type"), "multipart/form-data")

		require.NoError(t, r.ParseMultipartForm(1<<20))
		assert.Equal(t, "sora-2", r.FormValue("model"))
		assert.Equal(t, "A cat surfing", r.FormValue("prompt"))
		assert.Equal(t, "8", r.FormValue("seconds"))

		file, header, err := r.FormFile("input_reference")
		require.NoError(t, err)
		defer func() { _ = file.Close() }()
		assert.Equal(t, "ref.png", header.Filename)
		data, err := io.ReadAll(file)
		assert.NoError(t, err)
		assert.Equal(t, []byte("png-bytes"), data)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(VideoJob{
			ID: "video_123", Object: "video", Model: "sora-2", Status: "queued", CreatedAt: 1730419200,
		})
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})

	var ref openapi_types.File
	ref.InitFromBytes([]byte("png-bytes"), "ref.png")

	job, err := client.CreateVideo(context.Background(), Openai, CreateVideoRequest{
		Model:          "sora-2",
		Prompt:         new("A cat surfing"),
		Seconds:        new("8"),
		InputReference: &ref,
	})

	require.NoError(t, err)
	require.NotNil(t, job)
	assert.Equal(t, "video_123", job.ID)
	assert.Equal(t, VideoJobStatus("queued"), job.Status)
}

func TestCreateVideo_ReferenceImages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, r.ParseMultipartForm(1<<20))
		headers := r.MultipartForm.File["reference_images"]
		require.Len(t, headers, 2, "each reference image must be its own part")
		for i, want := range []string{"front.png", "side.png"} {
			assert.Equal(t, want, headers[i].Filename)
			file, err := headers[i].Open()
			require.NoError(t, err)
			data, err := io.ReadAll(file)
			_ = file.Close()
			require.NoError(t, err)
			assert.Equal(t, []byte(want+"-bytes"), data)
		}

		w.Header().Set("Content-Type", "application/json")
		assert.NoError(t, json.NewEncoder(w).Encode(VideoJob{ID: "video_123", Status: "queued"}))
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})

	var front, side openapi_types.File
	front.InitFromBytes([]byte("front.png-bytes"), "front.png")
	side.InitFromBytes([]byte("side.png-bytes"), "side.png")

	_, err := client.CreateVideo(context.Background(), Elevenlabs, CreateVideoRequest{
		Model:           "veo-3.1-generate-001",
		Prompt:          new("The same woman walking through a market"),
		ReferenceImages: &[]openapi_types.File{front, side},
	})
	require.NoError(t, err)
}

func TestRetrieveVideo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/videos/video_123", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "openai", r.URL.Query().Get("provider"))

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(VideoJob{
			ID: "video_123", Object: "video", Model: "sora-2", Status: "completed", CreatedAt: 1730419200, Progress: new(100),
		})
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})
	job, err := client.RetrieveVideo(context.Background(), Openai, "video_123")

	require.NoError(t, err)
	require.NotNil(t, job)
	assert.Equal(t, VideoJobStatus("completed"), job.Status)
	assert.Equal(t, 100, *job.Progress)
}

func TestDownloadVideoContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/videos/video_123/content", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Empty(t, r.URL.Query().Get("provider"))

		w.Header().Set("Content-Type", "video/mp4")
		_, err := w.Write([]byte("fake-mp4"))
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})
	video, err := client.DownloadVideoContent(context.Background(), "", "video_123")

	require.NoError(t, err)
	assert.Equal(t, []byte("fake-mp4"), video)
}

func TestDownloadVideoContent_NotReady(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(`{"error": "video not ready"}`))
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(&ClientOptions{BaseURL: server.URL + "/v1"})
	video, err := client.DownloadVideoContent(context.Background(), Openai, "video_123")

	require.Error(t, err)
	assert.Nil(t, video)
	assert.Contains(t, err.Error(), "video not ready")
}

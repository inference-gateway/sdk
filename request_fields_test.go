package sdk

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ptr is a small generic helper for building pointers to literals in tests.
func ptr[T any](v T) *T { return &v }

// TestStopUnion_RoundTrip verifies the `stop` oneOf union (string | []string)
// marshals and unmarshals without dropping or flattening either variant.
func TestStopUnion_RoundTrip(t *testing.T) {
	t.Run("string variant", func(t *testing.T) {
		var stop CreateChatCompletionRequest_Stop
		require.NoError(t, stop.FromCreateChatCompletionRequestStop0("\n\n"))

		data, err := json.Marshal(stop)
		require.NoError(t, err)
		assert.JSONEq(t, `"\n\n"`, string(data))

		var decoded CreateChatCompletionRequest_Stop
		require.NoError(t, json.Unmarshal(data, &decoded))
		got, err := decoded.AsCreateChatCompletionRequestStop0()
		require.NoError(t, err)
		assert.Equal(t, "\n\n", got)
	})

	t.Run("array variant", func(t *testing.T) {
		var stop CreateChatCompletionRequest_Stop
		require.NoError(t, stop.FromCreateChatCompletionRequestStop1([]string{"STOP", "END"}))

		data, err := json.Marshal(stop)
		require.NoError(t, err)
		assert.JSONEq(t, `["STOP","END"]`, string(data))

		var decoded CreateChatCompletionRequest_Stop
		require.NoError(t, json.Unmarshal(data, &decoded))
		got, err := decoded.AsCreateChatCompletionRequestStop1()
		require.NoError(t, err)
		assert.Equal(t, []string{"STOP", "END"}, got)
	})
}

// TestToolChoiceUnion_RoundTrip verifies the `tool_choice` oneOf union
// (string enum | named tool choice) round-trips through JSON.
func TestToolChoiceUnion_RoundTrip(t *testing.T) {
	t.Run("string enum variant", func(t *testing.T) {
		var tc ChatCompletionToolChoiceOption
		require.NoError(t, tc.FromChatCompletionToolChoiceOption0(ChatCompletionToolChoiceOption0Auto))

		data, err := json.Marshal(tc)
		require.NoError(t, err)
		assert.JSONEq(t, `"auto"`, string(data))

		var decoded ChatCompletionToolChoiceOption
		require.NoError(t, json.Unmarshal(data, &decoded))
		got, err := decoded.AsChatCompletionToolChoiceOption0()
		require.NoError(t, err)
		assert.Equal(t, ChatCompletionToolChoiceOption0Auto, got)
	})

	t.Run("named tool choice variant", func(t *testing.T) {
		named := ChatCompletionNamedToolChoice{Type: Function}
		named.Function.Name = "get_weather"

		var tc ChatCompletionToolChoiceOption
		require.NoError(t, tc.FromChatCompletionNamedToolChoice(named))

		data, err := json.Marshal(tc)
		require.NoError(t, err)
		assert.JSONEq(t, `{"type":"function","function":{"name":"get_weather"}}`, string(data))

		var decoded ChatCompletionToolChoiceOption
		require.NoError(t, json.Unmarshal(data, &decoded))
		got, err := decoded.AsChatCompletionNamedToolChoice()
		require.NoError(t, err)
		assert.Equal(t, Function, got.Type)
		assert.Equal(t, "get_weather", got.Function.Name)
	})
}

// TestResponseFormatUnion_RoundTrip verifies the `response_format` oneOf union
// (text | json_schema | json_object) round-trips through JSON with all three
// variants preserved.
func TestResponseFormatUnion_RoundTrip(t *testing.T) {
	t.Run("text variant", func(t *testing.T) {
		var rf CreateChatCompletionRequest_ResponseFormat
		require.NoError(t, rf.FromResponseFormatText(ResponseFormatText{Type: ResponseFormatTextTypeText}))

		data, err := json.Marshal(rf)
		require.NoError(t, err)
		assert.JSONEq(t, `{"type":"text"}`, string(data))

		var decoded CreateChatCompletionRequest_ResponseFormat
		require.NoError(t, json.Unmarshal(data, &decoded))
		got, err := decoded.AsResponseFormatText()
		require.NoError(t, err)
		assert.Equal(t, ResponseFormatTextTypeText, got.Type)
	})

	t.Run("json_object variant", func(t *testing.T) {
		var rf CreateChatCompletionRequest_ResponseFormat
		require.NoError(t, rf.FromResponseFormatJSONObject(ResponseFormatJSONObject{Type: JSONObject}))

		data, err := json.Marshal(rf)
		require.NoError(t, err)
		assert.JSONEq(t, `{"type":"json_object"}`, string(data))

		var decoded CreateChatCompletionRequest_ResponseFormat
		require.NoError(t, json.Unmarshal(data, &decoded))
		got, err := decoded.AsResponseFormatJSONObject()
		require.NoError(t, err)
		assert.Equal(t, JSONObject, got.Type)
	})

	t.Run("json_schema variant", func(t *testing.T) {
		var schema ResponseFormatJSONSchema
		schema.Type = JSONSchema
		schema.JSONSchema.Name = "weather"
		schema.JSONSchema.Strict = ptr(true)
		schema.JSONSchema.Schema = &ResponseFormatJSONSchemaSchema{"type": "object"}

		var rf CreateChatCompletionRequest_ResponseFormat
		require.NoError(t, rf.FromResponseFormatJSONSchema(schema))

		data, err := json.Marshal(rf)
		require.NoError(t, err)

		var decoded CreateChatCompletionRequest_ResponseFormat
		require.NoError(t, json.Unmarshal(data, &decoded))
		got, err := decoded.AsResponseFormatJSONSchema()
		require.NoError(t, err)
		assert.Equal(t, JSONSchema, got.Type)
		assert.Equal(t, "weather", got.JSONSchema.Name)
		require.NotNil(t, got.JSONSchema.Strict)
		assert.True(t, *got.JSONSchema.Strict)
	})
}

// TestCreateChatCompletionRequest_NewFields verifies the additive
// OpenAI-compatible request parameters added in schemas#71 are present on the
// public API and survive a JSON marshal/unmarshal round-trip.
func TestCreateChatCompletionRequest_NewFields(t *testing.T) {
	req := CreateChatCompletionRequest{
		Model:               "openai/gpt-4o",
		Messages:            []Message{{Role: User}},
		Temperature:         ptr(float32(0.7)),
		TopP:                ptr(float32(0.9)),
		N:                   ptr(2),
		FrequencyPenalty:    ptr(float32(0.5)),
		PresencePenalty:     ptr(float32(-0.5)),
		Seed:                ptr(42),
		Logprobs:            ptr(true),
		TopLogprobs:         ptr(5),
		LogitBias:           &map[string]int{"1234": 10},
		User:                ptr("end-user-1"),
		MaxCompletionTokens: ptr(256),
		ReasoningEffort:     ptr(High),
	}

	data, err := json.Marshal(req)
	require.NoError(t, err)

	var decoded CreateChatCompletionRequest
	require.NoError(t, json.Unmarshal(data, &decoded))

	require.NotNil(t, decoded.Temperature)
	assert.InDelta(t, 0.7, float64(*decoded.Temperature), 0.0001)
	require.NotNil(t, decoded.TopP)
	assert.InDelta(t, 0.9, float64(*decoded.TopP), 0.0001)
	require.NotNil(t, decoded.N)
	assert.Equal(t, 2, *decoded.N)
	require.NotNil(t, decoded.FrequencyPenalty)
	assert.InDelta(t, 0.5, float64(*decoded.FrequencyPenalty), 0.0001)
	require.NotNil(t, decoded.PresencePenalty)
	assert.InDelta(t, -0.5, float64(*decoded.PresencePenalty), 0.0001)
	require.NotNil(t, decoded.Seed)
	assert.Equal(t, 42, *decoded.Seed)
	require.NotNil(t, decoded.Logprobs)
	assert.True(t, *decoded.Logprobs)
	require.NotNil(t, decoded.TopLogprobs)
	assert.Equal(t, 5, *decoded.TopLogprobs)
	require.NotNil(t, decoded.LogitBias)
	assert.Equal(t, 10, (*decoded.LogitBias)["1234"])
	require.NotNil(t, decoded.User)
	assert.Equal(t, "end-user-1", *decoded.User)
	require.NotNil(t, decoded.MaxCompletionTokens)
	assert.Equal(t, 256, *decoded.MaxCompletionTokens)
	require.NotNil(t, decoded.ReasoningEffort)
	assert.Equal(t, High, *decoded.ReasoningEffort)
}

// TestReasoningEffort_Valid sanity-checks the generated enum validator for the
// new reasoning_effort field.
func TestReasoningEffort_Valid(t *testing.T) {
	for _, v := range []CreateChatCompletionRequestReasoningEffort{Minimal, Low, Medium, High} {
		assert.Truef(t, v.Valid(), "%q should be a valid reasoning_effort", v)
	}
	assert.False(t, CreateChatCompletionRequestReasoningEffort("extreme").Valid())
}

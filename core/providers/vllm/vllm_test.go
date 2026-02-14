package vllm_test

import (
	"os"
	"strings"
	"testing"

	"github.com/maximhq/bifrost/core/internal/llmtests"
	"github.com/maximhq/bifrost/core/schemas"
)

func TestVLLM(t *testing.T) {
	t.Parallel()
	baseURL := strings.TrimSpace(os.Getenv("VLLM_BASE_URL"))
	if baseURL == "" && strings.TrimSpace(os.Getenv("VLLM_API_KEY")) == "" {
		t.Skip("Skipping vLLM tests because VLLM_BASE_URL and VLLM_API_KEY are not set")
	}

	client, ctx, cancel, err := llmtests.SetupTest()
	if err != nil {
		t.Fatalf("Error initializing test setup: %v", err)
	}
	defer cancel()

	chatModel := getEnvWithDefault("VLLM_CHAT_MODEL", "Qwen/Qwen3-0.6B")
	textModel := getEnvWithDefault("VLLM_TEXT_MODEL", "Qwen/Qwen3-0.6B")
	embeddingModel := getEnvWithDefault("VLLM_EMBEDDING_MODEL", "BAAI/bge-m3")

	testConfig := llmtests.ComprehensiveTestConfig{
		Provider:       schemas.VLLM,
		ChatModel:      chatModel,
		TextModel:      textModel,
		EmbeddingModel: embeddingModel,
		Scenarios: llmtests.TestScenarios{
			TextCompletion:        true,
			TextCompletionStream:  true,
			SimpleChat:            true,
			CompletionStream:      true,
			MultiTurnConversation: true,
			ToolCalls:             true,
			ToolCallsStreaming:    true,
			MultipleToolCalls:     true,
			End2EndToolCalling:    true,
			AutomaticFunctionCall: true,
			ImageURL:              false,
			ImageBase64:           false,
			MultipleImages:        false,
			CompleteEnd2End:       true,
			Embedding:             true,
			ListModels:            true,
			Reasoning:             false,
			SpeechSynthesis:       false,
			SpeechSynthesisStream: false,
			Transcription:         true,
			TranscriptionStream:   false,
			ImageGeneration:       false,
			ImageGenerationStream: false,
			ImageEdit:             false,
			ImageEditStream:       false,
			ImageVariation:        false,
			ImageVariationStream:  false,
		},
	}

	t.Run("VLLMTests", func(t *testing.T) {
		llmtests.RunAllComprehensiveTests(t, client, ctx, testConfig)
	})
	client.Shutdown()
}

func getEnvWithDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

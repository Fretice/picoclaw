package agent

import (
	"testing"

	"github.com/sipeed/picoclaw/pkg/config"
)

func TestResolveModelCandidates_AutoFallbackFromModelListWhenUnset(t *testing.T) {
	cfg := &config.Config{
		ModelList: []*config.ModelConfig{
			{ModelName: "primary", Model: "openai/gpt-5.4"},
			{ModelName: "backup-anthropic", Model: "anthropic/claude-sonnet-4.6"},
			{ModelName: "backup-deepseek", Model: "deepseek/deepseek-chat"},
		},
	}

	candidates := resolveModelCandidates(cfg, "openai", "primary", nil)
	if len(candidates) != 3 {
		t.Fatalf("len(candidates) = %d, want 3", len(candidates))
	}

	if candidates[0].Provider != "openai" || candidates[0].Model != "gpt-5.4" {
		t.Fatalf("candidate[0] = %s/%s, want openai/gpt-5.4", candidates[0].Provider, candidates[0].Model)
	}
	if candidates[1].Provider != "anthropic" || candidates[1].Model != "claude-sonnet-4.6" {
		t.Fatalf(
			"candidate[1] = %s/%s, want anthropic/claude-sonnet-4.6",
			candidates[1].Provider,
			candidates[1].Model,
		)
	}
	if candidates[2].Provider != "deepseek" || candidates[2].Model != "deepseek-chat" {
		t.Fatalf("candidate[2] = %s/%s, want deepseek/deepseek-chat", candidates[2].Provider, candidates[2].Model)
	}
}

func TestResolveModelCandidates_ExplicitEmptyFallbacksDisableAutoFallback(t *testing.T) {
	cfg := &config.Config{
		ModelList: []*config.ModelConfig{
			{ModelName: "primary", Model: "openai/gpt-5.4"},
			{ModelName: "backup", Model: "anthropic/claude-sonnet-4.6"},
		},
	}

	candidates := resolveModelCandidates(cfg, "openai", "primary", []string{})
	if len(candidates) != 1 {
		t.Fatalf("len(candidates) = %d, want 1", len(candidates))
	}
	if candidates[0].Provider != "openai" || candidates[0].Model != "gpt-5.4" {
		t.Fatalf("candidate[0] = %s/%s, want openai/gpt-5.4", candidates[0].Provider, candidates[0].Model)
	}
}

func TestResolveModelCandidates_ExplicitFallbacksPreserved(t *testing.T) {
	cfg := &config.Config{
		ModelList: []*config.ModelConfig{
			{ModelName: "primary", Model: "openai/gpt-5.4"},
			{ModelName: "backup-anthropic", Model: "anthropic/claude-sonnet-4.6"},
			{ModelName: "backup-deepseek", Model: "deepseek/deepseek-chat"},
		},
	}

	candidates := resolveModelCandidates(
		cfg,
		"openai",
		"primary",
		[]string{"backup-anthropic"},
	)
	if len(candidates) != 2 {
		t.Fatalf("len(candidates) = %d, want 2", len(candidates))
	}

	if candidates[0].Provider != "openai" || candidates[0].Model != "gpt-5.4" {
		t.Fatalf("candidate[0] = %s/%s, want openai/gpt-5.4", candidates[0].Provider, candidates[0].Model)
	}
	if candidates[1].Provider != "anthropic" || candidates[1].Model != "claude-sonnet-4.6" {
		t.Fatalf(
			"candidate[1] = %s/%s, want anthropic/claude-sonnet-4.6",
			candidates[1].Provider,
			candidates[1].Model,
		)
	}
}


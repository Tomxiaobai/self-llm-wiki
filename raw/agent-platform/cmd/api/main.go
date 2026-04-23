package main

import (
	"log"
	"net/http"

	"agent-platform/internal/agents/gateway"
	"agent-platform/internal/agents/rag"
	"agent-platform/internal/agents/report"
	"agent-platform/internal/agents/search"
	"agent-platform/internal/api/handler"
	"agent-platform/internal/harness/eval"
	"agent-platform/internal/harness/hooks"
	"agent-platform/internal/harness/policy"
	"agent-platform/internal/harness/trace"
	"agent-platform/internal/mcp/registry"
	"agent-platform/internal/orchestrator/task"
	"agent-platform/internal/shared/config"
	"agent-platform/internal/store/memory"
)

func main() {
	cfg := config.Load()
	store := memory.NewStore()
	tracer := trace.NewNopTracer()

	ragService := rag.NewService(nil)
	searchService := search.NewService(nil)
	reportService := report.NewService("")
	reg := registry.NewMemoryRegistry()
	policyEngine := policy.NewAllowAllEngine()
	hookChain := hooks.NewChain(hooks.NopHook{})

	gatewayService := gateway.NewService(
		ragService,
		searchService,
		reportService,
		reg,
		policyEngine,
		hookChain,
		tracer,
	)

	taskService := task.NewService(store, store, gatewayService, eval.SimpleEvaluator{}, tracer)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("api ok"))
	})
	mux.Handle("/v1/chat", handler.NewChatHandler(taskService))

	log.Printf("api listening on %s", cfg.HTTP.APIAddr)
	log.Fatal(http.ListenAndServe(cfg.HTTP.APIAddr, mux))
}

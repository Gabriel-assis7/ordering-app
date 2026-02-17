.PHONY: run-order-service
run-order-service:
	@./scripts/run.sh order

.PHONY: lint
lint:
	@./scripts/lint.sh pkg
	@./scripts/lint.sh order

.PHONY: format
format:
	@./scripts/format.sh pkg
	@./scripts/format.sh order

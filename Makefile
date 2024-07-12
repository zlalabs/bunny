GOPATH=$(shell go env GOPATH)

.PHONY: dev
dev:
	$(MAKE) dev-api & $(MAKE) dev-ui

.PHONY: dev-api
dev-api:
	go run ./main.go

.PHONY: dev-ui
dev-ui:
	pnpm run dev

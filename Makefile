LINTER_VERSION := "v1.50.1"

build: build-client build-server

build-client:
	go build -o bin/gophkeeper-client cmd/client/main.go && chmod +x ./bin/gophkeeper-client

build-server:
	go build -o bin/gophkeeper-server cmd/server/main.go && chmod +x ./bin/gophkeeper-server

up:
	docker compose up -d

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:$(LINTER_VERSION) \
      golangci-lint run --modules-download-mode=mod --out-format code-climate --path-prefix src/ -c .golangci.yml | \
      jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'

lint-fix:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:$(LINTER_VERSION) \
	  golangci-lint run --modules-download-mode=mod --fix
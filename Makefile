LINTER_VERSION := "v1.48.0"

build:
	go build -o bin/gophkeeper-client cmd/client/main.go && chmod +x ./bin/gophkeeper-client

up:
	docker compose up -d

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:$(LINTER_VERSION) \
      golangci-lint run --out-format code-climate --path-prefix src/ -c .golangci.yml | \
      jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'

lint-fix:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:$(LINTER_VERSION) \
	  golangci-lint run --fix

vet:
	go vet ./...
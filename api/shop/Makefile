# =================================================================
# setup tools
# =================================================================
.PHONY: install-tools
install-tools:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install golang.org/x/tools/cmd/godoc@latest
	npm i -g @redocly/cli@latest
	cp -p .env_sample .env


# =================================================================
# oapi-codegen
# =================================================================
.PHONY: gen-api bundle-openapi
gen-api: bundle-openapi
	oapi-codegen -package gen -generate types -o internal/gen/types.gen.go ./spec/openapi.yaml
	oapi-codegen -package gen -generate strict-server,gin -templates ./spec/templates -o internal/gen/server.gen.go ./spec/openapi.yaml

bundle-openapi:
	redocly bundle ./spec/openapi_base.yaml --output ./spec/openapi.yaml


# =================================================================
# for CI
# =================================================================
.PHONY: lint deps build run

lint:
	golangci-lint run --timeout 3m

deps:
	go mod download && go mod tidy

build: deps
	go build -o build/shop_api cmd/main.go

run: build
	./build/shop_api


# =================================================================
# for Container
# =================================================================
.PHONY: up down logs
up:
	docker compose up -d shop-api opensearch opensearch-dashboards

down:
	docker compose down shop-api opensearch opensearch-dashboards

logs:
	docker compose logs -f shop-api


# =================================================================
# Go Doc
# =================================================================
.PHONY: go-doc
go-doc:
	# godoc サーバーを起動してブラウザで確認
	go doc controller.NewCorsConfig
	go doc controller.CreateProductComment
	@echo "Starting godoc server..."
	@godoc -http=:6060 &
	@echo "Godoc server started at http://localhost:6060"
	@echo "Open your browser to http://localhost:6060/pkg/github.com/tamaco489/elasticsearch_demo/api/shop/controller/"

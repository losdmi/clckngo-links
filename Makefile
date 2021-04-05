
build_output = ./clckngo-links

build:
	go build -o $(build_output) ./cmd/clckngo-links/main.go

run:
	export $$(grep -v '^#' .env | xargs -0) && $(build_output)

build-and-run: build
	$(MAKE) run

watch:
	go run github.com/cosmtrek/air

mod:
	go mod tidy
	go mod vendor

openapi:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config ./api/oapi-codegen_config_types.yml ./api/openapi.yml

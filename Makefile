run: build
	@./bin/chatai

install: ## Install packages and deps
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	#@npx tailwindcss init
	@npm install -D daisyui@latest

css: ## Generate css files and keep watching for changes
	@npx tailwindcss -i view/css/app.css -o public/styles.css --watch

templ: ## Generate go files from templ for auto-refresh
	@templ generate --watch --proxy=http://localhost:8080

build: ## Build the project
	@templ generate view
	@go build -ldflags "-s -w" -tags dev -o bin/chatai main.go static_dev.go

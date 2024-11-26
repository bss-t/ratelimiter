help: ## Show this help.
	@grep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

up:
	@go run cmd/main.go
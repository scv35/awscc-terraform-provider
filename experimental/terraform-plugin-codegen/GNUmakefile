GO_VER ?= go1.21.5

default: build

build: generate ## Build provider
	$(GO_VER) install

generate: ## Generate provider
	tfplugingen-framework generate resources \
		--input specification.json \
		--output internal/awscc \
		--package awscc

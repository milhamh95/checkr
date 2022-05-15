PROJECT_NAME := checkr
TEST_OPTS := -covermode=atomic $(TEST_OPTS)

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY:dev
dev: ## Run checkr in dev
	@go run main.go

.PHONY: unittest
unittest: ## Run unit test
	@go test $(TEST_OPTS) ./...

.PHONY: build
build: ## Build check binary
	@go build -o $(PROJECT_NAME)

.PHONY: start
start: ## Start checkr binary
	@./checkr

.PHONY: gen_bulkDiscountCartStorage_stub
gen_bulkDiscountCartStorage_stub:
	@counterfeiter -o ./counterfeiter ./service/ bulkDiscountCartStorage

.PHONY: gen_bundlingPromoProductStorage_stub
gen_bundlingPromoProductStorage_stub:
	@counterfeiter -o ./counterfeiter ./service/ bundlingPromoProductStorage

.PHONY: gen_bundlingPromoCartStorage_stub
gen_bundlingPromoCartStorage_stub:
	@counterfeiter -o ./counterfeiter ./service/ bundlingPromoCartStorage

.PHONY: gen_buyXPayYPromoCartStorage_stub
gen_buyXPayYPromoCartStorage_stub:
	@counterfeiter -o ./counterfeiter ./service/ buyXPayYPromoCartStorage

.PHONY: gen_cartProductStorage_stub
gen_cartProductStorage_stub:
	@counterfeiter -o ./counterfeiter ./service/ cartProductStorage

.PHONY: gen_cartStorageSource_stub
gen_cartStorageSource_stub:
	@counterfeiter -o ./counterfeiter ./service/ cartStorageSource

.PHONY: gen_cartPromoService_stub
gen_cartPromoService_stub:
	@counterfeiter -o ./counterfeiter ./service/ cartPromoService

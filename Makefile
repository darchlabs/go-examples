# load .env file
include .env
export $(shell sed 's/=.*//' .env)

# rm:
# 	@echo "[rm] Removing..."
# 	@rm -rf bin

# compile: rm
# 	@echo "[compile] Compiling..."

run:
	@echo "[run] Running..."
	@NODE_URL=$(NODE_URL) INTERVAL_SECONDS=$(INTERVAL_SECONDS) DATABASE_URL=$(DATABASE_URL) go run cmd/synchronizer/main.go
	

BUILD_PATH=build
CONTRACTS_PATH=contracts
PRIVATE_KEY=1b2eb64764d441e4c887238d5acc1f21df06ceeb7792b5e95f9b093ea2da5511
CONTRACT_ADDRESS=0x76a38d9B3B05f0f3D88837Fe0E3676E1F6264215

rm:
	@echo "[rm] Removing..."
	@rm -rf build
	@rm -rf internal/contract

compile: rm
	@echo "[compile] Compiling..."
	@solc --optimize --abi ./$(CONTRACTS_PATH)/ExampleContract.sol -o $(BUILD_PATH)
	@solc --optimize --bin ./$(CONTRACTS_PATH)/ExampleContract.sol -o $(BUILD_PATH) 
	@mkdir internal/contract
	@abigen --bin=./$(BUILD_PATH)/ExampleContract.bin --abi=./$(BUILD_PATH)/ExampleContract.abi --pkg=contract --out=./internal/contract/ExampleContract.go

run-example-deploy: compile
	@echo "[deploy-contract] Run example deploy..."
	@go run cmd/example/deploy/main.go --url $(NODE_URL) --privateKey $(PRIVATE_KEY)

run-example-set-event: compile
	@echo "[deploy-contract] Run example set event..."
	@go run cmd/example/set-event/main.go --url $(NODE_URL) --privateKey $(PRIVATE_KEY) --contractAddr $(CONTRACT_ADDRESS)

run-example-get-events: compile
	@echo "[deploy-contract] Run example get events..."
	@go run cmd/example/get-events/main.go --url $(NODE_URL) --privateKey $(PRIVATE_KEY) --contractAddr $(CONTRACT_ADDRESS)

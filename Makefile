REQ_GEN_DEPS := grpcurl protoc protoc-gen-go protoc-gen-go-grpc

GEN_PROTO := spacex/api/common/status/status.proto spacex/api/device/command.proto spacex/api/device/common.proto spacex/api/device/device.proto spacex/api/device/dish_config.proto spacex/api/device/dish.proto spacex/api/device/services/unlock/service.proto spacex/api/device/transceiver.proto spacex/api/device/wifi_config.proto spacex/api/device/wifi_util.proto spacex/api/device/wifi.proto spacex/api/satellites/network/ut_disablement_codes.proto spacex/api/telemetron/public/common/time.proto
GEN_MODEL_DIR := model

.ONESHELL:
.PHONY: generate

generate: ## Genertes Golang code based on GRPC reflection from Dishy
	$(foreach bin,$(REQ_GEN_DEPS),$(if $(shell which $(bin)),,$(error "Please install '$(bin)'")))

	TMPFILE=$$(mktemp)
	trap "rm $${TMPFILE}" EXIT

	grpcurl -plaintext -protoset-out "$${TMPFILE}" "192.168.100.1:9200" describe SpaceX.API.Device.Device

	protoc --descriptor_set_in="$${TMPFILE}" --go_out="$(GEN_MODEL_DIR)" --go-grpc_out="$(GEN_MODEL_DIR)" \
	--go_opt="module=spacex.com/api" --go-grpc_opt="module=spacex.com/api" \
	--go_opt="Mspacex/api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock" \
	--go-grpc_opt="Mspacex/api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock" \
	--go_opt="Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites" \
	--go-grpc_opt="Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites" \
	--go_opt="Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron" \
	--go-grpc_opt="Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron" \
	$(GEN_PROTO)

	find ./$(GEN_MODEL_DIR) -type f -name "*.go" -exec \
	sed -i "s|spacex.com/api|$$(awk '/module/{print $$2}' go.mod)/$(GEN_MODEL_DIR)|g" {} \;

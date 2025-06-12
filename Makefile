REQ_GEN_DEPS := grpcurl protoc protoc-gen-go protoc-gen-go-grpc

GEN_PROTO := spacex/api/common/status/status.proto spacex_api/device/command.proto spacex_api/device/common.proto spacex_api/device/device.proto spacex_api/device/dish_config.proto spacex_api/device/dish.proto spacex_api/device/rssi_scan.proto spacex_api/device/services/unlock/service.proto spacex_api/device/transceiver.proto spacex_api/device/wifi_config.proto spacex_api/device/wifi_util.proto spacex_api/device/wifi.proto spacex/api/satellites/network/ut_disablement_codes.proto spacex_api/telemetron/public/common/time.proto spacex_api/telemetron/public/integrations/ut_pop_link_report.proto
GEN_MODEL_DIR := starlink/model

.ONESHELL:
.PHONY: generate

generate: ## Genertes Golang code based on GRPC reflection from Dishy
	$(foreach bin,$(REQ_GEN_DEPS),$(if $(shell which $(bin)),,$(error "Please install '$(bin)'")))

	TMPFILE=$$(mktemp)
	trap "rm $${TMPFILE}" EXIT

	grpcurl -vv -plaintext -protoset-out "$${TMPFILE}" "192.168.100.1:9200" describe SpaceX.API.Device.Device

	protoc --descriptor_set_in="$${TMPFILE}" --go_out="$(GEN_MODEL_DIR)" --go-grpc_out="$(GEN_MODEL_DIR)" \
	--go_opt="module=spacex.com/api" \
	--go-grpc_opt="module=spacex.com/api" \
	--go_opt="Mspacex_api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock" \
	--go-grpc_opt="Mspacex_api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock" \
	--go_opt="Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites" \
	--go-grpc_opt="Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites" \
	--go_opt="Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron" \
	--go-grpc_opt="Mspacex_api/telemetron/public/common/time.proto=spacex.com/api/telemetron" \
	--go_opt="Mspacex_api/telemetron/public/integrations/ut_pop_link_report.proto=spacex.com/api/integrations" \
	--go-grpc_opt="Mspacex_api/telemetron/public/integrations/ut_pop_link_report.proto=spacex.com/api/integrations" \
	$(GEN_PROTO)

	find ./$(GEN_MODEL_DIR) -type f -name "*.go" -exec \
	sed -i "s|spacex.com/api|$$(awk '/module/{print $$2}' go.mod)/$(GEN_MODEL_DIR)|g" {} \;
	sed -i "/indirect/d" go.mod
	go mod tidy -v -e

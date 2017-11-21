
BUILD=build

default: binaries

binaries: $(BUILD)/server $(BUILD)/db-tester $(BUILD)/sqs-worker $(BUILD)/sqs-sender $(BUILD)/fkcli $(BUILD)/testing-random $(BUILD)/weather-proxy

all: binaries

SERVER_SOURCES = $(shell find server -type f -name '*.go' -not -path "server/vendor/*")

$(BUILD)/server: $(SERVER_SOURCES)
	go build -o $@ server/server.go

$(BUILD)/db-tester: server/db-tester/*.go $(SERVER_SOURCES)
	go build -o $@ server/db-tester/*.go

$(BUILD)/sqs-worker: server/sqs-worker/*.go $(SERVER_SOURCES)
	go build -o $@ server/sqs-worker/*.go

$(BUILD)/sqs-sender: server/sqs-sender/*.go $(SERVER_SOURCES)
	go build -o $@ server/sqs-sender/*.go

$(BUILD)/fkcli: server/api/tool/fieldkit-cli/*.go $(SERVER_SOURCES)
	go build -o $@ server/api/tool/fieldkit-cli/*.go

$(BUILD)/testing-random: testing/random/*.go $(SERVER_SOURCES)
	go build -o $@ testing/random/*.go

$(BUILD)/weather-proxy: testing/weather-proxy/*.go $(SERVER_SOURCES)
	go build -o $@ testing/weather-proxy/*.go

generate:
	mv server/vendor server/vendor-temp # See https://github.com/goadesign/goa/issues/923
	(cd $(GOPATH)/src/github.com/fieldkit/cloud/server && go generate) || true
	mv server/vendor-temp server/vendor

deps:
	go get -u github.com/goadesign/goa/...

clean:
	rm -rf build 

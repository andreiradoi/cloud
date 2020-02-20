UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
GOARCH ?= amd64
GOOS ?= linux
endif

ifeq ($(UNAME_S),Darwin)
GOARCH ?= amd64
GOOS ?= darwin
endif

GO ?= env GOOS=$(GOOS) GOARCH=$(GOARCH) go
BUILD ?= $(abspath build)
WORKING_DIRECTORY ?= `pwd`
DOCKER_TAG ?= master

SERVER_SOURCES = $(shell find server -type f -name '*.go' -not -path "server/vendor/*")
TESTING_SOURCES = $(shell find testing -type f -name '*.go' -not -path "server/vendor/*")

default: setup binaries

setup: legacy/src/js/secrets.js portal/src/secrets.js ocr-portal/src/js/secrets.js

legacy/src/js/secrets.js: legacy/src/js/secrets.js.template
	cp $^ $@

ocr-portal/src/js/secrets.js: ocr-portal/src/js/secrets.js.template
	cp $^ $@

portal/src/secrets.js: portal/src/secrets.js.template
	cp $^ $@

binaries: $(BUILD)/server $(BUILD)/ingester $(BUILD)/fktool

all: binaries

tests:
	cd legacy && yarn run flow

server: $(BUILD)/server

ingester: $(BUILD)/ingester

fktool: $(BUILD)/fktool

$(BUILD)/server: $(SERVER_SOURCES)
	cd server && $(GO) build -o $@ server.go

$(BUILD)/ingester: $(SERVER_SOURCES)
	cd server && $(GO) build -o $@ ingester.go

$(BUILD)/fktool: server/tools/fktool/*.go $(SERVER_SOURCES) $(TESTING_SOURCES)
	cd server/tools/fktool && $(GO) build -o $@ *.go

generate:
	./tools/goa-generate.sh

clean:
	rm -rf $(BUILD)

schema-production:
	@mkdir -p schema-production
	@rm -f schema-production/*.sql
	@echo "CREATE USER fk;" > schema-production/0.sql
	@echo "CREATE USER server;" >> schema-production/0.sql
	@echo "CREATE USER rdsadmin;" >> schema-production/0.sql
	@for f in `find schema-production -name "*.bz2"`; do              \
		echo $$f                                                     ;\
		bunzip2 < $$f > schema-production/1.sql                      ;\
		rm $$f                                                       ;\
	done
	$(MAKE) restart-postgres

.PHONY: schema-production

clean-postgres:
	docker-compose stop postgres
	rm -f active-schema

active-schema:
	ln -sf schema active-schema

restart-postgres: active-schema
	docker-compose stop postgres
	docker-compose rm -f postgres
	docker-compose up -d postgres

veryclean:

migrate-up:
	migrate -path migrations -database "postgres://fieldkit:password@127.0.0.1:5432/fieldkit?sslmode=disable&search_path=public" up

migrate-down:
	migrate -path migrations -database "postgres://fieldkit:password@127.0.0.1:5432/fieldkit?sslmode=disable&search_path=public" down

aws-image:
	cp ocr-portal/src/js/secrets.js.aws ocr-portal/src/js/secrets.js
	cp legacy/src/js/secrets.js.aws legacy/src/js/secrets.js
	cp portal/src/secrets.js.aws portal/src/secrets.js
	WORKING_DIRECTORY=$(WORKING_DIRECTORY) DOCKER_TAG=$(DOCKER_TAG) ./build.sh

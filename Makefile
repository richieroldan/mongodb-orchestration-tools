PLATFORM?=linux
GO_LDFLAGS?="-s -w"
ENABLE_MONGODB_TESTS?=false
MONGODB_DOCKER_UID?=1001

all: bin/mongodb-healthcheck-$(PLATFORM) bin/mongodb-controller-$(PLATFORM) bin/mongodb-executor-$(PLATFORM) bin/mongodb-watchdog-$(PLATFORM)

$(GOPATH)/bin/glide:
	go get github.com/Masterminds/glide

vendor: $(GOPATH)/bin/glide glide.yaml glide.lock
	$(GOPATH)/bin/glide install --strip-vendor

bin/mongodb-healthcheck-$(PLATFORM): vendor cmd/mongodb-healthcheck/main.go healthcheck/*.go common/*.go common/*/*.go
	CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=386 go build -ldflags=$(GO_LDFLAGS) -o bin/mongodb-healthcheck-$(PLATFORM) cmd/mongodb-healthcheck/main.go

bin/mongodb-controller-$(PLATFORM): vendor cmd/mongodb-controller/main.go controller/*.go controller/*/*.go common/*.go common/*/*.go
	CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=386 go build -ldflags=$(GO_LDFLAGS) -o bin/mongodb-controller-$(PLATFORM) cmd/mongodb-controller/main.go

bin/mongodb-executor-$(PLATFORM): vendor cmd/mongodb-executor/main.go executor/*.go executor/*/*.go common/*.go common/*/*.go
	CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=386 go build -ldflags=$(GO_LDFLAGS) -o bin/mongodb-executor-$(PLATFORM) cmd/mongodb-executor/main.go

bin/mongodb-watchdog-$(PLATFORM): vendor cmd/mongodb-watchdog/main.go watchdog/*.go watchdog/*/*.go common/*.go common/*/*.go
	CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=386 go build -ldflags=$(GO_LDFLAGS) -o bin/mongodb-watchdog-$(PLATFORM) cmd/mongodb-watchdog/main.go

test: vendor
	ENABLE_MONGODB_TESTS=$(ENABLE_MONGODB_TESTS) go test -v ./...

test-mongod.key:
	openssl rand -base64 512 >test-mongod.key
	chown $(MONGODB_DOCKER_UID):0 test-mongod.key
	chmod 0600 test-mongod.key

test-full: test-mongod.key
	docker-compose up -d
	scripts/init-test-replset-wait.sh
	ENABLE_MONGODB_TESTS=true go test -v ./...

clean:
	rm -rf bin test-mongod.key vendor 2>/dev/null || true

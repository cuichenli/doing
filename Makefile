cov = "false"
testargs =
common_args =
verbose =
ifeq ($(cov),  true)
	testargs += -coverprofile cover.out
endif

ifeq ($(verbose),  true)
	common_args += -v
endif

build: bin/doing
	go build $(common_args) -o bin/doing github.com/cuichenli/doing/cmd/doing

test:
	go test $(testargs) $(common_args) ./...

cover: testargs = -coverprofile cover.out
cover: test
	go tool cover -html=cover.out -o cover.html
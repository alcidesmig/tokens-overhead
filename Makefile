ifndef $(GOPATH)
	export GOPATH := $(shell go env GOPATH)
	export PATH := $(GOPATH)/bin:$(PATH)
endif


test:
	@go test ./...

test-cov:
	@go test -coverprofile=cover.out -covermode=count  ./...

test-cov-html:
	@go test -coverprofile=cover.out ./... && go tool cover -html=cover.out

mock:
	@go generate ./...

build:
	@go build -o app main.go

run-req-local:
	@MODE=requester POSTGRES_URI="user=postgres dbname=postgres host=localhost port=5432 sslmode=disable password=admin" go run main.go

# golang agnostic commands
get-gcloud-roles:
	@gcloud iam roles list --format="value(name.basename())" > gcloud_roles
	@echo "Roles collected: 'cat gcloud_roles' to verify"

gen-rsa:
	@openssl genrsa -out keys/rsa 1024
	@openssl rsa -in keys/rsa -pubout -out keys/rsa.pub

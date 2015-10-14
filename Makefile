build:
	@go build -o terraform-provider-cobbler

test:
	@go test -v .

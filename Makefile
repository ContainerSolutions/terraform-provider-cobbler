build:
	@go build -o terraform-provider-cobbler

test:
	@go test -v .

clean:
	@rm -f dist/*

dist:
	# Build for darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -o ./dist/terraform-provider-cobbler
	cd dist && tar cf terraform-provider-cobbler-osx.tar.gz terraform-provider-cobbler
	# Build for linux-amd64
	GOOS=linux GOARCH=amd64 go build -o ./dist/terraform-provider-cobbler
	cd dist && tar cf terraform-provider-cobbler-linux-amd64.tar.gz terraform-provider-cobbler
	# Build for linux-386
	GOOS=linux GOARCH=386 go build -o ./dist/terraform-provider-cobbler
	cd dist && tar cf terraform-provider-cobbler-linux-i386.tar.gz terraform-provider-cobbler

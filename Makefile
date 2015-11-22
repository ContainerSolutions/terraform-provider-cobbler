# Copyright 2015 Container Solutions
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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

.PHONY: gen-proto di serve clean-proto test

gen-proto: clean-proto
	protoc --proto_path=handler/proto --twirp_out=:handler/rpc --go_out=:handler/rpc handler/proto/*.proto

di:
	go generate di/wire_gen.go

serve:
	dev_appserver.py entrypoint/api/app.yaml

clean-proto:
	rm -rf handler/rpc/*.{pb,twirp}.go

test:
	go test ./...
	goapp test ./...

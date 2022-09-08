tidy:
	go fmt ./...
	go mod tidy

test:
	go test

testK8S:
	go test -run "K8S*"

testAWS:
	go test -run "AWS*"

setupTestResources:
	
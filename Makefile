tidy:
	go fmt ./...
	go mod tidy

test: setupTestResources
	go test
	$(MAKE) teardownTestResources

testK8S: setupTestResources
	go test -run "K8S*"
	$(MAKE) teardownTestResources

testAWS:
	go test -run "AWS*"

setupTestResources:
	k3d cluster create golang-utils-testing

teardownTestResources:
	k3d cluster delete golang-utils-testing

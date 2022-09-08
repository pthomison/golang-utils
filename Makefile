tidy:
	go fmt ./...
	go mod tidy

test: testK8S testAWS testDB testQueue

testK8S: setupTestResources
	cd k8s && \
	go test -v
	$(MAKE) teardownTestResources

testAWS:
	cd aws && \
	go test -v

testDB:
	cd db && \
	go test -v

testQueue:
	cd queue && \
	go test -v

setupTestResources:
	k3d cluster create golang-utils-testing

teardownTestResources:
	k3d cluster delete golang-utils-testing
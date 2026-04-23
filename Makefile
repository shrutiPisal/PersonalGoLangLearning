
functional-tests":
	ginkgo -v ./address-api-tests ./ginkgo-tests

deps:
	go mod download
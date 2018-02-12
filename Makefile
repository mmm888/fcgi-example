build: dep-ensure
	go build -o fcgiapi .

dep-install:
ifeq ($(shell type dep 2> /dev/null),)
	go get -u github.com/golang/dep/...
endif

dep-ensure: dep-install
	dep ensure

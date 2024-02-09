GIT_HASH ?= $(shell git rev-parse HEAD)
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64

url_shortner: 
	CGO_ENABLED=0 go build  -o url_shortner ./cmd/url_shortner

.PHONY: cross
cross:
	$(foreach GOOS, $(PLATFORMS),\
		$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); \
	$ go build -o url_shortner-$(GOOS)-$(GOARCH) ./cmd/url_shortner; \
	shasum -a 256 url_shortner-$(GOOS)-$(GOARCH) > url_shortner-$(GOOS)-$(GOARCH).sha256 ))) \

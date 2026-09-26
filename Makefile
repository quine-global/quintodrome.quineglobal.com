# SITE selects which site under sites/ to build/run: quintodrome or hyper.
SITE ?= quintodrome

.PHONY: benchmark
benchmark:
	go test -bench=.

.PHONY: build
build:
	docker build --platform linux/amd64 --build-arg SITE=$(SITE) . -t forge.quinefoundation.com/ironmagma/$(SITE)-quineglobal-com

.PHONY: push
push: build
	docker push forge.quinefoundation.com/ironmagma/$(SITE)-quineglobal-com

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: css
css:
	npm run css

.PHONY: css-watch
css-watch:
	npm run css:watch

.PHONY: dev
dev:
	@trap 'kill 0' SIGINT SIGTERM; npm run css:watch & go run ./sites/$(SITE)/cmd; wait

.PHONY: install
install:
	npm install

.PHONY: lint
lint:
	golangci-lint run

.PHONY: start
start:
	go run ./sites/$(SITE)/cmd

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...

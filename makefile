
OWNER=jakubknejzlik
IMAGE_NAME=surveyjs-graphql-api
QNAME=$(OWNER)/$(IMAGE_NAME)

GIT_TAG=$(QNAME):$(TRAVIS_COMMIT)
BUILD_TAG=$(QNAME):0.1.$(TRAVIS_BUILD_NUMBER)
LATEST_TAG=$(QNAME):latest

lint:
	docker run -it --rm -v "$(PWD)/Dockerfile:/Dockerfile:ro" redcoolbeans/dockerlint

build:
	# go get ./...
	# GOOS=linux GOARCH=amd64 go build -o bin/go-survey-alpine
	docker build -t $(GIT_TAG) .

tag:
	docker tag $(GIT_TAG) $(BUILD_TAG)
	docker tag $(GIT_TAG) $(LATEST_TAG)

login:
	@docker login -u "$(DOCKER_USER)" -p "$(DOCKER_PASS)"
push: login
	# docker push $(GIT_TAG)
	# docker push $(BUILD_TAG)
	docker push $(LATEST_TAG)

generate:
	GO111MODULE=on go run github.com/novacloudcz/graphql-orm

reinit:
	GO111MODULE=on go run github.com/novacloudcz/graphql-orm init

run:
	# DATABASE_URL=sqlite3://test.db PORT=8081 go run *.go
	DATABASE_URL=postgres://root:dXNoei3z2xXK6AxP@main.cjoxzdbtsn7e.eu-central-1.rds.amazonaws.com:5432/muni_surveys PORT=8080 go run *.go

voyager:
	docker run --rm -v `pwd`/gen/schema.graphql:/app/schema.graphql -p 8080:80 graphql/voyager

build-lambda-function:
	GO111MODULE=on GOOS=linux go build -o main lambda/main.go && zip lambda.zip main && rm main

test:
	go build -o app *.go && (DATABASE_URL=sqlite3://test.db PORT=8080 ./app& export app_pid=$$! && make test-godog || test_result=$$? && kill $$app_pid && exit $$test_result)
// TODO: add detection of host ip (eg. host.docker.internal) for other OS
test-godog:
	docker run --rm --network="host" -v "${PWD}/features:/godog/features" -e GRAPHQL_URL=http://$$(if [[ $${OSTYPE} == darwin* ]]; then echo host.docker.internal;else echo localhost;fi):8080/graphql jakubknejzlik/godog-graphql

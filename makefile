
OWNER=jakubknejzlik
IMAGE_NAME=surveyjs-graphql-api
QNAME=$(OWNER)/$(IMAGE_NAME)

GIT_TAG=$(QNAME):$(TRAVIS_COMMIT)
BUILD_TAG=$(QNAME):0.1.$(TRAVIS_BUILD_NUMBER)
LATEST_TAG=$(QNAME):latest

lint:
	docker run -it --rm -v "$(PWD)/Dockerfile:/Dockerfile:ro" redcoolbeans/dockerlint

build:
	go get ./...
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
	go run github.com/novacloudcz/graphql-orm

run:
	DATABASE_URL=sqlite3://test.db PORT=8080 go run *.go

voyager:
	docker run --rm -v `pwd`/gen/schema.graphql:/app/schema.graphql -p 8080:80 graphql/voyager

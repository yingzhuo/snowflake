TIMESTAMP   := $(shell /bin/date "+%F %T")
NAME		:= snowflake
VERSION		:= 1.1.1
LDFLAGS		:= -s -w \
			   -X 'main.BuildVersion=$(VERSION)' \
			   -X 'main.BuildGitBranch=$(shell git describe --all)' \
			   -X 'main.BuildGitRev=$(shell git rev-list --count HEAD)' \
			   -X 'main.BuildGitCommit=$(shell git rev-parse HEAD)' \
			   -X 'main.BuildDate=$(shell /bin/date "+%F %T")'

no_default:
	@echo "no default target"

fmt:
	@go fmt $(CURDIR)/...
	@go mod tidy

clean:
	@rm -rf $(CURDIR)/_bin/snowflake-* &> /dev/null
	@docker image rm registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest &> /dev/null || true
	@docker image rm registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) &> /dev/null || true
	@docker image prune -f &> /dev/null || true

install:
	@CGO_ENABLED=0 sudo go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o /usr/local/bin/$(NAME)

uninstall:
	@sudo rm -rf /usr/local/bin/$(NAME) &> /dev/null || true

build:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-linux-amd64-$(VERSION)
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-linux-macos-$(VERSION)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-linux-windows-$(VERSION).exe

release: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-linux-amd64-$(VERSION)
	docker login --username=yingzhor@gmail.com --password="${ALIYUN_PASSWORD}" registry.cn-shanghai.aliyuncs.com
	docker image build -t registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) --no-cache $(CURDIR)/_bin
	docker image push registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION)
	docker image tag  registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest
	docker image push registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest
	docker logout registry.cn-shanghai.aliyuncs.com &> /dev/null

github: clean fmt
	git add .
	git commit -m "$(TIMESTAMP)"
	git push

.PHONY: fmt clean github build install uninstall release
TIMESTAMP             	:= $(shell /bin/date "+%F %T")
NAME					:= snowflake
VERSION					:= 1.0.1

usage:
	@echo "------------------------------------------"
	@echo " 目标           | 功能"
	@echo "------------------------------------------"
	@echo " usage          | 显示本菜单"
	@echo " fmt            | 格式化代码"
	@echo " protoc         | 编译protobuf文件"
	@echo " build-linux    | 构建 (linux-amd64)"
	@echo " build-darwin   | 构建 (darwin-amd64)"
	@echo " build-windows  | 构建 (windows-amd64)"
	@echo " build-all      | 构建以上三者"
	@echo " clean          | 清理构建产物"
	@echo " release        | 发布"
	@echo " github         | 将代码推送到Github"
	@echo "------------------------------------------"

fmt:
	@go fmt $(CURDIR)/...

clean:
	@rm -rf $(CURDIR)/_bin/snowflake-* &> /dev/null
	@docker image rm registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest &> /dev/null || true
	@docker image rm registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) &> /dev/null || true
	@docker image prune -f &> /dev/null || true

protoc:
	protoc -I=$(CURDIR)/_proto/ --go_out=$(CURDIR) $(CURDIR)/_proto/snowflake.proto

build-linux: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(CURDIR)/_bin/$(NAME)-linux-amd64-$(VERSION)

build-darwin: protoc
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(CURDIR)/_bin/$(NAME)-darwin-amd64-$(VERSION)

build-windows: protoc
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(CURDIR)/_bin/$(NAME)-windows-amd64-$(VERSION).exe

build-all: build-linux build-darwin build-windows

release: build-linux
	docker login --username=yingzhor@gmail.com --password="${ALIYUN_PASSWORD}" registry.cn-shanghai.aliyuncs.com
	docker image build -t registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) --build-arg VERSION=$(VERSION) $(CURDIR)/_bin
	docker image push registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION)
	docker image tag  registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest
	docker image push registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest
	docker logout registry.cn-shanghai.aliyuncs.com &> /dev/null

github: clean fmt
	git add .
	git commit -m "$(TIMESTAMP)"
	git push

.PHONY: usage fmt clean protoc build-linux build-darwin build-windows build-all release github
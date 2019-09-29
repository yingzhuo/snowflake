# ----------------------------------------------------------------------------------------------------------------------
#  作者: 应卓
#  日期: 2019-09-16
# ----------------------------------------------------------------------------------------------------------------------
timestamp             = $(shell /bin/date "+%F %T")
gofiles               = `find $(CURDIR) -name "*.go" -type f -not -path "$(CURDIR)/src/vendor/*"`
packages              = `go list $(CURDIR)/... | grep -v /vendor/`
docker-image-tag      = quay.io/yingzhuo/snowflake:latest
docker-build-context  = $(CURDIR)/bin/

# 菜单
usage:
	@echo "------------------------------------------"
	@echo " 目标           | 功能"
	@echo "------------------------------------------"
	@echo " usage          | 显示本菜单"
	@echo " fmt            | 格式化代码"
	@echo " list           | 列出所有包和文件"
	@echo " proto          | 编译protobuf文件"
	@echo " clean          | 清理构建产物"
	@echo " release        | 发布"
	@echo " github         | 将代码推送到Github"
	@echo "------------------------------------------"

# 列出所有包
list:
	@echo "packages: "
	@echo $(packages)
	@echo ""
	@echo "go files: "
	@echo $(gofiles)

# 格式化代码
fmt:
	@gofmt -s -w $(gofiles)

# 清理
clean:
	@rm -rf $(CURDIR)/bin/snowflake-* &> /dev/null || true
	@docker image rm $(docker-image-tag) &> /dev/null || true
	@docker image prune -f &> /dev/null || true

# 编译protobuf
proto: clean
	protoc -I=$(CURDIR)/proto/ --go_out=$(CURDIR)/src/ $(CURDIR)/proto/snowflake.proto

# 发布
release: clean fmt proto
	GOOS=linux   GOARCH=amd64 go build -o $(CURDIR)/bin/snowflake-linux-amd64       github.com/yingzhuo/main
	docker image build -t $(docker-image-tag) $(docker-build-context)

	@cat $(CURDIR)/.github/quay.io.pwd | docker login --username=yingzhuo --password-stdin quay.io
	docker image push $(docker-image-tag)
	@docker logout quay.io &> /dev/null

# 推送源代码
github: clean fmt
	git add .
	git commit -m "$(timestamp)"
	git push

.PHONY: usage fmt list clean release github

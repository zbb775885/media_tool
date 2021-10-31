
GO=go

TARGET=media_tool

SRC=$(shell find ./ -name "*.go")

#设置cgo环境变量
#CGO_LIB=$(shell ls  ${OSP_LIB} |grep -E ".a"|xargs -I {} ls ${OSP_LIB}/{})
export CGO_LDFLAGS=-L $(OSP_LIB) -l:libavfilter.a -l:libswscale.a -l:libswresample.a -l:libavformat.a -l:libavcodec.a  -l:libavdevice.a -l:libavutil.a -lm -lpthread
export CGO_CFLAGS=-I $(OSP_INC)
export CGO_CXXFLAGS=-I $(OSP_INC)
export CGO_CPPFLAGS=-I $(OSP_INC)
export CGO_ENABLE=1

all:$(TARGET)

docker:
	$(call docker_env)

godefs:$(shell find ./pkg/cgo -name "*.go")
	$(GO) tool cgo -godefs $<

$(TARGET):$(SRC)
	$(GO) build -o $(TARGET) $<

clean:
	@rm -f $(TARGET)

.PHONY: $(TARGET) clean all

PROJECT_PATH=/go/path/src/github.com/zbb775885/media_tool

define  docker_env
	docker run --rm --name=media_tool \
	--env PROJECT_PATH=${PROJECT_PATH} \
	-v ${PWD}:${PROJECT_PATH} \
	zbb775885/ubuntu:v2 \
	bash -c "source /etc/profile && cd ${PROJECT_PATH} && make -j"
endef
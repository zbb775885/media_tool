package ffmpeg_mux // import "github.com/zbb775885/media_tool/pkg/imp/ffmpeg/mux"

import (
	log "github.com/sirupsen/logrus"
	mal_mux "github.com/zbb775885/media_tool/pkg/mal/format/mux"
)

type FFmpegMux struct {
	Malmux mal_mux.MalMuxerParam
}

//创建封装器
func (fd *FFmpegMux) MuxerInit(initParam mal_mux.MalMuxerParam) string {
	log.Info("xxxxxxxx")

	return ""
}

//销毁封装器
func (fd *FFmpegMux) MuxerDeinit() string {
	return ""
}

//配置封装器
func (fd *FFmpegMux) MuxerConfig(runParam mal_mux.MalMuxerParam) string {
	return ""
}

//启动封装器
func (fd *FFmpegMux) MuxerStart() string {
	return ""
}

//停止封装器
func (fd *FFmpegMux) MuxerStop() string {
	return ""
}

//解封装构造器构造器
type FFmpegMuxCreator struct {
}

//创建FFmpegMux对象
func (fd *FFmpegMuxCreator) Create() mal_mux.MalMuxerInterface {
	return &FFmpegMux{}
}

//mux 注册
func Register() {
	mal_mux.Register("ffmpeg_mux", &FFmpegMuxCreator{})
}

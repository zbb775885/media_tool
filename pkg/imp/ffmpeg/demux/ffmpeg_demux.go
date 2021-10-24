package ffmpeg_demux

import (
	log "github.com/sirupsen/logrus"
	mal_demux "github.com/zbb775885/media_tool/pkg/mal/format/demux"
)

type FFmpegDemux struct {
	MalDemux mal_demux.MalDemuxerParam
}

//创建封装器
func (fd *FFmpegDemux) DeMuxerInit(initParam mal_demux.MalDemuxerParam) string {
	log.Info("xxxxxxxx")

	return ""
}

//销毁封装器
func (fd *FFmpegDemux) DeMuxerDeinit() string {
	return ""
}

//配置封装器
func (fd *FFmpegDemux) DeMuxerConfig(runParam mal_demux.MalDemuxerParam) string {
	return ""
}

//启动封装器
func (fd *FFmpegDemux) DeMuxerStart() string {
	return ""
}

//停止封装器
func (fd *FFmpegDemux) DeMuxerStop() string {
	return ""
}

type CreateDemux struct {
}

//创建FFmpegDemux对象
func (fd *CreateDemux) Create() mal_demux.MalDemuxerInterface {
	return &FFmpegDemux{}
}

//Demux 注册
func Register() {
	mal_demux.Register("ffmpeg_demux", &CreateDemux{})
}

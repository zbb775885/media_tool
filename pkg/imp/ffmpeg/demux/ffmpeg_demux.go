package ffmpeg_demux // import "github.com/zbb775885/media_tool/pkg/imp/ffmpeg/demux"

import (
	log "github.com/sirupsen/logrus"
	"github.com/zbb775885/goav/avformat"
	mal_demux "github.com/zbb775885/media_tool/pkg/mal/format/demux"
)

type FFmpegDemux struct {
	MalDemux mal_demux.MalDemuxerParam
}

//创建封装器
func (fd *FFmpegDemux) DemuxerInit(initParam mal_demux.MalDemuxerParam) string {
	log.Info("AvformatVersion() is ", avformat.AvformatVersion())

	return ""
}

//销毁封装器
func (fd *FFmpegDemux) DemuxerDeinit() string {
	return ""
}

//配置封装器
func (fd *FFmpegDemux) DemuxerConfig(runParam mal_demux.MalDemuxerParam) string {
	return ""
}

//启动封装器
func (fd *FFmpegDemux) DemuxerStart() string {
	return ""
}

//停止封装器
func (fd *FFmpegDemux) DemuxerStop() string {
	return ""
}

//解封装构造器构造器
type FFmpegDemuxCreator struct {
}

//创建FFmpegDemux对象
func (fd *FFmpegDemuxCreator) Create() mal_demux.MalDemuxerInterface {
	return &FFmpegDemux{}
}

//Demux 注册
func Register() {
	mal_demux.Register("ffmpeg_demux", &FFmpegDemuxCreator{})
}

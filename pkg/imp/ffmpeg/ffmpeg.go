package ffmpeg

import (
	//cgo_avfilter "github.com/zbb775885/media_tool/pkg/cgo/ffmpeg"
	demux "github.com/zbb775885/media_tool/pkg/imp/ffmpeg/demux"
)

const (
	INIT_ERR = iota // 0
)

type FFmpeg struct {
}

//错误码
func (ff *FFmpeg) Error(errorCode int32) string {
	var err_str string

	switch errorCode {
	case INIT_ERR:
		err_str = "init fail"
	default:
		err_str = "unknow"
		err_str = "unknow"
	}

	return err_str
}

//初始化ffmpeg,打开ffmpeg动态库
func Init() string {
	// log.Info("filter version is ", cgo_avfilter.AvfilterVersion())
	// log.Info("filter license is ", cgo_avfilter.AvfilterLicense())
	// log.Info("filter configuration is ", cgo_avfilter.AvfilterConfiguration())
	//log.Info("filter pad type is ", cgo_avfilter.AvfilterPadGetType(nil, 0))

	//var data cgo_avfilter.AVFilter

	//log.Info("data is ", data.name)
	// ffmpeg := new(FFmpeg)
	// if nil == ffmpeg {
	// 	log.Error("new ffmpeg fail")
	// }

	//注册解封装器的实现
	demux.Register()

	//注册封装器的实现
	//mux.Register()

	return ""
}

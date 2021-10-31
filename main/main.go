package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	ffmpeg "github.com/zbb775885/media_tool/pkg/imp/ffmpeg"
	mal "github.com/zbb775885/media_tool/pkg/mal"
	mal_demux "github.com/zbb775885/media_tool/pkg/mal/format/demux"
)

//参考ffmpeg配置输入输出参数
var inputFile = flag.String("i", " ", "input file name")
var outputFile = flag.String("o", " ", "output file name")
var inputFormat = flag.String("if", " ", "input file format")
var outputFormat = flag.String("f", " ", "out file format")
var formatsSupported = flag.String("formats", "es,mp4,rtp,ps", "format support")

//init logrus log format
func InitLog() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}

func main() {
	//初始化日志模块
	InitLog()

	//输入参数解析
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	log.Info("input file is ", *inputFile)
	log.Info("output file is ", *outputFile)

	//ffmpeg实现层注册及初始化
	ffmpeg.Init()

	//媒体抽象层初始化
	mal.Init()

	mal_demux.Create("ffmpeg_demux").DemuxerInit(mal_demux.MalDemuxerParam{})
}

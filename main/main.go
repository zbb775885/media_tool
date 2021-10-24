package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/zbb775885/media_tool/pkg/ffmpeg"
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

	if 0 == len(flag.Args()) {
		flag.Usage()
		return
	}

	log.Info("input file is ", *inputFile)
	log.Info("output file is ", *outputFile)

	ffmpeg.Init()
}

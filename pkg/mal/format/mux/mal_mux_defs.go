// 定义媒体抽象层接口以及数据类型
package mal_format_mux

import (
	fmt "github.com/zbb775885/media_tool/pkg/mal/format"
)

//封装初始化参数
type MalMuxerParam struct {
	//输出的封装类型
	OutputFmt fmt.MediaFormatType
}

//封装的抽象接口
type MalMuxerInterface interface {
	//创建封装器
	MuxerInit(initParam MalMuxerParam) string
	//销毁封装器
	MuxerDeinit() string
	//配置封装器
	MuxerConfig(runParam MalMuxerParam) string
	//启动封装器
	MuxerStart() string
	//停止封装器
	MuxerStop() string
}

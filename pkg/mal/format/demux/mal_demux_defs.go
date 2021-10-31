// 定义媒体抽象层接口以及数据类型
package mal_format_demux // import "github.com/zbb775885/media_tool/pkg/mal/format/demux"

import (
	fmt "github.com/zbb775885/media_tool/pkg/mal/format"
)

//封装初始化参数
type MalDemuxerParam struct {
	//输出的封装类型
	InputFmt fmt.MediaFormatType
}

//封装的抽象接口
type MalDemuxerInterface interface {
	//创建封装器
	DemuxerInit(initParam MalDemuxerParam) string
	//销毁封装器
	DemuxerDeinit() string
	//配置封装器
	DemuxerConfig(runParam MalDemuxerParam) string
	//启动封装器
	DemuxerStart() string
	//停止封装器
	DemuxerStop() string
}

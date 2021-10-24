// 定义媒体抽象层接口以及数据类型
package media_abstract_layer_format

import (
	fmt "github.com/zbb775885/media_tool/pkg/mal/format"
)

//封装初始化参数
type MalDeMuxerParam struct {
	//输出的封装类型
	InputFmt fmt.MediaFormatType
}

//封装的抽象接口
type MalDeMuxer interface {
	//创建封装器
	DeMuxerCreate(initParam MalDeMuxerParam) error
	//销毁封装器
	DeMuxerDestroy() error
	//配置封装器
	DeMuxerConfig(runParam MalDeMuxerParam) error
	//启动封装器
	DeMuxerStart() error
	//停止封装器
	DeMuxerStop() error
}

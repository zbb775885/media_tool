// 定义媒体抽象层接口以及数据类型
package mal

//封装类型
type MediaFormatType int32

//封装格式
const (
	//es流
	MediaFormatEs MediaFormatType = iota
	//RTP流
	MediaFormatRtp
	//PS流
	MediaFormatPs
	//MP4流
	MediaFormatMp4
)

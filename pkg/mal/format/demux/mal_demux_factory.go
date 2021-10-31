// 定义媒体抽象层接口以及数据类型
package mal_format_demux

import (
	log "github.com/sirupsen/logrus"
)

//封装的抽象接口
type MalCreate interface {
	Create() MalDemuxerInterface
}

//全局函数表
var mapMethodNoArg = map[string]MalCreate{}

//注册函数
func Register(name string, function MalCreate) {
	if _, ret := mapMethodNoArg[name]; ret {
		mapMethodNoArg[name] = function
	} else {
		log.Errorf("function %s has been insert", name)
	}
}

//删除函数
func UnRegister(name string) {
	if _, ret := mapMethodNoArg[name]; ret {
		delete(mapMethodNoArg, name)
	} else {
		log.Errorf("function %s has not been insert", name)
	}
}

//删除函数
func Create(name string) MalDemuxerInterface {
	if _, ret := mapMethodNoArg[name]; ret {
		return mapMethodNoArg[name].Create()
	} else {
		log.Errorf("function %s has not been insert", name)
	}

	return nil
}

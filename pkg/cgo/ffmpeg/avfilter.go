package cgo_ffmpeg

// #include "libavfilter/avfilter.h"
import "C"

import (
	"unsafe"
)

//redefine filter struct
type AVFilterContext C.struct_AVFilterContext

type AVFilterLink C.struct_AVFilterLink

type AVFilterPad C.struct_AVFilterPad

type AVFilterFormats C.struct_AVFilterFormats

type AVFilterChannelLayouts C.struct_AVFilterChannelLayouts

type AVFilter C.struct_AVFilter

type AVFilterInternal C.struct_AVFilterInternal

type AVFilterFormatsConfig C.struct_AVFilterFormatsConfig

type AVFilterGraph C.struct_AVFilterGraph

//function definition

//return version of avfilter
func AvfilterVersion() uint32 {
	return uint32(C.avfilter_version())
}

//return the libavfilter build-time configuration.
func AvfilterConfiguration() string {
	return C.GoString(C.avfilter_configuration())
}

//return the libavfilter license.
func AvfilterLicense() string {
	return C.GoString(C.avfilter_license())
}

//Get the name of an AVFilterPad
func AvfilterPadGetName(avfilter_pad *AVFilterPad, pad_idx int32) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(unsafe.Pointer(avfilter_pad)), C.int(pad_idx)))
}

//Get the type of an AVFilterPad
func AvfilterPadGetType(avfilter_pad *AVFilterPad, pad_idx int32) AVMediaType {
	return AVMediaType(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(unsafe.Pointer(avfilter_pad)), C.int(pad_idx)))
}

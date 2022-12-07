package it8951

// #cgo LDFLAGS: -lbcm2835 -lm -lrt -lpthread
// #include "IT8951.h"
import "C"

func Init() {
	C.ext_IT8951_init()
}

func Width() int {
	return int(C.ext_IT8951_width())
}

func Height() int {
	return int(C.ext_IT8951_height())
}

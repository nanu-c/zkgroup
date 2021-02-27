package zkgroup

/*
#cgo CFLAGS: -I${SRCDIR}/lib
#cgo linux,arm64 LDFLAGS: ${SRCDIR}/lib/libzkgroup_linux_arm64.so -Wl,-rpath='$ORIGIN'
#cgo linux,arm LDFLAGS:  ${SRCDIR}/lib/libzkgroup_linux_armhf.so  -Wl,-rpath='$ORIGIN'
#cgo linux,amd64 LDFLAGS:  ${SRCDIR}/lib/libzkgroup_linux_amd64.so  -Wl,-rpath='$ORIGIN'
#include <zkgroup.h>
*/
import "C"
import (
	"crypto/rand"
	"unsafe"
)

func cBytes(b []byte) *C.uchar {
	return (*C.uchar)(unsafe.Pointer(&b[0]))
}

func cLen(b []byte) C.uint32_t {
	return C.uint32_t(len(b))
}

func randBytes(length int) []byte {
	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	return buf
}

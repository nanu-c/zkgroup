package zkgroup

/*
#cgo CFLAGS: -I${SRCDIR}/lib
#cgo linux,arm64 LDFLAGS: ${SRCDIR}/lib/libzkgroup_linux_arm64.so -Wl,-rpath='$ORIGIN'
#cgo linux,arm LDFLAGS:  ${SRCDIR}/lib/libzkgroup_linux_armhf.so  -Wl,-rpath='$ORIGIN'
#cgo linux,amd64 LDFLAGS:  ${SRCDIR}/lib/libzkgroup_linux_amd64.so  -Wl,-rpath='$ORIGIN'
#include <zkgroup.h>
*/
import "C"

const (
	profileKeyCommitmentSize = 97
	profileKeyVersionSize    = 64
)

func ProfileKeyGetCommitment(profileKey []byte, uuid []byte) (ServerSecretParams, error) {
	out := make([]byte, profileKeyCommitmentSize)
	if res := C.FFI_ProfileKey_getCommitment(cBytes(profileKey), cLen(profileKey), cBytes(uuid), cLen(uuid), cBytes(out), cLen(out)); res != C.FFI_RETURN_OK {
		return nil, errFromCode(res)
	}
	return ServerSecretParams(out), nil
}

// ProfileKeyGetProfileKeyVersion returns the profile key version
func ProfileKeyGetProfileKeyVersion(profileKey []byte, uuid []byte) (ServerSecretParams, error) {
	out := make([]byte, profileKeyVersionSize)
	if res := C.FFI_ProfileKey_getProfileKeyVersion(cBytes(profileKey), cLen(profileKey), cBytes(uuid), cLen(uuid), cBytes(out), cLen(out)); res != C.FFI_RETURN_OK {
		return nil, errFromCode(res)
	}
	return ServerSecretParams(out), nil
}

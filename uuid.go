/*
#######
##                  _    __
##       __ ____ __(_)__/ /
##      / // / // / / _  /
##      \_,_/\_,_/_/\_,_/
##
####### Copyright (c) 2020-2021 mls-361 ############################################################ MIT License #######
*/

package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	mathrand "math/rand"
)

func build(reader io.Reader, errorAllowed bool) (string, error) {
	buf := make([]byte, 16)

	if _, err := io.ReadFull(reader, buf); err != nil {
		if errorAllowed {
			return "", err
		}

		mathrand.Read(buf)
	}

	uuid := make([]byte, 36)

	hex.Encode(uuid, buf[:4])
	uuid[8] = '-'
	hex.Encode(uuid[9:13], buf[4:6])
	uuid[13] = '-'
	hex.Encode(uuid[14:18], buf[6:8])
	uuid[18] = '-'
	hex.Encode(uuid[19:23], buf[8:10])
	uuid[23] = '-'
	hex.Encode(uuid[24:], buf[10:])

	return string(uuid), nil

}

// New AFAIRE.
func New() string {
	uuid, _ := build(rand.Reader, false)
	return uuid
}

// Build AFAIRE.
func Build() (string, error) { return build(rand.Reader, true) }

/*
######################################################################################################## @(°_°)@ #######
*/

// You can edit this code!
// Click here and start typing.
package utilities

import (
	"math/rand"
	"time"

	"log"
	"fmt"
)

func GenerateOTP() (string) {
	// This was a an intrestng thing i found actually i was confused that its generating the same OTP again and again so checked that the default dataset will return predictable value that alwasy the same value that is Seed(1) so that concurrent go routines can work with same data but if we need randomness we need to seed with a new source which would help us generate randomness.
	// https://pkg.go.dev/math/rand
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%04d", rand.Intn(9999))
	log.Printf("Generated OTP: %s", otp)
	return otp
}

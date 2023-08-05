package color

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// Convert a string to one of the 12 colors of the color wheel.
func ConvertStrToColor(s string) string {
	colors := [12]string{"red", "red-orange", "orange", "yellow-orange", "yellow", "yellow-green", "green", "blue-green", "blue", "blue-violet", "violet", "red-violet"}

	// Compute the md5 hash of the string and only keep the first 8 digits
	hash := md5.Sum([]byte(s))
	hashStr := hex.EncodeToString(hash[:])[:8]

	// Convert the hash to a number
	n, _ := strconv.ParseInt(hashStr, 16, 64)
	return colors[n%12]
}

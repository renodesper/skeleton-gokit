package cursor

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"gitlab.com/renodesper/gokit-microservices/util/errors"
)

// EncodeCursor ...
func EncodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

// DecodeCursor ...
func DecodeCursor(encodedCursor string) (createdAt time.Time, uuid string, err error) {
	bytes_, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return
	}

	arrStr := strings.Split(string(bytes_), ",")
	if len(arrStr) != 2 {
		err = errors.InvalidCursor
		return
	}

	createdAt, err = time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return
	}

	uuid = arrStr[1]
	return
}

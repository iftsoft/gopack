package lla

import (
	"strings"
	"time"
	"io"
	"fmt"
	"crypto/rand"
)

// Format to camel string, xx_yy to XxYy
func FormatCamelString(s string) string {
	data := make([]byte, 0, len(s))
	flag, num := true, len(s)-1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' {
			flag = true
			continue
		} else if flag {
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
			flag = false
		}
		data = append(data, d)
	}
	return string(data[:])
}

// Format to snake string, XxYy to xx_yy , XxYY to xx_yy
func FormatSnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// Parse time to string with location
func TimeFormatParse(moment, format string) (time.Time, error) {
	tp, err := time.ParseInLocation(format, moment, time.Local)
	return tp, err
}

func DateTimeParse(moment string) (time.Time, error) {
	return TimeFormatParse(moment, "2006-01-02 15:04:05.999999")
}
func DateParse(moment string) (time.Time, error) {
	return TimeFormatParse(moment, "2006-01-02")
}
func TimeParse(moment string) (time.Time, error) {
	return TimeFormatParse(moment, "15:04:05")
}

// newUUID generates a random UUID according to RFC 4122
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}


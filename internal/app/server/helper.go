package server

import (
	"log"
	"strconv"
	"time"
)

func CheckErr(err error, msg string) {
	if msg == "" {
		msg = "something wrong"
	}
	if err != nil {
		log.Fatalf("[ERR] %s: %v", msg, err)
	}
}

func timeFromUnix(unixTime string) (time.Time, error) {
	s, err := strconv.ParseInt(unixTime, 10, 64)
	CheckErr(err, "timeFromUnix")
	return time.Unix(s, 0), err
}

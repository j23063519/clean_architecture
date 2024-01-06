package util

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/j23063519/clean_architecture/config"
)

// transform time.Duration to microseconds and string type
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// Remove file extension
func FileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// tansform string type of time to time.Time type of time
func StrTimeStampToTime(str string) (t time.Time) {
	if str == "" {
		return
	}
	secs, err := strconv.ParseInt(str[:10], 10, 64)
	if err != nil {
		log.Fatal("StrTimeStampToTime", "ParseInt", err)
		return
	}

	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Fatal("StrTimeStampToTime", "LoadLocation", err)
		return
	}

	microsStr := str[10:]
	micros, err := strconv.Atoi(microsStr)
	if err != nil {
		log.Fatal("StrTimeStampToTime", "Atoi", err)
		return
	}

	nanos := time.Duration(micros) * time.Microsecond
	nanosInt64 := int64(nanos)

	t = time.Unix(secs, nanosInt64).In(loc)

	return
}

// current time zone of time.Time
func TimeWithLocation(t time.Time) time.Time {
	timezone, _ := time.LoadLocation(config.Config.App.TIMEZONE)
	return t.In(timezone)
}

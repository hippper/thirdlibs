package util

import (
	"testing"
	"time"
)

func TestTimestamp2str(t *testing.T) {
	timestampNow := time.Now().Unix()
	timeStrNow := time.Unix(timestampNow, 0).Format(TIME_FORMAT)

	res := Timestamp2str(timestampNow, TIME_FORMAT)
	if timeStrNow != res {
		t.Error("Timestamp2str fail.")
	} else {
		t.Log("Timestamp2str success")
	}
}

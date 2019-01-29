package websms

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"strings"
	"testing"
)

var testDispatchStatus = map[string]DispatchStatus{
	`Accepted`:      StatusAccepted,
	`UNKNOWN`:       StatusUnknown,
	`Ready`:         StatusReady,
	`Spooled`:       StatusSpooled,
	`Sent`:          StatusSent,
	`Rejected`:      StatusRejected,
	`Delivered`:     StatusDelivered,
	`Undeliverable`: StatusUndeliverable,
	`Enroute`:       StatusEnroute,
	`Deleted`:       StatusDeleted,
	`Expired`:       StatusExpired,
	`Delayed`:       StatusDelayed,
	`Restricted`:    StatusRestricted,
	`Unroutable`:    StatusUnroutable,
	`NOT FOUND`:     StatusMessageIDNotFound,
}

func TestNewDispatchStatus(t *testing.T) {
	var key string
	var sta DispatchStatus

	for key = range testDispatchStatus {
		sta = NewDispatchStatus(key)
		if strings.ToLower(key) != sta.String() {
			t.Fatal("NewDispatchStatus error value")
		}
		if sta != testDispatchStatus[key] {
			t.Fatal("NewDispatchStatus error object")
		}
		if key == sta.String() {
			t.Fatal("NewDispatchStatus bat test")
		}
	}
	sta = NewDispatchStatus(``)
	if sta != StatusUnknown {
		t.Fatal("NewDispatchStatus error default value")
	}
	sta = NewDispatchStatus(`H9yM7c8AenA9583k27Xg`)
	if sta != StatusUnknown {
		t.Fatal("NewDispatchStatus error default value")
	}
}

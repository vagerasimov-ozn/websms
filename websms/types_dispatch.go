package websms // import "gopkg.in/webnice/websms.v1/websms"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"strings"
)

// DispatchStatus Статус отправки сообщения
type DispatchStatus string

// String Return dispatch status as string
func (ds DispatchStatus) String() string { return string(ds) }

// NewDispatchStatus Convert string value to status constant
func NewDispatchStatus(s string) (ret DispatchStatus) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case StatusAccepted.String():
		ret = StatusAccepted
	case StatusUnknown.String():
		ret = StatusUnknown
	case StatusReady.String():
		ret = StatusReady
	case StatusSpooled.String():
		ret = StatusSpooled
	case StatusSent.String():
		ret = StatusSent
	case StatusRejected.String():
		ret = StatusRejected
	case StatusDelivered.String():
		ret = StatusDelivered
	case StatusUndeliverable.String():
		ret = StatusUndeliverable
	case StatusEnroute.String():
		ret = StatusEnroute
	case StatusDeleted.String():
		ret = StatusDeleted
	case StatusExpired.String():
		ret = StatusExpired
	case StatusDelayed.String():
		ret = StatusDelayed
	case StatusRestricted.String():
		ret = StatusRestricted
	case StatusUnroutable.String():
		ret = StatusUnroutable
	case StatusMessageIDNotFound.String():
		ret = StatusMessageIDNotFound
	default:
		ret = StatusUnknown
	}

	return
}

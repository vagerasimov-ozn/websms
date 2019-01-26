/*

	Реализация XML протокола описанного в документе:
	http://websms.ru/content/doc/HTTPXMLsendmethod_v1.5.1.pdf?196
	Версия 1.5.1

*/

package websms // import "gopkg.in/webnice/websms.v1/websms"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"time"

	websmsTypes "gopkg.in/webnice/websms.v1/websms/types"
)

func init() {
	defaultConfiguration = &websmsTypes.DefaultConfiguration{}
}

// New Create new object and return interface
func New(cfg ...*Configuration) Interface {
	var wsm = new(impl)

	wsm.from = defaultConfiguration.From
	if len(cfg) > 0 {
		wsm.cfg = cfg[0]
	}
	if wsm.cfg == nil {
		wsm.cfg = &Configuration{
			Username: defaultConfiguration.Username,
			Password: defaultConfiguration.Password,
		}
	}

	return wsm
}

// DefaultFrom Set default sender name for new objects
func DefaultFrom(from string) { defaultConfiguration.From = from }

// DefaultUsername Set default username for new objects
func DefaultUsername(username string) { defaultConfiguration.Username = username }

// DefaultPassword Set default password for new objects
func DefaultPassword(password string) { defaultConfiguration.Password = password }

// Testing Set testing flag
func (wsm *impl) Testing(t bool) Interface { wsm.test = t; return wsm }

// From is an sender name, if not specified selects from service profile settings
func (wsm *impl) From(from string) Interface { wsm.from = from; return wsm }

// Message Send a single message at now to one address
func (wsm *impl) Message(msg *Message, to string) (*Status, error) {
	return wsm.messageToAnyAt(msg, 0, time.Time{}, to)
}

// MessageToAny Send a single message at now to multiple address
func (wsm *impl) MessageToAny(msg *Message, pkgID uint64, to ...string) (*Status, error) {
	return wsm.messageToAnyAt(msg, pkgID, time.Time{}, to...)
}

// MessageAt Send a single message at the specified time to one address
func (wsm *impl) MessageAt(msg *Message, start time.Time, to string) (*Status, error) {
	return wsm.messageToAnyAt(msg, 0, start, to)
}

// MessageToAnyAt Send a single message at the specified time to multiple address
func (wsm *impl) MessageToAnyAt(msg *Message, pkgID uint64, start time.Time, to ...string) (*Status, error) {
	return wsm.messageToAnyAt(msg, pkgID, start, to...)
}

// Messages Send a multiple message as a one packet at now to one address
func (wsm *impl) Messages(msgs []*MultiMessage, pkgID uint64) (*Status, error) {
	return wsm.messagesToAnyAt(msgs, pkgID, time.Time{})
}

// MessagesAt Send a multiple message as a one packet at the specified time to one address
func (wsm *impl) MessagesAt(msgs []*MultiMessage, pkgID uint64, start time.Time) (*Status, error) {
	return wsm.messagesToAnyAt(msgs, pkgID, start)
}

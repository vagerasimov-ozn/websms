package websms // import "gopkg.in/webnice/websms.v1/websms"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
//import ()

// ErrCode Type of error code
type ErrCode struct {
	code  uint16
	errEn string
	errRu string
}

// NewErrCode Convert uint16 value to error constant
func NewErrCode(code uint16) (ret ErrCode) {
	var ok bool
	if ret, ok = errCodeMap[code]; !ok {
		ret = errCodeMap[cErrCodeXX]
		return
	}
	return
}

// String Return as string
func (ec ErrCode) String() string { return ec.errEn }

// StringEN Return as string in english
func (ec ErrCode) StringEN() string { return ec.errEn }

// StringRU Return as string in russian
func (ec ErrCode) StringRU() string { return ec.errRu }

// Int Return as int
func (ec ErrCode) Int() int { return int(ec.code) }

// Uint16 Return as uint16
func (ec ErrCode) Uint16() uint16 { return ec.code }

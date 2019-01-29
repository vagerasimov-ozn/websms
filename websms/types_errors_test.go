package websms

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

func TestNewErrCode(t *testing.T) {
	err18 := NewErrCode(cErrCode18)
	err19 := NewErrCode(cErrCode19)
	if err18.Uint16() != err19.Uint16() && err18.Uint16() != cErrCodeXX {
		t.Fatalf("Error in NewErrCode()")
	}
	err00 := NewErrCode(0)
	if err00.StringEN() != errCodeMap[cErrCode00].errEn {
		t.Fatalf("Error in NewErrCode()")
	}
	if err00.StringRU() != errCodeMap[cErrCode00].errRu {
		t.Fatalf("Error in NewErrCode()")
	}
	if err00.String() != errCodeMap[cErrCode00].errEn {
		t.Fatalf("Error in NewErrCode()")
	}
	err23 := NewErrCode(23)
	if err23.Int() != int(err23.Uint16()) && err23.code != cErrCode23 {
		t.Fatalf("Error in NewErrCode()")
	}
}

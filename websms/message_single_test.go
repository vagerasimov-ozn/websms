package websms

//import "gopkg.in/webnice/debug.v1"
import "gopkg.in/webnice/log.v2"
import (
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gopkg.in/webnice/web.v1/status"
)

type (
	testSingleSrvHandler struct{ http.Handler }
)

func (tsh *testSingleSrvHandler) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	var err error
	var buf []byte
	var hsh hash.Hash
	var sum string

	if buf, err = ioutil.ReadAll(rq.Body); err != nil {
		wr.WriteHeader(status.InternalServerError)
		return
	}
	defer func() { _ = rq.Body.Close() }()
	//log.Debug(string(buf))
	hsh = sha512.New()
	if _, err = hsh.Write(buf); err != nil {
		wr.WriteHeader(status.InternalServerError)
		return
	}
	sum = hex.EncodeToString(hsh.Sum(nil))
	wr.WriteHeader(status.Ok)
	_, err = wr.Write([]byte(`<status id="1234567890" date="01.02.2019 3:45:17"><state error="` + sum + `" errcode="0">Accepted</state></status>`))
	if err != nil {
		log.Fatalf("response error: %s", err)
	}
}

func TestMessage(t *testing.T) {
	const (
		username  = `2qUTSpRYgZHcPYxm2zhq@Qx645rznnuZYWKr7HrbR.xz`
		password  = `bA7RufvpnEeKuF8NqGBn`
		fromname  = `JQNeFT6bbmedtUARTqFe`
		messageID = 12345677654321
		body      = `gMK2aH5ZRT93RBmxC2wC`
		toNumber1 = `11234567890`
		shaSumm   = `07c0dfe8bc7a37dcec95982874f8908533633e048dff7ea75fd7aabca5b6c1151dffe11033cc30fe1580134f05a9802707cedc451caf676f3fe345441541e0b3`
	)
	var err error
	var srvHndl *testSingleSrvHandler
	var srv *httptest.Server
	var obj Interface
	var st *Status

	srvHndl = &testSingleSrvHandler{}
	srv = httptest.NewServer(srvHndl)
	apiXMLURI = srv.URL
	obj = New(&Configuration{Username: username, Password: password}).
		From(fromname).
		Testing(true).
		Extended(true)
	st, err = obj.Message(&Message{ID: messageID, Body: body}, toNumber1)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if len(st.State) == 0 || st.State[0].ErrorString != shaSumm {
		t.Fatalf("error send single message")
	}
}

func TestMessageToAny(t *testing.T) {
	const (
		username  = `2qUTSpRYgZHcPYxm2zhq@Qx645rznnuZYWKr7HrbR.xz`
		password  = `bA7RufvpnEeKuF8NqGBn`
		fromname  = `JQNeFT6bbmedtUARTqFe`
		messageID = 12345677654321
		groupID   = 987654
		body      = `gMK2aH5ZRT93RBmxC2wC`
		toNumber1 = `11234567890`
		toNumber2 = `21234567890`
		shaSumm   = `a6b0662e560464f86f254d370b34916dbc6087f7270b725788d1d05176c2af2c12c76edc92fa10f0a84aa424016a1c1d1d5cbac2494b744f7eb1af4545824180`
	)
	var err error
	var srvHndl *testSingleSrvHandler
	var srv *httptest.Server
	var obj Interface
	var st *Status

	srvHndl = &testSingleSrvHandler{}
	srv = httptest.NewServer(srvHndl)
	apiXMLURI = srv.URL
	obj = New(&Configuration{Username: username, Password: password}).
		From(fromname).
		Testing(true).
		Extended(true)
	st, err = obj.MessageToAny(&Message{ID: messageID, Body: body}, groupID, toNumber1, toNumber2)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if len(st.State) == 0 || st.State[0].ErrorString != shaSumm {
		t.Fatalf("error send single message")
	}
}

func TestMessageAt(t *testing.T) {
	const (
		username  = `2qUTSpRYgZHcPYxm2zhq@Qx645rznnuZYWKr7HrbR.xz`
		password  = `bA7RufvpnEeKuF8NqGBn`
		fromname  = `JQNeFT6bbmedtUARTqFe`
		messageID = 12345677654321
		body      = `gMK2aH5ZRT93RBmxC2wC`
		toNumber1 = `11234567890`
		shaSumm   = `14a41b3c78c039d17123d124b520f554214fbdae37b68a94041f4b9907338d2a3617358202fd5ff62d4bd237350d45b11101066e0113c8a06e6a28ea35385ad5`
	)
	var err error
	var srvHndl *testSingleSrvHandler
	var srv *httptest.Server
	var obj Interface
	var st *Status
	var startAt time.Time

	srvHndl = &testSingleSrvHandler{}
	srv = httptest.NewServer(srvHndl)
	apiXMLURI = srv.URL
	obj = New(&Configuration{Username: username, Password: password}).
		From(fromname).
		Testing(true).
		Extended(true)
	startAt = time.Date(2019, 02, 01, 0, 1, 2, 0, time.UTC)
	st, err = obj.MessageAt(&Message{ID: messageID, Body: body}, startAt, toNumber1)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if len(st.State) == 0 || st.State[0].ErrorString != shaSumm {
		t.Fatalf(st.State[0].ErrorString)
		t.Fatalf("error send single message")
	}
}

func TestMessageToAnyAt(t *testing.T) {
	const (
		username  = `2qUTSpRYgZHcPYxm2zhq@Qx645rznnuZYWKr7HrbR.xz`
		password  = `bA7RufvpnEeKuF8NqGBn`
		fromname  = `JQNeFT6bbmedtUARTqFe`
		messageID = 12345677654321
		groupID   = 987654
		body      = `gMK2aH5ZRT93RBmxC2wC`
		toNumber1 = `11234567890`
		toNumber2 = `21234567890`
		shaSumm   = `fededd19aed5b09784afccf36b56d885d11a274948e7450bdc57f23a40fb635cd1a04b482f14f300afd85e54305a76eee2090150562ee8f717d9d9a760fb324e`
	)
	var err error
	var srvHndl *testSingleSrvHandler
	var srv *httptest.Server
	var obj Interface
	var st *Status
	var startAt time.Time

	srvHndl = &testSingleSrvHandler{}
	srv = httptest.NewServer(srvHndl)
	apiXMLURI = srv.URL
	obj = New(&Configuration{Username: username, Password: password}).
		From(fromname).
		Testing(true).
		Extended(true)
	startAt = time.Date(2019, 02, 01, 0, 1, 2, 0, time.UTC)
	st, err = obj.MessageToAnyAt(&Message{ID: messageID, Body: body}, groupID, startAt, toNumber1, toNumber2)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if len(st.State) == 0 || st.State[0].ErrorString != shaSumm {
		t.Fatalf(st.State[0].ErrorString)
		t.Fatalf("error send single message")
	}
}

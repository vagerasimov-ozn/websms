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

	"gopkg.in/webnice/web.v1/status"
)

type testSrvHandler struct {
	http.Handler
}

func (tsh *testSrvHandler) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	var err error
	var buf []byte
	var hsh hash.Hash
	var sum string

	if buf, err = ioutil.ReadAll(rq.Body); err != nil {
		wr.WriteHeader(status.InternalServerError)
		return
	}
	defer func() { _ = rq.Body.Close() }()
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

func TestRequest(t *testing.T) {
	const (
		username = `2qUTSpRYgZHcPYxm2zhq@Qx645rznnuZYWKr7HrbR.xz`
		password = `bA7RufvpnEeKuF8NqGBn`
		fromname = `JQNeFT6bbmedtUARTqFe`
	)
	var err error
	var srvHndl *testSrvHandler
	var srv *httptest.Server
	var obj Interface
	var st *Status

	srvHndl = &testSrvHandler{}
	srv = httptest.NewServer(srvHndl)
	apiXMLURI = srv.URL

	DefaultUsername(username)
	DefaultPassword(password)
	DefaultFrom(fromname)
	obj = New().Testing(true).Extended(true)
	st, err = obj.Message(&Message{ID: 12345677654321, Body: "gMK2aH5ZRT93RBmxC2wC"}, "11234567890")
	if err != nil {
		log.Fatalf("error: %s", err)
		return
	}
	if len(st.State) == 0 ||
		st.State[0].ErrorString != `07c0dfe8bc7a37dcec95982874f8908533633e048dff7ea75fd7aabca5b6c1151dffe11033cc30fe1580134f05a9802707cedc451caf676f3fe345441541e0b3` {
		t.Fatalf("error send single message")
	}
}

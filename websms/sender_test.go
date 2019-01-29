package websms

//import "gopkg.in/webnice/debug.v1"
import "gopkg.in/webnice/log.v2"
import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"gopkg.in/webnice/web.v1/status"
)

type (
	badSrvHandler struct{ http.Handler }
)

func (tsh *badSrvHandler) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	wr.WriteHeader(status.ImATeapot)
	if _, err := wr.Write([]byte(status.Text(status.ImATeapot))); err != nil {
		log.Fatalf("response error: %s", err)
	}
}

func TestSendRequestIncorrectInput(t *testing.T) {
	obj := New().(*impl)
	r, err := obj.SendRequest("test error")
	if err == nil {
		t.Fatalf("SendRequest error. incorrect input parameter")
	}
	if r != nil {
		t.Fatalf("SendRequest error. incorrect result")
	}
}

func TestSendRequestBadServer(t *testing.T) {
	obj := New()
	apiXMLURI = "http://localhost:65536/bad/port"
	_, err := obj.Message(&Message{}, "")
	if err == nil || !strings.Contains(err.Error(), "invalid port") {
		t.Fatalf("SendRequest error")
	}
}

func TestSendRequestServerResponse(t *testing.T) {
	var err error
	var badServer *badSrvHandler
	var srv *httptest.Server
	var obj Interface

	srv = httptest.NewServer(badServer)
	apiXMLURI = srv.URL
	obj = New()
	if _, err = obj.Message(&Message{}, ""); err == nil {
		t.Fatalf("SendRequest error")
	}
	if err == nil || !strings.Contains(err.Error(), strconv.Itoa(status.ImATeapot)) {
		t.Log(err.Error())
		t.Log(status.Text(status.ImATeapot))
		t.Fatalf("SendRequest error")
	}
}

func TestStatusByGroupID(t *testing.T) {
	if _, err := New().StatusByGroupID(0); err == nil {
		t.Fatalf(`not implemented tests, contribute please`)
	}
}

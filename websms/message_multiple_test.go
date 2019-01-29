package websms

//import "gopkg.in/webnice/debug.v1"
import "gopkg.in/webnice/log.v2"
import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gopkg.in/webnice/web.v1/status"
)

type (
	testMultipleSrvHandler struct{ http.Handler }
)

func (tsh *testMultipleSrvHandler) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
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
	if _, err = fmt.Fprintf(wr, `<?xml version='1.0' encoding='UTF-8'?>
<status groupid='9876543210' date='01.01.2019 10:11:12'>
  <uniq_key>987987001</uniq_key>
  <id>936412900</id>
  <state error='%s' errcode='0'>Accepted</state>
  <MESSAGE_PARTS>1</MESSAGE_PARTS>
  <MESSAGE_COST>2,15</MESSAGE_COST>
  <BALANCE>9999999001,75</BALANCE>
  <uniq_key>987987002</uniq_key>
  <id>936412900</id>
  <state error='%s' errcode='0'>Accepted</state>
  <MESSAGE_PARTS>1</MESSAGE_PARTS>
  <MESSAGE_COST>2,45</MESSAGE_COST>
  <BALANCE>9999998998,3</BALANCE>
</status>`, sum, sum); err != nil {
		log.Fatalf("response error: %s", err)
	}
}

func TestMessages(t *testing.T) {
	const (
		username   = `sH5tdhTVrbn9pnuKhrYy@EK9DpRn34UFt8KDrs23F.xz`
		password   = `amC3Nm672vG9JqFqPtxA`
		fromname   = `2pwa26BWDZcQN2bPWyZT`
		messageID1 = 987987001
		messageID2 = 987987002
		packetID   = 9876543210
		body1      = `Sdt6GypMN97W7dKmGV2c`
		body2      = `n3RcJd3khD4zFFsPDsCd`
		toNumber1  = `98765432101`
		toNumber2  = `98765432102`
		shaSumm    = `e0526b65cde5a17807e3ed8d05b6997757906cdc00745e1588d9320231e0871bea8052b37cb95ffd3e0141959f595f4b78a573a70f79cdc106bb9ba7771f4bff`
	)
	var err error
	var srvHndl *testMultipleSrvHandler
	var srv *httptest.Server
	var obj Interface
	var msgs []*MultiMessage
	var st *Status

	srvHndl = &testMultipleSrvHandler{}
	srv = httptest.NewServer(srvHndl)
	apiXMLURI = srv.URL
	obj = New(&Configuration{Username: username, Password: password}).
		From(fromname).
		Testing(true).
		Extended(true)
	msgs = []*MultiMessage{
		&MultiMessage{ID: messageID1, To: toNumber1, Body: body1},
		&MultiMessage{ID: messageID2, To: toNumber2, Body: body2},
	}
	st, err = obj.Messages(msgs, packetID)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if len(st.State) == 0 || st.State[0].ErrorString != shaSumm {
		t.Log(st.State[0].ErrorString)
		t.Fatalf("error send multiple messages")
	}
}

func TestMessagesAt(t *testing.T) {
	const (
		username   = `sH5tdhTVrbn9pnuKhrYy@EK9DpRn34UFt8KDrs23F.xz`
		password   = `amC3Nm672vG9JqFqPtxA`
		fromname   = `2pwa26BWDZcQN2bPWyZT`
		messageID1 = 987987001
		messageID2 = 987987002
		packetID   = 9876543210
		body1      = `Sdt6GypMN97W7dKmGV2c`
		body2      = `n3RcJd3khD4zFFsPDsCd`
		toNumber1  = `98765432101`
		toNumber2  = `98765432102`
		shaSumm    = `7212a3858c92a6b76b9f03a7933da7b88e417cb9c5649492ca41b45a9bdaca0d8a817353c9a94e39c227ec25b8addce2c97516bfd780ea2b73ebae9e71ed82d6`
	)
	var err error
	var srvHndl *testMultipleSrvHandler
	var srv *httptest.Server
	var obj Interface
	var msgs []*MultiMessage
	var st *Status
	var startAt time.Time

	srvHndl = &testMultipleSrvHandler{}
	srv = httptest.NewServer(srvHndl)
	apiXMLURI = srv.URL
	obj = New(&Configuration{Username: username, Password: password}).
		From(fromname).
		Testing(true).
		Extended(true)
	msgs = []*MultiMessage{
		&MultiMessage{ID: messageID1, To: toNumber1, Body: body1},
		&MultiMessage{ID: messageID2, To: toNumber2, Body: body2},
	}
	startAt = time.Date(2019, 02, 01, 0, 1, 2, 0, time.UTC)
	st, err = obj.MessagesAt(msgs, packetID, startAt)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if len(st.State) == 0 || st.State[0].ErrorString != shaSumm {
		t.Log(st.State[0].ErrorString)
		t.Fatalf("error send multiple messages")
	}
}

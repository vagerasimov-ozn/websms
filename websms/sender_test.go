package websms

//import "gopkg.in/webnice/debug.v1"
import "gopkg.in/webnice/log.v2"
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"gopkg.in/webnice/web.v1/status"
)

type uniSrvHandler struct {
	http.Handler
	Status int
	Body   *bytes.Buffer
}

func (ush *uniSrvHandler) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	wr.WriteHeader(ush.Status)
	if _, err := wr.Write(ush.Body.Bytes()); err != nil {
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
	_, err = obj.Balance()
	if err == nil || !strings.Contains(err.Error(), "invalid port") {
		t.Fatalf("SendRequest error")
	}
	_, err = obj.StatusByMessageID(0)
	if err == nil || !strings.Contains(err.Error(), "invalid port") {
		t.Fatalf("SendRequest error")
	}
}

func TestSendRequestServerResponse(t *testing.T) {
	var err error
	var badServer *uniSrvHandler
	var srv *httptest.Server
	var obj Interface

	badServer = &uniSrvHandler{
		Status: status.ImATeapot,
		Body:   bytes.NewBufferString(status.Text(status.ImATeapot)),
	}
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

func TestStatusConvertor(t *testing.T) {
	var err error
	var badServer *uniSrvHandler
	var srv *httptest.Server
	var obj Interface
	var st *Status

	badServer = &uniSrvHandler{
		Status: status.Ok,
		Body: bytes.NewBufferString(`<?xml version='1.0' encoding='UTF-8'?>
<status groupid='9876543210' date='01.01.2019 10:11:12'>
  <uniq_key>987987001</uniq_key>
  <id>936412900</id>
  <state error='OK' errcode='0'>Accepted</state>
  <MESSAGE_PARTS>1</MESSAGE_PARTS>
  <MESSAGE_COST>2-15</MESSAGE_COST>
  <BALANCE>9999999001_75</BALANCE>
  <uniq_key>987987002</uniq_key>
  <id>936412901</id>
  <state error='OK' errcode='0'>Accepted</state>
  <id>-936412902</id>
  <uniq_key>-987987003</uniq_key>
  <id>-</id>
  <MESSAGE_PARTS>1</MESSAGE_PARTS>
  <MESSAGE_COST>2:45</MESSAGE_COST>
  <BALANCE>9999998998%3</BALANCE>
</status>`),
	}
	srv = httptest.NewServer(badServer)
	apiXMLURI = srv.URL
	obj = New()
	if st, err = obj.Message(&Message{}, ""); err != nil {
		t.Fatalf("SendRequest error: %s", err)
	}
	if st.GroupID != 9876543210 {
		t.Fatalf("statusConvertor error")
	}
	if st.StateAt.Truncate(time.Second).UTC().String() != "2019-01-01 07:11:12 +0000 UTC" {
		t.Log(st.StateAt.Truncate(time.Second).UTC().String())
		t.Fatalf("statusConvertor error")
	}
	if len(st.State) != 4 {
		t.Fatalf("statusConvertor error")
	}
	if st.State[0].ID != 936412900 || st.State[0].UniqKey != 987987001 || st.State[0].ErrorString != "OK" || st.State[0].Value != StatusAccepted {
		t.Fatalf("statusConvertor error")
	}
	if st.State[1].ID != 936412901 || st.State[1].UniqKey != 987987002 || st.State[1].ErrorString != "OK" || st.State[1].Value != StatusAccepted {
		t.Fatalf("statusConvertor error")
	}
	if st.State[2].ID != 0 || st.State[2].UniqKey != 0 {
		t.Fatalf("statusConvertor error")
	}
}

func TestBalance(t *testing.T) {
	var err error
	var badServer *uniSrvHandler
	var srv *httptest.Server
	var balance float64

	badServer = &uniSrvHandler{
		Status: status.Ok,
		Body:   bytes.NewBufferString(`<BALANCE>9999998998,3</BALANCE>`),
	}
	srv = httptest.NewServer(badServer)
	apiXMLURI = srv.URL
	balance, err = New().Balance()
	if err != nil {
		t.Fatalf("Request balance error: %s", err)
	}
	if balance != 9999998998.3 {
		t.Fatalf("Balance value error")
	}
}

func TestBalanceIncorrect(t *testing.T) {
	var err error
	var badServer *uniSrvHandler
	var srv *httptest.Server
	var balance float64

	badServer = &uniSrvHandler{
		Status: status.Ok,
		Body:   bytes.NewBufferString(`<BALANCE>Some random text of a crazy API request architecture.</BALANCE>`),
	}
	srv = httptest.NewServer(badServer)
	apiXMLURI = srv.URL
	balance, err = New().Balance()
	if err == nil {
		t.Fatalf("Request balance error: %s", err)
	}
	if balance != 0 {
		t.Fatalf("Balance value error")
	}
}

func TestRequestStatusByMessageID(t *testing.T) {
	const (
		username  = `sH5tdhTVrbn9pnuKhrYy@EK9DpRn34UFt8KDrs23F.xz`
		password  = `amC3Nm672vG9JqFqPtxA`
		messageID = 987987001
		shaSumm   = `0cf02e0ef2b06ab55e4eddc00eee217deb82bc3ae0a6d204e3318754d5b2da3e107b6e56d1e8029eaaa9279af59d0a02ef342cc904e80217a532c781ba450773`
	)
	var err error
	var srvHndl *testMultipleSrvHandler
	var srv *httptest.Server
	var obj Interface
	var st *Status

	srvHndl = &testMultipleSrvHandler{}
	srv = httptest.NewServer(srvHndl)
	apiXMLURI = srv.URL
	obj = New(&Configuration{Username: username, Password: password}).Extended(true)
	st, err = obj.StatusByMessageID(messageID)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	if len(st.State) == 0 || st.State[0].ErrorString != shaSumm {
		t.Fatalf("Request status of message by ID error")
	}
}

func TestStatusByMessageID(t *testing.T) {
	var err error
	var badServer *uniSrvHandler
	var srv *httptest.Server
	var st *Status

	badServer = &uniSrvHandler{
		Status: status.Ok,
		Body: bytes.NewBufferString(`<?xml version='1.0' encoding='UTF-8'?>
<status id='536346304' date='29.01.2019 5:31:08'>
  <state>Delivered</state>
  <MESSAGE_ID>536346304</MESSAGE_ID>
  <REG_DATE>28.01.2019 10:56:49</REG_DATE>
  <SEND_ON>28.01.2019 10:56:00</SEND_ON>
  <DELIVERED_DATE>28.01.2019 10:57:32</DELIVERED_DATE>
  <MESSAGE_PARTS>18</MESSAGE_PARTS>
  <MESSAGE_COST>2,1531</MESSAGE_COST>
</status>`),
	}
	srv = httptest.NewServer(badServer)
	apiXMLURI = srv.URL
	st, err = New().Extended(true).StatusByMessageID(536346304)
	if err != nil {
		t.Fatalf("Request status of message by ID error: %s", err)
	}
	if st.ID != 536346304 {
		t.Fatalf("Request status of message by ID error")
	}
	if st.StateAt.Truncate(time.Second).UTC().String() != "2019-01-29 02:31:08 +0000 UTC" {
		t.Log(st.StateAt.Truncate(time.Second).UTC().String())
		t.Fatalf("Request status of message by ID error")
	}
	if st.RegistrationAt.Truncate(time.Second).Unix() != 1548662209 {
		t.Fatalf("Request status of message by ID error")
	}
	if st.SendAt.Truncate(time.Second).Unix() != 1548662160 {
		t.Fatalf("Request status of message by ID error")
	}
	if st.DeliveredAt.Truncate(time.Second).Unix() != 1548662252 {
		t.Fatalf("Request status of message by ID error")
	}
	if st.MessageParts != 18 || st.MessageCost != 2.1531 {
		t.Fatalf("Request status of message by ID error")
	}
	if len(st.State) == 0 {
		t.Fatalf("Request status of message by ID error")
	}
	if st.State[0].Value != StatusDelivered {
		t.Fatalf("Request status of message by ID error")
	}
}

func TestStatusByGroupID(t *testing.T) {
	if _, err := New().StatusByGroupID(0); err == nil {
		t.Fatalf(`not implemented tests, contribute please`)
	}
}

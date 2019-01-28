package websms

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

func TestNew(t *testing.T) {
	const (
		username  = `2qUTSpRYgZHcPYxm2zhq@Qx645rznnuZYWKr7HrbR.xz`
		password  = `bA7RufvpnEeKuF8NqGBn`
		fromname  = `JQNeFT6bbmedtUARTqFe`
		username2 = `A4xTD8nK7nn66WcSWAz7@VDf8ZdrG6ZtwQmEkEpzj.xz`
		password2 = `9jeNRPZ4QDDvH7xHn5e8`
		fromname2 = `A2Cftwc6YEYEMBKmdyRE`
	)
	var wsm Interface

	DefaultUsername(username)
	DefaultPassword(password)
	DefaultFrom(fromname)
	if defaultConfiguration.Username != username {
		t.Errorf("Error DefaultUsername")
	}
	if defaultConfiguration.Password != password {
		t.Errorf("Error DefaultPassword")
	}
	if defaultConfiguration.From != fromname {
		t.Errorf("Error DefaultFrom")
	}
	wsm = New()
	if wsm == nil {
		t.Errorf("Error New(), return nil")
	}
	if wsm.(*impl).extd {
		t.Errorf("Error New(), extended is true")
	}
	wsm = New().Extended(true)
	if !wsm.(*impl).extd {
		t.Errorf("Error New(), extended is false")
	}
	if wsm.(*impl).test {
		t.Errorf("Error New(), testing is true")
	}
	wsm = New().Testing(true)
	if !wsm.(*impl).test {
		t.Errorf("Error New(), testing is false")
	}
	if wsm.(*impl).from != fromname {
		t.Errorf("Error New(), fromname not set")
	}
	if wsm.(*impl).cfg.Username != username {
		t.Errorf("Error New(), username not set")
	}
	if wsm.(*impl).cfg.Password != password {
		t.Errorf("Error New(), password not set")
	}
	wsm = New(&Configuration{Username: username2, Password: password2}).From(fromname2)
	if wsm.(*impl).cfg.Username != username2 {
		t.Errorf("Error New(), username not set")
	}
	if wsm.(*impl).cfg.Password != password2 {
		t.Errorf("Error New(), password not set")
	}
	if wsm.(*impl).from != fromname2 {
		t.Errorf("Error New(), fromname not set")
	}
}

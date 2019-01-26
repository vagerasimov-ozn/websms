package websms // import "gopkg.in/webnice/websms.v1/websms"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"encoding/xml"
	"time"

	websmsTypes "gopkg.in/webnice/websms.v1/websms/types"
)

// Send a single message as a one packet, at the specified time, to multiple address
func (wsm *impl) messageToAnyAt(msg *Message, pkgID uint64, start time.Time, to ...string) (status *Status, err error) {
	var tnw websmsTypes.TimeRfc1123z
	var obj *websmsTypes.SingleRequest
	var toj *websmsTypes.To
	var j int

	tnw = websmsTypes.TimeRfc1123z(time.Now())
	obj = &websmsTypes.SingleRequest{
		Service: &websmsTypes.Service{
			ID:      websmsTypes.ServiceSingle.String(),
			UniqKey: msg.ID,
		},
		Body: msg.Body,
	}
	// Флаг тестирования
	if wsm.test {
		obj.Service.Testing = 1
	}
	// Авторизация
	obj.Service.Username, obj.Service.Password = wsm.cfg.Username, wsm.cfg.Password
	// Имя отправителя
	obj.Service.Source = wsm.from
	// Если указана дата и время, устанавливаем её
	if !start.IsZero() {
		obj.Service.StartAt = &tnw
	}
	// Если получателей множество
	if len(to) > 1 {
		obj.Service.ID = websmsTypes.ServiceBulk.String()
		obj.Service.UniqKey = pkgID
	}
	obj.To = make([]*websmsTypes.To, 0, len(to))
	for j = range to {
		toj = &websmsTypes.To{Value: to[j]}
		obj.To = append(obj.To, toj)
	}
	// Выполнение запроса отправки и получение результата запроса
	status = new(Status)
	*status, err = wsm.SendRequest(obj)

	return
}

// Send a multiple messages as a one packet, at the specified time, to multiple address
func (wsm *impl) messagesToAnyAt(msgs []*MultiMessage, pkgID uint64, start time.Time) (status *Status, err error) {
	const (
		tagTo   = `to`
		tagBody = `body`
	)
	var tnw websmsTypes.TimeRfc1123z
	var obj *websmsTypes.MultipleRequest
	var tbi *websmsTypes.ToBody
	var i int

	tnw = websmsTypes.TimeRfc1123z(time.Now())
	obj = &websmsTypes.MultipleRequest{
		Service: &websmsTypes.Service{
			ID:      websmsTypes.ServiceIndividual.String(),
			UniqKey: pkgID,
		},
	}
	// Флаг тестирования
	if wsm.test {
		obj.Service.Testing = 1
	}
	// Авторизация
	obj.Service.Username, obj.Service.Password = wsm.cfg.Username, wsm.cfg.Password
	// Имя отправителя
	obj.Service.Source = wsm.from
	// Если указана дата и время, устанавливаем её
	if !start.IsZero() {
		obj.Service.StartAt = &tnw
	}
	// Чередование to и body
	obj.ToBody = make([]*websmsTypes.ToBody, 0, len(msgs)*2)
	for i = range msgs {
		tbi = &websmsTypes.ToBody{
			XMLName: xml.Name{Local: tagTo},
			UniqKey: msgs[i].ID,
			Value:   msgs[i].To,
		}
		obj.ToBody = append(obj.ToBody, tbi)
		tbi = &websmsTypes.ToBody{
			XMLName: xml.Name{Local: tagBody},
			Value:   msgs[i].Body,
		}
		obj.ToBody = append(obj.ToBody, tbi)
	}
	// Выполнение запроса отправки и получение результата запроса
	status = new(Status)
	*status, err = wsm.SendRequest(obj)

	return
}

package websms // import "gopkg.in/webnice/websms.v1/websms"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	websmsTypes "gopkg.in/webnice/websms.v1/websms/types"

	"gopkg.in/webnice/transport.v2"
	"gopkg.in/webnice/transport.v2/request"
	"gopkg.in/webnice/web.v1/mime"
)

func init() {
	singletonTransport = transport.New().
		MaximumIdleConnections(defaultMaximumIdleConnections).               // Максимальное общее число бездействующих keepalive соединений
		MaximumIdleConnectionsPerHost(defaultMaximumIdleConnectionsPerHost). // Максимальное число бездействующих keepalive соединений для каждого хоста
		DialContextTimeout(defaultDialContextTimeout).                       // Таймаут установки соединения с хостом
		IdleConnectionTimeout(defaultIdleConnectionTimeout).                 // Таймаут keepalive соединения до обрыва связи
		TotalTimeout(defaultTotalTimeout).                                   // Общий таймаут на весь процесс связи, включает соединение, отправку данных, получение ответа
		RequestPoolSize(defaultRequestPoolSize)                              // Размер пула воркеров готовых для выполнения запросов к хостам
}

// Конвертирование статуса из больного во вменяемый формат
func (wsm *impl) statusConvertor(s *websmsTypes.Status) (status *Status, err error) {
	var i, maxCount int

	if maxCount < len(s.StateUniqKey) {
		maxCount = len(s.StateUniqKey)
	}
	if maxCount < len(s.StateID) {
		maxCount = len(s.StateID)
	}
	if maxCount < len(s.State) {
		maxCount = len(s.State)
	}
	status = &Status{
		ID:           s.ID,
		GroupID:      s.GroupID,
		StateAt:      s.StateAt.Time(),
		MessageParts: s.MessageParts,
		State:        make([]*State, maxCount),
	}
	if s.RegistrationAt != nil {
		status.RegistrationAt = s.RegistrationAt.Time()
	}
	if s.SendAt != nil {
		status.SendAt = s.SendAt.Time()
	}
	if s.DeliveredAt != nil {
		status.DeliveredAt = s.DeliveredAt.Time()
	}
	if status.MessageCost, err = strconv.ParseFloat(strings.Replace(s.MessageCost, ",", ".", -1), 64); err != nil {
		status.MessageCost, err = 0.0, nil
	}
	if status.Balance, err = strconv.ParseFloat(strings.Replace(s.Balance, ",", ".", -1), 64); err != nil {
		status.Balance, err = 0.0, nil
	}
	for i = 0; i < maxCount; i++ {
		status.State[i] = new(State)
	}
	for i = range s.StateUniqKey {
		if status.State[i].UniqKey, err = strconv.ParseUint(s.StateUniqKey[i].Value, 10, 64); err != nil {
			status.State[i].UniqKey, err = 0, nil
		}
	}
	for i = range s.StateID {
		if status.State[i].ID, err = strconv.ParseUint(s.StateID[i].Value, 10, 64); err != nil {
			status.State[i].ID, err = 0, nil
		}
	}
	for i = range s.State {
		status.State[i].ErrorCode = s.State[i].ErrorCode
		status.State[i].ErrorString = s.State[i].Error
		status.State[i].Error = NewErrCode(s.State[i].ErrorCode)
		status.State[i].Value = NewDispatchStatus(s.State[i].Value)
	}

	return
}

// SendRequest Выполнение запроса отправки и получение результата запроса
func (wsm *impl) SendRequest(obj interface{}) (status *Status, err error) {
	var req request.Interface
	var s *websmsTypes.Status

	switch obj.(type) {
	case *websmsTypes.SingleRequest, *websmsTypes.MultipleRequest:
	default:
		err = errors.New("incorrect call SendRequest(), wrong parameter type")
		return
	}
	// Запрос
	req = singletonTransport.RequestGet().
		UserAgent(userAgent).
		Method(singletonTransport.Method().Post()).
		ContentType(mime.ApplicationXMLCharsetUTF8).
		DataXML(obj).
		URL(apiXMLURI)
	defer singletonTransport.RequestPut(req)
	singletonTransport.Do(req)
	// Ожидание ответа
	if err = req.Done().Error(); err != nil {
		err = fmt.Errorf("Execute request error: %s", err)
		return
	}
	// Анализ результата
	if req.Response().StatusCode() < 200 || req.Response().StatusCode() > 299 {
		err = fmt.Errorf("Request %s %q error, HTTP code %d (%s)", req.Response().Response().Request.Method, req.Response().Response().Request.URL.String(), req.Response().StatusCode(), req.Response().Status())
		return
	}
	s = new(websmsTypes.Status)
	err = req.Response().Content().UnmarshalXML(s)
	// Конвертирование статуса во вменяемый формат
	status, err = wsm.statusConvertor(s)
	// DEBUG
	//req.Response().Content().BackToBegin()
	//log.Debug(req.Response().Content().String())
	// DEBUG

	return
}

// Balance Request account balance
func (wsm *impl) Balance() (ret float64, err error) {
	var obj *websmsTypes.BalanceRequest
	var rsp *websmsTypes.BalanceResponse
	var req request.Interface

	// Авторизация
	obj = &websmsTypes.BalanceRequest{
		Auth: websmsTypes.Auth{
			Username: wsm.cfg.Username,
			Password: wsm.cfg.Password,
		},
	}
	// Запрос
	req = singletonTransport.RequestGet().
		UserAgent(userAgent).
		Method(singletonTransport.Method().Post()).
		ContentType(mime.ApplicationXMLCharsetUTF8).
		DataXML(obj).
		URL(apiXMLURI)
	defer singletonTransport.RequestPut(req)
	singletonTransport.Do(req)
	// Ожидание ответа
	if err = req.Done().Error(); err != nil {
		err = fmt.Errorf("Execute request error: %s", err)
		return
	}
	// Анализ результата
	if req.Response().StatusCode() < 200 || req.Response().StatusCode() > 299 {
		err = fmt.Errorf("Request %s %q error, HTTP code %d (%s)", req.Response().Response().Request.Method, req.Response().Response().Request.URL.String(), req.Response().StatusCode(), req.Response().Status())
		return
	}
	rsp = new(websmsTypes.BalanceResponse)
	err = req.Response().
		Content().
		UnmarshalXML(rsp)
	if err != nil {
		return
	}
	ret, err = strconv.ParseFloat(strings.Replace(rsp.Value, ",", ".", -1), 64)
	if err != nil {
		err = fmt.Errorf(rsp.Value)
	}

	return
}

// StatusByMessageID Request status of dispatch by message ID
func (wsm *impl) StatusByMessageID(id uint64) (status *Status, err error) {
	const keyStatus = `status`
	var obj *websmsTypes.RequestStatus
	var req request.Interface
	var s *websmsTypes.Status

	// Авторизация
	obj = &websmsTypes.RequestStatus{
		Auth: websmsTypes.Auth{
			Username: wsm.cfg.Username,
			Password: wsm.cfg.Password,
		},
		ID:    id,
		Value: keyStatus,
	}
	// Флаги
	if wsm.extd {
		obj.Extended = 1
	}
	// Запрос
	req = singletonTransport.RequestGet().
		UserAgent(userAgent).
		Method(singletonTransport.Method().Post()).
		ContentType(mime.ApplicationXMLCharsetUTF8).
		DataXML(obj).
		URL(apiXMLURI)
	defer singletonTransport.RequestPut(req)
	singletonTransport.Do(req)
	// Ожидание ответа
	if err = req.Done().Error(); err != nil {
		err = fmt.Errorf("Execute request error: %s", err)
		return
	}
	// Анализ результата
	if req.Response().StatusCode() < 200 || req.Response().StatusCode() > 299 {
		err = fmt.Errorf("Request %s %q error, HTTP code %d (%s)", req.Response().Response().Request.Method, req.Response().Response().Request.URL.String(), req.Response().StatusCode(), req.Response().Status())
		return
	}
	s = new(websmsTypes.Status)
	err = req.Response().Content().UnmarshalXML(s)
	// Конвертирование статуса во вменяемый формат
	status, err = wsm.statusConvertor(s)
	// DEBUG
	//req.Response().Content().BackToBegin()
	//log.Debug(req.Response().Content().String())
	// DEBUG

	return
}

// StatusByGroupID Request status of dispatch by group ID
func (wsm *impl) StatusByGroupID(id uint64) (status *Status, err error) {
	return nil, errors.New(`not implemented, sorry, contribute please`)
}

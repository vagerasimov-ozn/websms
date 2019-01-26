package websms // import "gopkg.in/webnice/websms.v1/websms"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"errors"
	//"bytes"
	//"encoding/xml"

	websmsTypes "gopkg.in/webnice/websms.v1/websms/types"

	"gopkg.in/webnice/transport.v2"
)

func init() {
	singletonTransport = transport.New().
		MaximumIdleConnections(defaultMaximumIdleConnections).               // Максимальное общее число бездействующих keepalive соединений
		MaximumIdleConnectionsPerHost(defaultMaximumIdleConnectionsPerHost). // Максимальное число бездействующих keepalive соединений для каждого хоста
		DialContextTimeout(defaultDialContextTimeout).                       // Таймаут установки соединения с хостом
		IdleConnectionTimeout(defaultIdleConnectionTimeout).                 // Таймаут keepalive соединения до обрыва связи
		TotalTimeout(defaultTotalTimeout).                                   // Общий таймаут на весь процесс связи, включает соединение, отправку данных, получение ответа
		RequestPoolSize(defaultRequestPoolSize)                              // Размер пула воркеров готовых для выполнения запросов к хостам

	//ErrorFunc(singleton.ErrorStream).                                    // Функция получения ошибок работы транспорта
	//DialTLS(singleton.customDialTLSFunc)                                 // Кастомная функция установки соединения с TLS хостами для обслуживания самоподписанных сертификатов
}

// SendRequest Выполнение запроса отправки и получение результата запроса
func (wsm *impl) SendRequest(obj interface{}) (status *websmsTypes.Status, err error) {
	switch obj.(type) {
	case *websmsTypes.SingleRequest, *websmsTypes.MultipleRequest:
	default:
		err = errors.New("incorrect call SendRequest(), wrong parameter type")
		return
	}

	//var buf = &bytes.Buffer{}
	//enc := xml.NewEncoder(buf)
	////enc.Indent("", "  ")
	//if err = enc.Encode(obj); err != nil {
	//	return
	//}
	//log.Debug("msg: " + buf.String() + "\n")

	return
}

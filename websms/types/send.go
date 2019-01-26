package types // import "gopkg.in/webnice/websms.v1/websms/types"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"encoding/xml"
)

// SingleRequest Описание запроса отправки одиночного сообщения
type SingleRequest struct {
	XMLName xml.Name `xml:"message"`        // Наименование секции XML
	Service *Service `xml:"service"`        // Описание запроса и доступа к сервису
	To      []*To    `xml:"to"`             // Получатели сообщения
	Body    string   `xml:"body,omitempty"` // Тело сообщения
}

// MultipleRequest Описание запроса отправки множественного сообщения
type MultipleRequest struct {
	XMLName xml.Name  `xml:"message"` // Наименование секции XML
	Service *Service  `xml:"service"` // Описание запроса и доступа к сервису
	ToBody  []*ToBody `xml:""`        // Чередование to и body
}

// Service Описание запроса и доступа к сервису
type Service struct {
	Auth
	ID      string        `xml:"id,attr"`                 // Тип запроса к сервису
	Source  string        `xml:"source,attr"`             // Имя отправителя сообщения
	Testing uint8         `xml:"test,attr,omitempty"`     // Флаг тестирования
	StartAt *TimeRfc1123z `xml:"start,attr,omitempty"`    // Дата и время отправки сообщения сервисом
	UniqKey uint64        `xml:"uniq_key,attr,omitempty"` // Уникальный идентификатор сообщения или пользовательский номер пакета
}

// To Описание получателя
type To struct {
	UniqKey uint64 `xml:"uniq_key,attr,omitempty"` // Уникальный идентификатор сообщения
	Value   string `xml:",chardata"`               // Номер получателя
}

// ToBody Тэг с динамическим именем
type ToBody struct {
	XMLName xml.Name `xml:""`
	UniqKey uint64   `xml:"uniq_key,attr,omitempty"` // Уникальный идентификатор сообщения
	Value   string   `xml:",chardata"`               // Значение элемента
}

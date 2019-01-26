package types // import "gopkg.in/webnice/websms.v1/websms/types"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"encoding/xml"
)

// DefaultConfiguration Структура конфигурации по умолчанию
type DefaultConfiguration struct {
	From     string // Имя отправителя по умолчанию
	Username string // Имя пользователя по умолчанию
	Password string // Пароль пользователя по умолчанию
}

// ServiceType Тип запроса к сервису
type ServiceType string

// Auth Структура описания доступа к сервису
type Auth struct {
	Username string `xml:"login,attr"`    // Имя пользователя
	Password string `xml:"password,attr"` // пароль пользователя
}

// String Return type as string
func (st ServiceType) String() string { return string(st) }

// Status Статус запроса
type Status struct {
	XMLName xml.Name            `xml:"status"`                 // Наименование секции XML
	ID      int64               `xml:"id,attr,omitempty"`      // Уникальный идентификатор сообщения
	GroupID int64               `xml:"groupid,attr,omitempty"` // Номер пакета группы сообщений
	StateAt *TimeWebsmsShortOne `xml:"date,attr,omitempty"`    // Дата и время для которого был возвращён статус
	State   []*State            `xml:""`                       // Статус
}

// State Состояние статуса запроса
type State struct {
	XMLName   xml.Name `xml:""`             // Наименование тега XML
	Error     string   `xml:"error,attr"`   // Сообщение об ошибки
	ErrorCode int64    `xml:"errcode,attr"` // Код ошибки
	Value     string   `xml:",chardata"`    // Значение статуса или идентификатор сообщения
}

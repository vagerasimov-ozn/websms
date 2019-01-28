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
	XMLName        xml.Name            `xml:"status"`                   // Наименование секции XML
	ID             uint64              `xml:"id,attr,omitempty"`        // Уникальный идентификатор сообщения
	GroupID        uint64              `xml:"groupid,attr,omitempty"`   // Номер пакета группы сообщений
	StateAt        *TimeWebsmsShortOne `xml:"date,attr,omitempty"`      // Дата и время для которого был возвращён статус
	StateUniqKey   []*State            `xml:"uniq_key"`                 // Статус uniq_key
	StateID        []*State            `xml:"id"`                       // Статус id
	State          []*State            `xml:"state"`                    // Статус state
	RegistrationAt *TimeWebsmsShortOne `xml:"REG_DATE,omitempty"`       // Дата и время получения СМС-сообщения сервисом
	SendAt         *TimeWebsmsShortOne `xml:"SEND_ON,omitempty"`        // Дата и время отправки СМС-сообщения сервисом
	DeliveredAt    *TimeWebsmsShortOne `xml:"DELIVERED_DATE,omitempty"` // Дата и время доставки СМС-сообщения
	MessageParts   uint8               `xml:"MESSAGE_PARTS,omitempty"`  // Количество частей сообщения
	MessageCost    string              `xml:"MESSAGE_COST,omitempty"`   // Стоимость сообщения в книвом формате
	Balance        string              `xml:"BALANCE,omitempty"`        // Баланс аккаунта в кривом формате
}

// State Состояние статуса, секция state
type State struct {
	XMLName   xml.Name `xml:""`                       // Наименование тега XML
	Error     string   `xml:"error,attr,omitempty"`   // Сообщение об ошибки
	ErrorCode uint16   `xml:"errcode,attr,omitempty"` // Код ошибки
	Value     string   `xml:",chardata"`              // Значение статуса или идентификатор сообщения
}

// BalanceRequest Запрос баланса аккаунта
type BalanceRequest struct {
	Auth
	XMLName xml.Name `xml:"balance"`   // Наименование секции XML
	Value   string   `xml:",chardata"` // Значение
}

// BalanceResponse Ответ на запрос баланса
type BalanceResponse struct {
	XMLName xml.Name `xml:"BALANCE"`   // Наименование секции XML
	Value   string   `xml:",chardata"` // Значение
}

// RequestStatus Запрос статуса отправленного сообщения
type RequestStatus struct {
	XMLName  xml.Name `xml:"request"`                 // Наименование секции XML
	ID       uint64   `xml:"id,attr,omitempty"`       // Уникальный идентификатор сообщения
	Extended uint8    `xml:"extended,attr,omitempty"` // Флаг расширенного статуса
	Value    string   `xml:",chardata"`               // Значение запроса
	Auth
}

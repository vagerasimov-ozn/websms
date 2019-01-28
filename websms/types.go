package websms // import "gopkg.in/webnice/websms.v1/websms"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"time"

	websmsTypes "gopkg.in/webnice/websms.v1/websms/types"

	"gopkg.in/webnice/transport.v2"
)

const (
	defaultMaximumIdleConnections        = uint(1000)             // Максимальное общее число бездействующих keepalive соединений
	defaultMaximumIdleConnectionsPerHost = uint(150)              // Максимальное число бездействующих keepalive соединений для каждого хоста
	defaultDialContextTimeout            = time.Millisecond * 500 // Таймаут установки соединения с хостом
	defaultIdleConnectionTimeout         = time.Minute            // Таймаут keepalive соединения до обрыва связи
	defaultTotalTimeout                  = time.Second * 3        // Общий таймаут на весь процесс связи, включает соединение, отправку данных, получение ответа
	defaultRequestPoolSize               = uint16(64)             // Размер пула воркеров готовых для выполнения запросов к хостам
)

const (
	apiXmlUri = `http://cab.websms.ru/xml_in5.asp` // API URI
)

var defaultConfiguration *websmsTypes.DefaultConfiguration
var singletonTransport transport.Interface

// Interface is an interface of object
type Interface interface {
	// From Name of sender, if not specified selects from service profile settings
	From(from string) Interface

	// Testing Set testing flag
	Testing(t bool) Interface

	// Extended Выполнение запроса отправки с получением расширенной информации статуса
	Extended(t bool) Interface

	// Balance Request account balance
	Balance() (ret float64, err error)

	// Message Send a single message at now to one address
	Message(msg *Message, to string) (status *Status, err error)

	// MessageToAny Send a single message at now to multiple address
	MessageToAny(msg *Message, pkgID uint64, to ...string) (status *Status, err error)

	// MessageAt Send a single message at the specified time to one address
	MessageAt(msg *Message, start time.Time, to string) (status *Status, err error)

	// MessageToAnyAt Send a single message at the specified time to multiple address
	MessageToAnyAt(msg *Message, pkgID uint64, start time.Time, to ...string) (status *Status, err error)

	// Messages Send a multiple message as a one packet at now to one address
	Messages(msgs []*MultiMessage, pkgID uint64) (status *Status, err error)

	// MessagesAt Send a multiple message as a one packet at the specified time to one address
	MessagesAt(msgs []*MultiMessage, pkgID uint64, start time.Time) (status *Status, err error)

	// StatusByMessageID Request status of dispatch by message ID
	StatusByMessageID(id uint64) (status *Status, err error)

	// StatusByGroupID Request status of dispatch by group ID
	StatusByGroupID(id uint64) (status *Status, err error)
}

// is an implementation
type impl struct {
	cfg  *Configuration // Конфигурация
	test bool           // =true - выполнение запроса к сервису с флагом test
	extd bool           // =true - выполнение запроса отправки и проверки статуса с получением расширенной информации статуса
	from string         // Имя отправителя, если не указано, то выбирается в личном кабинете сервиса
}

// Configuration websms.ru service configuration structure
type Configuration struct {
	Username string // User name
	Password string // Password for access to API
}

// Message is an structure of single message
type Message struct {
	ID   uint64 // Unique ID of message
	Body string // Body of message
}

// MultiMessage is an structure of multiple messages
type MultiMessage struct {
	ID   uint64 // Unique ID of message
	To   string // Destination address
	Body string // Body of message
}

// Status Статус запроса
type Status struct {
	ID             uint64    // Уникальный идентификатор сообщения
	GroupID        uint64    // Номер пакета группы сообщений переданный при отправке сообщения
	StateAt        time.Time // Дата и время для которого был возвращён статус
	State          []*State  // Состояние
	RegistrationAt time.Time // +extend Дата и время получения СМС-сообщения сервисом
	SendAt         time.Time // +extend Дата и время отправки СМС-сообщения сервисом
	DeliveredAt    time.Time // +extend Дата и время доставки СМС-сообщения
	MessageParts   uint8     // +extend Количество частей сообщения
	MessageCost    float64   // +extend Стоимость сообщения в книвом формате
	Balance        float64   // +extend Баланс аккаунта в кривом формате
}

// State Состояние статуса запроса
type State struct {
	ID          uint64         // Уникальный идентификатор сообщения в системе websms.ru
	UniqKey     uint64         // Уникальный идентификатор сообщения переданный при отправке сообщения
	ErrorString string         // Сообщение об ошибки
	ErrorCode   uint16         // Код ошибки
	Error       ErrCode        // Ошибка в строгом формате
	Value       DispatchStatus // Статус сообщения
}

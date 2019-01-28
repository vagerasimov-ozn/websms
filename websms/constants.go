package websms // import "gopkg.in/webnice/websms.v1/websms"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
//import ()

const userAgent = `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:64.0) Gecko/20100101 Firefox/64.0`

const (
	// StatusAccepted Принято к отправке
	StatusAccepted = DispatchStatus(`accepted`)

	// StatusUnknown Неизвестный результат или не известный статус
	StatusUnknown = DispatchStatus(`unknown`)

	// StatusReady Готово к отправке
	StatusReady = DispatchStatus(`ready`)

	// StatusSpooled Поставлено в очередь
	StatusSpooled = DispatchStatus(`spooled`)

	// StatusSent Передано оператору
	StatusSent = DispatchStatus(`sent`)

	// StatusRejected Отказано оператором
	StatusRejected = DispatchStatus(`rejected`)

	// StatusDelivered Доставлено
	StatusDelivered = DispatchStatus(`delivered`)

	// StatusUndeliverable Не доставлено
	StatusUndeliverable = DispatchStatus(`undeliverable`)

	// StatusEnroute В процессе доставки
	StatusEnroute = DispatchStatus(`enroute`)

	// StatusDeleted Удалено пользователем
	StatusDeleted = DispatchStatus(`deleted`)

	// StatusExpired Время доставки сообщения истекло
	StatusExpired = DispatchStatus(`expired`)

	// StatusDelayed Отложено системой
	StatusDelayed = DispatchStatus(`delayed`)

	// StatusRestricted Закрытое направление
	StatusRestricted = DispatchStatus(`restricted`)

	// StatusUnroutable Неизвестное направление
	StatusUnroutable = DispatchStatus(`unroutable`)

	// StatusMessageIDNotFound Сообщение с указанным ID отсутствует
	StatusMessageIDNotFound = DispatchStatus(`not found`)
)

const (
	cErrCode00 uint16 = iota
	cErrCode01
	cErrCode02
	cErrCode03
	cErrCode04
	cErrCode05
	cErrCode06
	cErrCode07
	cErrCode08
	cErrCode09
	cErrCode10
	cErrCode11
	cErrCode12
	cErrCode13
	cErrCode14
	cErrCode15
	cErrCode16
	cErrCode17
	cErrCode18
	cErrCode19
	cErrCode20
	cErrCode21
	cErrCode22
	cErrCode23
)
const cErrCodeXX = uint16(65535)

var _ = cErrCode18 // Не используемые коды
var _ = cErrCode19 // Не используемые коды

var errCodeMap = map[uint16]ErrCode{
	cErrCodeXX: {cErrCodeXX, `Unknown error code`, `Не известный код ошибки`},
	cErrCode00: {cErrCode00, `Ok`, `Данные приняты системой`},
	cErrCode01: {cErrCode01, `Error login or password`, `Неверный логин или пароль`},
	cErrCode02: {cErrCode02, `Blocked user`, `Доступ заблокирован`},
	cErrCode03: {cErrCode03, `Insufficient funds`, `На счёте недостаточно средств`},
	cErrCode04: {cErrCode04, `Blocked IP`, `IP адрес заблокирован`},
	cErrCode05: {cErrCode05, `HTTP not enabled`, `Персональные настройки запрещают отправку по HTTP`},
	cErrCode06: {cErrCode06, `This server IP not enabled`, `IP-адрес не указан в персональных настройках`},
	cErrCode07: {cErrCode07, `E-mail sending not enabled`, `Персональные настройки запрещают отправку по SMTP`},
	cErrCode08: {cErrCode08, `This e-mail not enabled`, `Персональные настройки запрещают отправку по SMTP`},
	cErrCode09: {cErrCode09, `Blocked moderator ID`, `Доступ модератору закрыт`},
	cErrCode10: {cErrCode10, `Error manual phone list`, `Недопустимые символы в адресатах phone_list`},
	cErrCode11: {cErrCode11, `Empty message text`, `Не задан текст сообщения message`},
	cErrCode12: {cErrCode12, `Empty phone list`, `Не заданы адресаты phone_list`},
	cErrCode13: {cErrCode13, `Stop service`, `Сервис временно недоступен`},
	cErrCode14: {cErrCode14, `Error format date`, `Неверный формат даты send_on`},
	cErrCode15: {cErrCode15, `Double sent from web interface`, `Повторная отправка допускается через 10 секунд`},
	cErrCode16: {cErrCode16, `Error dealer off`, `Сервисы недоступны`},
	cErrCode17: {cErrCode17, `Error multi access`, `Процедура отправки занята`},
	cErrCode20: {cErrCode20, `Incorrect group`, `Неверный формат параметра group`},
	cErrCode21: {cErrCode21, `Empty password`, `Не указан пароль http_password`},
	cErrCode22: {cErrCode22, `Empty login`, `Не указан логин http_username`},
	cErrCode23: {cErrCode23, `Invalid from phone`, `Недозволительное имя отправителя`},
}

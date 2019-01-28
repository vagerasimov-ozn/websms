package types // import "gopkg.in/webnice/websms.v1/websms/types"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
//import ()

const (
	// ServiceSingle Отправка одиночного SMS сообщения
	ServiceSingle = ServiceType(`single`)

	// ServiceBulk Отправка одного текста на несколько номеров
	ServiceBulk = ServiceType(`bulk`)

	// ServiceIndividual Отправка индивидуального текста на несколько номеров
	ServiceIndividual = ServiceType(`individual`)
)

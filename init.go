package api

var EsamUrl, ServicesUrl string

func UrlInit(esamUrl, servicesUrl string) {
	EsamUrl, ServicesUrl = esamUrl, servicesUrl
}

package models

type MVC struct {
	Models, Views, Controllers, Config, Router string
}

type ENV struct {
	GOPATH , PATH string
}
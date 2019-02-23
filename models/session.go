package models

import "Nutrition/app"

type Session struct {
	Id     string
	UserId string
}

func (session *Session) Exists() bool {
	return session.Id != app.DefaultUUID
}

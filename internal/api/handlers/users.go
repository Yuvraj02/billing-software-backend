package handlers

import "restapi/billing-backend/internal/models"

var (
	usersMap = make(map[int]models.User)
	nextUser       = 0
)

func init() {

	usersMap[0] = models.User{Id: 0, Name: "Yuvraj", Email: "test@gmail.com", Password: "pass1"}
	nextUser++

}
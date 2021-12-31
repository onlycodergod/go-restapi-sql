package auth

import (
	"rest/api/database"
	"rest/api/models"
	"rest/api/security"
	"rest/api/utils/channels"

	"gorm.io/gorm"
)

func SignIn(login, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}

		err = db.Debug().Model(models.User{}).Where("login = ?", login).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return CreateToken(user.ID)
	}
	return "", err
}
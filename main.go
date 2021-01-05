package main

import (
	"example.com/service"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	service.DrawLayout()

	service.SetLogForUser("Успешно запустили прогрумму")

	// Сигнал по нажатию на кнопку
	service.MainInterface.Auth.Connect("clicked", func() {
		auth()
	})
	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	gtk.Main()
}

// Авторизация
func auth() {
	loginText, err1 := service.MainInterface.Login.GetText()
	passwordText, err2 := service.MainInterface.Password.GetText()
	if err1 == nil && err2 == nil{
		if !service.Login(loginText, passwordText) {
			service.SetLogForUser("Ошибка при авторизации в Instagram. Проверьте логин и пароль!")
		} else {
			service.SetLogForUser("Успешно авторизовались в Instagram!")
		}
	}
}
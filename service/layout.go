package service

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

var b, _ = gtk.BuilderNew()
var interfaceLogObject, _ = b.GetObject("interfaceLog")
var interfaceLog, _ = interfaceLogObject.(*gtk.TextView)

func DrawLayout() {
	// Инициализируем GTK.
	gtk.Init(nil)

	// Загружаем в билдер окно из файла Glade
	err := b.AddFromFile("./layout/layout.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("window_main")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Кнопка авторизации
	obj, _ = b.GetObject("auth")
	auth_button := obj.(*gtk.Button)
	// Сигнал по нажатию на кнопку
	auth_button.Connect("clicked", func() {
		auth(auth_button)
	})

	interfaceLogObject, _ := b.GetObject("interfaceLog")
	interfaceLog, _ := interfaceLogObject.(*gtk.TextView)

	// Отображаем все виджеты в окне
	win.ShowAll()

	buffer, err := interfaceLog.GetBuffer()
	buffer.SetText("Успешно запустили прогрумму")

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	gtk.Main()
}

func auth(button *gtk.Button) {
	// Поле логин
	login_text, err1 := getValue("login")
	// Поле паспорт
	password_text, err2 := getValue("password")
	if err1 == nil && err2 == nil{
		if !Login(login_text, password_text) {
			button.SetLabel("FAIL!")
		} else {
			button.SetLabel("OK!")
		}
	}
}

func getValue(fieldName string) (string, error) {
	obj, _ := b.GetObject(fieldName)

	return obj.(*gtk.Entry).GetText()
}

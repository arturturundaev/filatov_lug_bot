package service

import (
	"encoding/json"
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"reflect"
)

// Главный интерфейс
type MainInterfaceStruct struct {
	Login 		 *gtk.Entry
	Password 	 *gtk.Entry
	Auth 		 *gtk.Button
	AddNewUrl 	 *gtk.Button
	NameAddUrl 	 *gtk.Entry
	Start 		 *gtk.Button
	UrlList 	 *gtk.TextView
	InterfaceLog *gtk.TextView
}

var MainInterface MainInterfaceStruct


func DrawLayout() {
	// Инициализируем GTK.
	gtk.Init(nil)
	builder, _ := gtk.BuilderNew()

	// Загружаем в билдер окно из файла Glade
	err := builder.AddFromFile("./layout/layout.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := builder.GetObject("window_main")
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

	// Отображаем все виджеты в окне
	win.ShowAll()

	initMainInterface(builder)
}

func initMainInterface(builder *gtk.Builder) {
	obj, err := builder.GetObject("Login")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "Login"))
	}
	MainInterface.Login = obj.(*gtk.Entry)

	obj, err = builder.GetObject("Password")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "Password"))
	}
	MainInterface.Password = obj.(*gtk.Entry)

	obj, err = builder.GetObject("Auth")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "Auth"))
	}
	MainInterface.Auth = obj.(*gtk.Button)

	obj, err = builder.GetObject("AddNewUrl")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "AddNewUrl"))
	}
	MainInterface.AddNewUrl = obj.(*gtk.Button)

	obj, err = builder.GetObject("NameAddUrl")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "NameAddUrl"))
	}
	MainInterface.NameAddUrl = obj.(*gtk.Entry)

	obj, err = builder.GetObject("Start")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "Start"))
	}
	MainInterface.Start = obj.(*gtk.Button)

	obj, err = builder.GetObject("UrlList")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "UrlList"))
	}
	MainInterface.UrlList = obj.(*gtk.TextView)

	obj, err = builder.GetObject("InterfaceLog")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "InterfaceLog"))
	}
	MainInterface.InterfaceLog = obj.(*gtk.TextView)
	
	obj, err = builder.GetObject("Auth")
	if err != nil {
		SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", "Auth"))
	}
	MainInterface.Auth = obj.(*gtk.Button)



}
func initMainInterface2(builder *gtk.Builder) {
	list := make([]glib.IObject,  reflect.ValueOf(&MainInterface).Elem().NumField())

	e := reflect.ValueOf(&MainInterface).Elem()
	for i := 0; i < e.NumField(); i++ {
		propertyName := e.Type().Field(i).Name
		propertyType := e.Type().Field(i).Type.String()
		value, err := builder.GetObject(propertyName)

		if err != nil {
			SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", propertyName))
		}

		switch propertyType {
			case "*gtk.Entry":
				list[i] = value.(*gtk.Entry)
				break

			case "*gtk.Button":
				list[i] = value.(*gtk.Button)
				break

			case "*gtk.TextView":
				list[i] = value.(*gtk.TextView)
				break

			default:
				SetLogForUser(fmt.Sprintf("Ошибка при инициализации свойства %v\n", propertyName))
		}
	}
	data, _ := json.Marshal(list)
	fmt.Println(data)
}

func SetLogForUser(text string) {
	SetText(MainInterface.InterfaceLog, text)
}

func SetText(property *gtk.TextView, text string) {
	buffer, err := property.GetBuffer()

	oldText, err := GetText(property)

	if err != nil {
		oldText = ""
	}

	buffer.SetText(oldText + "\n" +  text + "\n")
}

func GetText(property *gtk.TextView) (string, error)  {
	buffer, err := property.GetBuffer()

	if err != nil {
		fmt.Print("Ошибка при получении текущего значения")
	}
	start, end := buffer.GetBounds()

	return buffer.GetText(start, end, true)
}


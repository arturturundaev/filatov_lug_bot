package main

import (
	"example.com/service"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"strings"
	goinsta "github.com/ahmdrz/goinsta/v2"
)

func main() {
	service.DrawLayout()

	service.SetLogForUser("Успешно запустили прогрумму")
	service.InitDataFromCache()
	service.SetText(service.MainInterface.UrlList, service.Cache.UrlList)
	// Сигнал по нажатию на кнопку
	service.MainInterface.Start.Connect("clicked", func() {
		start()
	})
	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	gtk.Main()
}

// Авторизация
func checkLoginAndPassword() bool {
	loginText, err1 := service.MainInterface.Login.GetText()
	passwordText, err2 := service.MainInterface.Password.GetText()
	if err1 == nil && err2 == nil{
		if service.Login(loginText, passwordText) {

			service.SetLogForUser("Успешно авторизовались в Instagram!")

			return true
		}
	}

	service.SetLogForUser("Ошибка при авторизации в Instagram. Проверьте логин и пароль!")

	return false
}

// Срабатывает после нажатия кнопки Старт
func start() bool {
	if checkLoginAndPassword() {
		likeByPost("sdf")

		listOfLinks, _ := service.GetText(service.MainInterface.UrlList)

		arrayOfLinks := strings.Split(listOfLinks, "\n")
		if len(arrayOfLinks) == 0 {
			service.SetLogForUser("Список ссылок пуст")

			return false
		}

		for _, link := range arrayOfLinks {
			if len(link) == 0 {
				continue
			}

			// Если это тег
			if string(link[0]) == "#" {
				likeByTag(link)
			} else {

			}


		}

		return true
	}

	return false
}

func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return ""
}

// Ставим лайки к постам по тегу
func likeByTag(tag string) {
	tag = trimFirstRune(tag)
	if len(tag) != 0 {
		feedTag, err := service.Instabot.Insta.Feed.Tags(tag)
		if err != nil {
			service.SetLogForUser(fmt.Sprintf("Ошибка при обработке тега %v", tag))
		}
		for _, item := range feedTag.RankedItems {
			err = item.Like()
			if err != nil {
				log.Printf("error on liking item %s, %v", item.ID, err)
			} else {
				log.Printf("item %s liked", item.ID)
			}
		}
	}
}

// Ставим лайки к посту и всем комментариям, в которым сожержатся указанные слова/фразы
func likeByPost(linkToPost string)  {
	mediaId, err := goinsta.MediaIDFromShortID("CJv5sY4rG9T")
	post, err := service.Instabot.Insta.GetMedia(mediaId)
	if err != nil {
		service.SetLogForUser(fmt.Sprintf("Не удалось найти пост по id %v", "CJx4ckfg9Ed"))
	}
	for _, item := range post.Items {
		item.Comments.Sync()
		item.Comments.Next()
		comments := item.Comments.Items
		for _, comment := range comments {
			comment.
		}
	}
	fmt.Print(post.Status)
}
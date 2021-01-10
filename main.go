package main

import (
	"encoding/json"
	"example.com/service"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
		//service.Instabot.Insta.GetToken()
		likeByPost("https://www.instagram.com/p/CJvEmvwnqAn&") // 2481222174512160807 2481222174512160807_25207636427  OK
		likeByPost("https://www.instagram.com/p/CJrX5AvHIyU") // 2480181092667919508  2480181092667919508_25207636427 FAIL
		likeByPost("https://www.instagram.com/p/CJoocmgnZhW") // OK
		likeByPost("https://www.instagram.com/p/CJk7swAi1SE") // FAIL

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
func likeByPost(linkToPost string) {
	service.GetFullMediaIdByShortId(linkToPost)
	post, err := service.Instabot.Insta.GetMedia(linkToPost)
	if err != nil {
		service.SetLogForUser(fmt.Sprintf("Не удалось найти пост по id %v", "CJx4ckfg9Ed"))

		return
	}

	for _, item := range post.Items {
		if !item.HasLiked {
			err = item.Like()
			if err != nil {
				log.Printf("error on liking item %s, %v", item.ID, err)
			} else {
				log.Printf("Post %s liked", item.ID)
			}
		}
		item.Comments.Sync()
		item.Comments.Next()
		comments := item.Comments.Items
		NextComment:
		for _, comment := range comments {
			if comment.HasLikedComment {
				continue NextComment
			}
			for _, whiteWord := range service.GetWhiteList() {
				if strings.Contains(strings.ToLower(comment.Text), strings.ToLower(whiteWord)) {
					comment.Like()
					log.Printf("Comment '%v' liked", comment.Text)
					continue NextComment
				}
			}
		}
	}
	fmt.Print(post.Status)

	return
}



type jsonResponseStruct struct {
	media_id string
}

func getMediaId(mainUrl string) jsonResponseStruct {
	links := []string{"https://api.instagram.com/oembed/?url=", mainUrl}
	url := strings.Join(links,"")

	resp, err := http.Get(url)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	spaceClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}


	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonResponse := jsonResponseStruct{}
	jsonErr := json.Unmarshal(body, &jsonResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return jsonResponse

}
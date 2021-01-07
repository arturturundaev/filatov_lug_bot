package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type cacheStruct struct {
	UrlList string `json:"UrlList"`
}

var Cache cacheStruct

var path = "./.cache"

// Функция обновляет данные в кеше
func SaveDataToCache() {
	bytes, err := json.Marshal(Cache)
	if err != nil {
		SetLogForUser("Ошибка при попытке обернуть данные в кеш")
	}

	err = ioutil.WriteFile(path, bytes, 0644)

	if err != nil {
		SetLogForUser("Ошибка при записе данных в файл")
	}
}

func InitDataFromCache()  {
	f, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		SetLogForUser("Ошибка при попытке создания файла с кешем")
	}
	if err != nil {
		SetLogForUser("Ошибка при попытке открытия файла с кешем")
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)

	if err != nil {
		SetLogForUser("Ошибка при попытке чтении файла с кешем")
	}

	if len(bytes) != 0 {
		config := cacheStruct{}
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			SetLogForUser("Ошибка при распарсивании данных из файла с кешем")
		}

		Cache.UrlList = config.UrlList
	}
}

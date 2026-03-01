package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sviatilnik/go-racer/internal/config"
	"github.com/sviatilnik/go-racer/internal/race"
	"github.com/sviatilnik/go-racer/internal/utils"
)

func main() {
	if !utils.CheckConnection() {
		fmt.Println("❌ Нет подключения к Интернету для начала гонки")
		return
	}

	urls := getURLs()

	if len(urls) == 0 {
		fmt.Println("❌ Некорректные адреса для начала гонки")
		return
	}

	conf, err := getConfig()
	if err != nil {
		fmt.Println("❌ Ошибка получения конфигурации " + err.Error())
		return
	}

	fmt.Printf("🏁 Стартуем гонку для %d URL с таймаутом %fs \n", len(urls), conf.GetTimeout().Seconds())

	ctx, cancel := context.WithTimeout(context.Background(), conf.GetTimeout())
	defer cancel()

	race.NewRace(ctx, urls).Start()
}

func getURLs() []string {
	if len(os.Args) < 2 {
		return nil
	}

	urls := make([]string, 0)
	for i := 1; i < len(os.Args); i++ {
		url := os.Args[i]
		if utils.IsURLValid(url) {
			urls = append(urls, url)
		}
	}

	return urls
}

func getConfig() (config.Config, error) {
	conf := config.NewFlagConfig()
	if err := conf.Init(); err != nil {
		return nil, err
	}

	return conf, nil
}

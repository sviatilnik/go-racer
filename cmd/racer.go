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

	conf, urls, err := initConfigAndURLs()
	if err != nil {
		fmt.Println("❌ Ошибка получения конфигурации: " + err.Error())
		return
	}

	if len(urls) == 0 {
		fmt.Println("❌ Некорректные адреса для начала гонки")
		return
	}

	fmt.Printf("🏁 Стартуем гонку для %d URL с таймаутом %fs\n", len(urls), conf.GetTimeout().Seconds())

	ctx, cancel := context.WithTimeout(context.Background(), conf.GetTimeout())
	defer cancel()

	race.NewRace(ctx, urls).Start()
}

func initConfigAndURLs() (config.Config, []string, error) {
	conf := config.NewFlagConfig()
	remaining, err := conf.Init(os.Args[1:])
	if err != nil {
		return nil, nil, err
	}

	urls := make([]string, 0, len(remaining))
	for _, arg := range remaining {
		if utils.IsURLValid(arg) {
			urls = append(urls, arg)
		}
	}

	return conf, urls, nil
}

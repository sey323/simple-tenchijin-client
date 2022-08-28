package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Tenchijin Poc: ") //ログのフォーマットの設定
}

func main() {
	username := flag.String("username", "", "発行されたユーザ名.")
	password := flag.String("password", "", "発行されたapikey.")
	flag.Parse()

	log.Printf("username: %s, password: %s", *username, *password)

	// ログイン処理
	tenchijinClient := NewTenchijinClient(*username, *password, "")
	tenchijinClient.Login()

	log.Printf("token: %s", tenchijinClient.Token)

	// 降水量を取得
	var weather_result = tenchijinClient.WeatherRainFall(
		35.660557,
		139.749682,
		"2022-01-01",
		"2022-01-07",
		"daily",
		"Asia/Tokyo",
	)
	log.Printf("降水量: %s", weather_result)

	// 地表面温度を取得
	var lst_result = tenchijinClient.WeatherLst(
		35.660557,
		139.749682,
		"2022-01-01",
		"2022-01-07",
		"daily",
	)
	log.Printf("地表面温度: %s", lst_result)
}

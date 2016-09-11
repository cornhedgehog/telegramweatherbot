package main

////////////////////////сообщения, отправляемые в мессенджер

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

func GetStartMessage(cityCode string) string {
	return "Добрый день! \xf0\x9f\x8c\xbb " +
		"Вас приветствует бот \"Погода в Калининграде\". \n " +
		"Данные о прогнозе погоды предоставляются сервисом Meteoservice.ru " +
		"Ваши замечания и предложения можно присылать сюда: @cornhedgehog \n " +
		"Для получения первого прогноза нажмите /help"
}

func GetWeatherData(cityCode string) string {
	resp, err := http.Get("http://xml.meteoservice.ru/export/gismeteo/point/" + cityCode + ".xml")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	v := []Forecast{}
	t := Town{Fo: v, Index: "none"}
	r := Report{To: t, Type: "none"}
	m := MMWeather{Rep: r}

	err = xml.Unmarshal(body, &m)
	if err != nil {
		return "Возникли неполадки в работе. Приносим извинения за доставленные неудобства."
	}

	var wData string = "Прогноз на " + m.Rep.To.Fo[0].Day + " " + GetMonthG(m.Rep.To.Fo[0].Month) + " " + m.Rep.To.Fo[0].Year + ".\n"

	for i := 0; i < 4; i++ {
		wData += "В " + m.Rep.To.Fo[i].Hour + ":00 будет " + m.Rep.To.Fo[i].Te.Min + "℃, " + GetPresipitation(m.Rep.To.Fo[i].Phe.Cloudiness,
			m.Rep.To.Fo[i].Phe.Precipitation) + PrecipitationPossibility(m.Rep.To.Fo[i].Phe.Rpower, m.Rep.To.Fo[i].Phe.Spower) + ".\n\n"
	}

	wData += "Для получения следующего прогноза нажмите /help"
	return wData
}

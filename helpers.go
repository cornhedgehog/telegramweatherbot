package main

import "strconv"

//типы осадков, получаемые из xml
type WeatherValue struct {
	Name  string
	Value string
}

//cloudiness -	облачность по градациям: 0 - ясно, 1- малооблачно, 2 - облачно, 3 - пасмурно
//precipitation -	тип осадков: 4 - дождь, 5 - ливень, 6,7 – снег, 8 - гроза, 9 - нет данных, 10 - без осадков
var WeatherValues = []WeatherValue{WeatherValue{"ясно", "0"}, WeatherValue{"малооблачно", "1"}, 
WeatherValue{"облачно", "2"}, WeatherValue{"пасмурно", "3"}, WeatherValue{"дождь", "4"}, WeatherValue{"ливень", "5"}, 
WeatherValue{"небольшой снег", "6"}, WeatherValue{"сильный снег", "7"}, WeatherValue{"гроза", "8"}, 
WeatherValue{"нет данных", "9"}, WeatherValue{"без осадков", "10"}}

func GetPresipitation(cloudiness string, presipitation string) string {
	cloudiness_int, _ := strconv.Atoi(cloudiness)
	presipitation_int, _ := strconv.Atoi(presipitation)
	wData := WeatherValues[cloudiness_int].Name + ", " + WeatherValues[presipitation_int].Name
	return wData
}

//месяцы в родительном падеже
var months = []string{"0", "января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "декабря"}

func GetMonthG(hour string) string {
	hour_int, _ := strconv.Atoi(hour)
	return months[hour_int]
}

//если есть вероятность дождя или грозы, сказать об этом
func PrecipitationPossibility(rpower string, spower string) string {
	reply := ""
	if rpower == "1" {
		reply += "возможен дождь, "
	}
	if spower == "1" {
		reply += "возможна гроза, "
	}
	return reply

}

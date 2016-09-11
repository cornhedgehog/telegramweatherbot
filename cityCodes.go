package main

import "errors"

type City struct {
	Name  string
	Value string
}

var Cities = []City{City{"Калининград", "105"}}

func GetCityByCode(cityCode string) string {
	for _, c := range Cities {
		if c.Value == cityCode {
			return c.Name
		}
	}
	return "0"
}

func GetCodeByCity(msg string) (string, error) {
	for _, c := range Cities {
		if c.Name == msg {
			return c.Value, nil
		}
	}
	return "0", errors.New("Нет такого города!")
}

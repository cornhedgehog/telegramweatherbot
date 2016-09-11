package main

import "encoding/xml"

type Phenomena struct {
	XMLName       xml.Name `xml:"PHENOMENA"`
	Cloudiness    string   `xml:"cloudiness,attr"`
	Precipitation string   `xml:"precipitation,attr"`
	Rpower        string   `xml:"rpower,attr"`
	Spower        string   `xml:"spower,attr"`
}

type Pressure struct {
	XMLName xml.Name `xml:"PRESSURE"`
	Max     string   `xml:"max,attr"`
	Min     string   `xml:"min,attr"`
}

type Temperature struct {
	XMLName xml.Name `xml:"TEMPERATURE"`
	Max     string   `xml:"max,attr"`
	Min     string   `xml:"min,attr"`
}

type Wind struct {
	XMLName   xml.Name `xml:"WIND"`
	MinWind   string   `xml:"min,attr"`
	MaxWind   string   `xml:"max,attr"`
	Direction string   `xml:"direction,attr"`
}

type Relwet struct {
	XMLName xml.Name `xml:"RELWET"`
	Max     string   `xml:"max,attr"`
	Min     string   `xml:"min,attr"`
}

type Heat struct {
	XMLName xml.Name `xml:"HEAT"`
	Min     string   `xml:"min,attr"`
	Max     string   `xml:"max,attr"`
}

type Forecast struct {
	XMLName xml.Name    `xml:"FORECAST"`
	Phe     Phenomena   `xml:"PHENOMENA"`
	Pre     Pressure    `xml:"PRESSURE"`
	Te      Temperature `xml:"TEMPERATURE"`
	Wi      Wind        `xml:"WIND"`
	Re      Relwet      `xml:"RELWET"`
	He      Heat        `xml:"HEAT"`

	Day     string `xml:"day,attr"`
	Month   string `xml:"month,attr"`
	Year    string `xml:"year,attr"`
	Hour    string `xml:"hour,attr"`
	Tod     string `xml:"tod,attr"`
	Predict string `xml:"predict,attr"`
	Weekday string `xml:"weekday,attr"`
}

type Town struct {
	XMLName xml.Name   `xml:"TOWN"`
	Fo      []Forecast `xml:"FORECAST"`
	Index   string     `xml:"index,attr"`
	SName   string     `xml:"sname,attr"`
	Lat     string     `xml:"latitude,attr"`
	Lon     string     `xml:"longitude,attr"`
}

type Report struct {
	XMLName xml.Name `xml:"REPORT"`
	Type    string   `xml:"type,attr"`
	To      Town     `xml:"TOWN"`
}

type MMWeather struct {
	XMLName xml.Name `xml:"MMWEATHER"`
	Rep     Report   `xml:"REPORT"`
}

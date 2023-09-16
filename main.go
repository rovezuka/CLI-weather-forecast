package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

// Структура для анмаршалинга json-дананых о погоде
type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	q := "Saint-Peterburg" // Текущее местоположение

	if len(os.Args) > 1 {
		q = os.Args[1] // Если пользователь ввел другой город для прогноза погоды
	}

	// Запрос к API
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=cab6116eaa9342a5882145255231609&q=" + q + "&days=1&lang=ru&aqi=no&alerts=no")

	// Произошла ошибка при запросе
	if err != nil {
		panic(err)
	}

	// Позже закрыть запрос
	defer res.Body.Close()

	// Невереный код при запросе к API
	if res.StatusCode != 200 {
		panic("Weather API is not available")
	}

	body, err := io.ReadAll(res.Body)

	// Произошла ошибка при чтении
	if err != nil {
		panic(err)
	}

	// Переменная структуры
	var weather Weather

	// Анмаршалинг json-данных
	err = json.Unmarshal(body, &weather)

	// Проверяем на ошибку
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	// Выводим погоду на данный момент
	fmt.Printf("%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text)

	// Выводим прогноз погоды в будущем
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		// Если текущее время позднее
		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf("%s - %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text)

		// Если вероятность дождя меньше 40% просто выводим, иначе красим в красный цвет
		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}

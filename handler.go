package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	owm "github.com/briandowns/openweathermap"
)

func handleResult(w http.ResponseWriter, r *http.Request) {
	citys := []string{
		r.FormValue("city0"), r.FormValue("city1"), r.FormValue("city2"),
		r.FormValue("city3"), r.FormValue("city4"),
	}

	key := os.Getenv("OWM_API_KEY")
	if key == "" {
		log.Println("key not valid: ", key)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
		return
	}

	wm, err := owm.NewCurrent("C", "EN", key)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
		return
	}

	var (
		ws []CityWeather  // collect all in this
		wg sync.WaitGroup // wait till all are done
		mu sync.Mutex     // no racing
	)

	for _, city := range citys {
		wg.Add(1)

		go func(c string) {
			defer wg.Done()

			err = wm.CurrentByName(c)
			if err != nil {
				log.Println("could not find weather data for: " + c)
				return
			}

			mu.Lock()
			ws = append(ws, CityWeather{
				City:     wm.Name,
				TempMaxC: wm.Main.TempMax,
				TempMinC: wm.Main.TempMin,
				Temp:     wm.Main.Temp,
				Humidity: wm.Main.Humidity,
				Pressure: wm.Main.Pressure,
				Desc:     wm.Weather[0].Main,
			})
			mu.Unlock()
		}(city)
	}

	wg.Wait()

	tmpl, err := template.New("").ParseFiles("tmpl/layout.tmpl", "tmpl/result.tmpl")
	if err != nil {
		log.Println("template not found: ", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", ws)
	if err != nil {
		log.Println("template not valid: ", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
		return
	}

	return
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("GM_API_KEY")
	if key == "" {
		log.Println("key not valid: ", key)
	}

	tmpl, err := template.New("").ParseFiles("tmpl/layout.tmpl", "tmpl/index.tmpl")
	if err != nil {
		log.Println("template not found: ", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", key)
	if err != nil {
		log.Println("template not valid: ", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
		return
	}
}

func handleError(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("tmpl/layout.tmpl", "tmpl/error.tmpl")
	if err != nil {
		log.Println("template not found: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Println("template not valid: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

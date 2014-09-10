package goconf

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/urlfetch"
	"io/ioutil"
	"encoding/json"
	"log"
	"sync"
	"net/url"
)

func init() {
	http.HandleFunc("/", main)
	http.HandleFunc("/result", result)
}

type City_weather struct{
	City	string
	TempMaxC	string
	TempMinC	string
	WeatherDesc string
	WeatherIconUrl string
}

func main(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		return
	}
	return
}

func result(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	var wg sync.WaitGroup

	wg.Add(5)

	citys := []string{
		r.FormValue("city1"),
		r.FormValue("city2"),
		r.FormValue("city3"),
		r.FormValue("city4"),
		r.FormValue("city5"),
	}

	resc, errc := make(chan string), make(chan error)

	log.Println("%v", citys)
	Maxs := []City_weather{}

	for _, city := range citys {
		go func(city string) {
			log.Println("Starting client for %v", city)
			ur := "http://api.worldweatheronline.com/free/v1/weather.ashx?format=json&num_of_days=1&cc=no&key=" + key + "&q=" + url.QueryEscape(city)

			resp, err := client.Get(ur)
			if err != nil {
				panic(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				errc <- err
				return
			}

			var jsontype jsonobject
			json.Unmarshal(body, &jsontype)

			city1 := City_weather{
				jsontype.Data.Request[0].Query,
				jsontype.Data.Weather[0].TempMaxC,
				jsontype.Data.Weather[0].TempMinC,
				jsontype.Data.Weather[0].WeatherDesc[0].Value,
				jsontype.Data.Weather[0].WeatherIconUrl[0].Value,
			}

			Maxs = append(Maxs, city1)

			resc <- string(body)
			log.Println("%v", jsontype.Data.Weather[0].TempMaxC)
			wg.Done()
		}(city)
	}

	for i := 0; i < len(citys); i++ {
		select {
		case res := <-resc:
			log.Println(res)
		case err := <-errc:
			log.Println(err)
		}
	}

	wg.Wait()

	err1 := templates.ExecuteTemplate(w, "result.html", Maxs)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
	}
}

var (
	templates = template.Must(template.ParseFiles(
		"result.html",
		"index.html",
	))
	key = "18b005b8af54d992cbaf31f8eabf54baa7260170"
)

type jsonobject struct {
	Data ObjectType `json:"data"`
}

type ObjectType struct {
	Request []request `json:"request"`
	Weather []weather `json:"weather"`
}

type request struct {
	Query string `json:"query"`
	Type string `json:"type"`
}

type weather struct {
	Date   string `json:"date"`
	TempMaxC   string `json:"tempMaxC"`
	TempMaxF   string `json:"tempMaxF"`
	TempMinC   string `json:"tempMinC"`
	TempMinF   string `json:"tempMinF"`
	WeatherDesc []weatherDesc `json:"weatherDesc"`
	WeatherIconUrl []weatherIconUrl `json:"weatherIconUrl"`
}

type weatherDesc struct {
	Value	string `json:"value"`
}

type weatherIconUrl struct {
	Value	string `json:"value"`
}

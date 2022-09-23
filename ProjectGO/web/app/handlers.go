package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	models "github.com/ZakirAvrora/kolesa-upgrade-homework-3/internal"
)

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.InfoLog.Println(r.URL, r.Method, r.RemoteAddr)

	if r.URL.Path != "/" || r.Method != "GET" {
		app.InfoLog.Println("Not found", r.URL)
		app.Errors(w, http.StatusNotFound, nil)
		return
	}

	mapBreed, err := ExtractBreeds()
	if err != nil {
		app.Errors(w, http.StatusServiceUnavailable, err)
	}

	tmpl, err := template.ParseFiles("web/ui/templates/index.html")
	if err != nil {
		app.Errors(w, http.StatusInternalServerError, err)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "index.html", mapBreed); err != nil {
		app.Errors(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *App) catView(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limit := query.Get("limit")

	breed := query.Get("breed")

	mapBreed, err := ExtractBreeds()
	if err != nil {
		app.Errors(w, http.StatusServiceUnavailable, err)
	}

	var CatPictures []models.CatPicture

	if _, ok := mapBreed[breed]; !ok && breed != "" {
		app.Errors(w, http.StatusBadRequest, nil)
		return
	}

	data, status := GetJsonData(fmt.Sprintf("%s?limit=%s&breed_ids=%s", models.Api, limit, mapBreed[breed]))
	if status != http.StatusOK {
		app.Errors(w, status, err)
		return
	}

	if err := json.Unmarshal(data, &CatPictures); err != nil {
		app.Errors(w, http.StatusInternalServerError, err)
		return
	}

	tmpl, err := template.ParseFiles("web/ui/templates/search.html")
	if err != nil {
		app.Errors(w, http.StatusInternalServerError, err)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "search.html", &CatPictures); err != nil {
		app.Errors(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *App) catDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.Errors(w, http.StatusNotFound, nil)
		return
	}

	data, status := GetJsonData(models.DetailsApi + id)
	if status != http.StatusOK {
		app.Errors(w, status, nil)
		return
	}

	var CatDetail models.CatDetail
	if err := json.Unmarshal(data, &CatDetail); err != nil {
		app.Errors(w, http.StatusInternalServerError, err)
		return
	}

	tmpl, err := template.ParseFiles("web/ui/templates/details.html")
	if err != nil {
		app.Errors(w, http.StatusInternalServerError, err)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "details.html", &CatDetail); err != nil {
		app.Errors(w, http.StatusInternalServerError, err)
		return
	}
}

func GetJsonData(url string) ([]byte, int) {
	res, err := http.Get(url)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	if res.StatusCode == http.StatusInternalServerError {
		return nil, http.StatusServiceUnavailable
	}

	defer func() {
		errClose := res.Body.Close()
		if errClose != nil {
			err = fmt.Errorf("error in closing response body: %w", errClose)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return body, res.StatusCode
}

func ExtractBreeds() (map[string]string, error) {
	breeds := make(map[string]string)

	data, status := GetJsonData(models.BreedApi)
	if status != http.StatusOK {
		return nil, fmt.Errorf(http.StatusText(http.StatusServiceUnavailable))
	}

	var CatBreed []models.CatBreed
	if err := json.Unmarshal(data, &CatBreed); err != nil {
		return nil, err
	}

	for i := range CatBreed {
		breeds[CatBreed[i].Name] = CatBreed[i].ID
	}

	return breeds, nil
}

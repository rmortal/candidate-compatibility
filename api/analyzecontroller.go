package api

import (
	"candidate-compatibility/domain/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type AnalyeController struct{}

func (ac AnalyeController) AnalyzeRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get(base, ac.HandleAnalyzeCompatibility)
	return r
}

func (ac AnalyeController) HandleAnalyzeCompatibility(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	var attributeSet model.AttributeSet
	err = json.Unmarshal(body, &attributeSet)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(attributeSet.Endurance)

	// TODO: Move into helper pattern
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResp, _ := json.Marshal(attributeSet)
	w.Write(jsonResp)

	//createdUrl, err := service.UrlService{}.CreateUrl(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//helper.HandleHttpOk(w, r, createdUrl, http.StatusCreated)
}

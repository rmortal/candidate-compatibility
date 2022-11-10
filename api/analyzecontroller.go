package api

import (
	"candidate-compatibility/domain/model"
	"candidate-compatibility/domain/service"
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

	var resources model.ResourcesDto
	err = json.Unmarshal(body, &resources)
	if err != nil {
		fmt.Println(err)
		return
	}

	resultDto := service.AnalysisService{}.AnalyzeCompatibility(resources)

	// TODO: Move into helper service
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResp, _ := json.Marshal(resultDto)
	w.Write(jsonResp)

	// TODO: Logger messages for service
}

package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Rail-KH/HTTP-Calculator/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Answer struct {
	Result float64 `json:"result"`
}

type ServerError struct {
	Error error `json:"error"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, calculation.UnprocessableEntity) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "err: %s", err.Error())
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json_bytes, _ := json.Marshal(ServerError{Error: err})
			w.Write(json_bytes)
		}

	} else {
		json_bytes, _ := json.Marshal(Answer{Result: result})
		w.Write(json_bytes)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":8080", nil)
}

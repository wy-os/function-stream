package server

import (
	"encoding/json"
	"fmt"
	"github.com/functionstream/functionstream/common"
	"github.com/functionstream/functionstream/common/model"
	"github.com/functionstream/functionstream/lib"
	"github.com/gorilla/mux"
	"io"
	"log/slog"
	"net/http"
)

func StartRESTHandlers() error {
	r := mux.NewRouter()
	manager := lib.NewFunctionManager()

	r.HandleFunc("/api/v1/function/{function_name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		functionName := vars["function_name"]

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(body) == 0 {
			http.Error(w, "The body is empty. You should provide the function definition", http.StatusBadRequest)
			return
		}

		var function model.Function
		err = json.Unmarshal(body, &function)
		if err != nil {
			http.Error(w, fmt.Errorf("failed to parse function definition: %w", err).Error(), http.StatusBadRequest)
			return
		}

		slog.Info("Starting function", slog.Any("name", functionName))
		manager.StartFunction(function)
	}).Methods("POST")

	return http.ListenAndServe(common.GetConfig().ListenAddr, r)
}

package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"lexicorn/pkg/impl/core/correct"
	"lexicorn/pkg/util/logger"
)

func ApiServer() {
	router := middleware(router())

	fmt.Println("Listening on 0.0.0.0:8000")
	logger.Error(http.ListenAndServe(":8000", router).Error())
}

func router() http.Handler {
	// DI
	_correct := NewCorrect(correct.CorrectService)

	// Set up router
	router := http.NewServeMux()

	router.HandleFunc("/correct-lang", _correct.CorrectLangHandler)

	return router
}

func middleware(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle CORS preflight access
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Session")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		router.ServeHTTP(w, r)
	})
}

/* Convenient tools */

type H map[string]interface{}

func writeResponse(w http.ResponseWriter, status int, body any) {
	w.WriteHeader(status)

	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(b))
}

func bindReqData(r *http.Request, v any) bool {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}
	defer r.Body.Close()

	return json.Unmarshal(data, &v) == nil
}

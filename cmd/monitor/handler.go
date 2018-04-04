package main

import (
	"encoding/json"
	"net/http"

	"github.com/andreipimenov/dmetric/model"
)

//WriteResponse - common helper function: marshal data to JSON and write into response
func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	j, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(j)
}

//WriteErrorResponse - helper function for errors: wrap errs in model.APIErrors, marshal to JSON and write into response
func WriteErrorResponse(w http.ResponseWriter, code int, errs ...*model.APIMessage) {
	j, _ := json.Marshal(&model.APIErrors{
		Errors: errs,
	})
	w.WriteHeader(code)
	w.Write(j)
}

//JSONCtx - setup all requests mime-type to application/json
func JSONCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

//NotAllowedHandler - handler for "Method Not Allowed" error
func NotAllowedHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, &model.APIMessage{
			Code: "NotAllowed", Message: "Method Not Allowed",
		})
	})
}

//NotFoundHandler - handler for "Not Found" error
func NotFoundHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteErrorResponse(w, http.StatusNotFound, &model.APIMessage{
			Code: "NotFound", Message: "Invalid API endpoint",
		})
	})
}

//PingHandler - health check handler
func PingHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteResponse(w, http.StatusOK, &model.APIMessage{
			Message: "pong",
		})
	})
}

//MetricsHandler - receive metrics from specific device
func MetricsHandler(a *Application) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &model.APIMetrics{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, &model.APIMessage{
				Code: "BadRequest", Message: "Cannot decode request body",
			})
			return
		}

		// if req.Key == "" {
		// 	WriteErrorResponse(w, http.StatusBadRequest, &model.APIMessage{
		// 		Code: "BadRequest", Message: "Key must being not-empty string",
		// 	})
		// 	return
		// }
		// err = s.Set(req.Key, req.Value)
		// if err != nil {
		// 	WriteErrorResponse(w, http.StatusBadRequest, &model.APIMessage{
		// 		Code: "BadRequest", Message: "Invalid value",
		// 	})
		// 	return
		// }
		// WriteResponse(w, http.StatusCreated, &model.APIMessage{
		// 	Message: "OK",
		// })
	})
}

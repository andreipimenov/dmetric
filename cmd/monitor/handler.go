package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andreipimenov/dmetric/model"
	"github.com/go-chi/chi"
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

//CheckMetric - helper function to check received metrics for type and value and collect alerts and errs
func CheckMetric(n int, metric interface{}, min int64, max int64, errs *[]*model.APIMessage, alerts *[]*model.APIMessage) {
	switch v := metric.(type) {
	case nil:
		return
	case float64:
		i := int64(v)
		if i >= min && i <= max {
			return
		}
		*alerts = append(*alerts, &model.APIMessage{
			Code:    "Alert",
			Message: fmt.Sprintf("Metric #%d with value %d is out of range [%d : %d]", n, i, min, max),
		})
	default:
		*errs = append(*errs, &model.APIMessage{
			Code:    "Error",
			Message: fmt.Sprintf("Metric #%d with value %v is not of type int64", n, metric),
		})
	}
	return
}

//MetricsHandler - receive metrics from specific device
func MetricsHandler(a *Application) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deviceID := chi.URLParam(r, "id")

		var deviceExists bool
		a.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM devices WHERE id=$1)", deviceID).Scan(&deviceExists)
		if !deviceExists {
			WriteErrorResponse(w, http.StatusNotFound, &model.APIMessage{
				Code: "NotFound", Message: fmt.Sprintf("Device with id %s not found", deviceID),
			})
			return
		}

		req := &model.APIMetrics{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, &model.APIMessage{
				Code: "BadRequest", Message: "Cannot decode request body",
			})
			return
		}

		var errs []*model.APIMessage
		var alerts []*model.APIMessage
		CheckMetric(1, req.Metric1, a.Config.MetricLimits.Metric1.Min, a.Config.MetricLimits.Metric1.Max, &errs, &alerts)
		CheckMetric(2, req.Metric2, a.Config.MetricLimits.Metric2.Min, a.Config.MetricLimits.Metric2.Max, &errs, &alerts)
		CheckMetric(3, req.Metric3, a.Config.MetricLimits.Metric3.Min, a.Config.MetricLimits.Metric3.Max, &errs, &alerts)
		CheckMetric(4, req.Metric4, a.Config.MetricLimits.Metric4.Min, a.Config.MetricLimits.Metric4.Max, &errs, &alerts)
		CheckMetric(5, req.Metric5, a.Config.MetricLimits.Metric5.Min, a.Config.MetricLimits.Metric5.Max, &errs, &alerts)
		if len(errs) > 0 {
			WriteErrorResponse(w, http.StatusBadRequest, errs...)
			return
		}

		_, err = a.DB.Exec(`INSERT INTO device_metrics(device_id, metric_1, metric_2, metric_3, metric_4, metric_5, local_time) VALUES($1, $2, $3, $4, $5, $6, $7)`,
			deviceID,
			req.Metric1,
			req.Metric2,
			req.Metric3,
			req.Metric4,
			req.Metric5,
			req.LocalTime,
		)
		if err != nil {
			log.Println(err)
		}

		if len(alerts) > 0 {
			alert, _ := json.Marshal(alerts)
			a.Cache.Set(fmt.Sprintf("DEVICE:%s:ALERT", deviceID), string(alert))
			a.DB.Exec(`INSERT INTO device_alerts(device_id, message) VALUES($1, $2)`, deviceID, string(alert))
			go a.Notifier.Send(a.Config.MailTo, fmt.Sprintf("Device %s - Notification", deviceID), string(alert))
		}

		WriteResponse(w, http.StatusCreated, &model.APIMessage{
			Message: "OK",
		})
	})
}

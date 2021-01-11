package main

import (
	"fmt"
	"github.com/5g-advanced-observability/policy-engine/src/api"
	"github.com/5g-advanced-observability/policy-engine/src/service/prometheus"
	"github.com/5g-advanced-observability/policy-engine/src/util"
	"net/http"
)

func main() {
	fmt.Println("serving server at http://" + util.Hostname() + ":" + util.Port() + " URL ...")
	prometheus.ChangeRetention("365d")

	http.HandleFunc("/", index)
	http.HandleFunc("/api/prometheus", api.PrometheusRetentionHandler)

	http.ListenAndServe(util.Port(), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http://" + util.Hostname() + ":8080 request invoked")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello policy engine ...")
}

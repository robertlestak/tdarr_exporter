package main

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robertlestak/tdarr_exporter/internal/prom"
	"github.com/robertlestak/tdarr_exporter/internal/tdarr"
	log "github.com/sirupsen/logrus"
)

func init() {
	ll, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		ll = log.InfoLevel
	}
	log.SetLevel(ll)
}

func main() {
	l := log.WithFields(log.Fields{
		"app": "tdarr_exporter",
		"fn":  "main",
	})
	l.Debug("starting tdarr_exporter")
	s := tdarr.NewServerFromEnv()
	prom.InitMetrics()
	go func() {
		for {
			l.Info("getting stats")
			stats, err := s.GetStats()
			if err != nil {
				l.WithError(err).Error("error getting stats")
				os.Exit(1)
			}
			l.Debug("got stats", stats)
			if err := stats.ExportProm(); err != nil {
				l.WithError(err).Error("error exporting stats")
				os.Exit(1)
			}
			time.Sleep(s.Interval)
		}
	}()
	l.Debug("starting http server")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9082"
	}
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		l.WithError(err).Error("error starting http server")
		os.Exit(1)
	}
}

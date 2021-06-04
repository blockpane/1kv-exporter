package kva

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

func Serve(port int) {
	log.Printf("serving metrics at 0.0.0.0:%d/metrics", port)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func ProcessResult(chain string, results chan *ValidatorInfo) {
	for {
		select {
		case r := <-results:
			l := map[string]string{
				"chain": chain,
				"validator_name": r.Name,
			}
			active.With(l).Set(r.isActive())
			bonded.With(l).Set(float64(r.Bonded) / 1_000_000_000_000)
			commission.With(l).Set(float64(r.Commission))
			faultEvents.With(l).Set(float64(len(r.FaultEvents)))
			faults.With(l).Set(float64(r.Faults))
			inclusion.With(l).Set(r.Inclusion)
			if label, values := r.isInvalid(chain); values != nil {
				for i := range values {
					invalidity.With(label[i]).Set(values[i])
				}
			}
			lastvalid.With(l).Set(time.Now().UTC().Sub(time.Unix(int64(r.LastValid/1000), 0)).Seconds())
			nominated.With(l).Set(time.Now().UTC().Sub(time.Unix(int64(r.NominatedAt/1000), 0)).Seconds())
			offlineAccum.With(l).Set(float64(r.OfflineAccumulated))
			offlineSince.With(l).Set(r.offlineSinceSeconds())
			rank.With(l).Set(float64(r.Rank))
			spanInclusion.With(l).Set(r.SpanInclusion)
			unclaimed.With(l).Set(float64(len(r.UnclaimedEras)))
		}
	}
}

var (
	active = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_active",
		Help: "Validator is active",
	}, []string{"chain", "validator_name"})
	bonded = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_bonded",
		Help: "Bonded tokens",
	}, []string{"chain", "validator_name"})
	commission = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_commission",
		Help: "Validator commission",
	}, []string{"chain", "validator_name"})
	faultEvents = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_fault_events",
		Help: "The total number of fault events",
	}, []string{"chain", "validator_name"})
	faults = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_faults",
		Help: "Validator faults",
	}, []string{"chain", "validator_name"})
	inclusion = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_inclusion",
		Help: "Validator inclusion",
	}, []string{"chain", "validator_name"})
	invalidity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_invalidity",
		Help: "Validator invalidity, seconds since invalidity even occurred",
	}, []string{"chain", "validator_name", "details"})
	lastvalid = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_last_valid",
		Help: "Seconds since last valid",
	}, []string{"chain", "validator_name"})
	nominated = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_last_nominated",
		Help: "Seconds since last nomination",
	}, []string{"chain", "validator_name"})
	offlineAccum = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_offline_accumulated",
		Help: "Accumulated offline time",
	}, []string{"chain", "validator_name"})
	offlineSince = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_offline_since",
		Help: "Current offline time",
	}, []string{"chain", "validator_name"})
	rank = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_rank",
		Help: "Current 1kv rank",
	}, []string{"chain", "validator_name"})
	// todo: rank events
	spanInclusion = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_span_inclusion",
		Help: "Span Inclusion",
	}, []string{"chain", "validator_name"})
	unclaimed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "onekv_unclaimed_eras",
		Help: "Unclaimed eras",
	}, []string{"chain", "validator_name"})

)

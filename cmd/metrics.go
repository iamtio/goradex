package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/iamtio/goradex/radexone"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

var requestSerialSleep time.Duration
var listenHTTP string

var (
	cpmMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "goradex_radiation_cpm",
		Help: "Counts per minutes.",
	})
	ambientMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "goradex_radiation_ambient_usv",
		Help: "Ambient uSv/h",
	})
	accumulatedMetric = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "goradex_radiation_accumulated_usv",
		Help: "Accumulated vaule μSv",
	})
)
var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Provide measures as prometheus metrics",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		http.Handle("/metrics", promhttp.Handler())

		go func() {
			handler := radexone.MeasureHandler{
				SerialPort: serialPort,
				SerialBaud: serialBaud,
			}
			var previousValues radexone.Measure
			log.Printf("starting handler for measures")
			for {
				values := handler.GetValues()

				cpmMetric.Set(float64(values.CPM))
				ambientMetric.Set(values.Ambient)
				accumulatedMetric.Add(values.Accumulated - previousValues.Accumulated)

				previousValues = values
				time.Sleep(requestSerialSleep)
			}
		}()
		log.Printf("starting prometheus metrics http on %s/metrics", listenHTTP)
		log.Fatal(http.ListenAndServe(listenHTTP, nil))
	},
}

func init() {
	rootCmd.AddCommand(metricsCmd)
	prometheus.MustRegister(cpmMetric)
	prometheus.MustRegister(ambientMetric)
	prometheus.MustRegister(accumulatedMetric)
	metricsCmd.PersistentFlags().DurationVarP(&requestSerialSleep, "sleep", "", time.Second*10, "sleep between requests on serial port")
	metricsCmd.PersistentFlags().StringVarP(&listenHTTP, "listen", "l", ":9090", "HTTP Listen addr")

}

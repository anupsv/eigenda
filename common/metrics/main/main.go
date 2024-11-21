package main

import (
	"fmt"
	"github.com/Layr-Labs/eigenda/common"
	"github.com/Layr-Labs/eigenda/common/metrics"
	"math/rand"
	"sync/atomic"
	"time"
)

// This is a simple test bed for validating the metrics server (since it's not straight forward to unit test).

func main() {

	metricsConfig := &metrics.Config{
		Namespace: "test",
		HTTPPort:  9101,
	}

	logger, err := common.NewLogger(common.DefaultLoggerConfig())
	if err != nil {
		panic(err)
	}

	metricsServer := metrics.NewMetrics(logger, metricsConfig)

	l1, err := metricsServer.NewLatencyMetric(
		"l1",
		"",
		metrics.NewQuantile(0.5),
		metrics.NewQuantile(0.9),
		metrics.NewQuantile(0.99))
	if err != nil {
		panic(err)
	}

	l1HALF, err := metricsServer.NewLatencyMetric(
		"l1",
		"HALF",
		metrics.NewQuantile(0.5),
		metrics.NewQuantile(0.9),
		metrics.NewQuantile(0.99))
	if err != nil {
		panic(err)
	}

	c1, err := metricsServer.NewCountMetric("c1", "")
	if err != nil {
		panic(err)
	}

	c1DOUBLE, err := metricsServer.NewCountMetric("c1", "DOUBLE")
	if err != nil {
		panic(err)
	}

	g1, err := metricsServer.NewGaugeMetric("g1", "")
	if err != nil {
		panic(err)
	}

	g2, err := metricsServer.NewGaugeMetric("g1", "previous")
	if err != nil {
		panic(err)
	}

	sum := atomic.Int64{}
	err = metricsServer.NewAutoGauge("g1", "autoPoll", 1*time.Second, func() float64 {
		return float64(sum.Load())
	})
	if err != nil {
		panic(err)
	}

	err = metricsServer.Start()
	if err != nil {
		panic(err)
	}

	prev := time.Now()
	previousElapsed := time.Duration(0)
	for i := 0; i < 100; i++ {
		fmt.Printf("Iteration %d\n", i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		now := time.Now()
		elapsed := now.Sub(prev)
		prev = now

		l1.ReportLatency(elapsed)
		l1HALF.ReportLatency(elapsed / 2)

		c1.Increment()
		c1DOUBLE.Add(2)
		g1.Set(float64(elapsed.Milliseconds()))
		g2.Set(float64(previousElapsed.Milliseconds()))

		sum.Store(sum.Load() + elapsed.Milliseconds())

		previousElapsed = elapsed
	}

	err = metricsServer.Stop()
	if err != nil {
		panic(err)
	}
}
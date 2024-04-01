package metrics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/kyverno/kyverno/pkg/config"
	cron "github.com/robfig/cron/v3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"k8s.io/client-go/kubernetes"
)

func InitMetrics(
	ctx context.Context,
	disableMetricsExport bool,
	otelProvider string,
	metricsAddr string,
	otelCollector string,
	metricsConfiguration config.MetricsConfiguration,
	transportCreds string,
	kubeClient kubernetes.Interface,
	logger logr.Logger,
) (MetricsConfigManager, *http.ServeMux, *sdkmetric.MeterProvider, error) {
	var err error
	var metricsServerMux *http.ServeMux
	if !disableMetricsExport {
		metricsServerMux, err = initMetricsHelper(ctx, otelProvider, metricsAddr, otelCollector, metricsConfiguration, transportCreds, kubeClient, logger)
		if err != nil {
			return nil, nil, nil, err
		}
		if metricsConfiguration.GetMetricsRefreshInterval() != 0 {
			err := setupMetricsRefreshCronJob(ctx, otelProvider, metricsAddr, otelCollector, metricsConfiguration, transportCreds, kubeClient, logger)
			if err != nil {
				logger.Error(err, "Failed refreshing metrics")
			}
		}
	}
	metricsConfig := MetricsConfig{
		Log:    logger,
		config: metricsConfiguration,
	}
	err = metricsConfig.initializeMetrics(otel.GetMeterProvider())
	if err != nil {
		logger.Error(err, "Failed initializing metrics")
		return nil, nil, nil, err
	}
	return &metricsConfig, metricsServerMux, nil, nil
}

func setupMetricsRefreshCronJob(ctx context.Context,
	otelProvider string,
	metricsAddr string,
	otelCollector string,
	metricsConfiguration config.MetricsConfiguration,
	transportCreds string,
	kubeClient kubernetes.Interface,
	logger logr.Logger) error {
	c := cron.New()
	_, err := c.AddFunc(fmt.Sprintf("@every %s", metricsConfiguration.GetMetricsRefreshInterval()), func() {
		initMetricsHelper(ctx, otelProvider, metricsAddr, otelCollector, metricsConfiguration, transportCreds, kubeClient, logger)
	})
	if err != nil {
		return err
	}
	c.Start()

	defer c.Stop()
	return nil
}

func initMetricsHelper(
	ctx context.Context,
	otelProvider string,
	metricsAddr string,
	otelCollector string,
	metricsConfiguration config.MetricsConfiguration,
	transportCreds string,
	kubeClient kubernetes.Interface,
	logger logr.Logger,
) (*http.ServeMux, error) {
	var meterProvider metric.MeterProvider
	var metricsServerMux *http.ServeMux
	var err error
	if otelProvider == "grpc" {
		endpoint := otelCollector + metricsAddr
		meterProvider, err = NewOTLPGRPCConfig(
			ctx,
			endpoint,
			transportCreds,
			kubeClient,
			logger,
			metricsConfiguration,
		)
		if err != nil {
			return nil, err
		}
	} else if otelProvider == "prometheus" {
		meterProvider, metricsServerMux, err = NewPrometheusConfig(ctx, logger, metricsConfiguration)
		if err != nil {
			return nil, err
		}
	}
	if meterProvider != nil {
		otel.SetMeterProvider(meterProvider)
	}
	return metricsServerMux, nil
}

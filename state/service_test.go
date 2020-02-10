package state

import (
	"bytes"
	"encoding/json"
	"expvar"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

type testCollector struct {
	testA *prometheus.Desc
	testB *prometheus.Desc
}

func (c *testCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.testA
	ch <- c.testB
}

func (c *testCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.testA, prometheus.GaugeValue, 1, "label_1")
	ch <- prometheus.MustNewConstMetric(c.testB, prometheus.GaugeValue, 2, "label_2")
}

func TestEncodeVariables(t *testing.T) {
	dump := make(map[string]interface{})

	expvar.NewString("test1").Set("test1")
	expvar.NewString("test2").Set("test2")

	res := EncodeVariables()

	require.NoError(t, json.Unmarshal(res.Variables, &dump))
	require.NotEmpty(t, dump)

	// dump should contains keys `test1` and `test2`
	require.Contains(t, dump, "test1")
	require.Equal(t, "test1", dump["test1"])

	require.Contains(t, dump, "test2")
	require.Equal(t, "test2", dump["test2"])
}

func TestEncodeConfig(t *testing.T) {
	v := viper.New()
	v.Set("test1", "test1")
	v.Set("test2", "test2")

	res, err := EncodeConfig(v)
	require.NoError(t, err)

	dump := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(res.Config, &dump))

	require.NotEmpty(t, dump)

	require.Contains(t, dump, "test1")
	require.Equal(t, dump["test1"], "test1")

	require.Contains(t, dump, "test2")
	require.Equal(t, dump["test2"], "test2")
}

func TestEncodeAndDecodeMetrics(t *testing.T) {
	registry := prometheus.NewRegistry()

	collector := &testCollector{
		testA: prometheus.NewDesc("test1", "test1", []string{"test1"}, prometheus.Labels{"label_1": "test1"}),
		testB: prometheus.NewDesc("test2", "test2", []string{"test2"}, prometheus.Labels{"label_2": "test2"}),
	}

	require.NoError(t, registry.Register(collector))

	gather, err := registry.Gather()
	require.NoError(t, err)

	res, err := EncodeMetrics(registry)
	require.NoError(t, err)

	metrics, err := DecodeMetrics(res)
	require.NoError(t, err)

	require.Len(t, metrics, len(gather))

	{ // Check that JSON encoded metrics are equal:
		expect := new(bytes.Buffer)
		actual := new(bytes.Buffer)

		require.NoError(t, json.NewEncoder(expect).Encode(gather))
		require.NoError(t, json.NewEncoder(actual).Encode(metrics))

		require.Equal(t, expect.Bytes(), actual.Bytes())
	}

	{ // Deep comparison of metrics:
		for i := range metrics {
			require.Equal(t, gather[i].Help, metrics[i].Help)
			require.Equal(t, gather[i].Name, metrics[i].Name)
			require.Equal(t, gather[i].Type, metrics[i].Type)

			require.Len(t, metrics[i].Metric, len(gather[i].Metric))

			for j := range metrics[i].Metric {
				require.Equal(t, gather[i].Metric[j].Gauge, metrics[i].Metric[j].Gauge)
				require.Len(t, metrics[i].Metric[j].Label, len(gather[i].Metric[j].Label))

				for k := range metrics[i].Metric[j].Label {
					require.Equal(t, gather[i].Metric[j].Label[k].Name, metrics[i].Metric[j].Label[k].Name)
					require.Equal(t, gather[i].Metric[j].Label[k].Value, metrics[i].Metric[j].Label[k].Value)
				}
			}
		}
	}
}

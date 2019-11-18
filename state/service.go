package state

import (
	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MetricFamily is type alias for proto.Message generated
// from github.com/prometheus/client_model/metrics.proto.
type MetricFamily = dto.MetricFamily

// EncodeMetrics encodes metrics from gatherer into MetricsResponse message,
// if something went wrong returns gRPC Status error (can be returned from service).
func EncodeMetrics(g prometheus.Gatherer) (*MetricsResponse, error) {
	metrics, err := g.Gather()
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	results := make([][]byte, 0, len(metrics))
	for _, mf := range metrics {
		item, err := proto.Marshal(mf)
		if err != nil {
			return nil, status.New(codes.Internal, err.Error()).Err()
		}

		results = append(results, item)
	}

	return &MetricsResponse{Metrics: results}, nil
}

// DecodeMetrics decodes metrics from MetricsResponse to []MetricFamily,
// if something went wrong returns error.
func DecodeMetrics(r *MetricsResponse) ([]*MetricFamily, error) {
	metrics := make([]*dto.MetricFamily, 0, len(r.Metrics))
	for i := range r.Metrics {
		mf := new(MetricFamily)
		if err := proto.Unmarshal(r.Metrics[i], mf); err != nil {
			return nil, err
		}
	}

	return metrics, nil
}

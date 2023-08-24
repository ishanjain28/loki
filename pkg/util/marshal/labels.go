package marshal

import (
	"github.com/grafana/loki/pkg/logproto"
	"github.com/prometheus/prometheus/promql/parser"

	"github.com/grafana/loki/pkg/loghttp"
)

// NewLabelSet constructs a Labelset from a promql metric list as a string
func NewLabelSet(s string) (loghttp.LabelSet, error) {
	labels, err := parser.ParseMetric(s)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string, len(labels))

	for _, l := range labels {
		ret[l.Name] = l.Value
	}

	return ret, nil
}

// NewLabelSetFromGroups constructs a Labelset from a a group of labels
func NewLabelSetFromGroups(g logproto.GroupedLabels) (loghttp.LabelSet, error) {
	size := len(g.Stream) + len(g.StructuredMetadata) + len(g.Parsed)
	ret := make(map[string]string, size)

	for _, l := range g.Stream {
		ret[l.Name] = l.Value
	}
	for _, l := range g.StructuredMetadata {
		ret[l.Name] = l.Value
	}
	for _, l := range g.Parsed {
		ret[l.Name] = l.Value
	}

	return ret, nil
}

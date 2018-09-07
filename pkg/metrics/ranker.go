package metrics

import (
	"sort"
)

type Ranker interface {
	sort.Interface
}

type RankResult struct {
	Key       string
	RealValue interface{}
	Rank      int
}

// ByIntValue implements sort.Interface for []MetricResult based on
// the RealValue field.
type ByValue []RankResult

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].RealValue < a[j].RealValue }

func RankByValue(keys []string, da DataAggregator) (results []RankResult, err error) {
	result = make([]RankResult, len(keys))
	for _, k := range keys {
		err = rm.getValue(k)
		if err != nil {
			return nil, err
		}
		results[k] = RankResult{Key: k, RealValue: da.Get(n), Rank: -1}
	}
	sort.Sort(ByValue(results))
	for i, r := range results {
		r[i].Rank = i + 1
	}

	return nil, nil
}

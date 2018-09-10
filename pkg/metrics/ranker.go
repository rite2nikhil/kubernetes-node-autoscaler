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
func (a ByValue) Less(i, j int) bool { return (a[i].RealValue).(int) < a[j].RealValue.(int) }

func RankByValue(keys []string, da DataAggregator) (results []RankResult, err error) {
	results = make([]RankResult, len(keys))
	for i, k := range keys {
		val, err := da.Get(k)
		if err != nil {
			return nil, err
		}
		results[i] = RankResult{Key: k, RealValue: val, Rank: -1}
	}
	sort.Sort(ByValue(results))
	for i, r := range results {
		r.Rank = i + 1
	}

	return nil, nil
}

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

func (a ByValue) Len() int      { return len(a) }
func (a ByValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool {
	switch a[i].RealValue.(type) {
	case int:
		// v is an int here, so e.g. v + 1 is possible.
		return (a[i].RealValue).(int) < a[j].RealValue.(int)
	case int64:
		return (a[i].RealValue).(int64) < a[j].RealValue.(int64)
	}

	return false
}

func RankByValue(keys []string, da DataAggregator) (results map[string]*RankResult, err error) {
	results = make(map[string]*RankResult)
	r := make([]RankResult, len(keys))
	for i, k := range keys {
		val, err := da.Get(k)
		if err != nil {
			return nil, err
		}
		r[i] = RankResult{Key: k, RealValue: val, Rank: -1}
	}
	sort.Sort(ByValue(r))
	last := r[0].RealValue
	curr := 1
	for i := 0; i < len(r); i++ {
		if r[i].RealValue != last {
			last = r[i]
			curr++
		}
		r[i].Rank = curr
		results[r[i].Key] = &r[i]
	}
	return results, nil
}

package kva

import "time"

type FaultEvent struct {
	PrevRank uint32 `json:"prevRank"`
	Reason   string `json:"reason"`
	When     uint64 `json:"when"`
}

type RankEvent struct {
	When      uint64 `json:"when"`
	StartEra  uint32 `json:"start_era"`
	ActiveEra uint32 `json:"active_era"`
}

type Invalid struct {
	Type    string `json:"type"`
	Valid   bool   `json:"valid"`
	Updated uint64 `json:"updated"`
	Details string `json:"details"`
}

type ValidatorInfo struct {
	Active             bool         `json:"active"`
	Bonded             uint64       `json:"bonded"`
	Commission         uint32       `json:"commission"`
	Controller         string       `json:"controller"`
	DiscoveredAt       uint64       `json:"discoveredAt"`
	FaultEvents        []FaultEvent `json:"faultEvents"`
	Faults             int          `json:"faults"`
	Inclusion          float64      `json:"inclusion"`
	Invalidity         []Invalid    `json:"invalidity"`
	InvalidityReasons  string       `json:"invalidityReasons"`
	LastValid          uint64       `json:"lastValid"`
	Name               string       `json:"name"`
	NextKeys           string       `json:"nextKeys"`
	NominatedAt        uint64       `json:"nominatedAt"`
	OfflineAccumulated uint64       `json:"offlineAccumulated"`
	OfflineSince       uint64       `json:"offlineSince"`
	OnlineSince        uint64       `json:"onlineSince"`
	QueuedKeys         string       `json:"queued_keys"`
	Rank               int          `json:"rank"`
	RankEvents         []RankEvent  `json:"rankEvents"`
	RewardDest         string       `json:"rewardDestination"`
	SpanInclusion      float64      `json:"spanInclusion"`
	Stash              string       `json:"stash"`
	UnclaimedEras      []uint32     `json:"unclaimedEras"`
	Version            string       `json:"version"`
}

func (vi ValidatorInfo) isActive() (active float64) {
	if vi.Active {
		active = 1
	}
	return
}

func (vi ValidatorInfo) isInvalid(chain string) ([]map[string]string, []float64) {
	labels := make([]map[string]string, 0)
	values := make([]float64, 0)
	for i := range vi.Invalidity {
		if vi.Invalidity[i].Valid {
			// not invalid....
			continue
		}
		labels = append(labels, map[string]string{
			"chain": chain,
			"validator_name": vi.Name,
			"details": vi.Invalidity[i].Details,
		})
		values = append(values, time.Now().UTC().Sub(time.Unix(int64(vi.Invalidity[i].Updated/1000), 0)).Seconds())
	}
	if len(values) == 0 {
		return nil, nil
	}
	return labels, values
}

func (vi ValidatorInfo) offlineSinceSeconds() float64 {
	if vi.OfflineSince == 0 {
		return 0
	}
	return time.Now().UTC().Sub(time.Unix(int64(vi.OfflineSince/1000), 0)).Seconds()
}

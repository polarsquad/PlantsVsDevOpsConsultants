package sink

import "time"

type Sink struct {
	init_time     time.Time
	mac_addresses []string
}

func Serve(m []string) *Sink {
	s := &Sink{
		init_time:     time.Now(),
		mac_addresses: m,
	}

	return s
}

package push

import "github.com/evcc-io/evcc/util"

var (
	eventC = make(chan Event, 1)
	sender *senderT
)

// Prepare initializes the push sender
func Prepare(hub *Hub, paramC chan util.Param, cache *util.Cache) {
	sender = &senderT{
		paramC: paramC,
		cache:  cache,
	}

	go hub.Run(eventC)
}

// Send sends an event to the push hub, adding current cache contents
func Send(ev Event) {
	sender.send(ev)
}

type senderT struct {
	paramC chan util.Param
	cache  *util.Cache
}

// send attaches cache contents to event and sends it to the hub
func (s *senderT) send(ev Event) {
	// let cache catch up, refs https://github.com/evcc-io/evcc/pull/445
	flushC := util.Flusher()
	s.paramC <- util.Param{Val: flushC}
	<-flushC

	// get data from cache at time of event
	ev.data = make(map[string]any)
	for _, p := range s.cache.All() {
		if p.Loadpoint == nil || ev.Loadpoint == p.Loadpoint {
			ev.data[p.Key] = p.Val
		}
	}

	eventC <- ev
}

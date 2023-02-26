package push

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/evcc-io/evcc/util"
)

// Event is a notification event
type Event struct {
	Loadpoint *int // optional loadpoint id
	Event     string
	data      map[string]any
}

// EventTemplateConfig is the push message configuration for an event
type EventTemplateConfig struct {
	Title, Msg string
}

// Hub subscribes to event notifications and sends them to client devices
type Hub struct {
	definitions map[string]EventTemplateConfig
	sender      []Messenger
}

// NewHub creates push hub with definitions and receiver
func NewHub(cc map[string]EventTemplateConfig) (*Hub, error) {
	// instantiate all event templates
	for k, v := range cc {
		if _, err := template.New("out").Funcs(template.FuncMap(sprig.FuncMap())).Parse(v.Title); err != nil {
			return nil, fmt.Errorf("invalid event title: %s (%w)", k, err)
		}
		if _, err := template.New("out").Funcs(template.FuncMap(sprig.FuncMap())).Parse(v.Msg); err != nil {
			return nil, fmt.Errorf("invalid event message: %s (%w)", k, err)
		}
	}

	h := &Hub{
		definitions: cc,
	}

	return h, nil
}

// Add adds a sender to the list of senders
func (h *Hub) Add(sender Messenger) {
	h.sender = append(h.sender, sender)
}

// apply applies the event template to the content to produce the actual message
func (h *Hub) apply(ev Event, tmpl string) (string, error) {
	// loadpoint id
	if ev.Loadpoint != nil {
		ev.data["loadpoint"] = *ev.Loadpoint + 1
	}

	return util.ReplaceFormatted(tmpl, ev.data)
}

// Run is the Hub's main publishing loop
func (h *Hub) Run(eventC <-chan Event) {
	log := util.NewLogger("push")

	for ev := range eventC {
		if len(h.sender) == 0 {
			continue
		}

		definition, ok := h.definitions[ev.Event]
		if !ok {
			continue
		}

		title, err := h.apply(ev, definition.Title)
		if err != nil {
			log.ERROR.Printf("invalid title template for %s: %v", ev.Event, err)
			continue
		}

		msg, err := h.apply(ev, definition.Msg)
		if err != nil {
			log.ERROR.Printf("invalid message template for %s: %v", ev.Event, err)
			continue
		}

		for _, sender := range h.sender {
			if strings.TrimSpace(title)+strings.TrimSpace(msg) != "" {
				go sender.Send(title, msg)
			} else {
				log.DEBUG.Printf("did not send empty message template for %s: %v", ev.Event, err)
			}
		}
	}
}

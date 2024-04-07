package logstash

import (
	"container/ring"
	"io"
	"slices"
	"strings"
	"sync"
	"unicode"

	"golang.org/x/exp/maps"
)

var DefaultHandler = New(10000)

func Areas() []string {
	return DefaultHandler.Areas()
}

func All(areas, levels []string) []string {
	return DefaultHandler.All(areas, levels)
}

func Size() int64 {
	return DefaultHandler.Size()
}

type logger struct {
	mu   sync.RWMutex
	data *ring.Ring
}

func New(size int) *logger {
	return &logger{
		data: ring.New(size),
	}
}

var _ io.Writer = (*logger)(nil)

func (l *logger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !strings.HasPrefix(string(p), "[cache ]") {
		l.data.Value = element(strings.TrimRightFunc(string(p), unicode.IsSpace))
		l.data = l.data.Next()
	}

	return len(p), nil
}

func (l *logger) Size() int64 {
	l.mu.RLock()
	defer l.mu.RUnlock()

	r := l.data
	var size int64

	for i := 0; i < r.Len(); i++ {
		if e, ok := r.Value.(element); ok {
			size += int64(len(e))
		}
		r = r.Next()
	}

	return size
}

func (l *logger) Areas() []string {
	l.mu.RLock()
	defer l.mu.RUnlock()

	r := l.data

	areas := make(map[string]struct{})
	for i := 0; i < r.Len(); i++ {
		if e, ok := r.Value.(element); ok {
			if a, _ := e.areaLevel(); a != "" {
				areas[a] = struct{}{}
			}
		}
		r = r.Next()
	}

	keys := maps.Keys(areas)
	slices.Sort(keys)
	return keys
}

func (l *logger) All(areas, levels []string) []string {
	l.mu.RLock()
	defer l.mu.RUnlock()

	r := l.data
	all := len(areas) == 0 && len(levels) == 0

	var res []string
	for i := 0; i < r.Len(); i++ {
		if e, ok := r.Value.(element); ok && e != "" && (all || e.match(areas, levels)) {
			res = append(res, string(e))
		}
		r = r.Next()
	}

	return res
}

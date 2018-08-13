package uuid

import (
	"sync"

	"math/rand"
)

// NewGenerator returns a new generator
func NewGenerator() *Generator {
	var g Generator
	g.rnd = rand.New(rand.NewSource(generateSeed()))
	return &g
}

// Generator is a thread-safe UUID generator
type Generator struct {
	mu  sync.Mutex
	rnd *rand.Rand
}

// New returns a new UUID
func (g *Generator) New() (u UUID) {
	g.mu.Lock()
	u = newUUID(g.rnd.Int63())
	g.mu.Unlock()
	return
}

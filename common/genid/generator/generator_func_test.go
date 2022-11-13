package generator

import (
	"testing"
	"time"
)

func TestGeneratorData_randDate(t *testing.T) {

	tests := []struct {
		name string
		want time.Time
	}{
		{"1", time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := new(GeneratorData)
			if got := g.randDate(); true {
				t.Logf("randDate() = %v \n", got)
			}
		})
	}
}

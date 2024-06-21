package utils

import (
	"reflect"
	"testing"
)

func TestInterpretSourceExportToGoMap(t *testing.T) {
	tests := []struct {
		name    string
		content string
		in      map[string]string
		want    map[string]string
	}{
		{
			name: "simple out",
			content: `
export a=1
`,
			in:   map[string]string{},
			want: map[string]string{"a": "1"},
		},
		{
			name: "use in",
			content: `
export a=$b
`,
			in:   map[string]string{"b": "221"},
			want: map[string]string{"a": "221"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterpretSourceExportToGoMap(tt.content, tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterpretSourceExportToGoMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

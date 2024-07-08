package utils

import (
	"reflect"
	"testing"
	"time"
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
			want: map[string]string{"a": "221", "b": "221"},
		},
		{
			name: "can deal with double quote",
			content: `
export a="$b"
`,
			in:   map[string]string{"b": "221"},
			want: map[string]string{"a": "221", "b": "221"},
		},
		{
			name: "can deal with big parentheses",
			content: `
export		a=${b}
`,
			in:   map[string]string{"b": "221"},
			want: map[string]string{"a": "221", "b": "221"},
		},
		{
			name: "can deal with /",
			content: `
export a=$b/$c
`,
			in:   map[string]string{"b": "221", "c": "222"},
			want: map[string]string{"a": "221/222", "b": "221", "c": "222"},
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

func TestReplaceStrUseEnvMapStrictWithBrace(t *testing.T) {
	tests := []struct {
		name    string
		content string
		envMap  map[string]string
		want    string
	}{
		{
			name:    "simple",
			content: `${a}`,
			envMap:  map[string]string{"a": "123"},
			want:    "123",
		},
		{
			name:    "double it",
			content: `${a}${a}`,
			envMap:  map[string]string{"a": "123"},
			want:    "123123",
		},
		{
			name:    "a and b",
			content: `${a}${b}`,
			envMap:  map[string]string{"a": "123", "b": "456"},
			want:    "123456",
		},
		{
			name:    "not found",
			content: `${a}123`,
			envMap:  map[string]string{},
			want:    "123",
		},
		{
			name:    "can deal with double quote",
			content: `"$a"123`,
			envMap:  map[string]string{"a": "456"},
			want:    "456123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceStrUseEnvMapStrictWithBrace(tt.content, tt.envMap); got != tt.want {
				t.Errorf("ReplaceStrUseEnvMapStrictWithBrace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name   string
		uptime time.Duration
		want   string
	}{
		{
			name:   "Years",
			uptime: time.Hour * 24 * 365 * 2,
			want:   "2 year ",
		},
		{
			name:   "Months",
			uptime: time.Hour * 24 * 30 * 3,
			want:   "3 month ",
		},
		{
			name:   "Days",
			uptime: time.Hour * 24 * 4,
			want:   "4 day ",
		},
		{
			name:   "Hours",
			uptime: time.Hour * 5,
			want:   "5 hour ",
		},
		{
			name:   "Minutes",
			uptime: time.Minute * 6,
			want:   "6 minute ",
		},
		{
			name:   "Seconds",
			uptime: time.Second * 7,
			want:   "7 second ",
		},
		{
			name:   "Mixed",
			uptime: time.Hour*24*365 + time.Hour*24*30 + time.Hour*24 + time.Hour + time.Minute + time.Second,
			want:   "1 year 1 month 1 day 1 hour 1 minute 1 second ",
		},
		{
			name:   "Zero",
			uptime: time.Duration(0),
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.uptime); got != tt.want {
				t.Errorf("FormatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		name  string
		total uint64
		want  string
	}{
		{
			name:  "TB",
			total: 1024 * 1024 * 1024 * 1024 * 2,
			want:  "2.00 TB",
		},
		{
			name:  "GB",
			total: 1024 * 1024 * 1024 * 3,
			want:  "3.00 GB",
		},
		{
			name:  "MB",
			total: 1024 * 1024 * 4,
			want:  "4.00 MB",
		},
		{
			name:  "KB",
			total: 1024 * 5,
			want:  "5.00 KB",
		},
		{
			name:  "B",
			total: 6,
			want:  "6 B",
		},
		{
			name:  "Zero",
			total: 0,
			want:  "0 B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBytes(tt.total); got != tt.want {
				t.Errorf("FormatBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

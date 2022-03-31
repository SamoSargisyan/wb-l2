package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnpackString(t *testing.T) {
	testTable := []struct {
		name string
		conf CutConfig
		in   []string
		out  []string
	}{
		{
			name: "simple cut",
			in: []string{"apple1	juice2	so3	tasty4"},
			conf: CutConfig{
				fields:    "1",
				delimiter: "\t",
			},
			out: []string{"apple1"},
		},
		{
			name: "cut with delimeter",
			in:   []string{"apple1;juice2;so3;tasty4"},
			conf: CutConfig{
				fields:    "2",
				delimiter: ";",
			},
			out: []string{"juice2"},
		},
		{
			name: "cut with delimeter several fields",
			in:   []string{"apple1;juice2;so3;tasty4"},
			conf: CutConfig{
				fields:    "2-4",
				delimiter: ";",
			},
			out: []string{"juice2;so3;tasty4"},
		},
		{
			name: "cut with delimeter several fields 2",
			in:   []string{"apple1;juice2;so3;tasty4"},
			conf: CutConfig{
				fields:    "3-4",
				delimiter: ";",
			},
			out: []string{"so3;tasty4"},
		},
		{
			name: "cut with delimeter several fields 3",
			in:   []string{"apple1;juice2;so3;tasty4"},
			conf: CutConfig{
				fields:    "1,2,4",
				delimiter: ";",
			},
			out: []string{"apple1;juice2;tasty4"},
		},
		{
			name: "cut with delimeter overlimited fields",
			in:   []string{"apple1;juice2;so3;tasty4"},
			conf: CutConfig{
				fields:    "10-13",
				delimiter: ";",
			},
			out: []string{""},
		},
		{
			name: "cut with delimeter overlimited fields",
			in:   []string{"apple1;juice2;so3;tasty4"},
			conf: CutConfig{
				fields:    "10-13",
				delimiter: ";",
			},
			out: []string{""},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result, err := cut(tt.in, &tt.conf)
			require.NoError(t, err)
			require.Equal(t, tt.out, result)
		})
	}
}

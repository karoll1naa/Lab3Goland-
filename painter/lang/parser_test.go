package lang

import (
	"errors"
	"strings"
	"testing"
)

func TestParser1(t *testing.T) {

	tests := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "Correct input: white command",
			input:       "white\n",
			expectedErr: nil,
		},
		{
			name:        "Correct input: green command",
			input:       "green\n",
			expectedErr: nil,
		},
		{
			name:        "Correct input: bgrect command",
			input:       "bgrect 0.1 0.1 0.5 0.5\n",
			expectedErr: nil,
		},
		{
			name:        "Correct input: figure command",
			input:       "figure 0.1 0.2\n",
			expectedErr: nil,
		},
		{
			name:        "Correct input: move command",
			input:       "move 0.1 0.2\n",
			expectedErr: nil,
		},
		{
			name:        "Correct input: reset command",
			input:       "reset\n",
			expectedErr: nil,
		},
		{
			name:        "Correct input: update command",
			input:       "update\n",
			expectedErr: nil,
		},
		{
			name:        "Incorrect input: unknown command",
			input:       "invalid\n",
			expectedErr: errors.New("invalid command invalid"),
		},
		{
			name:        "Incorrect input: missing parameters",
			input:       "bgrect 1\n",
			expectedErr: errors.New("wrong number of arguments for 'bgrect' command"),
		},
		{
			name:        "Incorrect input: incorrect parameter",
			input:       "bgrect a b c d\n",
			expectedErr: errors.New("invalid parameter for 'bgrect' command: 'a' is not a number"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{state: State{}}
			_, err := p.Parse(strings.NewReader(tt.input))
			if err == nil && tt.expectedErr != nil {
				t.Errorf("Expected error: %v, got nil", tt.expectedErr)

			}
			if err != nil && tt.expectedErr == nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestParser2(t *testing.T) {
	p := Parser{state: State{}}
	tests := []struct {
		name           string
		input          string
		expectedErr    error
		expectedOpsLen int
	}{
		{
			name:           "Correct input: white command",
			input:          "white",
			expectedErr:    nil,
			expectedOpsLen: 1,
		},
		{
			name:           "Incorrect input: wrong number of arguments for white command",
			input:          "white 1",
			expectedErr:    errors.New("wrong number of arguments for white command"),
			expectedOpsLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := p.parse(tt.input)
			if err != nil && tt.expectedErr == nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if err == nil && tt.expectedErr != nil {
				t.Errorf("Expected error: %v, got nil", tt.expectedErr)
			}
		})
	}
}

func TestCheckForErrorsInParameters(t *testing.T) {
	tests := []struct {
		name           string
		words          []string
		expectedParams []float32
		expectedErr    error
	}{
		{
			name:           "Correct input",
			words:          []string{"bgrect", "0.1", "0.4", "0.8", "0.4"},
			expectedParams: []float32{0.1, 0.4, 0.8, 0.4},
			expectedErr:    nil,
		},
		{
			name:           "Incorrect input - invalid parameter",
			words:          []string{"bgrect", "a", "b", "c", "d"},
			expectedParams: nil,
			expectedErr:    errors.New("invalid parameter for 'bgrect' command: 'a' is not a number"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := checkForErrorsInParameters(tt.words, len(tt.words))
			if err == nil && tt.expectedErr != nil {
				t.Errorf("Expected error: %v, got nil", tt.expectedErr)
			}
			if err != nil && tt.expectedErr == nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if len(params) != len(tt.expectedParams) {
				t.Errorf("Expected %v, got %v", tt.expectedParams, params)
			}
		})
	}
}

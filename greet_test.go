package hello_goo

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "regular name", input: "George", want: "Hey, George!"},
		{name: "empty string", input: "", want: "Hey, !"},
		{name: "with spaces", input: "Go Developer", want: "Hey, Go Developer!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.input)
			if got != tt.want {
				t.Errorf("Hello(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestFormal(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "regular name", input: "George", want: "Good day, George. Welcome."},
		{name: "empty string", input: "", want: "Good day, . Welcome."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Formal(tt.input)
			if got != tt.want {
				t.Errorf("Formal(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

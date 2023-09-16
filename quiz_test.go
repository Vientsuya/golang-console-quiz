package main

import (
	"reflect"
	"testing"
)

func TestGetRecordsFromFile(t *testing.T) {
	filePath := "elo.csv"
	want := [][]string{
		{"420+69", "489"},
		{"40+52", "92"},
		{"4+100", "104"},
		{"what is 10 + 12, sir?", "22"},
	}

	t.Run("Test reading records from a file", func(t *testing.T) {
		got := GetRecords(filePath)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGetQuestions(t *testing.T) {
	tests := []struct {
		name    string
		records [][]string
		want    []Question
	}{
		{
			name: "Test case 1",
			records: [][]string{
				{"Question 1", "Answer 1"},
				{"Question 2", "Answer 2"},
			},
			want: []Question{
				{"Question 1", "Answer 1"},
				{"Question 2", "Answer 2"},
			},
		},
		{
			name:    "Empty input",
			records: [][]string{},
			want:    []Question{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetQuestions(tt.records)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}

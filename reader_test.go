package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	t.Run("test case when yaml file valid", func(t *testing.T) {
		data, err := ReadFile("testdata/valid.yaml")
		expected := []Student{
			{
				Name: "Иванов Иван Иванович",
				SubjectGrades: map[string]int{
					"литература": 100, "математика": 67, "программирование": 91,
				},
			},
			{
				Name: "Петров Петр Петрович",
				SubjectGrades: map[string]int{
					"математика": 78, "социология": 61, "химия": 87,
				},
			},
		}
		require.Equal(t, expected, data)
		require.Nil(t, err)
	})

	t.Run("test case when yaml file invalid", func(t *testing.T) {
		_, err := ReadFile("testdata/invalid.yaml")
		require.Error(t, errors.Unwrap(err))
	})
	t.Run("test case when yaml file empty", func(t *testing.T) {
		_, err := ReadFile("testdata/empty.yaml")
		require.Error(t, err)
	})
	t.Run("test case when yaml file not exists", func(t *testing.T) {
		_, err := ReadFile("testdata/noname.yaml")
		require.Error(t, errors.Unwrap(err))
	})
}

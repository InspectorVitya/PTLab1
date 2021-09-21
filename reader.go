package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFile(fromPath string) ([]Student, error) {
	src, err := os.ReadFile(fromPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %q: %w", fromPath, err)
	}
	data := []Student{}
	err = yaml.Unmarshal(src, &data)
	if err != nil {
		return nil, fmt.Errorf("invalid yaml file %q: %w", fromPath, err)
	}
	if len(data) == 0 {
		return nil, errors.New("empty yaml file")
	}
	return data, nil
}

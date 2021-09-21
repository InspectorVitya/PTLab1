package main

type Student struct {
	Name          string
	SubjectGrades map[string]int `yaml:"grades,omitempty"`
}

func main() {
}

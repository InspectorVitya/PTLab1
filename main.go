package main

type Student struct {
	Name          string         `yaml:"name"`
	SubjectGrades map[string]int `yaml:"grades,omitempty"`
}

func main() {

}

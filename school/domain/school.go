package domain

type School struct {
	ID   int
	Name string
}

func (s *School) SetName(name string) {
	s.Name = name
}

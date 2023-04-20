package example

type Doer interface {
	Do(s string) string
}

type Character struct {
	Source string
	Doer   Doer
}

func (c Character) Output() string {
	return c.Doer.Do(c.Source)
}

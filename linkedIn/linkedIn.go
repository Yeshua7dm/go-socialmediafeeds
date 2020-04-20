package linkedIn

type LinkedIn struct {
	URL         string
	Name        string
	Connections int
}

//  LinkedIn struct
func (t *LinkedIn) Feed() []string {
	return []string{
		"LinkedIn Feeds",
		"Hey! Follow Please",
	}
}

func (f *LinkedIn) Fame() int {
	return f.Connections
}

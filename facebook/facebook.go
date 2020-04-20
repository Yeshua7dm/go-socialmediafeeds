package facebook

type Facebook struct {
	URL     string
	Name    string
	Friends int
}

// implement implicitly
// Facebook struct
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}

//method declaration : func (structName structType) methodName(args) return
func (f *Facebook) Fame() int {
	return f.Friends
}

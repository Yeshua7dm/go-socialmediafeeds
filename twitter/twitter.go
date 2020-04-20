package twitter

type Twitter struct {
	URL       string
	Username  string
	Followers int
}

//  Twitter struct
func (t *Twitter) Feed() []string {
	return []string{
		"Twittter Feeds",
		"Hey! Retweet Please",
	}
}

func (f *Twitter) Fame() int {
	return f.Followers
}

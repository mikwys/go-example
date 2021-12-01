package internal

import (
	"go.m3o.com/twitter"
)

type UserTimeline struct {
	Details  twitter.UserResponse     `json:"details"`
	Timeline twitter.TimelineResponse `json:"timeline"`
}

// TwitterClient is an interface that allows Loading user details and user timeline
type TwitterClient interface {
	Load(userName string) (*UserTimeline, error)
}

type mockTC struct {
}

func (m *mockTC) Load(userName string) (*UserTimeline, error) {
	panic("implement me")
}

///////

func NewM3OTwitterClient(apiToken string) TwitterClient {
	twitterService := twitter.NewTwitterService(apiToken)
	return &m3OTwitterClient{
		client: twitterService,
	}
}

// m3OTwitterClient is a concrete implementation of twitter client using m3o service
type m3OTwitterClient struct {
	client *twitter.TwitterService
}

func (m *m3OTwitterClient) Load(userName string) (*UserTimeline, error) {
	// load user details
	userDetails, err := m.client.User(&twitter.UserRequest{
		Username: userName,
	})
	if err != nil {
		return nil, err
	}
	// load recent tweets
	timeline, err := m.client.Timeline(&twitter.TimelineRequest{
		Username: userName,
	})
	if err != nil {
		return nil, err
	}
	return &UserTimeline{
		Details:  *userDetails,
		Timeline: *timeline,
	}, nil
}

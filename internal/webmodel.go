package internal

import "go.m3o.com/twitter"

type UserTimelineModel struct {
	Details  twitter.UserResponse
	Timeline twitter.TimelineResponse
}

package config

import "flag"

var (
	startMin       *int
	startMax       *int
	endMin         *int
	endMax         *int
	slackToken     *string
	slackChannelID *string
)

func init() {
	startMin = flag.Int("startMin", 8, "the minimum start time for your messages")
	startMax = flag.Int("startMax", 11, "the maximum start time for your messages")
	endMin = flag.Int("endMin", 18, "the minimum end time for your messages")
	endMax = flag.Int("endMax", 21, "the maximum end time for your messages")
	slackToken = flag.String("slackToken", "", "the slack token for your slack app, if you don't know read the readme")
	slackChannelID = flag.String("slackChannelID", "", "the slack channel id you want to send to, check readme for more info")
	flag.Parse()
}

func SlackChannelID() string {
	return *slackChannelID
}

func SlackToken() string {
	return *slackToken
}

func StartMin() int {
	return *startMin
}

func StartMax() int {
	return *startMax
}

func EndMin() int {
	return *endMin
}

func EndMax() int {
	return *endMax
}

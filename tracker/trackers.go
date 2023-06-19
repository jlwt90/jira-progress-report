package tracker

import "github.com/jlwt90/reportify/tracker/jira"

const Jira = "Jira"

var SupportedTrackers = []string{Jira}

type ProfileManager interface {
}

type Tracker interface {
	SetUpTracker() error
	ProfileManager
}

// NewTracker creates a tracker instance with the type specified.
func NewTracker(t string) (Tracker, bool) {
	if t == Jira {
		return jira.Tracker{}, true
	}
	return nil, false
}

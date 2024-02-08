package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	v1 "github.com/jplanckeel/events-tracker/pkg/apis/event/v1"
)

var loc = time.Now().Local().Location()

func TestValidateValues(t *testing.T) {

	testCases := []struct {
		name  string
		event *v1.Event
	}{
		{
			name: "OK - Test values deployment, P1 and start",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "deployment",
					Priority: "P1",
					Service:  "",
					Status:   "start",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test values incident, P2 and failure",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "incident",
					Priority: "P2",
					Service:  "",
					Status:   "failure",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test values operation, P3 and error",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "operation",
					Priority: "P3",
					Service:  "",
					Status:   "error",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test values P4 and warning",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "operation",
					Priority: "P4",
					Service:  "",
					Status:   "warning",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test values success",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "operation",
					Priority: "P3",
					Service:  "",
					Status:   "success",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test values user_update",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "operation",
					Priority: "P3",
					Service:  "",
					Status:   "user_update",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test values recommendation",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "operation",
					Priority: "P3",
					Service:  "",
					Status:   "recommendation",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test values snapshot",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "operation",
					Priority: "P3",
					Service:  "",
					Status:   "snapshot",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
	}

	for _, testCase := range testCases {
		e := ValidateValues(testCase.event)
		assert.NoError(t, e, testCase.name)
	}
}

func TestValidateValuesError(t *testing.T) {

	testCases := []struct {
		name  string
		event *v1.Event
	}{
		{
			name: "OK - Test error on attributes.type",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "error",
					Priority: "",
					Service:  "",
					Status:   "",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test error on attributes.priority",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "",
					Priority: "error",
					Service:  "",
					Status:   "",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
		{
			name: "OK - Test error on attributes.status",
			event: &v1.Event{
				Title: "",
				Attributes: v1.EventAttributes{
					Message:  "",
					Source:   "",
					Type:     "",
					Priority: "",
					Service:  "",
					Status:   "ko",
				},
				Links:    v1.EventLinks{},
				Metadata: v1.EventMetadata{},
			},
		},
	}

	for _, testCase := range testCases {
		e := ValidateValues(testCase.event)
		assert.Error(t, e, testCase.name)
	}
}

func TestParseDate(t *testing.T) {

	testCases := []struct {
		name string
		date string
	}{
		{
			name: "OK - Test layout 2024-01-20",
			date: "2024-01-20",
		},
		{
			name: "OK - Test layout 2024-01-20T15",
			date: "2024-01-20T15",
		},
		{
			name: "OK - Test layout 2024-01-20T15:01",
			date: "2024-01-20T15:01",
		},
		{
			name: "OK - Test layout 2024-01-20T15:01:05",
			date: "2024-01-20T15:01:05",
		},
		{
			name: "OK - Test layout 2024-01-21T12:09:52.496+00:00",
			date: "2024-01-21T12:09:52.496",
		},
		{
			name: "OK - Test layout 2024-01-21T12:09:52.496Z",
			date: "2024-01-21T12:09:52.496Z",
		},
	}

	for _, testCase := range testCases {
		_, e := parseDate(testCase.date)
		assert.NoError(t, e, testCase.name)
	}
}

func TestCheckDateInverted(t *testing.T) {

	testCases := []struct {
		name  string
		start time.Time
		end   time.Time
	}{
		{
			name:  "OK - Test layout 2024-01-20",
			start: time.Date(2022, time.March, 1, 0, 0, 0, 0, loc),
			end:   time.Date(2022, time.March, 2, 0, 0, 0, 0, loc),
		},
	}

	for _, testCase := range testCases {
		e := checkDateInverted(testCase.start, testCase.end)
		assert.NoError(t, e, testCase.name)
	}
}

func TestCheckDateInvertedError(t *testing.T) {

	testCases := []struct {
		name  string
		start time.Time
		end   time.Time
	}{
		{
			name:  "OK - Test layout 2024-01-20",
			start: time.Date(2022, time.March, 2, 0, 0, 0, 0, loc),
			end:   time.Date(2022, time.March, 1, 0, 0, 0, 0, loc),
		},
	}

	for _, testCase := range testCases {
		e := checkDateInverted(testCase.start, testCase.end)
		assert.Error(t, e, testCase.name)
	}
}

func TestCatchPullRequestId(t *testing.T) {

	testCases := []struct {
		name       string
		pull       string
		execptedId string
	}{
		{
			name:       "OK - Test Github Pull Request",
			pull:       "https://github.com/jplanckeel/events-tracker/pull/1543",
			execptedId: "1543",
		},
		{
			name:       "OK - Test Gitlab Merge Request",
			pull:       "https://gitlab.com/jplanckeel/events-tracker/-/merge_requests/1503",
			execptedId: "1503",
		},
	}

	for _, testCase := range testCases {
		e, _ := CatchPullRequestId(testCase.pull)
		assert.Equal(t, testCase.execptedId, e)
	}
}

func TestCatchPullRequestIdError(t *testing.T) {

	testCases := []struct {
		name string
		pull string
	}{
		{
			name: "OK - Url malformed",
			pull: "https://github.com/jplanckeel/events-tracker/pull/1543/test",
		},
		{
			name: "OK - Url is nill",
			pull: "",
		},
	}

	for _, testCase := range testCases {
		_, e := CatchPullRequestId(testCase.pull)
		assert.Error(t, e)
	}
}

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/jplanckeel/events-tracker/pkg/apis/event/v1"
)

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

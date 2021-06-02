package adalo

import (
	"reflect"
	"testing"
)

var (
	testAppID                 = "2bc38ea8-2eb1-4db0-974c-e9e031f2c0e0"
	testAudienceEmail         = "user@example.com"
	testNotificationTitleText = "Hello There"
	testNotificationBodyText  = "This is just a test..."
)

func TestMapSendNotificationInput(t *testing.T) {
	input := &SendNotificationInput{
		AppID:                 testAppID,
		AudienceEmail:         testAudienceEmail,
		NotificationTitleText: testNotificationTitleText,
		NotificationBodyText:  testNotificationBodyText,
	}

	mappedInput := mapSendNotificationInput(input)
	expectedMappedInput := &sendNotificationInput{
		AppID:    testAppID,
		Audience: &sendNotificationInputAudience{Email: testAudienceEmail},
		Notification: &sendNotificationInputNotification{
			TitleText: testNotificationTitleText,
			BodyText:  testNotificationBodyText,
		},
	}

	if !reflect.DeepEqual(expectedMappedInput, mappedInput) {
		t.Fatalf("expected: %v, got: %v", expectedMappedInput, mappedInput)
	}
}

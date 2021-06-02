package adalo

type SendNotificationInput struct {
	AppID                 string
	AudienceEmail         string
	NotificationTitleText string
	NotificationBodyText  string
}

type sendNotificationInput struct {
	AppID        string                             `json:"appId,omitempty"`
	Audience     *sendNotificationInputAudience     `json:"audience,omitempty"`
	Notification *sendNotificationInputNotification `json:"notification,omitempty"`
}

type sendNotificationInputAudience struct {
	Email string `json:"email,omitempty"`
}

type sendNotificationInputNotification struct {
	TitleText string `json:"titleText,omitempty"`
	BodyText  string `json:"bodyText,omitempty"`
}

type SendNotificationOutput struct {
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

func mapSendNotificationInput(input *SendNotificationInput) *sendNotificationInput {
	return &sendNotificationInput{
		AppID:    input.AppID,
		Audience: &sendNotificationInputAudience{Email: input.AudienceEmail},
		Notification: &sendNotificationInputNotification{
			TitleText: input.NotificationTitleText,
			BodyText:  input.NotificationBodyText,
		},
	}
}

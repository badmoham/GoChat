package enums

const (
	MessageTextOnly uint = 1

	// MessageImage and ... not implemented yet
	MessageImage        uint = 2
	MessageImageAndText uint = 3
	MessageVideo        uint = 4
	MessageVideoAndText uint = 5
	MessageFile         uint = 6
	MessageFileAndText  uint = 7
)

func MessageTypeText(code uint) string {
	switch code {
	case MessageTextOnly:
		return "text_only"
	case MessageImage:
		return "image"
	case MessageImageAndText:
		return "image_and_text"
	case MessageVideo:
		return "video"
	case MessageVideoAndText:
		return "video_and_text"
	case MessageFile:
		return "file"
	case MessageFileAndText:
		return "file_and_text"
	default:
		return ""
	}
}

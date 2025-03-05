package enums

const (
	ChatP2P     uint = 1
	ChatChannel      = 2
	ChatGroup        = 3
)

func ChatTypeText(code uint) string {
	switch code {
	case ChatP2P:
		return "personal"
	case ChatChannel:
		return "channel"
	case ChatGroup:
		return "group"
	default:
		return ""
	}
}

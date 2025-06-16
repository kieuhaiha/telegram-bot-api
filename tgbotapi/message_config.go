
package tgbotapi

type MessageConfig struct {
    ChatID           int64
    Text             string
    MessageThreadID  int // 🔥 добавлено для поддержки тредов
}

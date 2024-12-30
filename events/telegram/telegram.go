package telegram

import (
	"github.com/w212w/ReminderBot/events/telegram"
)

type Processor struct {
	tg     *telegram.Client
	offset int
}

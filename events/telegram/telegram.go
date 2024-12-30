package telegram

import "github.com/w212w/ReminderBot/enents/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}

package telegram

import "github.com/w212w/ReminderBot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}

package telegram

import (
	"github.com/pkg/errors"
	tb "gopkg.in/tucnak/telebot.v2"
	"razanaubot/pkg/config"
	"time"
)

type Service struct {
	bot *tb.Bot

	channelID tb.ChatID
}

func NewWithChannel(cfg config.Telegram) (*Service, error) {
	bot, err := tb.NewBot(tb.Settings{
		Token:     cfg.Token,
		Poller:    &tb.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tb.ModeMarkdown,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new bot")
	}
	return &Service{
		bot:       bot,
		channelID: tb.ChatID(cfg.ChannelID),
	}, nil
}

func (s *Service) SendMessage(msg string) error {
	_, err := s.bot.Send(s.channelID, msg)
	return err
}

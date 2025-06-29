package notification

import (
	"fmt"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
)

func sendTelegramNotification(message string) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		return fmt.Errorf("missing TELEGRAM_BOT_TOKEN or TELEGRAM_CHAT_ID")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	payload := strings.NewReader(fmt.Sprintf("chat_id=%s&text=%s", chatID, url2.QueryEscape(message)))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("telegram API returned status %s", resp.Status)
	}
	return nil
}

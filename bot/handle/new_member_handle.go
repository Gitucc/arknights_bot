package handle

import (
	bot "arknights_bot/bot/init"
	"arknights_bot/bot/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewMemberHandle(update tgbotapi.Update) (bool, error) {
	message := update.Message
	delMsg := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
	bot.Arknights.Send(delMsg)
	for _, member := range message.NewChatMembers {
		if member.ID == message.From.ID { // 自己加入群组
			go Verify(message)
			continue
		}
		// 邀请加入群组，无需进行验证
		utils.SaveInvite(message, &member)
		name := utils.GetFullName(message.From)
		newName := utils.GetFullName(&member)
		sendMessage := tgbotapi.NewMessage(message.Chat.ID,
			fmt.Sprintf("[%s](tg://user?id=%d)邀请了[%s](tg://user?id=%d)加入群组。",
				name, message.From.ID, newName, member.ID))
		sendMessage.ParseMode = tgbotapi.ModeMarkdownV2
		msg, _ := bot.Arknights.Send(sendMessage)
		utils.AddDelQueue(msg.Chat.ID, msg.MessageID, 2)
	}
	return true, nil
}
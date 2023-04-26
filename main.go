package main

import (
	"database/sql"
	"fmt"
	d "golang_bot/database"
	"log"
	"strings"

	_ "github.com/lib/pq"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard4 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад в меню"),
	),
)

func main() {

	data, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=testdb sslmode=disable password=qwerty")
	if err != nil {
		log.Fatal(err)
	}

	if err := data.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected 200")

	bot, err := tgbotapi.NewBotAPI("6282151514:AAGz-kpiRIpSx5JYyCRJc6oF_iQ90jPBGQs")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	manual := "\n\nОтправте только значение и на каждый вопрос по одной линиии !"

	isMessage := false
	which := ""

	sendMessege := func(update tgbotapi.Update) {
		parts := strings.SplitN(update.Message.Text, "\n", -1)
		for _, i := range parts {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, i)
			msg.ReplyMarkup = numericKeyboard4
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
	adminHandler := func(updates tgbotapi.UpdatesChannel) {
		var numericKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Создать"),
				tgbotapi.NewKeyboardButton("Привязать"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Видеть информацию"),
			),
		)

		var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Создать сущность"),
				tgbotapi.NewKeyboardButton("Создать карту"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Назад в меню"),
			),
		)

		var numericKeyboard3 = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Привязать подчинённых"),
				tgbotapi.NewKeyboardButton("Привязать карту к Даймё"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Назад в меню"),
			),
		)

		startHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главный меню")
			msg.ReplyMarkup = numericKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		createHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyMarkup = numericKeyboard2
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		createEntityHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста отправьте\n1. Имя пользователя телеграм.\n2. Прозвище.\n3. Тип сущности.\nТипы сущностей (Администратор, Согён, Даймё,Самурай и Инкассатор)"+manual)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		createCardHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста отправьте\n1. Номер карты.\n2. Банк эмитент.\n3. Дневной лимит (по умчанию 2000000).\n4. Прозвище Даймё."+manual)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		linkHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyMarkup = numericKeyboard3
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		linkSubordinateHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста отправьте\n(Кого)\n1. Прозвище пользователя.\n2. Вид его сущности.\n\n(Кому)\n3. Прозвище пользователя.\n4. Вид его сущности."+manual)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		linkCardHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста отправьте\n1. Номер карты.\n2. Прозвище даймё."+manual)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		infoAdminHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyMarkup = numericKeyboard4
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		for update := range updates {
			if update.Message != nil { // ignore non-Message updates
				if isMessage {
					sendMessege(update)
					parts := strings.SplitN(update.Message.Text, "\n", -1)
					switch which {
					case "entity":
						d.AddUser(data, d.User{TUsername: parts[0], Nickname: parts[1], UserType: parts[2]})
					}
					isMessage = false
					which = ""
				}
				switch update.Message.Text {
				case "start", "Назад в меню":
					startHandler(update)
				case "Создать":
					createHandler(update)
				case "Создать сущность":
					createEntityHandler(update)
					isMessage = true
					which = "entity"
				case "Создать карту":
					createCardHandler(update)
					isMessage = true
				case "Привязать":
					linkHandler(update)
				case "Привязать подчинённых":
					linkSubordinateHandler(update)
					isMessage = true
				case "Привязать карту к Даймё":
					linkCardHandler(update)
					isMessage = true
				case "Видеть информацию":
					infoAdminHandler(update)
				}
			}
		}
	}
	sogunHandler := func(updates tgbotapi.UpdatesChannel) {
		var numericKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Создать карту"),
				tgbotapi.NewKeyboardButton("Привязать карту"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Видеть информацию"),
			),
		)
		var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Список подчинённых"),
				tgbotapi.NewKeyboardButton("Информация про подчинённных"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Назад в меню"),
			),
		)
		menuHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyMarkup = numericKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		createCardHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста отправьте\n1. Номер карты.\n2. Банк эмитент.\n3. Дневной лимит (по умчанию 2000000).\n4. Прозвище Даймё."+manual)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		linkCardHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста отправьте\n1. Номер карты.\n2. Прозвище даймё."+manual)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		viewInfoHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Информации")
			msg.ReplyMarkup = numericKeyboard2
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		for update := range updates {
			if update.Message != nil { // ignore non-Message updates
				if isMessage {
					sendMessege(update)
					isMessage = false
				}
				switch update.Message.Text {
				case "/menu", "Назад в меню":
					menuHandler(update)
				case "Создать карту":
					createCardHandler(update)
					isMessage = true
				case "Привязать карту":
					linkCardHandler(update)
					isMessage = true
				case "Видеть информацию":
					viewInfoHandler(update)
				case "Список подчинённых":

				}
			}
		}
	}
	daymioHandler := func(updates tgbotapi.UpdatesChannel) {
		var numericKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Видеть информацию"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Создать заявку на пополнени карты"),
			),
		)
		var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Список карт"),
				tgbotapi.NewKeyboardButton("Список самураев"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Назад в меню"),
			),
		)

		menuHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню")
			msg.ReplyMarkup = numericKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		viewInfoHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Информации")
			msg.ReplyMarkup = numericKeyboard2
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		createRequestHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста отправьте\n1. Номер карты.\n2. Сумма для пополение."+manual)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		listCardsHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sooon")
			msg.ReplyMarkup = numericKeyboard4
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		listSamuraisHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sooon")
			msg.ReplyMarkup = numericKeyboard4
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		for update := range updates {
			if update.Message != nil { // ignore non-Message updates
				if isMessage {
					sendMessege(update)
					isMessage = false
				}
				switch update.Message.Text {
				case "/menu", "Назад в меню":
					menuHandler(update)
				case "Список подчинённых":
					viewInfoHandler(update)
				case "Видеть информацию":
					viewInfoHandler(update)
				case "Создать заявку на пополнени карты":
					createRequestHandler(update)
					isMessage = true
				case "Список карт":
					listCardsHandler(update)
				case "Список самураев":
					listSamuraisHandler(update)
				}
			}
		}
	}

	samuraiHandler := func(updates tgbotapi.UpdatesChannel) {
		var numericKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Вводить сумму оборота"),
			),
		)

		menuHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню")
			msg.ReplyMarkup = numericKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		enterAmountHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sooon")
			msg.ReplyMarkup = numericKeyboard4
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		for update := range updates {
			if update.Message != nil { // ignore non-Message updates
				if isMessage {
					sendMessege(update)
					isMessage = false
				}
				switch update.Message.Text {
				case "/menu", "Назад в меню":
					menuHandler(update)
				case "Вводить сумму оборота":
					enterAmountHandler(update)
				}
			}
		}
	}

	incassatorHandler := func(updates tgbotapi.UpdatesChannel) {
		var numericKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Видеть заявки"),
			),
		)
		var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Выполнять заявки"),
			),
		)
		menuHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню")
			msg.ReplyMarkup = numericKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		seeRequestsHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sooon")
			msg.ReplyMarkup = numericKeyboard2
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		fulfillRequestsHandler := func(update tgbotapi.Update) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста отправьте\n1. Идентификатор заявки.\n2. Номер карты."+manual)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		for update := range updates {
			if update.Message != nil { // ignore non-Message updates
				if isMessage {
					sendMessege(update)
					isMessage = false
				}
				switch update.Message.Text {
				case "/menu", "Назад в меню":
					menuHandler(update)
				case "Видеть заявки":
					seeRequestsHandler(update)
				case "Выполнять заявки":
					fulfillRequestsHandler(update)
					isMessage = true
				}
			}
		}

	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // ignore non-Message updates
			switch update.Message.Text {
			case "Admin":
				adminHandler(updates)
			case "Sogun":
				sogunHandler(updates)
			case "Daymio":
				daymioHandler(updates)
			case "Samurai":
				samuraiHandler(updates)
			case "Incassator":
				incassatorHandler(updates)
			}
		}
	}
}

package slack

import (
	"fmt"
	"strings"
	"time"

	"github.com/nlopes/slack"
)

/*
   TODO: Change @BOT_NAME to the same thing you entered when creating your Slack application.
   NOTE: command_arg_1 and command_arg_2 represent optional parameteras that you define
   in the Slack API UI
*/
const helpMessage = "You can say hi, ask me who's a goofy goober, and tell me to sing it."

var lyrics = [18]string{
	"I'm a Goofy Goober",
	"Rock!",
	"You're a Goofy Goober",
	"We're all Goofy Goobers",
	"Goofy Goofy Goober Goober",
	"\"Put your toys away\"",
	"When all I got to say",
	"When you tell me not to play",
	"I say no way no no no no way",
	"\"I'm a kid\" you say",
	"When you say I'm a kid I say",
	"\"Say it again\"",
	"And then I say \"thanks\"",
	"Thanks",
	"Thank you very much",
	"So if you're thinking that'd you'd like to be like me",
	"Go ahead ad try",
	"The kid inside will set you free",
}

var lyricsOrder = [27]int{0, 1, 2, 1, 3, 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 0, 1, 2, 1, 3, 1, 4}

/*
   CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
   initiating the socket connection and returning the client.
   DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

/*
   RespondToEvents waits for messages on the Slack client's incomingEvents channel,
   and sends a response when it detects the bot has been tagged in a message with @<botTag>.

   EDIT THIS FUNCTION IN THE SPACE INDICATED ONLY!
*/
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)

			// TODO: Make your bot do more than respond to a help command. See notes below.
			// Make changes below this line and add additional funcs to support your bot's functionality.
			// sendHelp is provided as a simple example. Your team may want to call a free external API
			// in a function called sendResponse that you'd create below the definition of sendHelp,
			// and call in this context to ensure execution when the bot receives an event.

			// START SLACKBOT CUSTOM CODE
			// ===============================================================
			sendResponse(slackClient, message, ev.Channel)
			sendHelp(slackClient, message, ev.Channel)
			// ===============================================================
			// END SLACKBOT CUSTOM CODE
		default:

		}
	}
}

// sendHelp is a working help message, for reference.
func sendHelp(slackClient *slack.RTM, message, slackChannel string) {
	if strings.ToLower(message) != "help" {
		return
	}
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

func sing(slackClient *slack.RTM, slackChannel string) {
	for i := 0; i < len(lyricsOrder); i++ {
		str := lyrics[lyricsOrder[i]]
		slackClient.SendMessage(slackClient.NewOutgoingMessage(str, slackChannel))
		time.Sleep(750 * time.Millisecond)
	}
}

func sendResponse(slackClient *slack.RTM, message, slackChannel string) {
	response := ""
	command := strings.ToLower(message)
	switch command {
	case "hello", "hey", "hi":
		response = "Hello fellow Goofy Goober."
		break

	case "who's a goofy goober?", "whos a goofy goober?", "who's a goofy goober", "whos a goofy goober":
		response = "I'm a Goofy Goober!"
		break

	case "sing", "sing it", "sing to me":
		sing(slackClient, slackChannel)
		break
	}

	if response != "" {
		slackClient.SendMessage(slackClient.NewOutgoingMessage(response, slackChannel))
	}
}

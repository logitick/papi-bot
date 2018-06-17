package main

import (
  "github.com/shomali11/slacker"
  "log"
)

func main() {
  bot := slacker.NewClient("xoxb-302568121895-fgvvyUITXElU2Vw0Mbo4P6gn")

  bot.Command("version", "Get the bot's version number", func(request slacker.Request, response slacker.ResponseWriter) {
    response.Reply("This is the alpha version")
  })

  bot.Help(func(request slacker.Request, response slacker.ResponseWriter) {
    response.Reply("Your own help function...")
  })

  bot.Command("repeat <word> <number>", "Repeat a word a number of times!", func(request slacker.Request, response slacker.ResponseWriter) {
    word := request.StringParam("word", "Hello!")
    number := request.IntegerParam("number", 1)
    for i := 0; i < number; i++ {
      response.Reply(word)
    }
  })

  err := bot.Listen()
  if err != nil {
    log.Fatal(err)
  }
}

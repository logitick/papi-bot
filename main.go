package main

import (
  "github.com/shomali11/slacker"
  "log"
  "github.com/nlopes/slack"
  "strings"
  "fmt"
  "os/exec"
)

func main() {
  bot := slacker.NewClient("")

  bot.Init(func() {
    log.Println("Connected!")
  })

  bot.Command("version", "Get the bot's version number", func(request slacker.Request, response slacker.ResponseWriter) {
    response.Reply("This is the alpha version")
  })

  bot.Command("tags", "Server time!", func(request slacker.Request, response slacker.ResponseWriter) {
    response.Typing()
    response.Reply(strings.Join(getTags(), "\n"))
  })

  bot.Command("release", "Echo a word!", func(request slacker.Request, response slacker.ResponseWriter) {
    response.Typing()
    println(request.Event().Channel)

    if strings.HasPrefix(request.Event().Channel, "D") {
      response.Reply(fmt.Sprintf("Hi %s, Please move this conversation to a channel", request.Event().Username))
      return
    }
    attachments := []slack.Attachment{}

    options := []slack.AttachmentActionOption{}

    for _, tag := range getTags() {
      options = append(options, slack.AttachmentActionOption{
        Text: tag,
        Value: tag,
        Description: fmt.Sprintf("Deploy version %s", tag),
      })
    }

    actions := []slack.AttachmentAction{}
    actions = append(actions, slack.AttachmentAction{
      Name: "Version deployment",
      Text: "Select version to release",
      Style: "primary",
      Type: "select",
      Options:options,
    })

    actions = append(actions, slack.AttachmentAction{
      Name: "env",
      Text: "Select environment",
      Style: "primary",
      Type: "select",
      Options:append(
        []slack.AttachmentActionOption{},
        slack.AttachmentActionOption{
          Text:"Production",
          Value: "production",
        },
        slack.AttachmentActionOption{
          Text:"Staging",
          Value: "staging",
        }),
    })

    attachments = append(attachments, slack.Attachment{
      Actions:actions,
    })

    response.Reply("Pick the tag to release", slacker.WithAttachments(attachments))
  })

  err := bot.Listen()
  if err != nil {
    log.Fatal(err)
  }
}

func getTags() []string  {
  out,err := exec.Command("git", "tag").Output()
  if err != nil {
    log.Fatal(err)
  }
  return strings.Split(fmt.Sprintf("%s", out), "\n")
}

package session

import (
	"fmt"
	"time"

	"github.com/kataras/golog"

	"github.com/qbarrand/evenings/internal/db"
	"github.com/qbarrand/evenings/internal/topic"

	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

var CliArgs = []cli.Command{
	{
		Name: "start",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "topic",
				Usage: "The session's topic",
			},
		},
		Action: func(c *cli.Context) {
			if c.NArg() < 1 {
				golog.Error("Please provide a topic for the session.")
				return
			}

			Start(c.Args().First())
		},
	},
	{
		Name:   "stop",
		Action: func(c *cli.Context) { Stop() },
	},
	{
		Name: "show",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "current",
				Usage: "Only show the current session",
			},
		},
		Action: func(c *cli.Context) {
			Show(c.Bool("current"))
		},
	},
}

type Session struct {
	gorm.Model
	Start *time.Time
	End   *time.Time
	Topic topic.Topic
}

// ToString returns a string representation of a session.
func (s *Session) ToString() string {
	return fmt.Sprintf(
		"Session #%d: %s from %s to %s",
		s.ID,
		s.Topic.Name,
		s.Start,
		s.End)
}

//
// Static
//

func isStarted(conn *gorm.DB) bool {
	var count int
	conn.Model(&Session{}).Where("start is not null and end is null").Count(&count)

	return count > 0
}

func getCurrent(conn *gorm.DB) *Session {
	var session Session

	conn.Where("start is not null and end is null").First(&session)

	return &session
}

// Show shows a list of the current sessions.
func Show(onlyCurrent bool) {
	conn := db.GetDb()
	defer conn.Close()

	var sessions []Session

	if onlyCurrent {
		s := getCurrent(conn)
		sessions = []Session{*s}
	} else {
		conn.Find(&sessions)
	}

	if len(sessions) == 0 {
		golog.Info("No session found.")
	}

	for _, s := range sessions {
		golog.Info(s.ToString())
	}
}

// Start starts a new session.
func Start(topicShortName string) {
	golog.Infof("starting %s session", topicShortName)

	conn := db.GetDb()
	defer conn.Close()

	if isStarted(conn) {
		golog.Error("A session is already active.")
		return
	}

	var topic topic.Topic
	conn.Where("short_name = ?", topicShortName).First(&topic)

	time := time.Now()

	s2 := Session{
		Start: &time,
		End:   nil,
		Topic: topic,
	}

	exists := conn.NewRecord(s2)
	golog.Infof("New record: %t", exists)

	conn.Create(&s2)

	exists = conn.NewRecord(s2)
	golog.Infof("New record: %t", exists)
}

// Stop stops the current session.
func Stop() {
	println("stopping session")

	conn := db.GetDb()
	defer conn.Close()

	if !isStarted(conn) {
		golog.Error("No session is currently started.")
		return
	}

	session := getCurrent(conn)
	time := time.Now()

	session.End = &time

	conn.Save(&session)
}

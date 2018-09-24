package topic

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/qbarrand/evenings/internal/db"
	"github.com/urfave/cli"
)

var CliArgs = []cli.Command{
	{
		Name:   "show",
		Action: Show,
	},
}

// Topic represents a study topic.
type Topic struct {
	gorm.Model
	Name      string
	ShortName string `gorm:"column:short_name"`
}

// ToString returns a string representation of a Topic.
func (t *Topic) ToString() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.ShortName)
}

//
// Static
//

// Show lists all the topics in the database.
func Show(c *cli.Context) {
	golog.Info("Showing all topics")

	conn := db.GetDb()
	defer conn.Close()

	var topics []*Topic

	conn.Find(&topics)

	for _, t := range topics {
		golog.Infof("%d %s", t.ID, t.ToString())
	}
}

package main

import (
	"fmt"
	"github.com/enjuus/gosho/models"
	flag "github.com/ogier/pflag"
	"log"
	"os"
)

var (
	help    bool
	list    bool
	add     bool
	del     int32
	update  bool
	name    string
	season  int32
	episode int32
	id      int32
)

type Env struct {
	db models.Datastore
}

func init() {
	flag.BoolVarP(&help, "help", "h", false, "Display this help message.")
	flag.BoolVarP(&list, "list", "l", false, "List all shows.")
	flag.BoolVarP(&add, "add", "a", false, "Add show, specifiend with --name, --season, --episode.")
	flag.BoolVarP(&update, "update", "u", false, "Update show, specifiend with --name, --season, --episode.")
	flag.StringVarP(&name, "name", "n", "", "The name of the show")
	flag.Int32VarP(&season, "season", "s", 0, "The season of the show")
	flag.Int32VarP(&episode, "episode", "e", 0, "The episode of the show")
	flag.Int32VarP(&id, "id", "i", 0, "The ID of the entry")
	flag.Int32VarP(&del, "del", "d", 0, "The ID of the entry to delete")
	flag.Parse()
}

func main() {
	// output help
	if help == true {
		PrintHelpMessage()
		os.Exit(0)
	}

	db, err := models.NewDB("test.db")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	if list == true {
		shws, err := env.db.AllShows()
		if err != nil {
			log.Panic(err)
		}

		for _, sh := range shws {
			fmt.Printf("%d - %s s%de%d\n", sh.ID, sh.Name, sh.Season, sh.Episode)
		}
	}

	if add == true {
		if name == "" || season <= 0 || episode <= 0 {
			fmt.Println("Please specify name, season and episode of the show")
			os.Exit(1)
		}

		err := env.db.AddShow(name, season, episode)
		if err != nil {
			log.Panic(err)
		}
	}

	if del > 0 {
		err := env.db.DeleteShow(del)
		if err != nil {
			log.Panic(err)
		}
	}

	if update == true {
		if id <= 0 {
			fmt.Println("Please specify the shows ID with --id")
			os.Exit(1)
		}

		sh, err := env.db.LoadShow(id)
		if err != nil {
			fmt.Println("No such show")
			os.Exit(1)
		}

		if name != "" {
			sh.Name = name
		}

		if season > 0 {
			sh.Season = season
		}

		if episode > 0 {
			sh.Episode = episode
		}

		err = env.db.UpdateShow(sh.ID, sh.Name, sh.Season, sh.Episode)
		if err != nil {
			log.Panic(err)
		}

	}
}

func PrintHelpMessage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}

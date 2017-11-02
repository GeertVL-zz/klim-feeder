package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/geertvl/klim-feeder/queue"
	"strconv"
	"github.com/geertvl/klim-feeder/util"
	"github.com/geertvl/klim-feeder/domain"
)

func main() {
	fmt.Println("Which one do you want to import:")
	fmt.Println("1. Actor")
	fmt.Println("2. PlanRequest")
	fmt.Println("3. Mails")
	fmt.Println("4. Plans")
	fmt.Print("Choose: ")

	scanner := bufio.NewScanner(os.Stdin)
 	scanner.Scan()
 	importType := scanner.Text()

 	fmt.Print("How many do you want to publish: ")
 	scanner.Scan()
 	publishCount, err := strconv.Atoi(scanner.Text())
 	util.FailOnError(err, "You did not enter a number")

 	switch importType {
	case "1":
		q := queue.Queue{ SyncObject: &domain.Actor{}, ImportFileName: "./resources/actor.csv", Exchange:"klim.actor", FeedCount: publishCount  }
		q.ReadAndPublish()
	case "2":
		q := queue.Queue{ SyncObject: &domain.PlanRequest{}, ImportFileName: "./resources/planrequest.csv", Exchange: "klim.planrequest", FeedCount: publishCount }
		q.ReadAndPublish()
	case "3":
		q := queue.Queue{ SyncObject: &domain.Mail{}, ImportFileName: "", Exchange: "klim.mail", FeedCount: publishCount }
		q.ReadAndPublish()
	default:
		fmt.Println("This option is not known")
	}
}
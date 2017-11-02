package queue

import (
	"io/ioutil"
	"github.com/geertvl/klim-feeder/util"
	"encoding/csv"
	"strings"
	"io"
	"log"
	"fmt"
	"github.com/geertvl/klim-feeder/domain"
	"encoding/json"
)

const exchange = "klim.actor"

func DoActor(feedCount int) {
	actorFile, err := ioutil.ReadFile("./resources/actor.csv")
	util.FailOnError(err,"Could not find the resource file")

	conn := Connect()
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	Declare(ch, exchange)

	r := csv.NewReader(strings.NewReader(string(actorFile)))
	i := 0
	for i < feedCount {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)

		actor := &domain.Actor{}
		actor.Import(record)
		body, err := json.Marshal(actor)
		util.FailOnError(err, "Could not convert the object to json")
		Publish(ch, exchange, body)

		i = i + 1
	}

	fmt.Println("Finished publishing actors (%d)", i + 1)
}

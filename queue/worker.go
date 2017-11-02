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

type Queue struct {
	SyncObject domain.Entity
	ImportFileName string
	Exchange string
	FeedCount int
}

func (q Queue) ReadAndPublish() {
	actorFile, err := ioutil.ReadFile(q.ImportFileName)
	util.FailOnError(err,"Could not find the resource file")

	conn := Connect()
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	Declare(ch, q.Exchange)

	r := csv.NewReader(strings.NewReader(string(actorFile)))
	i := 0
	for i < q.FeedCount {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)

		q.SyncObject.Import(record)
		body, err := json.Marshal(q.SyncObject)
		util.FailOnError(err, "Could not convert the object to json")
		Publish(ch, q.Exchange, body)

		i = i + 1
	}

	fmt.Println("Finished publishing actors (%d)", i + 1)
}

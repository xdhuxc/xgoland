package main

import (
	"context"
	"fmt"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	cfg := replication.BinlogSyncerConfig{
		ServerID: 1,
		Flavor:   "mysql",
		Host:     "52.221.216.74",
		Port:     3307,
		User:     "xdhuxc",
		Password: "1qaz1QAZ",
		Charset:  "utf8",
	}

	syncer := replication.NewBinlogSyncer(cfg)

	gtidSet, err := mysql.ParseGTIDSet("mysql", "d1623ded-124a-11e9-b46d-0242ac110004:1-9")
	if err != nil {
		logrus.Fatalf("Gets GTIDSet failed, ", err)
	}

	streamer, err := syncer.StartSyncGTID(gtidSet)

	if err != nil {
		logrus.Fatalf("Creates streamer failed, ", err)
	}

	for {
		ev, err := streamer.GetEvent(context.Background())

		if err != nil {
			logrus.Error("Gets event error, ", err)
		}

		// 输出 event
		ev.Dump(os.Stdout)
		fmt.Println("---------------------------------------------------")
	}

}

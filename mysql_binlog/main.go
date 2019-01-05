package main

import (
	"context"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	cfg := replication.BinlogSyncerConfig{
		ServerID: 65535,
		Flavor:   "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "19940423",
	}

	syncer := replication.NewBinlogSyncer(cfg)

	gtidSet, err := mysql.ParseGTIDSet("mysql", "")
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
	}

}

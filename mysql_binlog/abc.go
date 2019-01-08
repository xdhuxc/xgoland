package main

import (
	"context"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/sirupsen/logrus"
	"os"
	"time"
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

	s, err := syncer.StartSyncGTID(gtidSet)

	eventCount := 0
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		e, err := s.GetEvent(ctx)
		cancel()

		if err == context.DeadlineExceeded {
			eventCount += 1
			return
		}

		e.Dump(os.Stdout)
		os.Stdout.Sync()
	}
}

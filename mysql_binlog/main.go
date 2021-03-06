package main

import (
	"fmt"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func convert(b []uint8) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, "")
}

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func main() {

	cfg := replication.BinlogSyncerConfig{
		ServerID: 65535,
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
		events := streamer.DumpEvents()
		for _, v := range events {
			fmt.Println(v.Header)
			fmt.Println(string(v.RawData[:]))
			fmt.Println(convert(v.RawData))
			fmt.Println(v.RawData)
			fmt.Println(B2S(v.RawData))
		}
	}

}

package kayo

import (
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	log.Info("Create new cron")
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { log.Info("[Job 1]Every minute job\n") })

	// Start cron with one scheduled job
	log.Info("Start cron")
	c.Start()
	printCronEntries(c.Entries())
	time.Sleep(2 * time.Minute)

	// Funcs may also be added to a running Cron
	log.Info("Add new job to a running cron")
	entryID2, _ := c.AddFunc("*/2 * * * *", func() { log.Info("[Job 2]Every two minutes job\n") })
	printCronEntries(c.Entries())
	time.Sleep(5 * time.Minute)

	//Remove Job2 and add new Job2 that run every 1 minute
	log.Info("Remove Job2 and add new Job2 with schedule run every minute")
	c.Remove(entryID2)
	c.AddFunc("*/1 * * * *", func() { log.Info("[Job 2]Every one minute job\n") })
	time.Sleep(5 * time.Minute)
}

func StartKayo() {

}

func printCronEntries(cronEntries []cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}
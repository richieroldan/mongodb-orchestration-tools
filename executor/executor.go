package executor

import (
	"time"

	"github.com/percona/dcos-mongo-tools/common"
	"github.com/percona/dcos-mongo-tools/executor/metrics"
	"github.com/percona/dcos-mongo-tools/executor/pmm"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

// BackgroundJob is an interface for background backgroundJobs to be executed against the Daemon
type BackgroundJob interface {
	Name() string
	DoRun() bool
	IsRunning() bool
	Run() error
}

// Daemon is an interface for the mongodb (mongod or mongos) daemon
type Daemon interface {
	IsStarted() bool
	Start() error
	Wait()
}

type Executor struct {
	Config         *Config
	PMM            *pmm.PMM
	Metrics        *metrics.Metrics
	backgroundJobs []BackgroundJob
}

func New(config *Config) *Executor {
	e := &Executor{
		Config:         config,
		PMM:            pmm.New(config.PMM, config.FrameworkName),
		Metrics:        metrics.New(config.Metrics),
		backgroundJobs: make([]BackgroundJob, 0),
	}

	// Percona PMM
	if e.PMM.DoRun() {
		e.addBackgroundJob(e.PMM)
	} else {
		log.Info("Skipping Percona PMM client executor")
	}

	// DC/OS Metrics
	if e.Metrics.DoRun() {
		e.addBackgroundJob(e.Metrics)
	} else {
		log.Info("Skipping DC/OS Metrics client executor")
	}

	return e
}

func (e *Executor) waitForSession() (*mgo.Session, error) {
	return common.WaitForSession(
		e.Config.DB,
		e.Config.ConnectTries,
		e.Config.ConnectRetrySleep,
	)
}

func (e *Executor) addBackgroundJob(job BackgroundJob) {
	log.Debugf("Adding background job %s\n", job.Name())
	e.backgroundJobs = append(e.backgroundJobs, job)
}

func (e *Executor) backgroundJobRunner() {
	log.Info("Starting background job runner")

	log.WithFields(log.Fields{
		"delay": e.Config.DelayBackgroundJob,
	}).Info("Delaying the start of the background job runner")
	time.Sleep(e.Config.DelayBackgroundJob)

	for _, job := range e.backgroundJobs {
		log.Infof("Starting job %s\n", job.Name())
		err := job.Run()
		if err != nil {
			log.Errorf("Job %s failed: %s\n", job.Name(), err)
		}
	}

	log.Info("Completed background job runner")
}

func (e *Executor) Run(daemon Daemon) error {
	log.Infof("Running %s daemon\n", e.Config.NodeType)

	if len(e.backgroundJobs) > 0 {
		log.Infof("Waiting for %s daemon to become reachable\n", e.Config.NodeType)
		session, err := e.waitForSession()
		if err != nil {
			log.Errorf("Could not get connection to mongodb: %s\n", err)
			return err
		}
		log.Infof("Mongodb %s daemon is now reachable\n", e.Config.NodeType)
		session.Close()

		go e.backgroundJobRunner()
	} else {
		log.Info("Skipping start of background job runner, no backgroundJobs to run")
	}

	err := daemon.Start()
	if err != nil {
		return err
	}

	daemon.Wait()
	return nil
}

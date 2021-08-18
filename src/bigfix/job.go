package bigfix

import (
	"fmt"
	"sync"
)

type restable interface {
	canGet() bool
	canPut() bool
	canPost() bool
	canDel() bool
	hasPlural() bool
	canPluralPut() bool
	urlGet() string
	urlPut() string
	urlPost() string
	urlDel() string
	urlPlural() string
	updateXML(xml string)
	payloadXML() string
	updatePlural(xml string)
}

type job struct {
	id      int
	url     string
	payload string
}

type jobResult struct {
	id      int
	status  int
	result  string
	errMsg  string
	payload string
}

type jobQueue struct {
	// Default to a slice of 0 jobs
	jobs []job
	// Default to a max open job count of 5
	maxOpen int
}

var lock = &sync.Mutex{}

var jobQueueInstance *jobQueue

/*
 * Return a single instance of jobQueue
 */
func getJobQueue() *jobQueue {
	if jobQueueInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if jobQueueInstance == nil {
			fmt.Println("Creating Single Instance Now")
			jobQueueInstance = &jobQueue{make([]job, 0), 5}
			//			jobQueueInstance.jobs = make(job[], 0)
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return jobQueueInstance
}

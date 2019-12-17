package main

import (
	"Alagappan/compare"
	"Alagappan/file"
	"Alagappan/helpers"
	"Alagappan/request"
	"flag"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const noOfHttpExecutors  = 200
const noOfComparisonExecutors = 1000

var noOfCompletedComparision int64
var noOfLinesToBeExecuted int64
var wg sync.WaitGroup
var doneComparision = make(chan bool,noOfComparisonExecutors)
var doneHttp = make(chan bool,noOfHttpExecutors)

func main() {
	file,_:=os.Create("stdout.log")
	log.SetOutput(file)

	file1,file2:=readFileInputFromCommandLine(); // reading file paths from command line
	channel,noOfLinesToBeExecuted:=readFilesInBackground(file1,file2) //pushing files in channels in async mode (seperate thread)
	compChan:=make(chan compare.Compare,noOfLinesToBeExecuted)
	wg.Add(1)
	spawnHttpExecutor(channel, compChan) // spawn no of workers for Http request
	spawnComparisionExecutor(compChan) // spawn no of workers for comparing
	waitForTheTaskToComplete(&wg,noOfLinesToBeExecuted)
	wg.Wait()
}



func readFileInputFromCommandLine() (string, string){
	// reading file input from the command line
	file1 := flag.String("file1","","please specify path to file 1")
	file2 := flag.String("file2","","please specify path to file 2")

	flag.Parse()

	// if path to file1 or file2 is not specified exit the execution and print the flag defaults
	if *file1 == "" || *file2== "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return *file1,*file2
}

func readFilesInBackground(file1 string,  file2 string) (chan compare.Compare, int64) {
	channel,len,errorFile:=file.LoadFileInChannels(file1, file2)
	helpers.PrintError(errorFile)
	return channel,len
}

func spawnHttpExecutor(urlChannel chan compare.Compare, compChan chan compare.Compare) {
	for i:=0; i< noOfHttpExecutors; i++ {

		go func() {
			for {
				var respBytes1 []byte
				var respBytes2 []byte
				var err1 error
				var err2 error
				var url1 string
				var url2 string

				// getting value from channel
				urlObj:=<-urlChannel

				url1=urlObj.URL1
				url2=urlObj.URL2


				// making http request
				respBytes1,err1=request.MakeHTTPRequest(url1)
				respBytes2,err2=request.MakeHTTPRequest(url2)


				if err1 != nil {
					respBytes1=[]byte{}
				}
				if err2 != nil {
					respBytes2=[]byte{}
				}


				urlObj.Resp1 = respBytes1
				urlObj.Resp2 = respBytes2

				compChan <- urlObj

			}
		} ()
	}
}

func spawnComparisionExecutor(compChan chan compare.Compare) {
	for i:=0; i< noOfComparisonExecutors; i++ {
		go func(compChan chan compare.Compare) {
			for {
				compObj:=<- compChan
				result:=compare.Check(compObj)

				if result{
					log.Println(compObj.URL1 + " equals " + compObj.URL2)
					atomic.AddInt64(&noOfCompletedComparision,1)
				} else {
					log.Println(compObj.URL1 + " not equals " + compObj.URL2)
					atomic.AddInt64(&noOfCompletedComparision,1)
				}
			}
		} (compChan)
	}
}

func waitForTheTaskToComplete(wg *sync.WaitGroup, noOfLinesToBeExecuted int64) {
	go func() {
		tickChannel:=time.Tick(time.Second)
		for {
			<- tickChannel
			if atomic.LoadInt64(&noOfCompletedComparision) == noOfLinesToBeExecuted {
				wg.Done()
				break
			}
		}
	}()
}
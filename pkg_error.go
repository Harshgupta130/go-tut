package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"github.com/pkg/errors"
)
type config struct{

}

func readConfig(path string)(*config,error){
	file,err:=os.Open(path)
	if err!=nil{
		return nil, errors.Wrap("cant open configuration file")
	}
	defer file.Close()

	cfg:=&config{}

	return cfg,nil
}

func setupLogging(){
	out,err:=os.OpenFile("app.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY,064)
	if err!=nil{
		return
	}
	log.SetOutput(out)
}

func main(){
	setupLogging()
	cfg,err:=readConfig("/path/to/config.toml")

	if err!=nil{
		fmt.Fprintf(os.Stderr,"error:%s\n",err)
		log.Printf("error: %+v",err)
		os.
		Exit(1)
	}
	fmt.Println(cfg)

}
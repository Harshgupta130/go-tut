package main

import (
	"fmt"
	"log"
	"os"
)
func killServer(pidfile string) error{
	file,err:=os.Open(pidfile)

	if err!=nil{
		return err
	}
	defer file.Close()

	var pid int

	if _,err:=fmt.Fscanf(file,"%d",&pid);err!=nil{
		return fmt.Errorf("bad process id")

	}

	fmt.Printf("ilkling server with pid=%d\n",pid)
	

	if err:=os.Remove(pidfile);err!=nil{
		log.Printf("warning : can't remove pid file -%s",err)
	}
	return nil
}

func main(){
	if err:=killServer("server.pid");err!=nil{
		fmt.Fprintf(os.Stderr,"error: %s\n",err)
		os.Exit(1)
	}



}
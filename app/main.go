package main
import (
	"flag"
	"log"
	"rawmigration/dicommigration"
)

func main() {
	port := flag.Int("port", 104, "an int")
	server := flag.String("server", "127.0.0.1", "a string")
	aet := flag.String("aet","AE_ARCH1", "a string")
	target := flag.String("aet", "c:\\Temp", "a string")
	parallel:=flag.Int("parallel", 14, "an int")
	if err:= dicommigration.WalkAndSend(*target,*server,string(*port),*aet ,*parallel);err!=nil{
		log.Println(err)
	}
}

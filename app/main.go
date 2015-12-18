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
	target := flag.String("target", "c:\\Temp", "a string")
	parallel:=flag.Int("parallels",12, "an int")
	flag.Parse()
	if err:= dicommigration.WalkAndSend(*target,*server,string(*port),*aet ,*parallel);err!=nil{
		log.Println(err)
	}
}

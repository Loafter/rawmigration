package dicommigration
import (
	"os"
	"dicomsend/parralels"
	"path/filepath"
	"log"
)
func WalkAndSend(dir string, srv string, port string, aet string, par int) error {
	if _, err := os.Stat(dir); err != nil {
		return err
	}
	i:=0
	pb := parralels.ParralelsBallancer{MaxParralels:par, Pb:ParallelRawDicomSend{}}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		i++
		log.Println(info.Name())
		if !info.IsDir() {
			if pb.SleepedJobs() >= par * 40 {
				pb.WaitAll()
				log.Println("info: wait done")
			}
			dd := DicomSendData{Server:srv, Port:port, AET:aet, FileName:path}
			pb.StartNew(dd)
		}

		return nil
	})
	pb.WaitAll()
	log.Println(i)
	return nil
}
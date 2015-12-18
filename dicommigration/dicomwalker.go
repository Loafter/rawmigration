package dicommigration
import (
	"os"
	"dicomsend/parralels"
	"path/filepath"
	"log"
)

func WalkAndSend(dir string,srv string,port string,aet string ,par int)error{
	if _, err := os.Stat(dir); err != nil {
		return err
	}

	pb:=parralels.ParralelsBallancer{MaxParralels:par,Pb:ParallelRawDicomSend{}}
	filepath.Walk(dir,func(path string, info os.FileInfo, err error)error{
		if pb.ActiveJobs()<(par*4){
			if !info.IsDir() {
				log.Println("info: add job ----", pb.ActiveJobs())
				dd := DicomSendData{Server:srv, Port:port, AET:aet, FileName:path}
				pb.StartNew(dd)
			}
		}else {
			log.Println("info: wait done")
			pb.WaitAll()

		}
		return nil
	})
return nil
}
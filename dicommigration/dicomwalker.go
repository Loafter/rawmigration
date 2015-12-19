package dicommigration
import (
	"os"
	"dicomsend/parralels"
	"path/filepath"
)
func WalkAndSend(dir string, srv string, port string, aet string, par int,que int) error {
	if _, err := os.Stat(dir); err != nil {
		return err
	}
	pb := parralels.ParralelsBallancer{MaxParralels:par, Pb:ParallelRawDicomSend{},Done:make(chan bool),MaxQuied:que}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			dd := DicomSendData{Server:srv, Port:port, AET:aet, FileName:path}
			pb.StartNew(dd)
		}
		return nil
	})
	pb.WaitAll()
	return nil
}
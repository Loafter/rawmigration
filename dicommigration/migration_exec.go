package dicommigration
import (

	"dicomsend/parralels"
	//"os"
	"log"
	//"os/exec"
	"os"
)

type DicomSendData struct {
	Server   string
	Port     string
	AET      string
	FileName string
}
type ParallelRawDicomSend struct {

}
func (ParallelRawDicomSend)DoAction(pb* parralels.ParralelsBallancer, data interface{}) {
	ds := data.(DicomSendData)
	defer func() {
		if err := os.Remove(ds.FileName); err != nil {
			log.Println(err)
		}
	}()
	//dcmconv := os.Getenv("DCMCONV")
	log.Println("info: file done -->",ds)
	/*storescu := os.Getenv("STORESCU")
	if out, err := exec.Command(storescu, "--store", "-H", ds.Server, "-p", ds.Port,
		"--call", ds.AET,
		"--aetitle", "AE_WEBCLI",
		"-i", ds.FileName,
		"-D",
		"-D").Output(); err != nil {
		log.Printf("dicom send status : %s %s \n", out, err)
		return
	} else {
		log.Printf("success: %s\n", out)
	}*/

	return
}

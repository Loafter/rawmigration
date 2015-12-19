package dicommigration
import (

	"dicomsend/parralels"
	//"os"
	"log"
	//"os/exec"
	"os"
	"os/exec"
	"crypto/rand"
	"fmt"
	"strconv"
)

func sep() string {
	st := strconv.QuoteRune(os.PathSeparator)
	st = st[1 : len(st) - 1]
	return st
}

func genUid() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}


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
	gdcmconv := os.Getenv("GDCMCONV")
	convf:=os.TempDir() + sep() + genUid()
	defer func() {
		if err := os.Remove(convf); err != nil {
			log.Println(err)
		}
	}()
	if out, err := exec.Command(gdcmconv,"-i",ds.FileName, "-o",convf, "-w").Output(); err != nil {
		log.Printf("convert status : %s %s \n", out, err)
		return
	} else {
		log.Printf("success: %s\n", out)
	}
	storescu := os.Getenv("STORESCU")
	log.Println(storescu,ds.Server,ds.Port,convf,"-aet","AE_GDCMTK","-aec",ds.AET,"-xs")
	if out, err := exec.Command(storescu,ds.Server,ds.Port,convf,"-aet","AE_GDCMTK","-aec",ds.AET,"-xs").Output(); err != nil {
		log.Printf("dicom send status not null : %s %s \n", out, err)
		return
	} else {
		log.Printf("success: %s\n", out)
		os.Remove(ds.FileName)
	}
	return
}

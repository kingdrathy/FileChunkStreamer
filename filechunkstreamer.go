package filechunkstreamer

import (
	"loginit"
	"github.com/op/go-logging"
	"bufio"
	//"fmt"
	//"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil{
		//log.Error(e)
		panic(e)
	}
}
func StreamFileChunks() {
	loginit.InitializeLogger()
	var log = logging.MustGetLogger("example")
	log.Info("Logger initialized")

	scanWidth := 25

	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)

	f, err := os.Open("/tmp/dat")
	check(err)
	log.Info("File opened")

//	b1 := make([]byte, 20)
//	n1, err := f.Read(b1)
	buffRead := bufio.NewReader(f)
	n1, err := buffRead.Peek(40)
	check(err)
	log.Debug("40 bytes: ",string(n1))
//	fmt.Printf("%d bytes: %s\n", n1, string(b1))
//	o2, err := f.seek(0,0)

	scanSlice := make([]byte,scanWidth)

	log.Info("Preparing to build initial slices")

	for i := 0; i < scanWidth; i++ {
		temp, err := buffRead.ReadByte()
		check(err)
		log.Debug("InitLoad temp contents: ",string(temp))
		scanSlice[i] = temp//[0]
		log.Debug("InitLoad scanSlice contents: ",string(scanSlice[i]))
	}



	log.Debug("Initial scan slice contents: ",string(scanSlice))
			//len(dat)
	for i := scanWidth; i < len(dat); i++ {
		scanSlice = scanSlice[1:]
		//log.Debug("Shifted scan slice contents: ",string(scanSlice))
		temp, err := buffRead.ReadByte()
		check(err)
		scanSlice = append(scanSlice, temp)
		log.Debug("New scan slice snapshot: ",string(scanSlice))
	}

	f.Close()

}

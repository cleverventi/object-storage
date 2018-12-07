package objects

import (
	"../locate"
	"fmt"
	"io"
	"richie.com/richie/object-storage/src/lib/objectstream"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
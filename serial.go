package serial
/* Serial Interface Library
 * http://github.com/coreyshuman/serial
 * (C) 2016 Corey Shuman
 * 5/26/16
 *
 * License: MIT
 */
 
import (
	"time"
	"bufio"
	"bytes"
	"errors"
	"container/list"
	"fmt"
)



type SerialInterface struct {
	id int
	s *Port
}


var idSeed int = -1

var ifaceList *list.List = nil

func findIface(id int) *Port {
	for e := ifaceList.Front(); e != nil; e = e.Next() {
		if e.Value.(SerialInterface).id == id {
			return e.Value.(SerialInterface).s
		}
	}
	return nil
}

func removeIface(id int) {
	for e := ifaceList.Front(); e != nil; e = e.Next() {
		if e.Value.(SerialInterface).id == id {
			ifaceList.Remove(e)
			return
		}
	}
}

func Init() {
	if idSeed == -1 {
		idSeed = 1
	}
	if ifaceList == nil {
		ifaceList = list.New()
	}
}

func Connect(dev string, baud int, timeout int, rawMode bool) (id int, err error) {
	var serialIface SerialInterface
	c := &Config{Name: dev, Baud: baud, ReadTimeout: time.Millisecond * time.Duration(timeout)}
	serialIface.id = idSeed
	idSeed = idSeed + 1
	serialIface.s, err = OpenPort(c)
	if err != nil {
		id = -1
		return
	}
	if rawMode {
	   err = serialIface.s.RawMode()
	   fmt.Println("Raw Mode")
	}
	ifaceList.PushBack(serialIface)
	id = serialIface.id
	return
}

func Disconnect(id int) {
	iface := findIface(id)
	if iface != nil {
		iface.Close()
		removeIface(id)
	}
}

func Send(id int, str string) (n int, err error) {
	iface := findIface(id)
	if iface != nil {
		n, err = iface.Write([]byte(str))
		return
	} else {
		n = -1
		err = errors.New("Device id not found")
		return
	}
}

func SendBytes(id int, data []byte) (n int, err error) {
	iface := findIface(id)
	if iface != nil {
		n, err = iface.Write(data)
		return
	} else {
		n = -1
		err = errors.New("Device id not found")
		return
	}
}

func ReadLine(id int) (string, error) {
	iface := findIface(id)
	if iface != nil {
		reader := bufio.NewReader(iface)
		
		d, err := reader.ReadBytes('\n')
			
			if err != nil {
				return "", err
			}
		n := bytes.IndexByte(d, '\n')
		
		return string(d[:n]), nil
	} else {
		return "", errors.New("Device id not found")
	}
}

func ReadBytes(id int, d []byte) (n int, err error) {
	iface := findIface(id)
	if iface != nil {
		n, err = iface.Read(d)
			
			if err != nil {
				return 0, err
			}
		
		return n, nil
	} else {
		return 0, errors.New("Device id not found")
	}
}


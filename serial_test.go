package serial

import (
    "testing"
    "fmt"
    "time"
)

func TestNonblockingRead(t *testing.T) {
    defer timeTrack(time.Now(), "TestNonblockingRead", t)
    Init()
    t.Log("Connecting...")
    sid, cerr := Connect("COM1", 115200, 1)
    if cerr != nil || sid == -1 {
        t.Error("Failed to Connect.")
    }
    t.Log("Start read...")
    var d []byte
    n, err := ReadBytes(sid, d)
    if err != nil  {
        t.Error("Failed to Read.")
    }
    t.Log("End read...")
    t.Log(fmt.Sprintf("Count= (%d)", n))
}

func timeTrack(start time.Time, name string, t *testing.T) {
    elapsed := time.Since(start)
    t.Log(fmt.Sprintf("%s took %s", name, elapsed))
}
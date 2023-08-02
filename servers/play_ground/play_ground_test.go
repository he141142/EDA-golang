package play_ground

import (
	"fmt"
	"testing"
)

func TestDrakePlayGround(t *testing.T) {
	d := NewStruct[string]("hihi")
	fmt.Println(d.getHead())
}

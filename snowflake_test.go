package snowflake

import (
	"strconv"
	"testing"
)

func TestNewSnowFlake(t *testing.T) {
	_, got := NewSnowFlake(-1, 0)
	want := "data center id can't be greater than 31 or less than 0"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := NewSnowFlake(32, 0)
	want2 := "data center id can't be greater than 31 or less than 0"
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, got3 := NewSnowFlake(0, -1)
	want3 := "worker id can't be greater than 31 or less than 0"
	if got3 == nil || got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	_, got4 := NewSnowFlake(0, 32)
	want4 := "worker id can't be greater than 31 or less than 0"
	if got4 == nil || got4.Error() != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}
}

func TestDataCenterId(t *testing.T) {
	var dataCenterId int64
	for dataCenterId = 0; dataCenterId <= MaxDataCenterId; dataCenterId++ {
		snowFlake, _ := NewSnowFlake(dataCenterId, 0)
		id, _ := snowFlake.NextId()
		bits := strconv.FormatInt(id, 2)
		num := len(bits)
		want, _ := strconv.ParseInt(bits[num-22:num-17], 2, 64)
		if dataCenterId != want {
			t.Errorf("got %d; want %d", dataCenterId, want)
		}
	}
}

func TestWorkerId(t *testing.T) {
	var workerId int64
	for workerId = 0; workerId <= MaxWorkerId; workerId++ {
		snowFlake, _ := NewSnowFlake(0, workerId)
		id, _ := snowFlake.NextId()
		bits := strconv.FormatInt(id, 2)
		num := len(bits)
		want, _ := strconv.ParseInt(bits[num-17:num-12], 2, 64)
		if workerId != want {
			t.Errorf("got %d; want %d", workerId, want)
		}
	}
}

func TestNextId(t *testing.T) {
	snowFlake, _ := NewSnowFlake(0, 0)
	for i := 0; i < 8192; i++ {
		id, _ := snowFlake.NextId()
		bits := strconv.FormatInt(id, 2)
		num := len(bits)
		sequence, _ := strconv.ParseInt(bits[num-12:], 2, 64)
		if sequence < 0 || sequence > 4095 {
			t.Errorf("sequence wrong")
		}
	}
}

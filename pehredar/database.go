package pehredar

import (
	"math/rand"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/xyproto/randomstring"
)

type USB struct {
	PID           int
	VID           int
	Name          string
	FirstSeen     time.Time
	FirstSeenUser string
	LastSeen      time.Time
	LastSeenUser  string
}

type PehredarDatabase struct {
	DeviceID string
	// AllUSBs []USB
	WhitelistUSBs []USB
}

func NewPehredarDatabase() PehredarDatabase {
	pdb := parseDBFromTOMLFile()

	return pdb
}

func parseDBFromTOMLFile() PehredarDatabase {
	path := GetOrCreateDefaultDatabasePath()
	pdb := PehredarDatabase{}

	raw_db_content, err := os.ReadFile(path)
	PanicIfNotNil(err)

	if len(raw_db_content) == 0 {
		pdb.Clear()
	}

	err = toml.Unmarshal(raw_db_content, &pdb)
	PanicIfNotNil(err)

	return pdb
}

func (pdb PehredarDatabase) Clear() {
	path := GetOrCreateDefaultDatabasePath()
	empty_toml_object, err := toml.Marshal(PehredarDatabase{})
	PanicIfNotNil(err)

	err = os.WriteFile(path, empty_toml_object, os.ModePerm)
	PanicIfNotNil(err)
}

func (pdb PehredarDatabase) Save() {
	path := GetOrCreateDefaultDatabasePath()

	toml_object, err := toml.Marshal(pdb)
	PanicIfNotNil(err)

	err = os.WriteFile(path, toml_object, os.ModePerm)
	PanicIfNotNil(err)
}

func (pdb PehredarDatabase) Refresh() PehredarDatabase {
	new_pdb := PehredarDatabase{}
	path := GetOrCreateDefaultDatabasePath()

	raw_db_content, err := os.ReadFile(path)
	PanicIfNotNil(err)

	err = toml.Unmarshal(raw_db_content, &new_pdb)
	PanicIfNotNil(err)

	return new_pdb
}

func (pdb PehredarDatabase) RandomFill() {
	pdb.DeviceID = "Some Device ID"
	for i := 1; i <= rand.Intn(100); i++ {
		a := USB{
			PID:           rand.Intn(0xffff-0x0000) + 0x0000,
			VID:           rand.Intn(0xffff-0x0000) + 0x0000,
			Name:          randomstring.HumanFriendlyString(rand.Intn(36-10) + 10),
			FirstSeen:     time.Unix(rand.Int63n(time.Now().Unix()-94608000)+94608000, 0),
			FirstSeenUser: randomstring.HumanFriendlyString(rand.Intn(36-10) + 10),
			LastSeen:      time.Unix(rand.Int63n(time.Now().Unix()-94608000)+94608000, 0),
			LastSeenUser:  randomstring.HumanFriendlyString(rand.Intn(36-10) + 10),
		}
		pdb.WhitelistUSBs = append(pdb.WhitelistUSBs, a)
	}
	pdb.Save()
}

package entity

type MountainType string

const (
	MountainTypeAlpental MountainType = "Alpental"
)

type Mountain struct {
	Type  MountainType
	StIDs []string
}

var Alpental = Mountain{
	Type:  MountainTypeAlpental,
	StIDs: []string{"1", "2", "3"},
}

func GetMountain(mountainType MountainType) (Mountain, bool) {
	switch mountainType {
	case MountainTypeAlpental:
		return Alpental, true
	}
	return Mountain{}, false
}

package mock

var UserData = map[uint64]struct {
	ID      uint64
	Name    string
	Address string
}{
	1: {ID: 1, Name: "User A", Address: "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"},
}

var ManagerData = map[uint64]struct {
	ID   uint64
	Name string
}{
	2: {ID: 2, Name: "Manager A"},
	3: {ID: 3, Name: "Manager B"},
}

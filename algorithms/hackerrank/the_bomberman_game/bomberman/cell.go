package bomberman

import (
	"crypto/md5"
	"encoding/json"
)

type BombState uint8

const (
	PLANTED BombState = iota
	WAIT1
	WAIT2
	EXPLODING
	EMPTY
)

type Cell struct {
	Value rune      `json:"value"`
	State BombState `json:"bombTime"`
}

func (self *Cell) UpdateState(state BombState) {
	self.State = state
	self.Value = self.GetChar()
}

func (self *Cell) TimeProgress() {
	if self.State != EMPTY {
		self.State += 1
	}
}

func (self *Cell) PlantBombIfEmpty() {
	if self.State == EMPTY {
		self.State = PLANTED
	}
}

func (self *Cell) Hash(i, j int) (string, error) {
	// FIXME: to fix
	jsonBytes, err := json.Marshal(self)
	if err != nil {
		return "", err
	}
	bytes := md5.Sum(jsonBytes)
	return string(bytes[:]), nil
}

func (self *Cell) GetChar() rune {
	if self.State == EMPTY {
		return '.'
	}
	return 79
}

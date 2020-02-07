package core

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"math"
	"math/big"
	"strings"
)

var (
	SscMap = map[string]SscMoneyGroup {
		"wei": {Money{Id: "wei", Power: int64(0)}},
		"ssc": {Money{Id: "ssc", Power: int64(18)}},
	}
)

type Money struct{
	Id    string
	Power int64
	Value *big.Int
}

type SscMoneyGroup struct {
	Money Money
}

func NewSsc(id string, value *big.Int) SscMoneyGroup {
	id = strings.ToLower(id)
	_, exist := SscMap[id]

	if !exist {
		log.Fatalf("key %s is not within SSC", id)
	}

	moneyType := SscMap[id]

	return SscMoneyGroup{
		Money: Money {
			moneyType.Money.Id,
			moneyType.Money.Power,
			value,
		},
	}
}

func (m *SscMoneyGroup) Convert(n SscMoneyGroup) SscMoneyGroup {
	if m.Money.Power < 0 || n.Money.Power < 0 {
		log.Fatalf("We only allow positive integers and zero for POW().")
	}

	var countedValue *big.Int
	distance := m.Money.Power - n.Money.Power
	absDistance := math.Abs(float64(distance))
	valuePointer := m.Money.Value
	countedInterval := big.NewInt(int64(math.Pow(float64(10), absDistance)))

	if distance > 0 {
		countedValue = m.Money.Value.Mul(valuePointer, countedInterval)
	}

	if distance < 0 {
		countedValue = m.Money.Value.Quo(valuePointer, countedInterval)
	}

	if distance == 0 {
		countedValue = valuePointer
	}

	return NewSsc(n.Money.Id, countedValue)
}

func (m *SscMoneyGroup) PrintValue() {
	fmt.Printf("%s %s \n", m.Money.Value, m.Money.Id)
}

//Prints values only in WEI
func (m *SscMoneyGroup) PrintHex() {
	convertVal := NewSsc("wei", big.NewInt(0))
	converted := m.Convert(convertVal)
	fmt.Println(hexutil.EncodeBig(converted.Money.Value))
}

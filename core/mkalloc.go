// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// +build none

/*

   The mkalloc tool creates the genesis allocation constants in genesis_alloc.go
   It outputs a const declaration that contains an RLP-encoded list of (address, balance) tuples.

       go run mkalloc.go genesis.json

*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	total = allocSum{core.NewSsc("wei", big.NewInt(0))}
)

type allocSum struct {core.SscMoneyGroup}
type allocItem struct{ Addr, Balance *big.Int }
type allocList []allocItem

func (a allocList) Len() int           { return len(a) }
func (a allocList) Less(i, j int) bool { return a[i].Addr.Cmp(a[j].Addr) < 0 }
func (a allocList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func makelist(g *core.Genesis) allocList {
	a := make(allocList, 0, len(g.Alloc))
	for addr, account := range g.Alloc {
		if len(account.Storage) > 0 || len(account.Code) > 0 || account.Nonce != 0 {
			panic(fmt.Sprintf("can't encode account %x", addr))
		}
		bigAddr := new(big.Int).SetBytes(addr.Bytes())
		a = append(a, allocItem{bigAddr, account.Balance})
	}
	sort.Sort(a)
	return a
}

func makealloc(g *core.Genesis) string {
	a := makelist(g)
	data, err := rlp.EncodeToBytes(a)
	if err != nil {
		panic(err)
	}
	return strconv.QuoteToASCII(string(data))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkalloc genesis.json")
		os.Exit(1)
	}

	g := new(core.Genesis)
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(file).Decode(g); err != nil {
		panic(err)
	}

	executeAlloc(g)
}

func executeAlloc(g *core.Genesis) {
	if len(os.Args) < 2 {
		return;
	}

	if len(os.Args) == 2 {
		fmt.Println("const allocData =", makealloc(g))

		return;
	}

	executeCounterCmd(os.Args[2], g)
}

func executeCounterCmd(cmdName string, g *core.Genesis) {
	var list allocList = makelist(g)

	switch cmdName {
	case "count":
		list.countTotalAlloc()
	case "toHex":
		countHashFromSsc()
	case "fromHex":
		countSscFromHash()
	case "countFor":
		list.countSscForAccount()
	default:
		fmt.Println("didn't match any case")
		fmt.Println("avalible cases: `count, toHex, fromHex, countFor`")
	}
}

func countSscFromHash() {
	if len(os.Args) < 4 {
		log.Fatalf("not enough arguments to count")
	}

	decodedBig, err := hexutil.DecodeBig(os.Args[3])

	if err != nil {
		log.Fatalf("could not decode provided val")
	}

	sscWei := core.NewSsc("wei", decodedBig)
	ssc := sscWei.Convert(core.NewSsc("ssc", big.NewInt(0)))
	ssc.PrintValue()
}

func countHashFromSsc() {
	if len(os.Args) < 4 {
		log.Fatalf("not enough arguments to count")
	}

	valueString := os.Args[3]
	valueInt, err := strconv.Atoi(valueString)

	if err != nil {
		log.Fatal(err)
	}

	valueSsc := core.NewSsc("ssc", big.NewInt(int64(valueInt)))
	valueSsc.PrintHex()
}

func (list allocList) countSscForAccount() {
	if len(os.Args) < 4 {
		log.Fatalf("not enough arguments to count")
	}

	addr := common.HexToAddress(os.Args[3])
	fmt.Println(addr.Hex())

	for _, item := range list {
		if common.BigToAddress(item.Addr).Hex() == addr.Hex() {
			val := core.NewSsc("wei", item.Balance)
			val.PrintHex()
			converted := val.Convert(core.NewSsc("ssc", big.NewInt(0)))
			converted.PrintValue()
		}
	}
}

func (list allocList) countTotalAlloc() {
	for _, item := range list {
		var item allocItem = item
		total.addFromAllocItem(item)
	}

	conversion := core.NewSsc("ssc", big.NewInt(0))
	sscSum := total.Convert(conversion)
	sscSum.PrintValue()
}

func (s allocSum) addFromAllocItem(item allocItem) {
	balance := item.Balance
	s.Money.Value.Add(balance, s.Money.Value)
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type Stats struct {
	Count       int32   `json:"c"`
	Probability float32 `json:"p"`
}

type Nodes map[byte]map[byte]Stats

var nodes Nodes

func main() {
	nodes = Nodes{}

	// loader()
	// saveNodes()
	loadNodes()

	score("password", true)
	score("asdgawegwmgaf", true)
	score("00000000", true)
	score("12345678", true)
	score("shotokan", true)
	score("fubgbxna", true)
	score("sometimes I eat cabbage", true)
	score("started", true)
	score("october", true)
	score("octanuary", true)
	score("fishsticks", true)
	score("cucumbers", true)
	score("chesterfield", true)
	score("46I2skN/", true)
	score("LXu^f4h57", true)
	score("VztaJ055~", true)
	score("VztaJ055*", true)
	score("s,r9BzN94e", true)
	score("3Scx2yoW8b^", true)
	score("!!!!!!!", true)
	score("!!!!!!!!", true)
	score("!!!!!!!!!", true)
	score("!!!!!!!!!!", true)
	score("&&&&&&&", true)
	score("&&&&&&&&", true)
	score("&&&&&&&&&", true)
	score("&&&&&&&&&&", true)

	w, err := os.Open("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	scanner := bufio.NewScanner(w)
	for scanner.Scan() {
		s := scanner.Text()
		f := score(s, false)
		if f > 60 {
			fmt.Printf("%s has a score of %f\n", s, f)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

var maxEntropyPerChar float64 = 6

func score(pass string, print bool) float64 {
	entropyBits := 0.0
	lastChar := byte(0)
	for i, char := range strings.Split(pass, "") {
		b := byte(char[0])
		if i > 0 {
			if nodes[lastChar][b].Probability == 0 {
				entropyBits += maxEntropyPerChar
			} else {
				entropyBits += math.Log2(1.0 / float64(nodes[lastChar][b].Probability))
			}
		}
		lastChar = b
	}

	if print {
		fmt.Printf("%s has an estimated %f bits of entropy.\n", pass, entropyBits)
	}
	return entropyBits
}

// save the node tree
func saveNodes() {
	f, err := os.Create("nodes.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := json.Marshal(nodes)
	if err != nil {
		panic(err)
	}
	if _, err := f.Write(b); err != nil {
		panic(err)
	}
}

// load the node tree
func loadNodes() {
	f, err := os.Open("nodes.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &nodes)
	if err != nil {
		panic(err)
	}
}

// loader loads the 10-mil-pws.txt file (carriage-return-separated passwords)
// into the markov probability tree structure
func loader() {
	f, err := os.Open("10-mil-pws.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		prevChar := byte(0)
		for i, char := range strings.Split(scanner.Text(), "") {
			b := byte(char[0])
			if i > 0 {
				if nodes[prevChar] == nil {
					nodes[prevChar] = map[byte]Stats{}
				}
				s := nodes[prevChar][b]
				s.Count++
				nodes[prevChar][b] = s
			}
			prevChar = b
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	// go through and turn them all into probabilities?
	for k, v := range nodes {
		total := int32(0)
		for _, stats := range v {
			total += stats.Count
		}

		for k1 := range v {
			s := nodes[k][k1]
			s.Probability = float32(s.Count) / float32(total)
			nodes[k][k1] = s
		}
	}
}

// convert converts from the original user\tpass\r\n line format to a pass\n line format.
func convert() {
	f, err := os.Open("10-million-combos.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pws, err := os.Create("10-mil-pws.txt")
	if err != nil {
		panic(err)
	}
	defer pws.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "\t")
		pw := parts[len(parts)-1]
		if len(parts) != 2 {
			switch pw {
			case "markcgilberteternity2969":
				pw = "eternity2969"
			case "sailer1216soccer1216":
				pw = "soccer1216"
			default:
				fmt.Println("Couldn't read line: ", parts)
			}
		}
		fmt.Fprintln(pws, pw)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

package main

import "fmt"

type Key byte

type Player struct {
	// Name string
	Keys []Key
}

const (
	Silver Key = iota + 1
	Gold
	Platinum
)

func (p *Player) foundKey(k Key) error {
	if k != Silver && k != Gold && k != Platinum {
		return fmt.Errorf("unknown key: %d", k)
	}

	for _, existing := range p.Keys {
		if existing == k {
			return nil
		}
	}

	p.Keys = append(p.Keys, k)
	return nil
}

func main() {
	p := Player{}
	if err := p.foundKey(Silver); err != nil {
		fmt.Println("error: ", err)
	}
	if err := p.foundKey(Gold); err != nil {
		fmt.Println("error: ", err)
	}
	if err := p.foundKey(Silver); err != nil {
		fmt.Println("error: ", err)
	}
	if err := p.foundKey(69); err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("Player has %d keys: %v\n", len(p.Keys), p.Keys)
}

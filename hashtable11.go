package main

import (
	"errors"
	"fmt"
)

type HashMap struct {
	table [512] *Pair
}

type Pair struct {
	key   string
	value string
}

func HashTable(key string, size int) (int, error) {
	if len(key) == 0 {
		return 0, errors.New("keySize 0")
	}
	hashSum := 0
	for i := 0; i < len(key); i++ {
		hashSum += int(key[i])
	}
	return hashSum % size, nil
}

func (hmap *HashMap) insert(key string, value string) error {
	p := &Pair{key: key, value: value}
	hash, err := HashTable(key, len(hmap.table))
	if err != nil {
		return err
	}
	if hmap.table[hash] == nil {
		hmap.table[hash] = p
		return nil
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) {
		if hmap.table[i] == nil {
			hmap.table[i] = p
			return nil
		}
	}
	return errors.New("full")
}

func (hmap *HashMap) get(key string) (string, error) {
	hash, err := HashTable(key, len(hmap.table))
	if err != nil {
		return "", err
	}

	for i := hash; hmap.table[i] != nil; i = (i + 1) % len(hmap.table) {
		if hmap.table[i].key == key {
			return hmap.table[i].value, errors.New("")
		}
	}

	return "", errors.New("not found")
}

func (hmap *HashMap) del(key string) error {
	hash, err := HashTable(key, len(hmap.table))
	if err != nil {
		return err
	}

	for i := hash; hmap.table[i] != nil; i = (i + 1) % len(hmap.table) {
		if hmap.table[i].key == key { 	
			hmap.table[i] = nil
			return nil
		}
	}

	return errors.New("not found")
}



func main() {
	// Create a new HashMap
	hmap := &HashMap{}

	// Insert some key-value pairs
	hmap.insert("2", "value1")
	hmap.insert("1", "value2")
	hmap.insert("3", "value3")

	// Retrieve values by keys
	val, err := hmap.get("1")
 	if err == nil {
		fmt.Printf("1: %s\n", val)
	} else {
		fmt.Println("key1 not found")
	}

	val, err = hmap.get("key2")
	if err == nil {
		fmt.Printf("key2: %s\n", val)
	} else {
		fmt.Println("key2 not found")
	}

	// Delete a key-value pair
	err = hmap.del("key2")
	if err == nil {
		fmt.Println("key2 deleted")
	} else {
		fmt.Println("key2 not found, couldn't delete")
	}
}

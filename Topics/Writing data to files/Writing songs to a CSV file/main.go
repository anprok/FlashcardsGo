package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// DO NOT modify the contents of the `data` variable!
	data := `Title,Artist,Album,Duration
"Fix You","Coldplay","X&Y",4:56
"Clocks","Coldplay","A Rush of Blood to the Head",5:08
"Yellow","Coldplay","Parachutes",4:27
"Summertime Sadness","Lana Del Rey","Born to Die",4:25
"Young and Beautiful","Lana Del Rey","Born to Die",3:56
"Pumped Up Kicks","Foster the People","Torches",3:59
`
	// Write the data to the "songs.csv" file below
	// You can use the os.WriteFile() to write the data directly to "songs.csv"
	// Or you can use the os.Create() to create "songs.csv" and then use fmt.Fprintln() to write the data.
	file, err := os.Create("songs.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := fmt.Fprintln(file, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes written successfully!", b)
}

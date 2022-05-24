package main

import (
	"bufio"
	"log"
	os "os"
	"toyrobot/cli"
	"toyrobot/core"
	"toyrobot/simulation"
)

func main() {
	cli := &cli.Cli{}
	board, _ := core.NewBoard(5, 5)
	simulation, _ := simulation.NewSimulation(board)
	if (len(os.Args)) == 1 {
		log.Fatalf("There should be an argument specifying the commands file")
	}

	commands := os.Args[1]

	file, err := os.Open(commands)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		cmd, err := cli.ParseLine(line)
		if err != nil {
			log.Println(err.Error())
		} else {
			if err := cmd.Execute(simulation); err != nil {
				log.Println(err.Error())
			}
		}
	}

}

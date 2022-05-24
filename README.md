# ToyRobot Simulation
### Build
The application is written in Go, so to build it just needs `go build`. This has been tested with Go 1.18. 
To run tests: `go test -v ./...`

### Run

The compiler outputs a binary `toyrobot`, and there's a file called `commands.txt` in the same directory, that contains
a list of moves. To run it, simply: 

```
$ ./toyrobot commands.txt
2022/05/24 11:02:44 Executing PLACE command
2022/05/24 11:02:44 Executing MOVE command
2022/05/24 11:02:44 Executing MOVE command
2022/05/24 11:02:44 Executing RIGHT command
2022/05/24 11:02:44 Executing MOVE command
2022/05/24 11:02:44 Executing LEFT command
2022/05/24 11:02:44 Executing MOVE command
2022/05/24 11:02:44 Executing REPORT command
2022/05/24 11:02:44 Current position of robot is 1, 3, NORTH
```

### Some design choices
I've decided to go mainly with a "commands" pattern. The main driver for such decision was that I wanted to keep the `Cli`, 
the `Commands`, `Core`, and the `Simulation` packages as independent as possible, having a single resposibility for each of them.
This makes testing easier, and also reading and understanding the project.

Most of the "business logic" you'll find in the `Simulation` package. `Cli` deals with how to parse lines to commands. 
`Cmd` has all the commands and how to interact with `Simulation`, and `Core` has all the core structs such as `Board`, `Position`, etc. 

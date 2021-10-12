# Grainfuck

Golang compiler and interpreter for Brainfuck. Because.. why not? It's called Grainfuck because it's Brainfuck but simulated and compiled with Go, so Grainfuck (yaeh).
I was looking for a little cool project, and I heard about this beautiful language: the Brainfuck. It was looking interesting, so I thought about making a compiler and an interpreter for it!

## Todo

- [ ] Simulation
- [ ] Compilation

Note that simulation and emulation are inspired

## How to build

This is Go, so easy stuff
```shell
# clone repo
$ git clone https://github.com/stevancorre/grainfuck.git
$ cd grainfuck

# build
$ mkdir build
$ go build -o build

# run
$ cd build
$ ./grainfuck
```

## Brainfuck ?

It's in its name! This language is here to destroy your brain. Honestly, I think it looks kinda cool, a bit sus, but kinda cool ngl.
Here are the instructions according to [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck): 


| Op |                                                                                      Meaning                                                                                      |
|----|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| >  | Increment the data pointer                                                                                                                                                        |
| <  | Decrement the data pointer                                                                                                                                                        |
| +  | Increment the byte at the data pointer                                                                                                                                            |
| -  | Decrement the byte at the data pointer                                                                                                                                            |
| .  | Output the byte at the data pointer                                                                                                                                               |
| ,  | Accept one byte of input and store its value in the byte at the data pointer                                                                                                      |
| [  | If the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching ] command. |
| ]  | If the byte at the data pointer is nonzero,  then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching [ command |


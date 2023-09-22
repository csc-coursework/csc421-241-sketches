### Accumulate.java

This is a java program to illustrate Java Thread programming, and the need for synchronization. 

Simple programs are single threaded, and reasoning about such programs only need to consider the code
flow in a single data activity line. A program can branch, loop, have subroutine calls, or exceptions. However,
the basic notion is a well known flow of instructions, in which it is possible always to answer which instruction
preceeded and followed any given instruction.

With mutliple threads, two or more code lines are proceeding at the same time, and it might not be possible to
establish between the threads a necessary time ordering. If the threads share data, it might not be possible to
establish the order in which the threads operate on the data. Accumulate is an example which shows that such 
indeterminancy will affect the outcome of the program. 

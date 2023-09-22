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


##### Threads

The Java thread implementation is based on a Thread object, that must have a method called `run` of the proper
signature (`public void run(void)`). In practice, objects `implement` the `Runnable` interface by providing the run method.
1. The creator of the thread instantiates an object of Runnable type,
2. Then calls the `start` method on the object.
From that point on, the run code is executed in virtual parallelism with the calling thread.

Various structures allow threads to share variables. Object references can be passed during the instantiation of the Thread
object. In this case, we pack everything into a single class and use Class variables to share data. 

##### Synchronization and Concurrency.

We can place all events across all threads on a time line. Pairs of events may or may not have a necessary before-after
relationship on this time line. If event A must come before B, we write A&lt;B, and the events are called *synchronous*.
Events that cannot be shown to be synchronous are called *concurrent*. If the two events are in a single thread, they 
are synchronous. To have concurrent events, the events must be separate threads. 

The entire art of multi-threaded programming, whether it be on a CPU or on the massive conccurrencty scale of GPU's, is
how concurrency is handled to achieve efficiency and correctness.


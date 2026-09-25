### How to Accumulate


Simple programs are single threaded, and reasoning about such programs only need to consider the code
flow in a single data activity line. A program can branch, loop, have subroutine calls, or exceptions. However,
the basic notion is a well known flow of instructions, in which it is possible always to answer which instruction
preceeded and followed any given instruction.

With mutliple threads, two or more code lines are proceeding independently, and it might not be possible to
establish between the threads a necessary time ordering. If the threads share data, it might not be possible to
establish the order in which the threads operate on the data. Accumulate is an example which shows that such 
indeterminancy will affect the outcome of the program. 


##### Threads

There are both *hardware threads* and *software threads*. A hardware thread is the physical ability 
of the CPU to follow a stream of instructions. A hardware that supports multiple threads, supports
*parallelism*. A software thread is a routine, a bit of code that operates sequentially. A software
thread needs a hardware thread to execute; but having multiple software thread implies *concurrency*, which 
is different that parallelism.


##### Synchronization and Concurrency.

We can place all events across all threads on a time line. Pairs of events may or may not have a necessary before-after
relationship on this time line. If event A must come before B, we write A&lt;B, and the events are called *synchronous*.
Events that cannot be shown to be synchronous are called *concurrent*. If the two events are in a single thread, they 
are synchronous. To have concurrent events, the events must be separate threads. 

The entire art of multi-threaded programming, whether it be on a CPU or on the massive conccurrencty scale of GPU's, is
how concurrency is handled to achieve efficiency and correctness.

##### Theoretical Elements of Concurrency.

A collection of possibly concurrent processes is analyzed for three mathematical properties.
1. *Safety*. This property is also *correctness*. The collection has this 
property if the constraints are guaranteed satisfied. For instance, that never
more than one thread is running the critical section
2. *Liveness*. This property is the opposite of *deadlock*. It can be defined
as the absence of a situtation where no thread can ever make progress. For instance,
a trivial way to assure safety is to never allow any thread into the critical section. However
such a solution might not have the liveness property.
3. *Fairness*. This property can be difficult to define and assure. It is at least
the absence of *starvation*. At its weakest, 
it ensures that any process that is waiting to enter a critical section, will eventually
enter that critical section.


##### Concurrency in C


It is an abuse of words to say *concurrency in C* because the C use of
threads and synchronization is not part of the language, whereas for Java
and Go it is. However, neither is `malloc` part of the C language, nor
`printf`. There are library routines, written in C and dealing intimately
with a particular operating system, that become C through convention and
tradition.

Threads can be implemented using the [`Pthreads`](https://hpc-tutorials.llnl.gov/posix/) library,
where `P` stands for POSIX, a Portable Operating System Interface.
The program demonstrates just three Pthread operations,

1. `pthread_create` to create and start a thread;
2. `pthread_mutex_lock` and `pthread_mutex_unlock` to acquire and release a
   mutex, attached to the lock object as in Java;
3. `pthread_join` to wait for a thread to exit, so that the main program can proceed.

Once again, the program uses lots of static variables, which is not a proper programming
style, because while it makes things the simplest for small programs, it becomes
a hazard with large programs. And any useful program will become large.

```
% make c
accumulate in c: no mutex
cc -o accumulate accumulate.c -lpthread
./accumulate
thread sleeps
thread sleeps
thread sleeps
thread sleeps
thread sleeps
thread exits
thread exits
thread exits
thread exits
thread exits
the final value of accumulator is 1

accumulate in c: with mutex
cc -DUSE_MUTEX -o accumulate accumulate.c -lpthread
./accumulate
thread sleeps
thread exits
thread sleeps
thread exits
thread sleeps
thread exits
thread sleeps
thread exits
thread sleeps
thread exits
the final value of accumulator is 5
```

##### Concurrency in Go

[Go concurrency](https://go.dev/tour/concurrency/1) uses a different
principle than Java. Java uses monitors, developed independently by
C. A. R. Hoare and Per Brinch Hansen. Go is based on Communicating
Sequential Processes (CSP). There are *channels* that communicate much
like pipes.

The program `go-unsynced-accumulate.go` reiterates for us the problem of
unsynchronized access to common data. I provide two solutions.

The solution `go-mutex-accumulate.go` creates a mutex out of the
`lock` channel. It is initialized with a single element of content. To
take the lock, the channel is read. Now the channel is empty, so any
other read blocks. To release the lock, an element is returned to the
channel.

In that program, to keep track of when the goroutines are done, they also
write into a `completion` channel.

A more native version is `go-csp-accumulate.go`. This follows the CSP
paradigm, where the `passing` channel passes the current value of
`accumulate` from one goroutine to another.

```
% make go
go-unsynced-accumulate
go run go-unsynced-accumulate.go
thread sleeping on lock
thread sleeping on lock
thread sleeping on lock
thread sleeping on lock
thread sleeping on lock
thread completing
thread completing
thread completing
thread completing
thread completing
final accumulate value: 1

go-csp-accumulate
go run go-csp-accumulate.go
thread sleeping 0
an accumulator thread wrote to the completion channel
thread sleeping 1
thread sleeping 2
an accumulator thread wrote to the completion channel
an accumulator thread wrote to the completion channel
thread sleeping 3
thread sleeping 4
an accumulator thread wrote to the completion channel
an accumulator thread wrote to the completion channel
final accumulate value: 5

go-mutex-accumulate
go run go-mutex-accumulate.go
thread sleeping on lock
thread sleeping on lock
an accumulator thread wrote to the completion channel
an accumulator thread wrote to the completion channel
thread sleeping on lock
thread sleeping on lock
an accumulator thread wrote to the completion channel
an accumulator thread wrote to the completion channel
thread sleeping on lock
an accumulator thread wrote to the completion channel
final accumulate value: 5

```

##### Concurrency in Java

Java uses a [thread object](https://docs.oracle.com/javase/tutorial/essential/concurrency/procthread.html) to
handle concurrency, and a series of methods that implement a [Mesa Monitor](https://pages.mtu.edu/~shene/NSF-3/e-Book/MONITOR/monitor-types.html)
to handle the problem of synchronization. We will only use the locking aspect of the
Monitor, and for this it does not matter the details of the Monitor's type.

The Java thread implementation is based on a Thread object, that must have a method called `run` of the proper
signature (`public void run(void)`). In practice, objects `implement` the `Runnable` interface by providing the run method.
1. The creator of the thread instantiates an object of Runnable type,
2. Then calls the `start` method on the object.
3. Other calls on the thread can pause, stop or synchronize the thread.
From the start method call,  the run code is executed in virtual parallelism with the calling thread.

Various structures allow threads to share variables. Object references can be passed during 
the instantiation of the Thread
object. In this case, we pack everything into a single class and use Class variables to share data. 

One version of Accumulate is not safe. It does not assure that five increments 
of a counter, when started at the value 0, ends with the value of 1. 
Its lack of safety is deliberate, to illustrate how one must
consider concurrency. 

A second version AccumulateSync uses the java `synchronized` block to achieve
safety. This solution is also live and fair. The synchronized block introduces the concurrency tool
of the lock. In the abstract language of a high level programming language, a lock is a software
entity that can be *acquired*, *held*, and *released*, with the safety property that at any moment, 
only one thread can hold the lock. The the lock is held, another thread attempting the acquire the
lock will block until the lock is released.

In Java, every object has a lock. The synchronized code new an object simply for its lock. 
Before the Accumulator code operates on the accumulation variable, it acquires the lock, and holds
the lock until it has completed the sequence of instructions that reason over the values
of accumulation variable.

A feature of Java is that the lock is asserted for a block, and released at the end of the block.
The Java languages enforces a paradigm that safety is achieved if sequences of instructions are
*transactions*, run all or nothing, in isolation datawise from any other thread. 

There is a subtle but important notion in this code. The lock at various times by various threads.
What we do not know is the order in which the threads hold the lock. This is not possible and not 
necessary. What the lock assures is that there is some order in which the lock is held, and in that
order, all of the instructions on one thread are executed before any of the instructions of the 
other thread are executed.

<pre>
make[1]: Entering directory '/home/ubuntu/csc421/csc421-241-sketches/threads'
critical race
javac Accumulate.java 
java Accumulate
[0.0040] reading i
[0.0050] reading i
[0.0030] reading i
[0.0000] reading i
[0.0090] reading i
[2.0680] writing i
[2.0710] thread exits
[2.0740] writing i
[2.0730] writing i
[2.0760] thread exits
[2.0730] writing i
[2.0780] thread exits
[2.0720] writing i
[2.0780] thread exits
[2.0790] thread exits
accumulator = 1

synchronized
javac AccumulateSync.java
java AccumulateSync
[0.0000] reading i
[2.0900] writing i
[2.1010] reading i
[2.1010] thread exits
[4.1100] writing i
[4.1110] reading i
[4.1120] thread exits
[6.1850] writing i
[6.1860] reading i
[6.1870] thread exits
[8.1870] writing i
[8.1880] reading i
[8.1900] thread exits
[10.2610] writing i
[10.2630] thread exits
accumulator = 5
make[1]: Leaving directory '/home/ubuntu/csc421/csc421-241-sketches/threads'
</pre>

##### References

You might enjoy listening to Rob Pike and his talk [Concurrency is not Parallelism](https://www.youtube.com/watch?v=oV9rvDllKEg).

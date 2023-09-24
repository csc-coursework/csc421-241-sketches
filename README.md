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
3. Other calls on the thread can pause, stop or synchronize the thread.
From the start method call,  the run code is executed in virtual parallelism with the calling thread.

Various structures allow threads to share variables. Object references can be passed during the instantiation of the Thread
object. In this case, we pack everything into a single class and use Class variables to share data. 

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

##### Concurrency in Java

The version Accumulate is not safe. It does not assure that five increments 
of a counter, when started at the value 0, ends with the value of 1. The
version AccumulateSync is safe. It uses the java `synchronized` block to achieve
safety. The solution is also live and fair.

The Accumulator is written so that there are three events: two of reading the accumulator
variable and one writing the accumulation variable. They are all synchronus within each thread,
but might be concurrent across threads.

Deleteing code line L will make the code concurrent, and as a result the final value of 
accumulator will be 1. With line L in, and the lock caused by the synchronization block, the code is synchronous, and the final value of the accumulator will be 5.

Note that when synchronized, for events A and B in two different threads, it is not possible
to know if A&lt;B or B&lt;A. But in this case, this is an unnecessary knowledge. We only 
need to know that one or the other is true.

##### Concurrency in Go

(Go concurrency)[https://go.dev/tour/concurrency/1] uses a different 
principle than Java. Java uses a Monitor, by computer scientists Hoare. Go 
is based on Communicating Sequential Processes (CSP). There are *channels* that 
communicate just like pipes.

This example cheats the paradigm a bit by having CSP emulate a Monitor. In 
this case we just take a lock from the Monitor structure and that's all we need. 
There is more to a Monitor but here we do not need that.

Using a pair of
channels, I emulate in Go a lock with a *token passing model*, where a token
is passed into on channel by the lock meister, to be claimed by some go routine at the other 
end of the channel. It then proceeds, and when done, returns the token
by passing it into the second channel, were it is redeemed by a lock meister.
There are parallels with the lock meister routine in Go being the lock object in Java.


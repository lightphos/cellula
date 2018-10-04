# cellula

Inspired by this book

https://books.google.co.uk/books/about/Biomimicry.html?id=4XybQgAACAAJ&redir_esc=y&hl=en

## DNA

Code assembly and deployment that behaves similar to living cells. 

Cell nucleus is all the code (as in DNA).

Cell system that interprets and runs the code, cloning and deploying automatically.

## Technology.

Docker

Java/Vertx (Hazelcast)

Prefer Go

## Concepts.

Nucleus 
   Holds the complete instruction set
   
Cell
   Interprets the instructions in the nucleus and runs them

Deployer
   Should be part of the cell/nuclear
   
Communications/Event Bus

## Deployment Logic

-- All code to be run are verticles (extends AbstractVerticle) http://vertx.io

All code to do absolutely anything is in the nucleus, including the deployment code

Deployment runs as many cells as needed

## Leader Election
For N cells, each sends N messages (every cell has all N cells code). Each  receives an announcement from all cells for the 1st once and checks the timestamp of all (nanosecond time) for the 1st cell. The shortest time becomes the master.



Eg 3 cells

[cell1] 

        --> announce -> [cell2]

        --> announce -> [cell3]
        
[cell2] has shortest time. So its starts itself. Then sends the service commands to the rest

[cell2] 

        -- serve -> [cell 1]

        -- serve -> [cell 3]
        
cells 1 and 3 now realise they are not the master and stop checking to become master and start themselves. How do they know who they are as they all started with the first cell. Each has its own unique id when they started. The master will have built a set of all the cells details from the announcement and then distributes the remain cells to the rest. It then sends this cell details to each. Each one only runs the cell that matches its own ID to ensure only one is run.


What if the master dies. Or indeed any cell dies.  What about a heartbeat by all?

This is only one kind of election - what about others?

https://en.wikipedia.org/wiki/Leader_election
http://www.goodmath.org/blog/2015/01/30/paxos-a-really-beautiful-protocol-for-distributed-consensus/
https://en.wikipedia.org/wiki/Gossip_protocol
https://raft.github.io

We need simple leader election - plus something to make sure? - Do we even need a leader?

## Blueprint
The blueprint contains what to build - also in the nucleus

## Cell EcoSystem
The cells run in an ecosystem/environment (virtual machine)



        

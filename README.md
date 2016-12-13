# cellula

Inspired by this book

https://books.google.co.uk/books/about/Biomimicry.html?id=4XybQgAACAAJ&redir_esc=y&hl=en

Code assembly and deployment that behaves similar to living cells. 

Cell nucleus is all the code (as in DNA).

Cell system that interprets and runs the code, cloning and deploying automatically.

Technology.

Docker

Vertx (Hazelcast)

Java 9 modules


Concepts.

Nucleus 
   Holds the complete instruction set
   
Cell
   Interprets the instructions in the nucleus and runs them

Deployer
   Should be part of the cell?
   
Deployer logic...

All code to be run are verticles (extends AbstractVerticle) http://vertx.io

Deploy runs as many as there are verticles (translated to cells).

For N cells, each cell sends N messages (each cell has all N cells code). Each cell receives an announcement from all the cells for the 1st once and checks the timestamp of all (nanosecond time) for the 1st cell. The shortest time becomes the master.



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


        

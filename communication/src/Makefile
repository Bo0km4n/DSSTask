JAVAC = javac
JAVA = java

SRC = client/client.c client/net.c
GCC = gcc
OBJ = bin/client

run.server:
	$(JAVAC) Server.java
	$(JAVA) Server

run.client:
	$(GCC) $(SRC) -l curl -o $(OBJ)
	$(OBJ)

clean:
	rm -rf bin
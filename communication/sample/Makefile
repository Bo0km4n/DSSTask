
RMI = Hello.java Task.java Server.java Client.java MyTask.java
SEL = Moo.java MooMain.java

rmi: $(RMI)
	javac $(RMI)

sel: $(SEL)
	javac $(SEL)

runsel: sel
	java MooMain

runsrv: rmi
	java Server 

runcnt: rmi
	java Client


clean:
	rm -rf *.class *~ *.hex

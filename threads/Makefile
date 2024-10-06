


AA= Accumulate
AS= AccumulateSync

all:
	@echo "critical race"
	javac ${AA}.java 
	java ${AA}
	@echo ""
	@echo "synchronized" 
	javac ${AS}.java
	java ${AS}
	make clean

clean:
	-rm *.class

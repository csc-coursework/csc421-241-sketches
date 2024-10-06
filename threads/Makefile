


AA= Accumulate
AS= AccumulateSync

java:
	@echo "critical race"
	javac ${AA}.java 
	java ${AA}
	@echo ""
	@echo "synchronized" 
	javac ${AS}.java
	java ${AS}
	make clean

go:
	go run go-accumulate.go
	go run go-native-accumulate.go

c:
	cc -o accumulate accumulate.c -lpthread
	./accumulate

install-go:
	sudo apt  install golang-go

clean:
	-rm *.class accumulate



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

java-capture:
	@echo building out file ...
	make java > accumulate.out

go:
	@echo "go-native-accumulate"
	go run go-native-accumulate.go
	@echo "go-accumulate"
	go run go-accumulate.go
	@echo "go-accumulate-mutex"
	go run go-accumulate-mutex.go

c:
	cc -o accumulate accumulate.c -lpthread
	./accumulate

hello:
	go run hello.go

rendez-vous:
	go run rendez-vous.go

install-go:
	sudo apt install golang-go
	# brew update ; brew install go 

clean:
	-rm *.class accumulate accumulate.out

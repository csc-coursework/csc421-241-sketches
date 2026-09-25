

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
	@echo "go-unsynced-accumulate"
	go run go-unsynced-accumulate.go
	@echo
	@echo "go-csp-accumulate"
	go run go-csp-accumulate.go
	@echo
	@echo "go-mutex-accumulate"
	go run go-mutex-accumulate.go

c:
	@echo "accumulate in c: no mutex"
	cc -o accumulate accumulate.c -lpthread
	./accumulate
	@echo
	@echo "accumulate in c: with mutex"
	cc -DUSE_MUTEX -o accumulate accumulate.c -lpthread
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

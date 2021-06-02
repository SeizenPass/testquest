gen:
	go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1
mem:
	go tool pprof testquest.test.exe mem.out
cpu:
	go tool pprof testquest.test.exe cpu.out
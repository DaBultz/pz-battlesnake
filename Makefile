# This assumes that the current OS is always Linux
# TODO: Compile based on the current OS
compile:
	make clean
	make compile_linux

compile_linux:
	cd ./battlesnake; GOOS=linux go build -o ../bin/battlesnake.so -buildmode=c-shared .

compile_windows:
	cd ./battlesnake; GOOS=windows go build -o ../bin/battlesnake.dll -buildmode=c-shared .


clean:
	rm -rf ./bin/**

run:
	make compile
	@echo ""
	poetry run python ./pz_battlesnake/main.py

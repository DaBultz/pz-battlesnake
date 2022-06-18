compile:
	make clean
	cd ./battlesnake && go build -o ../bin/battlesnake.so -buildmode=c-shared .

clean:
	rm -rf ./battlesnake.h
	rm -rf ./battlesnake.so


run:
	make compile
	@echo ""
	poetry run python ./pz_battlesnake/main.py

example:
	make compile
	@echo ""
	poetry run python ./example/main.py

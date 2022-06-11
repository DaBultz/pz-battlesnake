compile:
	make clean
	cd ./battlesnake && go build -o ../battlesnake.so -buildmode=c-shared .

clean:
	rm -f battlesnake.so
	rm -f battlesnake.h

run:
	make compile
	@echo ""
	poetry run python ./pz-battlesnake/main.py

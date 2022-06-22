compile:
	make clean
	cd ./battlesnake; go build -o ../bin/battlesnake -buildmode=c-shared .

clean:
	rm -rf ./bin/**

run:
	make compile
	@echo ""
	poetry run python ./pz_battlesnake/main.py

build-docs:
	@sphinx-build -b dirhtml -v docs/source build

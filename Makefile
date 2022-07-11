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
	@poetry run sphinx-build -b dirhtml -v docs/source build

serve-docs:
	@rm -rf build/
	@poetry run sphinx-autobuild docs/source build/ -b dirhtml

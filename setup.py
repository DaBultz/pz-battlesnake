#!/usr/bin/env python
from distutils.errors import CompileError
from distutils.core import setup

from subprocess import call
from setuptools import Extension, find_packages
from setuptools.command.build_ext import build_ext

import os


class build_battlesnake(build_ext):
    def build_extension(self, ext):
        # Get current setup.py location
        here = os.path.abspath(os.path.dirname(__file__))
        # CD into battlesnake
        os.chdir(f"{here}/battlesnake")

        # Init go modules
        cmd = ["go", "get"]
        success = call(cmd)
        if success != 0:
            raise CompileError("Failed to get dependencies")

        # Build go modules
        ext_path = f"../{self.get_ext_fullpath(ext.name)}/../bin/battlesnake"
        cmd = ["go", "build", "-buildmode=c-shared", "-o", ext_path]
        cmd += ext.sources

        print(cmd)

        success = call(cmd)
        if success != 0:
            raise CompileError("Failed to build extension")

        # CD back to setup.py
        os.chdir(here)


setup(
    name="pz_battlesnake",
    version="0.1.0",
    description="PettingZoo Environment for BattleSnake",
    author="DaBultz",
    url="https://github.com/DaBultz/pz-battlesnake",
    py_modules=["pz_battlesnake"],
    packages=["pz_battlesnake", "pz_battlesnake.env", "pz_battlesnake.spaces"],
    python_requires=">=3.7, <=3.11",
    install_requires=["PettingZoo>=1.18"],
    ext_modules=[
        Extension(
            "_battlesnake",
            [
                "main.go",
                "battlesnake.go",
                "json.go",
                "types.go",
            ],
        )
    ],
    cmdclass={"build_ext": build_battlesnake},
    include_package_data=True,
    zip_safe=False,
)

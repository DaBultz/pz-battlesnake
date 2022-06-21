# PZ-BattleSnake

PettingZoo/Gym Multi-Agent Environment For [BattleSnake](https://play.battlesnake.com/)

Documentation can be found here: [dabultz.github.io/pz-battlesnake/](https://dabultz.github.io/pz-battlesnake/)

You are able to contribute to this project, read the [contributing guide](CONTRIBUTING.md) for more info

## Environments

- [X] Solo
- [X] Standard
- [X] Duels
- [ ] Wrapped
- [ ] Constrictor
- [X] Arcade Maze (Used in Summer 2022 League)

## TO-DO:

- [ ] Compare [CTypes](https://docs.python.org/3/library/ctypes.html) to [CFFI](https://cffi.readthedocs.io/en/latest/) and [gRPC](https://grpc.io/)
- [ ] Move Rendering Logic to Python (maybe)
- [ ] Add Support For MacOS (Unable to test)
- [ ] Make custom BaseEnv
- [ ] Improve setup.py
- [ ] Improve Release Workflow ([googleapis/release-please](https://github.com/googleapis/release-please))


## Project Organization

```
├── battlesnake         <-- Go Library Code (compiled into a C library)
├── docs                <-- Documentation
├── example             <-- Example of how to use the enviorment
└── pz_battlesnake      <-- PettingZoo enviorment
```


## Useful links:

- [PettingZoo](https://github.com/Farama-Foundation/PettingZoo)
- [BattleSnake Rules](https://github.com/BattlesnakeOfficial/rules/)
- [Textualize](https://www.textualize.io/)

## Credits

- [Battlesnake Inc.](https://play.battlesnake.com/) for making the game and [open sourcing the logic & rules](https://github.com/BattlesnakeOfficial/rules)

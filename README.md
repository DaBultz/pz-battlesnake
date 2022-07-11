# PZ-BattleSnake

PettingZoo/Gym Multi-Agent Environment For [BattleSnake](https://play.battlesnake.com/)

Documentation can be found here: [dabultz.github.io/pz-battlesnake/](https://dabultz.github.io/pz-battlesnake/)

You are able to contribute to this project, read the [contributing guide](CONTRIBUTING.md) for more info

## Environments

- [X] Solo
- [X] Standard
- [X] Duels
- [X] Wrapped
- [X] Arcade Maze (Used in Summer 2022 League)
- [X] Easy to add new environments

## Plans:

To see the future plans/features, please see the following issue: [Future Plans (#12)](https://github.com/DaBultz/pz-battlesnake/issues/12)

### Next Stage

In the next stage, the are 2 focuses for the project

**Focus 1: Environments**

A base environment will be developed and all environments will inherit from this, there are minor changes between the environments. This would also included updates to the examples, so they are more accurate and up to date.

**Focus 2: Documentation**

The documentation is really lacking, in this stage there will be a focus on the following
- [ ] Documentation for the environments
- [ ] General API Documentation

## Project Organization

```
├── battlesnake         <-- Go Library Code (compiled into a C library)
├── docs                <-- Documentation
└── pz_battlesnake      <-- PettingZoo enviorment
```


## Useful links:

- [PettingZoo](https://github.com/Farama-Foundation/PettingZoo)
- [BattleSnake Rules](https://github.com/BattlesnakeOfficial/rules/)

## Credits

- [Battlesnake Inc.](https://play.battlesnake.com/) for making the game and [open sourcing the logic & rules](https://github.com/BattlesnakeOfficial/rules)

# PZ-BattleSnake
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FDaBultz%2Fpz-battlesnake.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FDaBultz%2Fpz-battlesnake?ref=badge_shield)


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

## Plans:

The project will be expanded in stages, 

### Next Stage

In the next stage, the are 2 focuses for the project

**Focus 1: Environments**

A base environment will be developed and all environments will inherit from this, there are minor changes between the environments. This would also included updates to the examples, so they are more accurate and up to date.

**Focus 2: Documentation**

The documentation is really lacking, in this stage there will be a focus on the following
- [ ] Documentation for the environments
- [ ] General API Documentation

### Future Stages

- Vectorization of the environments (allow multiple games run at the same time)
- Tests, it would allow us to keeo a certain quality

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

## Credits

- [Battlesnake Inc.](https://play.battlesnake.com/) for making the game and [open sourcing the logic & rules](https://github.com/BattlesnakeOfficial/rules)


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FDaBultz%2Fpz-battlesnake.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FDaBultz%2Fpz-battlesnake?ref=badge_large)
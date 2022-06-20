---
title: Maze (Summer League 2022)
---

# Maze (Summer League 2022)

## Enviorment Creation

```python
from pz_battlesnake.env import maze_v0

env = maze_v0.env()

env.reset()
for agent in env.agent_iter():
    observation, reward, done, info = env.last()
    action = env.action_space(agent).sample()
    env.step(action)
```

## Parameters

| Parameter  | Type | Description      | Default                                      | Note                               |
| ---------- | ---- | ---------------- | -------------------------------------------- | ---------------------------------- |
| num_agents | int  | number of agents | 4                                            |                                    |
| colors     | list | list of colors   | ['#00FF00', '#0000FF', '#FF00FF', '#FFFF00'] | Colors from this list will be used |

## Observation Space

The observation space matches the API request from battlesnake, you can find this in their official [API Docs](https://docs.battlesnake.com/references/api#post-move).


## Action Space

The Action space consists of a Move, which has 4 diffrent strings as options:
- `up`
- `down`
- `left`
- `right`

## Rewards

| Win | Draw | Loss |
| --- | ---- | ---- |
| 1   | 0    | -1   |

## Verion History

v0 - Initial Release

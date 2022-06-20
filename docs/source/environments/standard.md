---
title: Standard/Duel
---

# Standard/Duel

## Enviorment Creation

**Standard Game of 4 Players:**
```python
from pz_battlesnake.env import standard_v0

env = standard_v0.env()

env.reset()
for agent in env.agent_iter():
    observation, reward, done, info = env.last()
    action = policy(observation, agent)
    env.step(action)

```

**Duel Game of 2 Players:**
```python
from pz_battlesnake.env import standard_v0

env = standard_v0.env(num_agents=2)

env.reset()
for agent in env.agent_iter():
    observation, reward, done, info = env.last()
    action = env.action_space(agent).sample()
    env.step(action)
```

## Parameters

| Parameter  | Type | Description         | Default                                      | Note                               |
| ---------- | ---- | ------------------- | -------------------------------------------- | ---------------------------------- |
| width      | int  | width of the board  | 11                                           |                                    |
| height     | int  | height of the board | 11                                           |                                    |
| num_agents | int  | number of agents    | 4                                            |                                    |
| colors     | list | list of colors      | ['#00FF00', '#0000FF', '#FF00FF', '#FFFF00'] | Colors from this list will be used |

## Observation Space

The observation space matches the API request from battlesnake, you can find this in their official [API Docs](https://docs.battlesnake.com/references/api#post-move).


## Action Space

The Action space consists of a Move, which has 4 diffrent options:
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

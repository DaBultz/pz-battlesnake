---
title: Lava Bridges
---

# Lava Bridges

## Enviorment Creation

```{note}
This will create an enviorment where all agents take a random action.
```

```python
from pz_battlesnake.env import lava_bridges_v0

env = lava_bridges_v0.env()

env.reset()
for agent in env.agent_iter():
    observation, reward, done, info = env.last()
    action = env.action_space(agent).sample() if not done else None
    env.step(action)
```


## Parameters

| Parameter | Type      | Description        | Default                                      | Note                               |
| --------- | --------- | ------------------ | -------------------------------------------- | ---------------------------------- |
| num_agent | int       | width of the board | 4                                            |                                    |
| colors    | List[str] | list of colors     | ['#00FF00', '#0000FF', '#FF00FF', '#FFFF00'] | Colors from this list will be used |

## Observation Space

The observation space matches the API request from battlesnake, you can find this in their official [API Docs](https://docs.battlesnake.com/references/api#post-move).

## Action Space

The Action space consists of a Move, which has 4 diffrent strings as options:
- `up`
- `down`
- `left`
- `right`

See more: [Action Spaces](../api/spaces.md)

## Rewards

| Win | Draw | Loss |
| --- | ---- | ---- |
| 1   | 0    | -1   |

## Verion History

v0 - Initial Release

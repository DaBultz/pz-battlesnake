---
title: Getting Started
---

# Getting Started

## Prerequisites

- Installed pz-battlesnake (see [Installation](./install.md))

## Initializating an Enviorment

We can create a new enviorment by importing one of the available enviorments, for simplicity we will use the [solo_v0](../environments/solo.md) enviorment.

```python
# First we need to import the enviorment
from pz_battlesnake.env import solo_v0

# Then we can create a new enviorment
env = solo_v0.env() 
```

the solo_v0 enviorment has parameters which can be passed to the enviorment, for example:

```python
# Create a 15x15 solo enviorment 
env = solo_v0.env(width=15, height=15)
```

There are more parameters available, see the [solo_v0](../environments/solo.md) enviorment for more information.

## Interacting with the Enviorment

We now have an enviorment, we can interact with. We need to reset the enviorment before we can interact with it.

```python
env.reset()
```

Once the enviorment is reset, we can iterate over the agents in the enviorment by looping over all `env.agents`

```python
# Iterate over the agents in the enviorment
for agent in env.agents:
    # Get Last Observation, Reward, Done, Info
    observation, reward, done, info = env.last()
    # Pick a random action, if agent is done pick None
    action = env.action_space(agent).sample() if not done else None
    # Step through the enviorment
    env.step(action)
```

```{danger}
if an agent is done, it needs to be `None`, otherwise the enviorment will throw an error.
```

## Running Multiple Games 

The current code only runs through 1 game, to run through more than 1 game, we need to reset the enviorment after each game.

```python
# Run for 10 games
for _ in range(10):
    # Reset enviorment, before each game
    env.reset() 
    done = False
    while not done:
        for agent in env.agents:
            # Get Last Observation, Reward, Done, Info
            observation, reward, done, info = env.last()
            # Pick a random action, if agent is done pick None
            action = env.action_space(agent).sample() if not done else None
            # Step through the enviorment
            env.step(action)
        # This is a shortcut to set the done to be true, 
        # since when all agents are done the env.agents array will be empty
        done = not env.agents
```
## Useful links

- [PettingZoo API Docs](https://www.pettingzoo.ml/api#interacting-with-environments)
```{warning}
Some of the things listed on PettingZoo's Guide are not available in pz-battlesnake, as we implement an AEC Enviorment and not an MARL Enviorment. Below is listed what you can't do.
    - PettingZoo Parrallel API
    - Raw Enviorments 
```
## Complete Code

```{note}
The code below implements a random agent, which will take a random action each step.
```

```python
# First we need to import the enviorment
from pz_battlesnake.env import solo_v0

# Then we can create a new enviorment
env = solo_v0.env() 

# Create a 15x15 solo enviorment 
# env = solo_v0.env(width=15, height=15) # uncomment this to create a 15x15 solo enviorment

# Run for 10 games
for _ in range(10):
    # Reset enviorment, before each game
    env.reset()
    done = False
    while not done:
        for agent in env.agents:
            # Get Last Observation, Reward, Done, Info
            observation, reward, done, info = env.last()
            # Pick an action, if agenmt is done, pick None
            action = env.action_space(agent).sample() if not done else None
            # Step through the enviorment
            env.step(action)
        # Code below runs, when all agents has taken an action
        # Render the enviorment
        # env.render() # uncomment this to render
        # This is a shortcut to set the done to be true, 
        # since when all agents are done the env.agents array will be empty
        done = not env.agents
```

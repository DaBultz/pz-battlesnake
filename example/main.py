# from pz_battlesnake import solo
from pz_battlesnake.env import solo_v0

import random

from pz_battlesnake.wrapper import get_request

env = solo_v0.env()

for _ in range(2):
    env.reset()
    done = False
    while not done:
        for agent in env.agent_iter():
            obs = get_request()

            action = env.action_space(agent).sample()
            env.step(action)
            obs, reward, done, info = env.last()

            if done:
                print("done")
                break

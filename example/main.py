# from pz_battlesnake import solo
from pz_battlesnake.env import solo_v0

import random

env = solo_v0.env()

for _ in range(2):
    env.reset()
    done = False
    while not done:
        for agent in env.agent_iter():
            action = env.action_space(agent).sample()
            env.step(action)
            obs, reward, done, info = env.last()
            print(obs["you"]["body"])

            if done:
                print("\n")
                break

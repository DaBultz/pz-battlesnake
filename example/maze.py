from os import system
from pz_battlesnake.env import maze_v0
from pz_battlesnake.wrapper import env_done

env = maze_v0.env()


for _ in range(1):
    env.reset()
    done = False
    while not done:
        env.render(mode="human_color")

        for agent in env.agent_iter():
            obs, rewards, _, info = env.last()

            action = env.action_space(agent).sample()
            env.step(action)

            if env_done():
                done = True
                break

from pz_battlesnake.env import solo_v0
from pz_battlesnake.wrapper import env_done, env_render

env = solo_v0.env()


for _ in range(100):
    env.reset()
    done = False
    while not done:
        env.render()

        for agent in env.agent_iter():
            obs, rewards, _, info = env.last()

            action = env.action_space(agent).sample()
            env.step(action)

            if env_done():
                done = True
                break

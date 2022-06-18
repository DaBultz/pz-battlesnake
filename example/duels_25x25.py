from os import system
from pz_battlesnake.env import standard_v0
from pz_battlesnake.wrapper import env_done

# Set the map to be 25x25
env = standard_v0.env(width=25, height=25, num_agents=2)

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

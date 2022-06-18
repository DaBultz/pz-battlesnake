# from pz_battlesnake import solo
from timeit import default_timer as timer
import timeit
from pz_battlesnake.env import solo_v0
from pz_battlesnake.wrapper import env_done

env = solo_v0.env()

start = timer()

for _ in range(1000):
    env.reset()
    done = False
    while not done:
        for agent in env.agent_iter():
            obs, rewards, _, info = env.last()

            action = env.action_space(agent).sample()
            env.step(action)
            env.render()

            if env_done():
                done = True
                break
            # obs, reward, done, info = env.last()

end = timer()
print(f"Took: {end - start} seconds")

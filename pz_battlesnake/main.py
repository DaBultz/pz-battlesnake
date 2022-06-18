import ctypes
import json

from click import option

battlesnake = ctypes.CDLL("./battlesnake.so")

# Setup
env_setup = battlesnake.setup
env_setup.argtypes = [ctypes.c_char_p]

# Step
_env_step = battlesnake.step
_env_step.argtypes = [ctypes.c_char_p]
_env_step.restype = ctypes.c_char_p

options = {
    "width": 17,
    "height": 17,
    "names": ["agent_0", "agent_1", "agent_2", "agent_3"],
    "map": "standard",
    "game_type": "standard",
}

env_setup(json.dumps(options).encode("utf-8"))

actions = {
    "agent_0": "down",
    "agent_1": "up",
    "agent_2": "down",
    "agent_3": "up",
}


def env_step(actions):
    # Convert actions to string
    actions = json.dumps(actions).encode("utf-8")

    # Call step in go
    res = _env_step(actions)

    # convert result to python object
    res = json.loads(res.decode("utf-8"))

    # return result
    return res


for i in range(10):
    res = env_step(actions)
    turn = res["agent_0"]["observation"]["turn"]

    print(f"=============== Turn {turn} ================")
    for name in options["names"]:
        snake = res[name]["observation"]
        name = snake["you"]["name"]
        body = snake["you"]["body"][0]
        print(f"{name}: {body}")

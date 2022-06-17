import ctypes
import json

battlesnake = ctypes.CDLL("./battlesnake.so")

# Setup
env_setup = battlesnake.setup

# Reset
env_reset = battlesnake.reset

# step
_step = battlesnake.step
_step.argtypes = [ctypes.c_char_p]
_step.restype = ctypes.c_char_p


def env_step(actions):
    # Convert actions to string
    actions = json.dumps(actions).encode("utf-8")

    # Call step in go
    res = _step(actions)

    # convert result to python object
    res = json.loads(res.decode("utf-8"))

    # return result
    return res["observation"], res["reward"], res["done"], res["info"]

import ctypes
import json

battlesnake = ctypes.CDLL("./battlesnake.so")

# Setup
_setup = battlesnake.setup
_setup.argtypes = [ctypes.c_char_p]

# step
_step = battlesnake.step
_step.argtypes = [ctypes.c_char_p]
_step.restype = ctypes.c_char_p

# isGameOver
_done = battlesnake.isGameOver
_done.restype = ctypes.c_int


def env_done():
    return _done() == 1


def env_setup(options):
    # Convert options to string
    options = json.dumps(options).encode("utf-8")

    # Call setup in go
    _setup(options)


def env_step(actions):
    # Convert actions to string
    actions = json.dumps(actions).encode("utf-8")

    # Call step in go
    res = _step(actions)

    # convert result to python object
    res = json.loads(res.decode("utf-8"))

    # return result
    return res["observation"], res["reward"], res["done"], res["info"]

import ctypes
import json

battlesnake = ctypes.CDLL("./battlesnake.so")

# fromJSON
from_json = battlesnake.fromJSON
from_json.argtypes = [ctypes.c_char_p]

doc = {
    "width": 20,
    "height": 20,
}

from_json(json.dumps(doc).encode("utf-8"))

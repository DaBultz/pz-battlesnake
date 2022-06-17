from typing import List
from gym.spaces import Dict, MultiDiscrete
import numpy as np

ob = Dict(
    {
        "body": MultiDiscrete(np.repeat([[10, 10]], 10, axis=0)),
        "head": MultiDiscrete([6, 6]),
    }
)

print(ob.sample())

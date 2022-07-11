from typing import List
from pettingzoo.utils import parallel_to_aec, OrderEnforcingWrapper

# Local Imports
from pz_battlesnake.constants import DEFAULT_COLORS
from pz_battlesnake.env.base_env import BaseEnv


def env(
    width: int = 19,
    height: int = 19,
    colors: List[str] = DEFAULT_COLORS,
):
    env = BaseEnv(
        wdith=width,
        height=height,
        game_map="standard",
        game_type="wrapped",
        num_agents=8,
        colors=colors,
    )

    # Set the metadata enviorment name
    env.metadata["name"] = "battlesnake-duels_v0"

    # Convert from MARL to AEC API
    env = parallel_to_aec(env)
    # Provides a wide vareity of helpful error checks
    env = OrderEnforcingWrapper(env)

    return env

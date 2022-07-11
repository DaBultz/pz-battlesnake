from typing import List
from pettingzoo.utils import parallel_to_aec, OrderEnforcingWrapper

# Local Imports
from pz_battlesnake.constants import DEFAULT_COLORS
from pz_battlesnake.env.base_env import BaseEnv


def env(
    num_agent: int = 4,
    colors: List[str] = DEFAULT_COLORS,
):
    env = BaseEnv(
        wdith=21,
        height=19,
        game_map="arcade_maze",
        game_type="wrapped",
        num_agents=num_agent,
        colors=colors,
    )

    # Set the metadata enviorment name
    env.metadata["name"] = "battlesnake-maze_v0"

    # Convert from MARL to AEC API
    env = parallel_to_aec(env)

    return env

import functools
from typing import List
from pettingzoo import ParallelEnv
from pettingzoo.utils import parallel_to_aec, OrderEnforcingWrapper
from gym import spaces

# Local Imports
from pz_battlesnake.constants import DEFAULT_COLORS
from pz_battlesnake.spaces.move import Move
from pz_battlesnake.types.battlesnake_options import BattlesnakeOptions
from pz_battlesnake.wrapper import env_done, env_render, env_reset, env_setup, env_step


def make_env(**kwargs):
    env = BaseEnv(**kwargs)

    # Set the metadata enviorment name
    env.metadata[
        "name"
    ] = f"battlesnake-{env._options.game_map}-{env._options.game_type}_v0"

    # Convert from MARL to AEC API
    env = parallel_to_aec(env)
    # Provides a wide vareity of helpful error checks
    env = OrderEnforcingWrapper(env)

    return env


class BaseEnv(ParallelEnv):
    """
    Implements a BaseEnv for the Battlesnake environment, in which all environments based on
    """

    metadata = {
        "render_modes": ["ascii", "color"],
    }

    def __init__(
        self,
        width: int = 11,
        height: int = 11,
        num_agents: int = 4,
        colors: List[str] = DEFAULT_COLORS,
        game_map: str = "standard",
        game_type: str = "standard",
    ):
        self.possible_agents = ["agent_" + str(i) for i in range(num_agents)]
        self.agent_name_mapping = dict(
            zip(self.possible_agents, list(range(len(self.possible_agents))))
        )

        self.agent_selection = self.possible_agents[0]

        self._options = BattlesnakeOptions(
            width=width,
            height=height,
            colors=colors,
            game_map=game_map,
            game_type=game_type,
            names=self.possible_agents,
        )

        self.options = {
            "width": width,
            "height": height,
            "map": "standard",
            "game_type": "standard",
            "names": self.possible_agents,
            "colors": colors,
        }

    @functools.lru_cache(maxsize=0)
    def observation_space(self, agent=None):
        # Check if agent is provided
        assert agent, "Agent must be provided"

        # Check if agent is valid
        assert agent in self.possible_agents, "agent must be one of {}".format(
            self.possible_agents
        )

        return spaces.Dict()

    @functools.lru_cache(maxsize=0)
    def action_space(self, agent=None):
        # Check if agent is provided
        assert agent, "Agent must be provided"

        # Check if agent is valid
        assert agent in self.possible_agents, "agent must be one of {}".format(
            self.possible_agents
        )

        # assert False, "observation_space() is not implemented yet"
        return Move()

    def render(self, mode="ascii"):
        """
        Renders the environment. In human mode, it can print to terminal, open
        up a graphical window, or open up some other display that a human can see and understand.

        Args:
            mode: The mode to render the environment in. Can be "ascii" or "color"
        """
        if mode == "ascii" or mode == "color":
            env_render(True if mode == "color" else False)
        else:
            assert False, "Valid render modes are 'ascii' and 'color'"

    def reset(self, seed=None):
        """
        Reset needs to initialize the `agents` attribute and must set up the
        environment so that render(), and step() can be called without issues.

        Returns the observations for each agent
        """
        self.agents = self.possible_agents[:]

        if seed:
            self.options["seed"] = seed

        return env_reset(self._options.options)

    def step(self, action):
        """
        step(action) takes in an action for each agent and should return the
            - observations
            - rewards
            - dones
            - infos
        dicts where each dict looks like {agent_0: item_1, agent_1: item_2}
        """
        if not action:
            self.agents = []
            return {}, {}, {}, {}

        agents = env_step(action)

        observations = {}
        rewards = {}
        dones = {}
        infos = {}

        for agent in agents:
            observations[agent] = agents[agent]["observation"]
            rewards[agent] = agents[agent]["reward"]
            dones[agent] = agents[agent]["done"]
            infos[agent] = agents[agent]["info"]

        if env_done():
            self.agents = []

        return observations, rewards, dones, infos

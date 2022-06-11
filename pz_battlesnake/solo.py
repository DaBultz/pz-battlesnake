import functools
from gym import spaces
from pettingzoo import ParallelEnv
from pettingzoo.utils import parallel_to_aec, wrappers


def env():
    env = parallel_env()
    # This wrapper is only for environments which print results to the terminal
    env = wrappers.CaptureStdoutWrapper(env)
    # this wrapper helps error handling for discrete action spaces
    env = wrappers.AssertOutOfBoundsWrapper(env)
    # Provides a wide vareity of helpful user errors
    # Strongly recommended
    env = wrappers.OrderEnforcingWrapper(env)
    return env


class parallel_env(ParallelEnv):
    metadata = {
        "render_modes": ["none", "human"],
        "name": "battlesnake-solo_v0",
    }

    def __init__(self):
        self.possible_agents = ["player_" + str(i) for i in range(1)]
        self.agent_name_mapping = dict(
            zip(self.possible_agents, list(range(len(self.possible_agents))))
        )

    # this cache ensures that same space object is returned for the same agent
    # allows action space seeding to work as expected
    @functools.lru_cache(maxsize=0)
    def observation_space(self, agent=None):
        # Gym spaces are defined and documented here: https://gym.openai.com/docs/#spaces

        # Check if agent is provided
        assert agent, "Agent must be provided"

        # Check if agent is valid
        assert agent in self.possible_agents, "agent must be one of {}".format(
            self.possible_agents
        )

        # Not implemented yet
        assert False, "observation_space() is not implemented yet"
        return spaces.Discrete(4)

    @functools.lru_cache(maxsize=0)
    def action_space(self, agent=None):
        # Gym spaces are defined and documented here: https://gym.openai.com/docs/#spaces
        # Check if agent is provided
        assert agent, "Agent must be provided"

        # Check if agent is valid
        assert agent in self.possible_agents, "agent must be one of {}".format(
            self.possible_agents
        )

        return spaces.Discrete(4)

    def render(self, mode="none"):
        """
        Renders the environment. In human mode, it can print to terminal, open
        up a graphical window, or open up some other display that a human can see and understand.
        """
        pass

    def close(self):
        """
        Close should release any graphical displays, subprocesses, network connections
        or any other environment data which should not be kept around after the
        user is no longer using the environment.
        """
        pass

    def reset(self, seed=None):
        """
        Reset needs to initialize the `agents` attribute and must set up the
        environment so that render(), and step() can be called without issues.

        Returns the observations for each agent
        """
        pass

    def step(self, actions):
        """
        step(action) takes in an action for each agent and should return the
            - observations
            - rewards
            - dones
            - infos
        dicts where each dict looks like {agent_1: item_1, agent_2: item_2}
        """
        pass

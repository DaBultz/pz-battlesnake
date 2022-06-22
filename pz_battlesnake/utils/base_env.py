from pettingzoo import ParallelEnv


class BaseEnv(ParallelEnv):
    """
    The BaseEnv implements all of the shared functionality of all environments using the parallel epi.
    this means that it steps all agents at once rather than one at a time.
    """

    def __init__(self, **kwargs):
        super().__init__(**kwargs)

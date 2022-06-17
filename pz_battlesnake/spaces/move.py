from typing import List

from gym.spaces.space import Space

import random


class Move(Space):
    """
    This class respresents the move action in BattleSnake

    Example:
     >>> Move() # ["up", "down", "left", "right"]
    """

    def __init__(self):
        self.moves = ["up", "down", "left", "right"]
        super().__init__()

    def sample(self):
        return random.choice(self.moves)

    def contains(self, x) -> bool:
        return x in self.moves

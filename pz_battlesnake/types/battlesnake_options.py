import json
from typing import Any, Dict, List

from pz_battlesnake.constants import DEFAULT_COLORS


class BattlesnakeOptions:
    """BattlesnakeOptions represents the options for a battlesnake game.

    Args:
        width (int, optional): width of the board. Defaults to 11.
        height (int, optional): height of the board. Defaults to 11.
        map (str, optional): name of the map. Defaults to "standard".
        game_type (str, optional): game type. Defaults to "solo".
        seed (int, optional): the seed to use. Defaults to None.
        names (List[str], optional): list over all names. Defaults to None.
        colors (List[str], optional): list of all colors. Defaults to DEFAULT_COLORS.
    """

    width: int
    height: int
    game_map: str
    game_type: str
    seed: int
    names: List[str]
    colors: List[str]

    def __init__(
        self,
        width: int = 11,
        height: int = 11,
        game_map: str = "standard",
        game_type: str = "solo",
        seed: int = None,
        names: List[str] = None,
        colors: List[str] = DEFAULT_COLORS,
    ):
        self.width = width
        self.height = height
        self.game_map = game_map
        self.game_type = game_type
        self.seed = seed
        self.names = names
        self.colors = colors

    @property
    def options(self) -> Dict[str, Any]:
        """Returns Battlesnake Options as a dictionary

        Todo:
            * Add example of return

        Returns:
            Dict[str, Any]: returns a dictionary of the options

        """
        return {
            "width": self.width,
            "height": self.height,
            "map": self.game_map,
            "game_type": self.game_type,
            "seed": self.seed,
            "names": self.names,
            "colors": self.colors,
        }

    def __repr__(self) -> str:
        return f"BattlesnakeOptions(width={self._width}, height={self._height}, map={self._map}, game_type={self._game_type}, seed={self._seed}, names={self._names}, colors={self._colors})"

    def __str__(self) -> str:
        return json.dumps(self.options, indent=2)

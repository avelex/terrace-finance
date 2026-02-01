# pragma version ^0.4.3
# @license MIT
from ethereum.ercs import IERC20

from snekmate.extensions import erc4626

initializes: erc4626
exports: erc4626.__interface__


@deploy
def __init__(asset: IERC20):
    erc4626.__init__(
        "Staked tUSD", "stUSD", asset, 0, "Terrace Finance", "0x01"
    )

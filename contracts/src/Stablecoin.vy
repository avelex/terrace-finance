# pragma version ^0.4.3
# @license MIT
from snekmate.auth import ownable as ow

initializes: ow

from snekmate.tokens import erc20

exports: erc20.__interface__
initializes: erc20[ownable := ow]


@deploy
def __init__():
    erc20.__init__("Terrace USD", "tUSD", 18, "Terrace Finance", "0x01")
    ow.__init__()

# pragma version ^0.4.3
"""
@notice Mock ERC20 for testing
"""

from snekmate.auth import ownable as ow

initializes: ow

from snekmate.tokens import erc20

initializes: erc20[ownable := ow]

exports: erc20.__interface__


@deploy
def __init__(
    _name: String[25],
    _symbol: String[5],
    _decimals: uint8,
    _initial_supply: uint256,
):
    erc20.__init__(_name, _symbol, _decimals, _name, "0x01")
    ow.__init__()
    erc20._mint(msg.sender, _initial_supply * 10**convert(_decimals, uint256))

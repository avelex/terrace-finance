# pragma version ^0.4.3
# @license MIT
from src.modules import roles
from snekmate.auth import access_control as access

initializes: access

from snekmate.auth import ownable as ow

initializes: ow

from src.modules import ledger

initializes: ledger

from ethereum.ercs import IERC20
from snekmate.tokens import erc20

initializes: erc20[ownable := ow]

from src.modules import withdrawal_queue

initializes: withdrawal_queue[access_control := access]

from src.modules import strategy_registry

initializes: strategy_registry[access_control := access, ledger := ledger]

from src.modules import cctp

initializes: cctp[access_control := access, ledger := ledger]

from src.modules import vault

initializes: vault[access_control := access, erc20 := erc20, ledger := ledger, withdrawal_queue := withdrawal_queue]

from src.modules import gateway

initializes: gateway[vault := vault, cctp := cctp]

exports: (
    erc20.__interface__,
    strategy_registry.__interface__,
    access.__interface__,
    withdrawal_queue.__interface__,
    ledger.__interface__,
    cctp.__interface__,
    vault.__interface__,
    gateway.__interface__,
    roles.__interface__,
)

@deploy
def __init__(
    usdc: address,
    withdrawal_delay: uint256,
    messanger: address,
    transmitter: address,
    hub_domain: uint32,
    gateway_minter: address,
):
    ow.__init__()
    access.__init__()
    ledger.__init__(usdc)
    erc20.__init__("Terrace USD", "tUSD", 18, "Terrace Finance", "0x01")
    withdrawal_queue.__init__(withdrawal_delay)
    strategy_registry.__init__()
    cctp.__init__(hub_domain, messanger, transmitter)
    vault.__init__()
    gateway.__init__(gateway_minter)
    # renouncing ownership
    ow._transfer_ownership(empty(address))

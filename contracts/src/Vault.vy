# pragma version ^0.4.3
# @license MIT
from src.modules import roles
from snekmate.auth import access_control as access

initializes: access

from src.modules import ledger

initializes: ledger

from src.modules import strategy_registry

initializes: strategy_registry[access_control := access, ledger := ledger]

from src.modules import cctp

initializes: cctp[access_control := access, ledger := ledger]

exports: (
    strategy_registry.__interface__,
    access.__interface__,
    ledger.__interface__,
    cctp.__interface__,
    roles.__interface__,
)


@deploy
def __init__(
    usdc: address,
    messanger: address,
    transmitter: address,
    hub_domain: uint32,
):
    access.__init__()
    ledger.__init__(usdc)
    strategy_registry.__init__()
    cctp.__init__(hub_domain, messanger, transmitter)

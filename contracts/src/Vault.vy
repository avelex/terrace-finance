# pragma version ^0.4.3
# @license MIT
from snekmate.auth import access_control as access

uses: access

from src.modules import strategy_registry

initializes: strategy_registry[access_control := access]

exports: (
    strategy_registry.strategy_root_hash,
    strategy_registry.update_root_hash,
    strategy_registry.is_strategy_allowed,
)

OPERATOR_ROLE: public(constant(bytes32)) = keccak256("OPERATOR_ROLE")


@deploy
def __init__():
    strategy_registry.__init__()


@external
def execute_strategy(
    target: address, selector: Bytes[4], data: Bytes[256], proof: DynArray[bytes32, 32]
):
    access._check_role(OPERATOR_ROLE, msg.sender)
    assert strategy_registry._is_allowed(target, selector, proof)
    calldata: Bytes[260] = concat(selector, data)
    raw_call(
        target,
        calldata,
        revert_on_failure=True,
    )

@view
@external
def encode_supply_aave_v3(asset: address, amount: uint256, onBehalfOf: address, referalCode: uint16) -> Bytes[128]:
    return abi_encode(asset, amount, onBehalfOf, referalCode)
    
@view
@external
def encode_erc20_approve(account: address, amount: uint256) -> Bytes[128]:
    return abi_encode(account, amount)

@external
def fill():
    pass


@external
def receive_fees():
    pass

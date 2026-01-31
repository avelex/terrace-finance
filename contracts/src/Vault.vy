# pragma version ^0.4.3
# @license MIT
from snekmate.auth import access_control as access

uses: access

from src import strategy_registry

initializes: strategy_registry[access_control := access]

exports: strategy_registry.__interface__

OPERATOR_ROLE: public(constant(bytes32)) = keccak256("OPERATOR_ROLE")


@deploy
def __init__():
    strategy_registry.__init__()


@external
def execute_strategy(
    target: address, data: Bytes[256], proof: DynArray[bytes32, 32]
):
    access._check_role(OPERATOR_ROLE, msg.sender)
    strategy_hash: bytes32 = keccak256(concat(convert(target, bytes32), data))
    assert strategy_registry._is_allowed(proof, strategy_hash)
    raw_call(
        target,
        data,
        revert_on_failure=True,
    )


@external
def fill():
    pass


@external
def receive_fees():
    pass

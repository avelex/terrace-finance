# pragma version ^0.4.3
# @license MIT
from ethereum.ercs import IERC20
from snekmate.auth import access_control

uses: access_control

from snekmate.utils import merkle_proof_verification

initializes: merkle_proof_verification

from src import ledger

uses: ledger

OPERATOR_ROLE: public(constant(bytes32)) = keccak256("OPERATOR_ROLE")
_MAX_BATCH_SIZE: constant(uint256) = 32

strategy_root_hash: public(bytes32)

# @dev Emitted when the root hash is updated
event StrategyRootHashUpdated:
    root: indexed(bytes32)


@deploy
def __init__():
    merkle_proof_verification.__init__()


@external
def update_root_hash(root: bytes32):
    access_control._check_role(access_control.DEFAULT_ADMIN_ROLE, msg.sender)
    self.strategy_root_hash = root
    log StrategyRootHashUpdated(root=root)


@external
def execute_strategy(
    target: address, selector: Bytes[4], data: Bytes[256], proof: DynArray[bytes32, 32]
):
    access_control._check_role(OPERATOR_ROLE, msg.sender)
    self._execute_strategy(target, selector, data, proof)


@external
def batch_execute(
    targets: DynArray[address, _MAX_BATCH_SIZE],
    selectors: DynArray[Bytes[4], _MAX_BATCH_SIZE],
    data: DynArray[Bytes[256], _MAX_BATCH_SIZE],
    proofs: DynArray[DynArray[bytes32, 32], _MAX_BATCH_SIZE]
):
    access_control._check_role(OPERATOR_ROLE, msg.sender)
    for i: uint256 in range(len(targets), bound=_MAX_BATCH_SIZE):
        self._execute_strategy(targets[i], selectors[i], data[i], proofs[i])


@internal
def _execute_strategy(_target: address, _selector: Bytes[4], _data: Bytes[256], _proof: DynArray[bytes32, 32]):
    assert self._is_allowed(_target, _selector, _proof)

    # TODO: ensure that we work with 6 decimals USDC, to make this convertion safety
    balance_before: int256 = convert(staticcall IERC20(ledger.USDC).balanceOf(self), int256)
    raw_call(
        _target,
        concat(_selector, _data),
        revert_on_failure=True,
    )
    balance_after: int256 = convert(staticcall IERC20(ledger.USDC).balanceOf(self), int256)

    diff: int256 = balance_after - balance_before
    if diff > 0:
        ledger._repay(_target, convert(diff, uint256))
    elif diff < 0:
        ledger._borrow(_target, convert(-diff, uint256))


@external
@view
def is_strategy_allowed(target: address, selector: Bytes[4], proof: DynArray[bytes32, merkle_proof_verification._DYNARRAY_BOUND]) -> bool:
    return self._is_allowed(target, selector, proof)

@internal
@view
def _is_allowed(target: address, selector: Bytes[4], proof: DynArray[bytes32, merkle_proof_verification._DYNARRAY_BOUND]) -> bool:
    encoded: Bytes[128] = abi_encode(target, convert(selector, bytes4))
    strategy_hash: bytes32 = keccak256(keccak256(encoded))
    return merkle_proof_verification._verify(
        proof, self.strategy_root_hash, strategy_hash
    )

@external
@pure
def encode_erc20_approve(target: address, amount: uint256) -> Bytes[256]:
    return abi_encode(target, amount)

@external
@pure
def encode_supply_aave_v3(asset: address, amount: uint256, onBehalfOf: address, referrer: uint16) -> Bytes[256]:
    return abi_encode(asset, amount, onBehalfOf, referrer)
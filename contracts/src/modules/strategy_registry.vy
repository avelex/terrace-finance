# pragma version ^0.4.3
# @license MIT
from snekmate.auth import access_control

uses: access_control

from snekmate.utils import merkle_proof_verification

initializes: merkle_proof_verification

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
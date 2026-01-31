# pragma version ^0.4.3
# @license MIT
from snekmate.utils import merkle_proof_verification

initializes: merkle_proof_verification

rootHash: public(bytes32)

# @dev Emitted when the root hash is updated
event RootHashUpdated:
    root: indexed(bytes32)


@deploy
def __init__():
    pass


@external
def update_root_hash(root: bytes32):
    self.rootHash = root
    log RootHashUpdated(root=root)


@external
@view
def is_allowed() -> bool:
    return True

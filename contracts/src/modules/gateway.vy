# pragma version ^0.4.3
# @license MIT
from ethereum.ercs import IERC4626

from src.modules import vault

uses: vault

from src.modules import cctp

uses: cctp


interface IGatewayMinter:
    def gatewayMint(attestation: Bytes[1024], signature: Bytes[65]): nonpayable


# @dev Byte encoding (big-endian):
#     FIELD                      OFFSET   BYTES   NOTES
#     magic                           0       4   Always 0xff6fb334
#     max block height                4      32
#     transfer spec length           36       4   In bytes, may vary based on hook data length
#     encoded transfer spec          40       ?   Must be the length indicated above
struct Attestation:
    magic: Bytes[4]
    max_block_height: uint256
    transfer_spec_length: uint256
    encoded_transfer_spec: Bytes[1024]


# @dev Byte encoding (big-endian):
#     FIELD                   OFFSET   BYTES   NOTES
#     magic                        0       4   Always 0xca85def7
#     version                      4       4   Always 1
#     source domain                8       4
#     destination domain          12       4
#     source contract             16      32
#     destination contract        48      32
#     source token                80      32
#     destination token          112      32
#     source depositor           144      32
#     destination recipient      176      32
#     source signer              208      32
#     destination caller         240      32   May be 0, to allow any caller
#     value                      272      32
#     salt                       304      32   Any random value that makes the transfer spec hash unique
#     hook data length           336       4   In bytes, 0 to indicate no hook data
#     hook data                  340       ?   Must be the length indicated above if present
struct TransferSpec:
    version: uint32
    source_domain: uint32
    destination_domain: uint32
    source_contract: address
    destination_contract: address
    source_token: address
    destination_token: address
    source_depositor: address
    destination_recipient: address
    source_signer: address
    destination_caller: address
    value: uint256
    salt: bytes32
    hook_data_length: uint32
    # hook_data: Bytes[1024]


GATEWAY_MINTER: public(immutable(address))

_ATTESTATION_MAGIC: constant(Bytes[4]) = x"ff6fb334"
_TRANSFER_SPEC_MAGIC: constant(Bytes[4]) = x"ca85def7"


@deploy
def __init__(gateway_minter: address):
    assert gateway_minter != empty(address)
    GATEWAY_MINTER = gateway_minter


@external
def gateway_deposit_and_stake(attestation: Bytes[1024], signature: Bytes[65]):
    transfer: TransferSpec = self._decode_transfer_spec(attestation)
    extcall IGatewayMinter(GATEWAY_MINTER).gatewayMint(attestation, signature)

    if cctp._is_hub():
        deposited: uint256 = vault._deposit(transfer.value)
        extcall IERC4626(vault.staking).deposit(
            deposited, transfer.destination_recipient
        )


def _decode_transfer_spec(attestation: Bytes[1024]) -> TransferSpec:
    attestation_magic: Bytes[4] = slice(attestation, 0, 4)
    assert attestation_magic == _ATTESTATION_MAGIC

    transfer_spec_length: uint256 = convert(slice(attestation, 36, 4), uint256)
    encoded_transfer_spec: Bytes[1024] = slice(
        attestation, 40, transfer_spec_length
    )

    transfer_spec_magic: Bytes[4] = slice(encoded_transfer_spec, 0, 4)
    assert transfer_spec_magic == _TRANSFER_SPEC_MAGIC

    transfer_spec_version: uint32 = convert(
        slice(encoded_transfer_spec, 4, 4), uint32
    )
    assert transfer_spec_version == 1

    source_domain: uint32 = convert(slice(encoded_transfer_spec, 8, 4), uint32)
    destination_domain: uint32 = convert(
        slice(encoded_transfer_spec, 12, 4), uint32
    )
    source_contract: address = abi_decode(
        slice(encoded_transfer_spec, 16, 32), address
    )
    destination_contract: address = abi_decode(
        slice(encoded_transfer_spec, 48, 32), address
    )
    source_token: address = abi_decode(
        slice(encoded_transfer_spec, 80, 32), address
    )
    destination_token: address = abi_decode(
        slice(encoded_transfer_spec, 112, 32), address
    )
    source_depositor: address = abi_decode(
        slice(encoded_transfer_spec, 144, 32), address
    )
    destination_recipient: address = abi_decode(
        slice(encoded_transfer_spec, 176, 32), address
    )
    source_signer: address = abi_decode(
        slice(encoded_transfer_spec, 208, 32), address
    )
    destination_caller: address = abi_decode(
        slice(encoded_transfer_spec, 240, 32), address
    )
    value: uint256 = convert(slice(encoded_transfer_spec, 272, 32), uint256)
    salt: bytes32 = convert(slice(encoded_transfer_spec, 304, 32), bytes32)
    hook_data_length: uint32 = convert(
        slice(encoded_transfer_spec, 336, 4), uint32
    )

    return TransferSpec(
        version=transfer_spec_version,
        source_domain=source_domain,
        destination_domain=destination_domain,
        source_contract=source_contract,
        destination_contract=destination_contract,
        source_token=source_token,
        destination_token=destination_token,
        source_depositor=source_depositor,
        destination_recipient=destination_recipient,
        source_signer=source_signer,
        destination_caller=destination_caller,
        value=value,
        salt=salt,
        hook_data_length=hook_data_length,
    )

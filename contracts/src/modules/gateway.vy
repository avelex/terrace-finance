# pragma version ^0.4.3
# @license MIT
from ethereum.ercs import IERC4626
from snekmate.tokens import erc20

uses: erc20

from src.modules import vault

uses: vault

from src.modules import cctp

uses: cctp


interface IGatewayMinter:
    def gatewayMint(attestation: Bytes[1528], signature: Bytes[65]): nonpayable


# @dev AttestationSet struct
# Byte encoding (big-endian):
#     FIELD                      OFFSET   BYTES   NOTES
#     magic                           0       4   Always 0x1e12db71
#     number of attestations          4       4
#     attestations                    8       ?   Concatenated one after another


# @dev Byte encoding (big-endian):
#     FIELD                      OFFSET   BYTES   NOTES
#     magic                           0       4   Always 0xff6fb334
#     max block height                4      32
#     transfer spec length           36       4   In bytes, may vary based on hook data length
#     encoded transfer spec          40       ?   Must be the length indicated above
struct Attestation:
    magic: Bytes[4]
    max_block_height: uint256
    transfer_spec_length: uint32
    encoded_transfer_spec: Bytes[340]


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


GATEWAY_MINTER: public(immutable(address))

# bytes4(keccak256("circle.gateway.AttestationSet"))
_ATTESTATION_SET_MAGIC: constant(Bytes[4]) = x"1e12db71"
# bytes4(keccak256("circle.gateway.Attestation"))
_ATTESTATION_MAGIC: constant(Bytes[4]) = x"ff6fb334"
# bytes4(keccak256("circle.gateway.TransferSpec"))
_TRANSFER_SPEC_MAGIC: constant(Bytes[4]) = x"ca85def7"
# max number of attestations
_MAX_ATTESTATIONS: constant(uint256) = 4
_ATTESTATION_LENGTH: constant(uint256) = 380


@deploy
def __init__(gateway_minter: address):
    assert gateway_minter != empty(address)
    GATEWAY_MINTER = gateway_minter


@external
def gateway_deposit_and_stake(attestation: Bytes[1528], signature: Bytes[65]):
    transfers: DynArray[
        TransferSpec, _MAX_ATTESTATIONS
    ] = self._decode_attestation(attestation)
    extcall IGatewayMinter(GATEWAY_MINTER).gatewayMint(attestation, signature)

    for transfer: TransferSpec in transfers:
        if cctp._is_hub() and transfer.destination_recipient == self:
            deposited: uint256 = vault._deposit(self, transfer.value)
            erc20._approve(self, vault.staking, deposited)
            extcall IERC4626(vault.staking).deposit(
                deposited, transfer.source_depositor
            )


@external
def decode_attestation(attestation: Bytes[1528]) -> DynArray[TransferSpec, _MAX_ATTESTATIONS]:
    return self._decode_attestation(attestation)


def _decode_attestation(
    attestation: Bytes[1528],
) -> DynArray[TransferSpec, _MAX_ATTESTATIONS]:
    magic: Bytes[4] = slice(attestation, 0, 4)
    transfers: DynArray[TransferSpec, _MAX_ATTESTATIONS] = empty(
        DynArray[TransferSpec, _MAX_ATTESTATIONS]
    )

    if magic == _ATTESTATION_SET_MAGIC:
        number_of_attestations: uint256 = convert(
            slice(attestation, 4, 4), uint256
        )
        #attestation_set: Bytes[1520] = slice(attestation, 8, _MAX_ATTESTATIONS * _ATTESTATION_LENGTH)

        for i: uint256 in range(
            number_of_attestations, bound=_MAX_ATTESTATIONS
        ):
            # 8 bytes is offset of attestation set metadata.
            attestation_payload: Bytes[380] = slice(attestation, 8 + i * _ATTESTATION_LENGTH, _ATTESTATION_LENGTH)
            transfers.append(self._decode_transfer_spec(attestation_payload))
    elif magic == _ATTESTATION_MAGIC:
        attestation_payload: Bytes[380] = slice(attestation, 0, _ATTESTATION_LENGTH)
        transfers.append(self._decode_transfer_spec(attestation_payload))
    else:
        raise "unknown circle attestation payload"

    return transfers


def _decode_transfer_spec(attestation: Bytes[380]) -> TransferSpec:
    attestation_magic: Bytes[4] = slice(attestation, 0, 4)
    assert attestation_magic == _ATTESTATION_MAGIC

    # transfer spec length is 340 bytes, without hook data
    encoded_transfer_spec: Bytes[340] = slice(attestation, 40, 340)

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

# pragma version ^0.4.3
# @license MIT
from snekmate.auth import access_control

uses: access_control

from ethereum.ercs import IERC20

from src.modules import ledger

uses: ledger

from src.modules import roles

#  MessageV2 format
#  * Field                        Bytes      Type       Index
#  * version                      4          uint32     0
#  * sourceDomain                 4          uint32     4
#  * destinationDomain            4          uint32     8
#  * nonce                        32         bytes32    12
#  * sender                       32         bytes32    44
#  * recipient                    32         bytes32    76
#  * destinationCaller            32         bytes32    108
#  * minFinalityThreshold         4          uint32     140
#  * finalityThresholdExecuted    4          uint32     144
#  * messageBody                  dynamic    bytes      148
struct MessageV2:
    version: uint32
    sourceDomain: uint32
    destinationDomain: uint32
    nonce: bytes32
    sender: bytes32
    recipient: bytes32
    destinationCaller: bytes32
    minFinalityThreshold: uint32
    finalityThresholdExecuted: uint32
    messageBody: Bytes[512]


#  BurnMessageV2 format
#  * Field                 Bytes      Type       Index
#  * version               4          uint32     0
#  * burnToken             32         bytes32    4
#  * mintRecipient         32         bytes32    36
#  * amount                32         uint256    68
#  * messageSender         32         bytes32    100
#  * maxFee                32         uint256    132
#  * feeExecuted           32         uint256    164
#  * expirationBlock       32         uint256    196
#  * hookData              dynamic    bytes      228
struct BurnMessageV2:
    version: uint32
    burnToken: bytes32
    mintRecipient: bytes32
    amount: uint256
    messageSender: bytes32
    maxFee: uint256
    feeExecuted: uint256
    expirationBlock: uint256
    hookData: Bytes[64]


interface ITokenMessanger:
    def depositForBurn(
        amount: uint256,
        destinationDomain: uint32,
        mintRecipient: bytes32,
        burnToken: address,
        destinationCaller: bytes32,
        maxFee: uint256,
        minFinalityThreshold: uint32,
    ): nonpayable
    def depositForBurnWithHook(
        amount: uint256,
        destinationDomain: uint32,
        mintRecipient: bytes32,
        burnToken: address,
        destinationCaller: bytes32,
        maxFee: uint256,
        minFinalityThreshold: uint32,
        hookData: Bytes[64],
    ): nonpayable


interface IMessageTransmitter:
    def receiveMessage(
        message: Bytes[1024], attestation: Bytes[1024]
    ) -> bool: nonpayable
    def localDomain() -> uint32: view


MESSANGER: public(immutable(address))
TRANSMITTER: public(immutable(address))
HUB_DOMAIN: public(immutable(uint32))

whitelisted_domains: public(HashMap[uint32, address])


@deploy
def __init__(hub_domain: uint32, messanger: address, transmitter: address):
    assert messanger != empty(address)
    assert transmitter != empty(address)

    MESSANGER = messanger
    TRANSMITTER = transmitter
    HUB_DOMAIN = hub_domain


@external
def update_whitelisted_domain(domain: uint32, terrace: address):
    access_control._check_role(access_control.DEFAULT_ADMIN_ROLE, msg.sender)
    self.whitelisted_domains[domain] = terrace


@external
def receive_message(message: Bytes[1024], attestation: Bytes[1024]):
    assert extcall IMessageTransmitter(TRANSMITTER).receiveMessage(
        message, attestation
    )
    bridge_message: MessageV2 = abi_decode(message, MessageV2)
    burn_message: BurnMessageV2 = abi_decode(
        bridge_message.messageBody, BurnMessageV2
    )

    if self._is_hub():
        # from terrace to hub
        whitelsited_terrace: address = self.whitelisted_domains[
            bridge_message.sourceDomain
        ]
        assert whitelsited_terrace != empty(address)
        # TODO: check that sender is whitelsited_terrace
        ledger._repay(whitelsited_terrace, burn_message.amount)
    else:
        # from hub to terrace
        assert bridge_message.sourceDomain == HUB_DOMAIN
        ledger._fill(burn_message.amount)


@external
def send_to_terrace(
    amount: uint256,
    destination_domain: uint32,
    max_fee: uint256,
    min_finality_threshold: uint32,
):
    access_control._check_role(roles.OPERATOR_ROLE, msg.sender)
    whitelsited_terrace: address = self.whitelisted_domains[destination_domain]

    assert whitelsited_terrace != empty(address)
    assert amount > 0
    assert destination_domain != HUB_DOMAIN
    assert min_finality_threshold > 0

    ledger._borrow(whitelsited_terrace, amount)

    encoded_terrace: bytes32 = convert(whitelsited_terrace, bytes32)

    extcall IERC20(ledger.USDC).approve(MESSANGER, amount)
    extcall ITokenMessanger(MESSANGER).depositForBurn(
        amount,
        destination_domain,
        encoded_terrace,
        ledger.USDC,
        empty(bytes32),  # can call any address
        max_fee,
        min_finality_threshold,
    )


@external
def send_to_hub(
    amount: uint256, max_fee: uint256, min_finality_threshold: uint32
):
    access_control._check_role(roles.OPERATOR_ROLE, msg.sender)
    assert amount > 0
    assert min_finality_threshold > 0

    hub_address: address = self.whitelisted_domains[HUB_DOMAIN]
    assert hub_address != empty(address)

    encoded_hub: bytes32 = convert(hub_address, bytes32)

    ledger._deplete(amount)

    extcall IERC20(ledger.USDC).approve(MESSANGER, amount)
    extcall ITokenMessanger(MESSANGER).depositForBurn(
        amount,
        HUB_DOMAIN,
        encoded_hub,
        ledger.USDC,
        empty(bytes32),  # can call any address
        max_fee,
        min_finality_threshold,
    )


@internal
@view
def _is_hub() -> bool:
    return self._local_domain() == HUB_DOMAIN


@internal
@view
def _local_domain() -> uint32:
    return staticcall IMessageTransmitter(TRANSMITTER).localDomain()

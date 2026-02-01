# pragma version ^0.4.3
# @license MIT
from src.modules import ledger

uses: ledger

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

    
interface IMessageTransmitter:
    def receiveMessage(message: Bytes[1024], attestation: Bytes[1024]) -> bool: nonpayable

MESSAGE_TRANSMITTER: public(immutable(address))

@deploy
def __init__(transmitter: address):
    assert transmitter != empty(address)
    MESSAGE_TRANSMITTER = transmitter


@external
def receive_message(message: Bytes[1024], attestation: Bytes[1024]):
    assert extcall IMessageTransmitter(MESSAGE_TRANSMITTER).receiveMessage(message, attestation)
    bridge_message: MessageV2 = abi_decode(message, MessageV2)
    burn_message: BurnMessageV2 = abi_decode(bridge_message.messageBody, BurnMessageV2)
    ledger._fill(burn_message.amount)
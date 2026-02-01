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
    mintRecipient: address
    amount: uint256
    messageSender: address
    maxFee: uint256
    feeExecuted: uint256
    expirationBlock: uint256
    nonce: uint256


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
        hookData: Bytes[32],
    ): nonpayable


interface IMessageTransmitter:
    def receiveMessage(
        message: Bytes[1024], attestation: Bytes[1024]
    ) -> bool: nonpayable
    def localDomain() -> uint32: view

# id is for mapping SendFunds <-> ReceivedFunds events
# id = keccak256(from_domain,to_domain,nonce)
event SendFunds:
    id: bytes32
    from_terrace: address
    from_domain: uint32
    to_terrace: address
    to_domain: uint32
    amount: uint256

event ReceivedFunds:
    id: bytes32
    from_terrace: address
    from_domain: uint32
    to_terrace: address
    to_domain: uint32
    amount: uint256

MESSANGER: public(immutable(address))
TRANSMITTER: public(immutable(address))
HUB_DOMAIN: public(immutable(uint32))

whitelisted_domains: public(HashMap[uint32, address])
bridge_nonce: public(uint256)

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
    ), "Failed to receive message"

    source_domain: uint32 = self._get_source_domain(message)
    burn_message: BurnMessageV2 = self._decode_message(message)
    whitelsited_terrace: address = self.whitelisted_domains[source_domain]
    actual_amount: uint256 = burn_message.amount - burn_message.feeExecuted

    if self._is_hub():
        # from terrace to hub
        if burn_message.messageSender == whitelsited_terrace:
            ledger._repay(whitelsited_terrace, actual_amount)
    else:
        # from hub to terrace
        if source_domain == HUB_DOMAIN and burn_message.messageSender == whitelsited_terrace:
            ledger._fill(actual_amount)
    
    self._emit_received_funds_event(source_domain, burn_message.messageSender, actual_amount, burn_message.nonce)


@external
def send_all_to_terrace(
    destination_domain: uint32, max_fee: uint256, min_finality_threshold: uint32
):
    balance: uint256 = staticcall IERC20(ledger.USDC).balanceOf(self)
    self._send_to_terrace(
        balance, destination_domain, max_fee, min_finality_threshold
    )


@external
def send_to_terrace(
    amount: uint256,
    destination_domain: uint32,
    max_fee: uint256,
    min_finality_threshold: uint32,
):
    self._send_to_terrace(
        amount, destination_domain, max_fee, min_finality_threshold
    )


@external
def send_all_to_hub(max_fee: uint256, min_finality_threshold: uint32):
    balance: uint256 = staticcall IERC20(ledger.USDC).balanceOf(self)
    self._send_to_hub(balance, max_fee, min_finality_threshold)


@external
def send_to_hub(
    amount: uint256,
    max_fee: uint256,
    min_finality_threshold: uint32,
):
    self._send_to_hub(amount, max_fee, min_finality_threshold)


@internal
def _send_to_hub(
    amount: uint256, max_fee: uint256, min_finality_threshold: uint32
):
    access_control._check_role(roles.OPERATOR_ROLE, msg.sender)
    assert amount > 0
    assert min_finality_threshold > 0

    hub_address: address = self.whitelisted_domains[HUB_DOMAIN]
    assert hub_address != empty(address)

    nonce: uint256 = self.bridge_nonce + 1
    encoded_nonce: Bytes[32] = abi_encode(nonce)
    encoded_hub: bytes32 = convert(hub_address, bytes32)

    ledger._deplete(amount)

    extcall IERC20(ledger.USDC).approve(MESSANGER, amount)
    extcall ITokenMessanger(MESSANGER).depositForBurnWithHook(
        amount,
        HUB_DOMAIN,
        encoded_hub, # mint recipient
        ledger.USDC,
        encoded_hub,  # only hub can call receive message
        max_fee,
        min_finality_threshold,
        encoded_nonce,
    )

    self.bridge_nonce = nonce
    self._emit_send_funds_event(HUB_DOMAIN, hub_address, amount, nonce)


@internal
def _send_to_terrace(
    amount: uint256,
    destination_domain: uint32,
    max_fee: uint256,
    min_finality_threshold: uint32,
):
    access_control._check_role(roles.OPERATOR_ROLE, msg.sender)
    assert self._is_hub()

    destination_terrace: address = self.whitelisted_domains[destination_domain]

    assert destination_terrace != empty(address)
    assert amount > 0
    assert destination_domain != HUB_DOMAIN
    assert min_finality_threshold > 0

    ledger._borrow(destination_terrace, amount)

    nonce: uint256 = self.bridge_nonce + 1
    encoded_nonce: Bytes[32] = abi_encode(nonce)
    encoded_terrace: bytes32 = convert(destination_terrace, bytes32)

    extcall IERC20(ledger.USDC).approve(MESSANGER, amount)
    extcall ITokenMessanger(MESSANGER).depositForBurnWithHook(
        amount,
        destination_domain,
        encoded_terrace, # mint recipient
        ledger.USDC,
        encoded_terrace,  # only terrace can call receive message
        max_fee,
        min_finality_threshold,
        encoded_nonce,
    )
    self.bridge_nonce = nonce
    self._emit_send_funds_event(destination_domain, destination_terrace, amount, nonce)

@internal
@view
def _is_hub() -> bool:
    return self._local_domain() == HUB_DOMAIN


@internal
@view
def _local_domain() -> uint32:
    return staticcall IMessageTransmitter(TRANSMITTER).localDomain()


@internal
@pure
def _get_source_domain(message: Bytes[1024]) -> uint32:
    # sourceDomain: 4 bytes at index 4
    return convert(slice(message, 4, 4), uint32)

@internal
@pure
def _decode_message(message: Bytes[1024]) -> BurnMessageV2:
    # messageBody - dynamic bytes starting at index 148
    message_body_start: uint256 = 148
    message_body_len: uint256 = len(message) - message_body_start
    message_body: Bytes[1024] = slice(message, message_body_start, message_body_len)

    # parse BurnMessageV2 from messageBody (packed format)
    burn_version: uint32 = convert(slice(message_body, 0, 4), uint32)
    burn_token: bytes32 = convert(slice(message_body, 4, 32), bytes32)
    mint_recipient: address = abi_decode(slice(message_body, 36, 32), address)
    amount: uint256 = convert(slice(message_body, 68, 32), uint256)
    message_sender: address = abi_decode(slice(message_body, 100, 32), address)
    max_fee: uint256 = convert(slice(message_body, 132, 32), uint256)
    fee_executed: uint256 = convert(slice(message_body, 164, 32), uint256)
    expiration_block: uint256 = convert(slice(message_body, 196, 32), uint256)
    nonce: uint256 = abi_decode(slice(message_body, 228, 32), uint256)

    return BurnMessageV2(
        version=burn_version,
        burnToken=burn_token,
        mintRecipient=mint_recipient,
        amount=amount,
        messageSender=message_sender,
        maxFee=max_fee,
        feeExecuted=fee_executed,
        expirationBlock=expiration_block,
        nonce=nonce
    )

@internal
def _emit_send_funds_event(_to_domain: uint32, _to_terrace: address, _amount: uint256, _nonce: uint256):
    from_domain: uint32 = self._local_domain()
    from_terrace: address = self
    event_id: bytes32 = self._bridge_event_id(from_domain, _to_domain, _nonce)
    log SendFunds(
        id=event_id,
        from_terrace=from_terrace,
        from_domain=from_domain,
        to_terrace=_to_terrace,
        to_domain=_to_domain,
        amount=_amount,
    )

@internal
def _emit_received_funds_event(_from_domain: uint32, _from_terrace: address, _amount: uint256, _nonce: uint256):
    to_domain: uint32 = self._local_domain()
    to_terrace: address = self
    event_id: bytes32 = self._bridge_event_id(_from_domain, to_domain, _nonce)
    log ReceivedFunds(
        id=event_id,
        from_terrace=_from_terrace,
        from_domain=_from_domain,
        to_terrace=to_terrace,
        to_domain=to_domain,
        amount=_amount,
    )
    

@internal
@pure
def _bridge_event_id(_from_domain: uint32, _to_domain: uint32, nonce: uint256) -> bytes32:
    return keccak256(abi_encode(_from_domain, _to_domain, nonce))
    

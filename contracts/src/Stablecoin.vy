# pragma version ^0.4.3
# @license MIT
from snekmate.auth import access_control as access

initializes: access

from snekmate.auth import ownable as ow

initializes: ow

from ethereum.ercs import IERC20
from snekmate.tokens import erc20

initializes: erc20[ownable := ow]

from src.modules import withdrawal_queue

initializes: withdrawal_queue[access_control := access]

from src import Vault

initializes: Vault[access := access]

from src.modules import strategy_registry

exports: (
    erc20.__interface__,
    Vault.__interface__,
    access.__interface__,
    withdrawal_queue.__interface__,
)

USDC: public(immutable(address))
staking: public(address)


# @dev Emitted when the user deposits
event Deposit:
    user: indexed(address)
    amount: uint256


# @dev Emitted when the user requests to withdraw
event RequestWithdraw:
    user: indexed(address)
    amount: uint256


event Withdraw:
    user: indexed(address)
    amount: uint256

@deploy
def __init__(usdc: address, withdrawal_delay: uint256):
    assert usdc != empty(address)
    assert withdrawal_delay > 0

    USDC = usdc

    ow.__init__()
    access.__init__()
    withdrawal_queue.__init__(withdrawal_delay)
    Vault.__init__()
    erc20.__init__("Terrace USD", "tUSD", 18, "Terrace Finance", "0x01")


@external
def set_staking(staking: address):
    access._check_role(access.DEFAULT_ADMIN_ROLE, msg.sender)
    assert staking != empty(address)
    self.staking = staking


@external
def deposit(amount: uint256):
    assert amount > 0
    assert extcall IERC20(USDC).transferFrom(msg.sender, self, amount)
    scaled_amount: uint256 = amount * 10**12
    erc20._mint(msg.sender, scaled_amount)
    log Deposit(user=msg.sender, amount=scaled_amount)


@external
def withdraw(amount: uint256):
    assert amount > 0
    erc20._transfer(msg.sender, self, amount)
    withdrawal_queue._add_withdrawal(amount)
    log RequestWithdraw(user=msg.sender, amount=amount)


@external
def execute_withdraw():
    req: withdrawal_queue.WithdrawalRequest = (
        withdrawal_queue._remove_withdrawal()
    )
    erc20._burn(self, req.amount)
    real_amount: uint256 = req.amount // 10**12
    assert extcall IERC20(USDC).transfer(msg.sender, real_amount)
    log Withdraw(user=msg.sender, amount=real_amount)

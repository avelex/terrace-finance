# pragma version ^0.4.3
# @license MIT
from snekmate.auth import access_control

uses: access_control

from src.modules import withdrawal_queue

uses: withdrawal_queue

from ethereum.ercs import IERC20
from snekmate.tokens import erc20

uses: erc20

from src.modules import ledger

uses: ledger

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
def __init__():
    pass


@external
def set_staking(staking: address):
    access_control._check_role(access_control.DEFAULT_ADMIN_ROLE, msg.sender)
    assert staking != empty(address)
    self.staking = staking


@external
def deposit(amount: uint256):
    assert amount > 0
    assert extcall IERC20(ledger.USDC).transferFrom(msg.sender, self, amount)
    self._deposit(amount)


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
    assert extcall IERC20(ledger.USDC).transfer(msg.sender, real_amount)
    ledger._deplete(real_amount)
    log Withdraw(user=msg.sender, amount=real_amount)


@external
def distribute_fees():
    fees: uint256 = ledger._pull_fees()
    assert fees > 0
    assert self.staking != empty(address)

    scaled_amount: uint256 = fees * 10**12
    erc20._mint(self.staking, scaled_amount)


@internal
def _deposit(amount: uint256) -> uint256:
    ledger._fill(amount)
    scaled_amount: uint256 = amount * 10**12
    erc20._mint(msg.sender, scaled_amount)
    log Deposit(user=msg.sender, amount=scaled_amount)
    return scaled_amount

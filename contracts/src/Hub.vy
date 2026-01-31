# pragma version ^0.4.3
# @license MIT
from snekmate.tokens import erc20
from ethereum.ercs import IERC20
from src import withdrawal_queue

initializes: withdrawal_queue

tUSD: public(immutable(erc20.__interface__))
stUSD: public(immutable(address))
USDC: public(immutable(address))

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
def __init__(
    tUSD: address, stUSD: address, usdc: address, withdrawal_delay: uint256
):
    assert tUSD != empty(address)
    assert stUSD != empty(address)
    assert usdc != empty(address)
    assert withdrawal_delay > 0

    tUSD = erc20.__at__(tUSD)
    stUSD = stUSD
    USDC = usdc

    withdrawal_queue.__init__(withdrawal_delay)


@external
def deposit(amount: uint256):
    assert amount > 0
    assert extcall IERC20(USDC).transferFrom(msg.sender, self, amount)
    extcall tUSD.mint(msg.sender, amount)
    log Deposit(user=msg.sender, amount=amount)


@external
def withdraw(amount: uint256):
    assert amount > 0
    assert extcall tUSD.transferFrom(msg.sender, self, amount)
    withdrawal_queue._add_withdrawal(amount)
    log RequestWithdraw(user=msg.sender, amount=amount)


@external
@nonreentrant
def execute_withdraw():
    req: withdrawal_queue.WithdrawalRequest = (
        withdrawal_queue._remove_withdrawal()
    )
    extcall tUSD.burn(req.amount)
    assert extcall IERC20(USDC).transferFrom(self, msg.sender, req.amount)
    log Withdraw(user=msg.sender, amount=req.amount)

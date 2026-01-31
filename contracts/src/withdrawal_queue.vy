# pragma version ^0.4.3
# @license MIT
struct WithdrawalRequest:
    amount: uint256
    timestamp: uint256


MAX_DELAY: constant(uint256) = 7 * 24 * 60 * 60

# @dev The delay between withdrawal request and withdrawal execution in seconds.
# max value is 7 days.
delay: public(uint256)

pendingWithdrawals: public(HashMap[address, WithdrawalRequest])


@deploy
def __init__(delay: uint256):
    assert delay <= MAX_DELAY
    self.delay = delay


@internal
def _add_withdrawal(amount: uint256):
    request: WithdrawalRequest = self.pendingWithdrawals[msg.sender]
    assert amount > 0
    assert request.timestamp == 0

    withdraw_at: uint256 = block.timestamp + self.delay

    self.pendingWithdrawals[msg.sender] = WithdrawalRequest(
        amount=amount, timestamp=withdraw_at
    )


@internal
def _remove_withdrawal() -> WithdrawalRequest:
    request: WithdrawalRequest = self.pendingWithdrawals[msg.sender]
    assert request.timestamp > 0
    assert block.timestamp >= request.timestamp
    self.pendingWithdrawals[msg.sender] = empty(WithdrawalRequest)
    return request

# pragma version ^0.4.3
# @license MIT
from snekmate.auth import access_control

uses: access_control


struct WithdrawalRequest:
    amount: uint256
    timestamp: uint256


MAX_WITHDRAWAL_DELAY: constant(uint256) = 7 * 24 * 60 * 60

# @dev The delay between withdrawal request and withdrawal execution in seconds.
# max value is 7 days.
withdrawal_delay: public(uint256)

# @dev The mapping of pending withdrawals by user address.
pending_withdrawals: public(HashMap[address, WithdrawalRequest])

# @dev Emitted when the withdrawal delay is updated.
event DelayUpdated:
    new_delay: uint256


@deploy
def __init__(delay: uint256):
    assert delay > 0 and delay <= MAX_WITHDRAWAL_DELAY
    self.withdrawal_delay = delay


@external
def update_withdrawal_delay(delay: uint256):
    access_control._check_role(access_control.DEFAULT_ADMIN_ROLE, msg.sender)
    assert delay > 0 and delay <= MAX_WITHDRAWAL_DELAY
    self.withdrawal_delay = delay
    log DelayUpdated(new_delay=delay)


@internal
def _add_withdrawal(amount: uint256):
    request: WithdrawalRequest = self.pending_withdrawals[msg.sender]
    assert amount > 0
    assert request.timestamp == 0

    withdraw_at: uint256 = block.timestamp + self.withdrawal_delay

    self.pending_withdrawals[msg.sender] = WithdrawalRequest(
        amount=amount, timestamp=withdraw_at
    )


@internal
def _remove_withdrawal() -> WithdrawalRequest:
    request: WithdrawalRequest = self.pending_withdrawals[msg.sender]
    assert request.timestamp > 0
    assert block.timestamp >= request.timestamp
    self.pending_withdrawals[msg.sender] = empty(WithdrawalRequest)
    return request

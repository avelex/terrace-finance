# pragma version ^0.4.3
# @license MIT

USDC: public(immutable(address))

usdc_reserve: public(uint256)
usdc_usage: public(uint256)
usdc_fees: public(uint256)
usage: HashMap[address, uint256]

@deploy
def __init__(usdc: address):
    assert usdc != empty(address)
    USDC = usdc

@internal
def _fill(amount: uint256):
    self.usdc_reserve += amount

@internal
def _deplete(amount: uint256):
    self.usdc_reserve -= amount

@internal
def _borrow(who: address, amount: uint256):
    self._deplete(amount)
    self.usdc_usage += amount
    self.usage[who] += amount
    
@internal
def _repay(who: address, amount: uint256):
    current_usage: uint256 = self.usage[who]
    
    assert current_usage > 0
    assert amount > 0

    if amount > current_usage:
        self.usdc_usage -= current_usage
        self.usage[who] = 0
        self.usdc_fees += amount - current_usage
    else:
        self.usdc_usage -= amount
        self.usage[who] -= amount
    
    self._fill(amount)

    

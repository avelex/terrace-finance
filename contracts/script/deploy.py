from src import Stablecoin
from src import Staking
from src.mocks import erc20_mock
from moccasin.boa_tools import VyperContract

def deploy_stablecoin(usdc, withdrawal_delay) -> VyperContract:
    return Stablecoin.deploy(usdc, withdrawal_delay)

def deploy_staking(stablecoin) -> VyperContract:
    return Staking.deploy(stablecoin)

def deploy_usdc_mock() -> VyperContract:
    return erc20_mock.deploy("USD Coin", "USDC", 6, 1_000_000)
import pytest
import boa
from script.deploy import deploy_stablecoin, deploy_staking, deploy_usdc_mock

@pytest.fixture(scope="session")
def admin():
    return boa.env.generate_address()

@pytest.fixture(scope="session")
def accounts():
    return [boa.env.generate_address() for _ in range(10)]


@pytest.fixture(scope="session")
def usdc(admin):
    with boa.env.prank(admin):
        return deploy_usdc_mock()

@pytest.fixture(scope="session")
def withdrawal_delay():
    return 3600

@pytest.fixture
def stablecoin_contract(admin, usdc, withdrawal_delay):
    with boa.env.prank(admin):
        return deploy_stablecoin(usdc, withdrawal_delay)

@pytest.fixture
def staked_stablecoin_contract(admin, stablecoin):
    with boa.env.prank(admin):
        return deploy_staking(stablecoin)
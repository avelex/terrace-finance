import boa
from script.deploy import deploy_stablecoin

def test_deploy_stablecoin(stablecoin_contract, admin, usdc, withdrawal_delay):
    assert stablecoin_contract.name() == "Terrace USD"
    assert stablecoin_contract.symbol() == "tUSD"
    assert stablecoin_contract.decimals() == 18
    assert stablecoin_contract.owner() == admin

    default_admin_role = bytes.fromhex("0000000000000000000000000000000000000000000000000000000000000000")
    assert stablecoin_contract.hasRole(default_admin_role, admin)

    assert stablecoin_contract.USDC() == usdc.address
    assert stablecoin_contract.withdrawal_delay() == withdrawal_delay
    assert stablecoin_contract.strategy_root_hash().hex() == "0000000000000000000000000000000000000000000000000000000000000000"


def test_deposit_happy_path(stablecoin_contract, admin, usdc, accounts):
    alice = accounts[0]
    
    with boa.env.prank(admin):
        usdc.transfer(alice, 5_000_000)

    want_tusd = 5 * 10**18
    # Alice deposits 5 USDC (all of her balance)
    with boa.env.prank(alice):
        usdc.approve(stablecoin_contract.address, 5_000_000)
        stablecoin_contract.deposit(5_000_000) 

    # Ensure 5 tUSD minted to Alice
    assert stablecoin_contract.balanceOf(alice) == want_tusd
    # Ensure 5 tUSD minted in total
    assert stablecoin_contract.totalSupply() == want_tusd
    # Ensure Alice has no USDC left
    assert usdc.balanceOf(alice) == 0 
    # Ensure tUSD contract keeps 5 USDC
    assert usdc.balanceOf(stablecoin_contract.address) == 5_000_000

def test_withdraw_happy_path(stablecoin_contract, admin, usdc, accounts):
    alice = accounts[0]
    
    with boa.env.prank(admin):
        usdc.transfer(alice, 5_000_000)

    # Alice deposits 5 USDC and withdraws 5 tUSD
    with boa.env.prank(alice): 
        usdc.approve(stablecoin_contract.address, 5_000_000)
        stablecoin_contract.deposit(5_000_000) 
        stablecoin_contract.withdraw(5 * 10**18)

    # Ensure Alice has no tUSD left
    assert stablecoin_contract.balanceOf(alice) == 0
    # Ensure Alice has no USDC yet
    assert usdc.balanceOf(alice) == 0

    pending = stablecoin_contract.pending_withdrawals(alice)
    assert pending.amount == 5 * 10**18
    assert pending.timestamp == boa.env.timestamp + stablecoin_contract.withdrawal_delay()

    with boa.reverts("User has pending withdrawal"):
        with boa.env.prank(alice):
            stablecoin_contract.execute_withdraw()

    # Travel to the future, after the withdrawal delay
    boa.env.timestamp += stablecoin_contract.withdrawal_delay() + 1

    want_usdc = 5_000_000
    with boa.env.prank(alice):
        stablecoin_contract.execute_withdraw()
    
    # Ensure Alice has 5 USDC
    assert usdc.balanceOf(alice) == want_usdc
    # Ensure tUSD contract has no USDC left
    assert usdc.balanceOf(stablecoin_contract.address) == 0
    # Ensure 5 tUSD burned
    assert stablecoin_contract.totalSupply() == 0

def test_is_allowed_strategy(stablecoin_contract, admin):
    with boa.env.prank(admin):
        new_hash = bytes.fromhex("9c208074e7b8aea28a4fbc10ffe54a25f4e5eb297d16411ea7e9990b6b02e2f3")
        stablecoin_contract.update_root_hash(new_hash)

    proof = [
        bytes.fromhex("220f2c6e9fc386a9244e977253160f39168afa8282a15c86cd4ffa9586f67d52"),
        bytes.fromhex("ea1fe3f259049170acb189a30de0a1fbc2869e99e00703b8c24fd3d8824eb54b"),
    ]

    target = "0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"
    selector = bytes.fromhex("617ba037")

    assert stablecoin_contract.is_strategy_allowed(target, selector, proof)

    wrong_selector = bytes.fromhex("eeeeeeee")
    assert not stablecoin_contract.is_strategy_allowed(target, wrong_selector, proof)


def test_execute_strategy_supply_aave(admin, accounts):
    erc20_mock = boa.load_partial("src/mocks/erc20_mock.vy")
    
    real_usdc_address = "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"
    real_usdc = erc20_mock.at(real_usdc_address)
    usdc_whale = "0x20FE51A9229EEf2cF8Ad9E89d91CAb9312cF3b7A"

    with boa.env.prank(admin):
        stablecoin_contract = deploy_stablecoin(real_usdc_address, 3600)

    alice = accounts[0]
    amount = 10_000 * 10**6  # 10,000 USDC 
    
    with boa.env.prank(usdc_whale):
        real_usdc.transfer(alice, amount)
    
    assert real_usdc.balanceOf(alice) >= amount

    with boa.env.prank(alice):
        real_usdc.approve(stablecoin_contract.address, amount)
        stablecoin_contract.deposit(amount)
    
    assert real_usdc.balanceOf(stablecoin_contract.address) == amount


    approve_proof = [
        bytes.fromhex("ea1fe3f259049170acb189a30de0a1fbc2869e99e00703b8c24fd3d8824eb54b"),
        bytes.fromhex("dcced49bcc2022f18d087b37325862e1d8ddf620b4b2f07db1061b85900ef3b0"),
    ]

    approve_target = "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"
    approve_selector = bytes.fromhex("095ea7b3")
    approve_encoded_data = stablecoin_contract.encode_erc20_approve("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5", amount)

    supply_proof = [
        bytes.fromhex("220f2c6e9fc386a9244e977253160f39168afa8282a15c86cd4ffa9586f67d52"),
        bytes.fromhex("149c0b63d7c8c95e4f6e03dc273ed695e1066e34dd2d078723b1b372fcbf586a"),
    ]

    supply_target = "0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"
    supply_selector = bytes.fromhex("617ba037")
    supply_encoded_data = stablecoin_contract.encode_supply_aave_v3(real_usdc_address, amount, stablecoin_contract.address, 0)

    with boa.env.prank(admin):
        new_hash = bytes.fromhex("27bbcef950c4fa0db57bd9883b6009c1c041ffefca6b98e82c8abd79cae2bad7")
        stablecoin_contract.update_root_hash(new_hash)
        stablecoin_contract.grantRole(stablecoin_contract.OPERATOR_ROLE(), admin)

        stablecoin_contract.execute_strategy(approve_target, approve_selector, approve_encoded_data, approve_proof)
        stablecoin_contract.execute_strategy(supply_target, supply_selector, supply_encoded_data, supply_proof)
    
    assert real_usdc.balanceOf(stablecoin_contract.address) == 0
    
    aave_usdc_address = "0x4e65fE4DbA92790696d040ac24Aa414708F5c0AB"
    aave_usdc = erc20_mock.at(aave_usdc_address)

    assert aave_usdc.balanceOf(stablecoin_contract.address) > 0


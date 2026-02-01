from src import Stablecoin
from src import Staking
from moccasin.config import get_config
from moccasin.boa_tools import VyperContract

def deploy() -> VyperContract:
    config = get_config()
    active_network = config.get_active_network()
    print("Deploying Stablecoin on", active_network.name, active_network.chain_id)

    usdc = active_network.get_named_contract("usdc")
    messanger = active_network.get_named_contract("messanger")
    transmitter = active_network.get_named_contract("transmitter")
    gateway_minter = active_network.get_named_contract("gateway_minter")
    withdrawal_delay = active_network.extra_data["withdrawal_delay"]
    hub_domain = active_network.extra_data["hub_domain"]

    stablecoin = Stablecoin.deploy(usdc.address, withdrawal_delay, messanger.address, transmitter.address, hub_domain, gateway_minter.address)
    print("Stablecoin deployed at", stablecoin.address)

    result = active_network.moccasin_verify(stablecoin)
    result.wait_for_verification()
    print("Stablecoin verified")

    staking = Staking.deploy(stablecoin.address)
    print("Staking deployed at", staking.address)

    result = active_network.moccasin_verify(staking)
    result.wait_for_verification()
    print("Staking verified")

    stablecoin.set_staking(staking.address)
    print("Staking set in Stablecoin")
    
    return stablecoin

def moccasin_main():
    deploy()

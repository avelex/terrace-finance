from src import Vault
from moccasin.config import get_config
from moccasin.boa_tools import VyperContract

def deploy() -> VyperContract:
    config = get_config()
    active_network = config.get_active_network()
    print("Deploying Vault on", active_network.name, active_network.chain_id)

    usdc = active_network.get_named_contract("usdc")
    messanger = active_network.get_named_contract("messanger")
    transmitter = active_network.get_named_contract("transmitter")
    hub_domain = active_network.extra_data["hub_domain"]

    vault = Vault.deploy(usdc.address, messanger.address, transmitter.address, hub_domain)
    print("Vault deployed at", vault.address)

    result = active_network.moccasin_verify(vault)
    result.wait_for_verification()
    print("Vault verified")
    
    return vault

def moccasin_main():
    deploy()
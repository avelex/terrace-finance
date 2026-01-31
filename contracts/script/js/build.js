import { StandardMerkleTree } from "@openzeppelin/merkle-tree";
import fs from "fs";

// Map of function selectors to their bytes4 representation
const selectors = new Map([
    ["approve(address,uint256)", "0x095ea7b3"],
    ["supply(address,uint256,address,uint16)", "0x617ba037"],
    ["withdraw(address,uint256,address)", "0x69328dec"],
    ["depositForBurn(uint256,uint32,bytes32,address,bytes32,uint256,uint32)", "0x8e0250ee"],
]);

const values = [
    ["0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913", selectors.get("approve(address,uint256)")],
    ["0xA238Dd80C259a72e81d7e4664a9801593F98d1c5", selectors.get("supply(address,uint256,address,uint16)")],
    ["0xA238Dd80C259a72e81d7e4664a9801593F98d1c5", selectors.get("withdraw(address,uint256,address)")],
    ["0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d", selectors.get("depositForBurn(uint256,uint32,bytes32,address,bytes32,uint256,uint32)")],
];



const tree = StandardMerkleTree.of(values, ["address", "bytes4"]);

console.log('Merkle Root:', tree.root);

fs.writeFileSync("tree.json", JSON.stringify(tree.dump()));
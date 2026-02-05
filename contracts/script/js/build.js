import { StandardMerkleTree } from "@openzeppelin/merkle-tree";
import fs from "fs";

// Map of function selectors to their bytes4 representation
const selectors = new Map([
    ["approve(address,uint256)", "0x095ea7b3"],
    ["supply(address,uint256,address,uint16)", "0x617ba037"],
    ["withdraw(address,uint256,address)", "0x69328dec"],
]);


const domains = new Map([
    [0, [
        ["0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238", selectors.get("approve(address,uint256)")],
        ["0x6Ae43d3271ff6888e7Fc43Fd7321a503ff738951", selectors.get("supply(address,uint256,address,uint16)")],
        ["0x6Ae43d3271ff6888e7Fc43Fd7321a503ff738951", selectors.get("withdraw(address,uint256,address)")],
    ]],
])

domains.forEach((values, key) => {
    const tree = StandardMerkleTree.of(values, ["address", "bytes4"]);
    console.log('Merkle Root for domain', key, ':', tree.root);
    fs.writeFileSync(key.toString() + "_tree.json", JSON.stringify(tree.dump()));
})
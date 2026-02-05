import { StandardMerkleTree } from "@openzeppelin/merkle-tree";
import fs from "fs";

const args = process.argv.slice(2);

// Parse --path flag from command line arguments
let targetPath = null;

for (let i = 0; i < args.length; i++) {
    if (args[i] === '--path' && args[i + 1]) {
        targetPath = args[i + 1];
        break;
    }
}

if (targetPath === null) {
    console.error('Usage: node get_proof.js --path <path>');
    process.exit(1);
}

// Parse --index flag from command line arguments
let targetIndex = null;

for (let i = 0; i < args.length; i++) {
    if (args[i] === '--index' && args[i + 1]) {
        targetIndex = parseInt(args[i + 1]);
        break;
    }
}

if (targetIndex === null) {
    console.error('Usage: node get_proof.js --index <number>');
    process.exit(1);
}

const tree = StandardMerkleTree.load(JSON.parse(fs.readFileSync(targetPath, "utf8")));

for (const [i, v] of tree.entries()) {
    if (i === targetIndex) {
        const proof = tree.getProof(i);
        console.log('Proof:', proof);
        break;
    }
}
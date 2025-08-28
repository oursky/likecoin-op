// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

// https://www.alchemy.com/docs/create2-an-alternative-to-deriving-contract-addresses#factory-deploy-script
// - DeterministicDeployFactory.sol:DeterministicDeployFactory
abstract contract Create2able {
    function create2(
        bytes memory bytecode,
        uint salt
    ) public returns (address) {
        address addr;
        assembly {
            addr := create2(0, add(bytecode, 0x20), mload(bytecode), salt)
            if iszero(extcodesize(addr)) {
                revert(0, 0)
            }
        }
        return addr;
    }

    function _getBytecode(
        bytes memory bytecode,
        bytes memory arguments
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(bytecode, arguments);
    }
}

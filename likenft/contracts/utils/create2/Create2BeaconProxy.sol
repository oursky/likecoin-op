// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {Create2able} from "./Create2able.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";

contract Create2BeaconProxy is Create2able {
    function _create2BeaconProxy(
        address beacon,
        bytes memory initData,
        uint salt
    ) internal returns (address) {
        bytes memory bytecode = type(BeaconProxy).creationCode;
        bytes memory arguments = abi.encode(beacon, initData);
        return create2(_getBytecode(bytecode, arguments), salt);
    }
}

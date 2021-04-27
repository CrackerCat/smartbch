// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import "./interfaces/ISEP101.sol";

contract SEP101ProxyForGasTest is ISEP101 {

    bytes4 public constant _SELECTOR_SET = bytes4(keccak256(bytes("set(bytes,bytes)")));
    bytes4 public constant _SELECTOR_GET = bytes4(keccak256(bytes("get(bytes)")));

    address constant public agent = address(0x2712);

    function set(bytes calldata key, bytes calldata value) override external {
        agent.delegatecall(abi.encodeWithSelector(_SELECTOR_SET, key, value));
    }
    function get(bytes calldata key) override external returns (bytes memory) {
        agent.delegatecall(abi.encodeWithSelector(_SELECTOR_GET, key));
    }

}
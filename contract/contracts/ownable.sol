// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;

import "@openzeppelin/contracts/utils/Context.sol";

contract Ownable is Context {

  address payable owner;

  constructor() {
    owner = _msgSender();
  }

  modifier onlyOwner {
    require(
      _msgSender() == owner,
      "Only owner can call this function."
    );
    _;
  }
}
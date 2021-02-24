// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;

import "@openzeppelin/contracts/utils/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/EnumerableMap.sol";

contract ERC721 {
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableMap for EnumerableMap.UintToAddressMap;


  mapping (address => EnumerableSet.UintSet) private _holderTokens;
  EnumerableMap.UintToAddressMap private _tokenOwners;

  mapping (uint256 => address) private _tokenApprovals;

   // Token name
  string private _name;

  // Token symbol
  string private _symbol;

  // Optional mapping for token URIs
  mapping (uint256 => string) private _tokenURIs;

  // Base URI
  string private _baseURI;

  event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
  event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
  event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

  constructor (string memory name_, string memory symbol_) {
    _name = name_;
    _symbol = symbol_;
  }


  function _transfer(address from, address to, uint256 tokenId) internal virtual {
    require(ERC721.ownerOf(tokenId) == from, "ERC721: transfer of token that is not own"); // internal owner
    require(to != address(0), "ERC721: transfer to the zero address");
    // Clear approvals from the previous owner
    _approve(address(0), tokenId);

    _holderTokens[from].remove(tokenId);
    _holderTokens[to].add(tokenId);

    _tokenOwners.set(tokenId, to);

    emit Transfer(from, to, tokenId);
  }

}
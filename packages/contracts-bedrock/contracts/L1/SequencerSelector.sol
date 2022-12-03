// pragma solidity >= 0.8.0;
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import { Initializable } from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

/**
 * @title SequencerSelector
 * @notice The SequencerSelector contract is used to select sequencers on the L2 network.
 */
contract SequencerSelector is Initializable {
    uint256 public constant SEQUENCER_COUNT = 4;
    uint256 public startingBlockNumber; // The number of the first L2 block.
    string[] public sequencers; // The list of sequencers.

    event SequencerChanged (
        uint256 blockNumber,
        string current_sequencer_pubkey,
        string next_sequencer_pubkey
    );

    constructor(uint256 _startingBlockNumber) {
        sequencers = [
        "8a646f71cf55fabe61788b7989a125b0c543e679c28f07e704300c90f7b9c2d91f0a13e4f7741cd35a2b71ead34f3292",
        "ab05d6e17adab92a0252981331ea22c09cb33099135eba68913ca654709de594ca0b6316941add0139f452af884deb53",
        "af30a194a6b532703beb1b88081cc5b6f2dd357a739cd89ff400119d27d8d560a3492f86967ec967884bea47913995ac",
        "852efb403792fdb389b80c826608294834d5406b528c66964f480d8a86c0a561208a2815f0173f1a0b38962c30afbbaf"
        ];

        initialize(_startingBlockNumber);
    }

    /**
 * @notice Initializer.
     *
     * @param _startingBlockNumber Block number for the first recoded L2 block.
     */
    function initialize(uint256 _startingBlockNumber)
    public
    initializer
    {
        startingBlockNumber = _startingBlockNumber;
    }

    // genesisEpoch = 5
    // first (0) > 5
    // second (1) > 6
    // last (15) > 20
    // GetSequencer returns the pubkey of the current sequencer for the given epoch using round-robin strategy.
    function GetSequencer(uint256 blockNumber) public view returns (uint256) {
        uint256 offset = blockNumber - startingBlockNumber;
        uint256 index = offset % SEQUENCER_COUNT;
        return index;
    }

    // TODO: We currently don't allow new sequencers into the set.
    function AddSequencer() public returns (bool) {}
}

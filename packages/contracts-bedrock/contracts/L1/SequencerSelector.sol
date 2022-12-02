// pragma solidity >= 0.8.0;
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

/**
 * @title SequencerSelector
 * @notice The SequencerSelector contract is used to select sequencers on the L2 network.
 */
contract SequencerSelector {
    uint256 public constant SEQUENCER_COUNT = 16;
    uint256 public genesisEpoch; // The L2 epoch from which distributed sequencer selection started
    uint256[] public sequencerIDs; // The list of sequencerIDs.

    event SequencerChanged (
        uint256 epoch,
        uint256 current_sequencer,
        uint256 next_sequencer
    );

    constructor(uint256 epoch) {
        genesisEpoch = epoch;
        for (uint256 i = 0; i < SEQUENCER_COUNT; i++) {
            sequencerIDs.push(i);
        }
    }

    // genesisEpoch = 5
    // first (0) > 5
    // second (1) > 6
    // last (15) > 20
    // GetSequencer returns the current sequencer for the given epoch using round-robin strategy.
    function GetSequencer(uint256 epoch) public returns (uint) {
        uint256 offset = epoch - genesisEpoch;
        uint index = offset % SEQUENCER_COUNT;

        uint256 current_sequencerID = sequencerIDs[index];
        uint256 next_sequencerID = sequencerIDs[(index+1)%SEQUENCER_COUNT];
        emit SequencerChanged(
            epoch,
            current_sequencerID,
            next_sequencerID
        );

        return current_sequencerID;
    }

    // TODO: We currently don't allow new sequencers into the set.
    function AddSequencer() public returns (bool) {}
}

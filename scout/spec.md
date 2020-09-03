---
title: Scout
tags: spec, draft, scout
date: 25 august 2020
authors: Pro Gevorgyan, Ben Bollen
---

# Problem statement for Scout

The global view of the entire data stream processing is orchestrated
by the contracts on chain. As a node in the processing stream, we should
only concern ourselves with the events and data relevant to our computational
steps. This global view is indexed by the Landscape.
The scout subscribes to event updates given by the Landscape, with a policy
(initially just a filter) of which contract events are relevant for the node's
operation.

The scout reports back to the node (specifically the BoardsManager) on events
it received from the Landscape.

# Approach

# Learnings

# Notes
(1) Explorer: indexes all contracts (relating to Mosaic application),
    ie. the full circuit of boards and memberships, and MWL logs, etc
    Explorer does not care where "we"/"the node" operates inside this
    circuit.

(2) Scout: maintains a policy which defines the scope of the threads that are
    relevant to the operation of the boards the node is active on,
    ie. threads from fellow members to construct an (objective) view, or
    threads that are the input on which the board operates (from
    the preceding layer).

(3) Landscape: maintains an queryable datastore of the state of the contracts
    of interest and the chain (block height and blocktrie). The Explorer sends
    update events on the broadcaster for the chain and contracts it tracks.
    The Landscape persists these updates to a queryable datastore.
    Scouts can subscribe to the same event stream from the Landscape, with a filter.

(4) BoardsManager: maintains the active boards upon which the node operates.
    It relies on the Scout to receive updates from the board contracts,
    among others which sources (logIDs) are authorised the post messages,
    also, which boards the node is active on.

---
title: Boards package specification
tags: spec, draft, board
date: 25 aug 2020
authors: Pro Gevorgyan, Ben Bollen
---

# Boards package specification

A board is first an on-chain contract representation. The contract code governs
which sources (logIDs) have write access to the board.(1)

In the boards package boards are referenced with a BoardID.
The Threads network only has to concern itself with those threadIDs which are
locally relevant to any of the boards on which the node is active.
Boards relies on the Scout and the Landscape to be informed which threadIDs,
logIDs are authorized sources or inputs to a board. Hence, it only needs to have
a local data structure of the global data stream processing.

Each active board has a funnel which subscribes (without ThreadID filter) to the Threads network for new records.




# Notes
(1) Write access to a board is understood as the logID should be taken into
    consideration. It is possible for an arbitrary key (logID) to create a log
    which is encrypted with the servive key (sk) and read key (rk)
    of the same ThreadID, however, the mosaic node must ignore any logIDs which
    have not joined the board (or those who have logged out of the board).
    Board access can be guarded, eg. by putting forward a stake to join a column
    as a member, or opening a payment channel for usage payments as a user.

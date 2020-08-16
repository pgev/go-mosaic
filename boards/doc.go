/*
Package boards manages boards and associated threads for the node.

Mosaic boards build on the Threads protocol introduced by textilo.io and Mozilla.
We have tried to base our naming off of the december 2019 draft of the whitepaper,
augmented by the github.com/textileio/go-threads ongoing implementation.
The (last updated) whitepaper can be found here: https://docsend.com/view/gu3ywqi
and the archived december whitepaper is (not guaranteed) available on IPFS as https://ipfs.io/ipfs/QmRjgbB5pTxUnoLFbXtDHs1ph5t5pUhRnYn1wFiYEpP5s1?filename=201912TextileThreads.pdf

As a minimal summary, Threads is a protocol for a decentralised database which runs on IPFS,
which uses authenticated, append-only logs for synchronising state across
collaborating peers on a network.

Boards provide the following functionality. (draft version 0.1, august 2020)

1. A source of a board is authenticated to write a log (under a thread) on a board.

Because the stream processing is effectively a public process open to any
staked node to participate, the service (sk) and read key (rk) are effectively public too.
Rather mosaic defaults to assuming all threads are publicly readable, and
the key pair (sk, rk) functions as a data hygiene policy,
both separating unrelated data, and purging any CIDs the node lacks the sk for.
As a result, knowledge of the (sk, rk) is not sufficient to assume an Identity Key
writes a log the node should take into account.
An additional ACL layer (on top of the service and read key from threads)
is therefore needed to constrain which Identity Keys are permitted
to write on a board.

Access to write to a board is governed by a board smart contract on a chain, and
the scout will inform the board manager which Identity Keys participate on the board,
and other logs (even if encrypted with (sk, rk)) must be ignored when written by
non-participating Identity Keys.

def: a Source is a pair of an Identity Key and a BoardID, where the Identity Key is
authenticated to participate on the board.

2. the BoardManager tracks all the boards which the node needs to read from,
   and/or write to to function correctly within the circuit.

3. a board can contain one or more threads, each with a key pair (sk, rk).
   The BoardsManager internally manages threads, and the policy depends on the
   board type.
   For a gate board, each assigned column overlay has its own thread.
   For a databus board, all multi-writer-logs share one thread.
   For an input board, each end-user has their own thread.
   For an output board, all multi-writer-logs share one thread.
   Note, the implementation starts out simplified with one thread per board;
   later, threads should rotate over time as to rotate (sk, rk) pairs.

4. an Overlay groups sources on a board such that while valid participants of a board,
   sources can be filtered out. A source can be in multiple overlays.
   Messages from a source are passed on to the switch only if the source is in
   at least one overlay, or no overlays are set for a board.

5. the body of a log's record is a protobuf encoded (list of) message(s),
   where a message has a source, a topicID and a payload.
   (Note, the message will be be extended with additional payment information
   in the header).
   The payload is passed on to the reactor registered
   for the tuple (boardID, topicID), a "locus", which can additionally pass
   (an accumulation of processed) messages on to the application proxy.

(needs to be continued)

*/
package boards

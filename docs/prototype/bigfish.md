# Mosaic Prototype A: Big Fish

## Outset

Have a set of users represented by `logId`s.

Have a `column` with `members` which can:
- announce new users (inlcude an initial record of a user; user is given a fixed initial amount on the balance)
- update subjective view for new records of users
- publish subjective view (under batching conditions) to column members
- evaluate subjective views of members
    - try pull records of users if not present
- publish signature for preferred subjective views
- compute on all threshold-reached objective views (in txt msg app example, compute will simply update UI with messages received)
- publish signature for resulting batch to MWL (K_2)
- (there is no confirmation on-chain contract in this prototype)
- select result batch with "lowest" uint(cid of result batch) (prototype hack)
- rinse and repeat

Test: assert that all members robustly acquire the same result batch under various imperfect conditions

Have a state view in the column for all users

```
type ViewNode struct {
    ViewInfoNode
    PreviousViewNode
    Balance
    logHead
}
```

```
type ViewInfoNode struct {
    LogIdPublicKey
    PaymentChannelId (what is the identifier?)
}
```

K_1: type `avatar:logId` user
V_1: type `string` text message written by user

K_1:(view):V_1

## code breakdown

### Mosaic deamon

`go-mosaic/mosaicd/main.go`

hosts API to control Mosaic daemon on localhost `6500`

hosts ThreadsDB bind address on `4006`

runs public message board as demo example application

### Mosaic API client


### BigFish prototype client

example to broadcast unencrypted text messages to public group.

## notes

objective: all members end up seeing all the messages from all the users
    ie, also user threads which were unknown

two mains:

- member of column
    - initial member starts the column (column name);
    - or, on start join a column (get ma skrk from initial member);
    - add a user (multiaddrss servicekeyreadkey)

- chat app for end user
    - should start with new thread (user name)

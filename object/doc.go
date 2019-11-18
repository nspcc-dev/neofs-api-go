/*
Package object manages main storage structure in the system. All storage
operations are performed with the objects. During lifetime object might be
transformed into another object by cutting its payload or adding meta
information. All transformation may be reversed, therefore source object
will be able to restore.

Object structure

Object consists of Payload and Header. Payload is unlimited but storage nodes
may have a policy to store objects with a limited payload. In this case object
with large payload will be transformed into the chain of objects with small
payload.

Headers are simple key-value fields that divided into two groups: system
headers and extended headers. System headers contain information about
protocol version, object id, payload length in bytes, owner id, container id
and object creation timestamp (both in epochs and unix time). All these fields
must be set up in the correct object.

    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
    | System Headers                                                   |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
    | Version        : 1                                               |
    | Payload Length : 21673465                                        |
    | Object ID      : 465208e2-ba4f-4f99-ad47-82a59f4192d4            |
    | Owner ID       : AShvoCbSZ7VfRiPkVb1tEcBLiJrcbts1tt              |
    | Container ID   : FGobtRZA6sBZv2i9k4L7TiTtnuP6E788qa278xfj3Fxj    |
    | Created At     : Epoch#10, 1573033162                            |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
    | Extended Headers                                                 |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
    | User Header         : <user-defined-key>, <user-defined-value>   |
    | Verification Header : <session public key>, <owner's signature>  |
    | Homomorphic Hash    : 0x23d35a56ae...                            |
    | Payload Checksum    : 0x1bd34abs75...                            |
    | Integrity Header    : <header checksum>, <session signature>     |
    | Transformation      : Payload Split                              |
    | Link-parent         : cae08935-b4ba-499a-bf6c-98276c1e6c0b       |
    | Link-next           : c3b40fbf-3798-4b61-a189-2992b5fb5070       |
    | Payload Checksum    : 0x1f387a5c36...                            |
    | Integrity Header    : <header checksum>, <session signature>     |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
    | Payload                                                          |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
    | 0xd1581963a342d231...                                            |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-

There are different kinds of extended headers. A correct object must contain
verification header, homomorphic hash header, payload checksum and
integrity header. The order of headers is matter. Let's look through all
these headers.

Link header points to the connected objects. During object transformation, large
object might be transformed into the chain of smaller objects. One of these
objects drops payload and has several "Child" links. We call this object as
zero-object. Others will have "Parent" link to the zero-object, "Previous"
and "Next" links in the payload chain.

    [ Object ID:1 ] = > transformed
    `- [ Zero-Object ID:1 ]
        `- Link-child    ID:2
        `- Link-child    ID:3
        `- Link-child    ID:4
        `- Payload [null]
    `- [ Object ID:2 ]
        `- Link-parent   ID:1
        `- Link-next     ID:3
        `- Payload [ 0x13ba... ]
    `- [ Object ID:3 ]
        `- Link-parent   ID:1
        `- Link-previous ID:2
        `- Link-next     ID:4
        `- Payload [ 0xcd34... ]
    `- [ Object ID:4 ]
        `- Link-parent   ID:1
        `- Link-previous ID:3
        `- Payload [ 0xef86... ]

Storage groups are also objects. They have "Storage Group" links to all
objects in the group. Links are set by nodes during transformations and,
in general, they should not be set by user manually.

Redirect headers are not used yet, they will be implemented and described
later.

User header is a key-value pair of string that can be defined by user. User
can use these headers as search attribute. You can store any meta information
about object there, e.g. object's nicename.

Transformation header notifies that object was transformed by some pre-defined
way. This header sets up before object is transformed and all headers after
transformation must be located after transformation header. During reverse
transformation, all headers under transformation header will be cut out.

     +-+-+-+-+-+-+-+-+-+-       +-+-+-+-+-+-+-+-+-+-+     +-+-+-+-+-+-+-+-+-+-+
     | Payload checksum |       | Payload checksum  |     | Payload checksum  |
     | Integrity header |   =>  | Integrity header  |  +  | Integrity header  |
     +-+-+-+-+-+-+-+-+-+-       | Transformation    |     | Transformation    |
     | Large payload    |       | New Checksum      |     | New Checksum      |
     +-+-+-+-+-+-+-+-+-+-       | New Integrity     |     | New Integrity     |
                                +-+-+-+-+-+-+-+-+-+-+     +-+-+-+-+-+-+-+-+-+-+
                                | Small payload     |     | Small payload     |
                                +-+-+-+-+-+-+-+-+-+-+     +-+-+-+-+-+-+-+-+-+-+

For now, we use only one type of transformation: payload split transformation.
This header set up by node automatically.

Tombstone header notifies that this object was deleted by user. Objects with
tombstone header do not have payload, but they still contain meta information
in the headers. This way we implement two-phase commit for object removal.
Storage nodes will eventually delete all tombstone objects. If you want to
delete object, you must create new object with the same object id, with
tombstone header, correct signatures and without payload.

Verification header contains session information. To put the object in
the system user must create session. It is required because objects might
be transformed and therefore must be re-signed. To do that node creates
a pair of session public and private keys. Object owner delegates permission to
re-sign objects by signing session public key. This header contains session
public key and owner's signature of this key. You must specify this header
manually.

Homomorphic hash header contains homomorphic hash of the source object.
Transformations do not affect this header. This header used by data audit and
set by node automatically.

Payload checksum contains checksum of the actual object payload. All payload
transformation must set new payload checksum headers. This header set by node
automatically.

Integrity header contains checksum of the header and signature of the
session key. This header must be last in the list of extended headers.
Checksum is calculated by marshaling all above headers, including system
headers. This header set by node automatically.

Storage group header is presented in storage group objects. It contains
information for data audit: size of validated data, homomorphic has of this
data, storage group expiration time in epochs or unix time.


*/
package object

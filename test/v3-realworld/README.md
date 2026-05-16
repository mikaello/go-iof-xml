# Real-world IOF v3 examples

These XML files are taken from real orienteering events and event tools,
sourced from
[mikaello/iof-list-manipulator](https://github.com/mikaello/iof-list-manipulator/tree/main/examples).

Unlike the spec examples in `../v3-examples`, these are produced by real
software and exercise edge cases that the canonical IOF examples skip,
e.g. timezone-less `<Time>` elements and UTF-8 byte-order marks. They are
covered by `TestDecodeRealWorldExamples` in `pkg/marshallers`.

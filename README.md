# bitquery-protobuf-schema

bitquery-protobuf-schema is an NPM package by Bitquery that returns Protocol Buffers (Protobuf) schemas for the Bitquery Protobuf Kafka streams by providing the topic name. The main purpose of this package is to avoid the hastle of downloading the `.proto` files again and again.

## Installation

To install the package, use npm:

``` sh
npm install streaming_protobuf
```

## Usage

The usage for the following package is as follows.

``` js
    let ParsedIdlBlockMessage;
    let topic = '<topic>';

    ParsedIdlBlockMessage = await loadProto(topic); 
```

This `ParsedIdlBlockMessage` variable can then be used for decoding the recieved Protobuf message as shown below, where message is a singular message recieved from the Bitquery Protobuf stream.

``` js
    const buffer = message.value;
    const decoded = ParsedIdlBlockMessage.decode(buffer);
```

To get started with the Bitquery Protobuf Kaka streams contact - sales@bitquery.io. 

To read more about Bitquery Protobuf Kafka solutions checkout their official [documentation](https://docs.bitquery.io/docs/streams/kafka-streaming-concepts/).

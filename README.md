# bitquery-protobuf-schema

**bitquery-protobuf-schema** is an NPM package by **Bitquery**, designed to simplify working with **on-chain data streaming** pipelines powered by **Protobuf** and **Kafka**.
This package automatically returns Protocol Buffer (`.proto`) schemas for **Bitquery’s Protobuf Kafka streams** when you provide the stream's topic name.

Bitquery is a leading **on-chain data provider**, offering blockchain intelligence via **GraphQL APIs**, **WebSocket streaming**, and **high-performance Kafka streams**.
This package removes the hassle of repeatedly downloading `.proto` files manually and allows developers to integrate directly with Bitquery’s **real-time blockchain data streams**.

---

## 🚀 Features

* Easy loading of **Protobuf schemas** for Bitquery Kafka stream topics
* Ideal for applications using **real-time streaming on-chain data**
* Automatically decodes **blockchain events**, **transactions**, **logs**, **shreds**, and more
* Useful for building analytics dashboards, trading bots, monitoring tools, and ingestion pipelines
* Zero need to manually store or update `.proto` files

---

## 📦 Installation

Install with npm:

```sh
npm i bitquery-protobuf-schema
```

---

## 🧠 Usage

Here's how to load a Protobuf schema for any **Bitquery Kafka stream topic**:

```js
const { loadProto } = require('bitquery-protobuf-schema');

let ParsedIdlBlockMessage;
let topic = '<topic>';

ParsedIdlBlockMessage = await loadProto(topic);
```

Once loaded, use the parsed schema to **decode incoming Protobuf messages** from the Kafka stream:

```js
const buffer = message.value;
const decoded = ParsedIdlBlockMessage.decode(buffer);
```

This gives you fully structured **on-chain data** decoded directly from Bitquery's **Protobuf streaming infrastructure**.

---

## 📡 Access Bitquery Protobuf Kafka Streams

To start streaming real-time blockchain data using Protobuf + Kafka, contact:
📧 **[sales@bitquery.io](mailto:sales@bitquery.io)**

---

## 📚 Documentation

Explore Bitquery’s on-chain data streaming docs:

* [**Kafka Streaming Concepts**](https://docs.bitquery.io/docs/streams/kafka-streaming-concepts/?utm_source=github&utm_medium=npm&utm_campaign=proto_schema)

* [**Bitcoin Protobuf Streams**](https://docs.bitquery.io/docs/streams/protobuf/chains/Bitcoin-protobuf/?utm_source=github&utm_medium=npm&utm_campaign=proto_schema)

* [**EVM Protobuf Streams**](https://docs.bitquery.io/docs/streams/protobuf/chains/EVM-protobuf/?utm_source=github&utm_medium=npm&utm_campaign=proto_schema)

* [**Solana Shreds Stream (High-performance Solana data)**](https://docs.bitquery.io/docs/streams/protobuf/chains/Solana-protobuf/?utm_source=github&utm_medium=npm&utm_campaign=proto_schema)

* [**Tron Protobuf Stream**](https://docs.bitquery.io/docs/streams/protobuf/chains/Tron-protobuf/?utm_source=github&utm_medium=npm&utm_campaign=proto_schema)

* [**Trading Bot Tutorial (Kafka Sniper Bot)**](https://docs.bitquery.io/docs/streams/sniper-trade-using-bitquery-kafka-stream/?utm_source=github&utm_medium=npm&utm_campaign=proto_schema)

---

## 🏁 Summary

`bitquery-protobuf-schema` helps developers decode **streaming, real-time on-chain data** from Bitquery’s **Kafka Protobuf infrastructure** without managing `.proto` files manually.
Whether you're building blockchain analytics, DeFi trading bots, monitoring systems, or ingestion pipelines, this package streamlines your integration with Bitquery’s high-performance data streams.

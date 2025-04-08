const protobuf = require('protobufjs');
const path = require('path');

const loadProto = async (topic) => {
    
    let evmFiles = [
        'evm/dex_block_message.proto', 
        'evm/token_block_message.proto', 
        'evm/block_message.proto', 
        'evm/parsed_abi_block_message.proto'
    ];

    let evmMessages = [
        'evm_messages.DexBlockMessage', 
        'evm_messages.TokenBlockMessage',
        'evm_messages.BlockMessage',
        'evm_messages.ParsedAbiBlockMessage',
    ];

    let tronFiles = [
        'tron/dex_block_message.proto', 
        'tron/token_block_message.proto', 
        'tron/block_message.proto', 
        'tron/parsed_abi_block_message.proto'
    ]

    let tronMessages = [
        'tron_messages.DexBlockMessage', 
        'tron_messages.TokenBlockMessage',
        'tron_messages.BlockMessage',
        'tron_messages.ParsedAbiBlockMessage',
    ];

    let topicToPath = {
        'bsc.dextrades.proto': evmFiles[0],
        'bsc.tokens.proto': evmFiles[1],
        'bsc.raw.proto': evmFiles[2],
        'bsc.transactions.proto': evmFiles[3],
        'bsc.broadcasted.dextrades.proto': evmFiles[0],
        'bsc.broadcasted.tokens.proto': evmFiles[1],
        'bsc.broadcasted.raw.proto': evmFiles[2],
        'bsc.broadcasted.transactions.proto': evmFiles[3],
        'eth.dextrades.proto': evmFiles[0],
        'eth.tokens.proto': evmFiles[1],
        'eth.raw.proto': evmFiles[2],
        'eth.transactions.proto': evmFiles[3],
        'eth.broadcasted.dextrades.proto': evmFiles[0],
        'eth.broadcasted.tokens.proto': evmFiles[1],
        'eth.broadcasted.raw.proto': evmFiles[2],
        'eth.broadcasted.transactions.proto': evmFiles[3],
        'base.dextrades.proto': evmFiles[0],
        'base.tokens.proto': evmFiles[1],
        'base.raw.proto': evmFiles[2],
        'base.transactions.proto': evmFiles[3],
        'solana.transactions.proto': 'solana/parsed_idl_block_message.proto',
        'solana.tokens.proto': 'solana/token_block_message.proto',
        'solana.dextrades.proto': 'solana/dex_block_message.proto',
        'tron.dextrades.proto': tronFiles[0],
        'tron.tokens.proto': tronFiles[1],
        'tron.raw.proto': tronFiles[2],
        'tron.transactions.proto': tronFiles[3],
        'tron.broadcasted.dextrades.proto': tronFiles[0],
        'tron.broadcasted.tokens.proto': tronFiles[1],
        'tron.broadcasted.raw.proto': tronFiles[2],
        'tron.broadcasted.transactions.proto': tronFiles[3],
    }

    let topicToMessage = {
        'bsc.dextrades.proto': evmMessages[0],
        'bsc.tokens.proto': evmMessages[1],
        'bsc.raw.proto': evmMessages[2],
        'bsc.transactions.proto': evmMessages[3],
        'bsc.broadcasted.dextrades.proto': evmMessages[0],
        'bsc.broadcasted.tokens.proto': evmMessages[1],
        'bsc.broadcasted.raw.proto': evmMessages[2],
        'bsc.broadcasted.transactions.proto': evmMessages[3],
        'eth.dextrades.proto': evmMessages[0],
        'eth.tokens.proto': evmMessages[1],
        'eth.raw.proto': evmMessages[2],
        'eth.transactions.proto': evmMessages[3],
        'eth.broadcasted.dextrades.proto': evmMessages[0],
        'eth.broadcasted.tokens.proto': evmMessages[1],
        'eth.broadcasted.raw.proto': evmMessages[2],
        'eth.broadcasted.transactions.proto': evmMessages[3],
        'base.dextrades.proto': evmMessages[0],
        'base.tokens.proto': evmMessages[1],
        'base.raw.proto': evmMessages[2],
        'base.transactions.proto': evmMessages[3],
        'solana.transactions.proto': 'solana_messages.ParsedIdlBlockMessage',
        'solana.tokens.proto': 'solana_messages.TokenBlockMessage',
        'solana.dextrades.proto': 'solana_messages.DexParsedBlockMessage',
        'tron.dextrades.proto': tronMessages[0],
        'tron.tokens.proto': tronMessages[1],
        'tron.raw.proto': tronMessages[2],
        'tron.transactions.proto': tronMessages[3],
        'tron.broadcasted.dextrades.proto': tronMessages[0],
        'tron.broadcasted.tokens.proto': tronMessages[1],
        'tron.broadcasted.raw.proto': tronMessages[2],
        'tron.broadcasted.transactions.proto': tronMessages[3],
    }

    const filePath = path.join(__dirname, topicToPath[topic])
    const root = await protobuf.load(filePath);
    return root.lookupType(topicToMessage[topic]);
}; 

module.exports = {loadProto};
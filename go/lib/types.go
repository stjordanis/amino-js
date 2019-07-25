package lib

const (
	TendermintPrivKeyLedgerSecp256k1             = "tendermint/PrivKeyLedgerSecp256k1"
	CryptoKeysHdBIP44Params                      = "crypto/keys/hd/BIP44Params"
	CryptoKeysLocalInfo                          = "crypto/keys/localInfo"
	CryptoKeysLedgerInfo                         = "crypto/keys/ledgerInfo"
	CryptoKeysOfflineInfo                        = "crypto/keys/offlineInfo"
	CryptoKeysMultiInfo                          = "crypto/keys/multiInfo"
	AuthAccount                                  = "auth/Account"
	AuthBaseVestingAccount                       = "auth/BaseVestingAccount"
	AuthContinuousVestingAccount                 = "auth/ContinuousVestingAccount"
	AuthDelayedVestingAccount                    = "auth/DelayedVestingAccount"
	AuthStdTx                                    = "auth/StdTx"
	CosmosSdkBaseAccount                         = "cosmos-sdk/BaseAccount"
	CosmosSdkBaseVestingAccount                  = "cosmos-sdk/BaseVestingAccount"
	CosmosSdkContinuousVestingAccount            = "cosmos-sdk/ContinuousVestingAccount"
	CosmosSdkDelayedVestingAccount               = "cosmos-sdk/DelayedVestingAccount"
	TerraMsgSend                                 = "pay/MsgSend"
	TerraMsgMultiSend                            = "pay/MsgMultiSend"
	TerraMsgVerifyInvariant                      = "crisis/MsgVerifyInvariant"
	TerraMsgWithdrawDelegationReward             = "distribution/MsgWithdrawDelegationReward"
	TerraMsgWithdrawValidatorCommission          = "distribution/MsgWithdrawValidatorCommission"
	TerraMsgModifyWithdrawAddress                = "distribution/MsgModifyWithdrawAddress"
	TerraMsgSubmitProposal                       = "gov/MsgSubmitProposal"
	TerraMsgDeposit                              = "gov/MsgDeposit"
	TerraMsgVote                                 = "gov/MsgVote"
	TerraTextProposal                            = "gov/TextProposal"
	TerraSoftwareUpgradeProposal                 = "gov/SoftwareUpgradeProposal"
	TerraParameterChangeProposal                 = "gov/ParameterChangeProposal"
	TerraMsgIBCTransfer                          = "ibc/MsgIBCTransfer"
	TerraMsgIBCReceive                           = "ibc/MsgIBCReceive"
	TerraMsgUnjail                               = "staking/MsgUnjail"
	TerraMsgCreateValidator                      = "staking/MsgCreateValidator"
	TerraMsgEditValidator                        = "staking/MsgEditValidator"
	TerraMsgDelegate                             = "staking/MsgDelegate"
	TerraMsgUndelegate                           = "staking/MsgUndelegate"
	TerraMsgBeginRedelegate                      = "staking/MsgBeginRedelegate"
	TendermintBlockchainBlockRequest             = "tendermint/blockchain/BlockRequest"
	TendermintBlockchainBlockResponse            = "tendermint/blockchain/BlockResponse"
	TendermintBlockchainNoBlockResponse          = "tendermint/blockchain/NoBlockResponse"
	TendermintBlockchainStatusResponse           = "tendermint/blockchain/StatusResponse"
	TendermintBlockchainStatusRequest            = "tendermint/blockchain/StatusRequest"
	TendermintNewRoundStepMessage                = "tendermint/NewRoundStepMessage"
	TendermintNewValidBlockMessage               = "tendermint/NewValidBlockMessage"
	TendermintProposal                           = "tendermint/Proposal"
	TendermintProposalPOL                        = "tendermint/ProposalPOL"
	TendermintBlockPart                          = "tendermint/BlockPart"
	TendermintVote                               = "tendermint/Vote"
	TendermintHasVote                            = "tendermint/HasVote"
	TendermintVoteSetMaj23                       = "tendermint/VoteSetMaj23"
	TendermintVoteSetBits                        = "tendermint/VoteSetBits"
	TendermintWalEventDataRoundState             = "tendermint/wal/EventDataRoundState"
	TendermintWalMsgInfo                         = "tendermint/wal/MsgInfo"
	TendermintWalTimeoutInfo                     = "tendermint/wal/TimeoutInfo"
	TendermintWalEndHeightMessage                = "tendermint/wal/EndHeightMessage"
	TendermintPubKeyEd25519                      = "tendermint/PubKeyEd25519"
	TendermintPrivKeyEd25519                     = "tendermint/PrivKeyEd25519"
	TendermintPubKeySecp256k1                    = "tendermint/PubKeySecp256k1"
	TendermintPrivKeySecp256k1                   = "tendermint/PrivKeySecp256k1"
	TendermintPubKeyMultisigThreshold            = "tendermint/PubKeyMultisigThreshold"
	TendermintEvidenceEvidenceListMessage        = "tendermint/evidence/EvidenceListMessage"
	TendermintMempoolTxMessage                   = "tendermint/mempool/TxMessage"
	TendermintP2pPacketPing                      = "tendermint/p2p/PacketPing"
	TendermintP2pPacketPong                      = "tendermint/p2p/PacketPong"
	TendermintP2pPacketMsg                       = "tendermint/p2p/PacketMsg"
	TendermintP2pPexRequestMessage               = "tendermint/p2p/PexRequestMessage"
	TendermintP2pPexAddrsMessage                 = "tendermint/p2p/PexAddrsMessage"
	TendermintRemotesignerPubKeyRequest          = "tendermint/remotesigner/PubKeyRequest"
	TendermintRemotesignerPubKeyResponse         = "tendermint/remotesigner/PubKeyResponse"
	TendermintRemotesignerSignVoteRequest        = "tendermint/remotesigner/SignVoteRequest"
	TendermintRemotesignerSignedVoteResponse     = "tendermint/remotesigner/SignedVoteResponse"
	TendermintRemotesignerSignProposalRequest    = "tendermint/remotesigner/SignProposalRequest"
	TendermintRemotesignerSignedProposalResponse = "tendermint/remotesigner/SignedProposalResponse"
	TendermintRemotesignerPingRequest            = "tendermint/remotesigner/PingRequest"
	TendermintRemotesignerPingResponse           = "tendermint/remotesigner/PingResponse"
	TendermintEventNewBlock                      = "tendermint/event/NewBlock"
	TendermintEventNewBlockHeader                = "tendermint/event/NewBlockHeader"
	TendermintEventTx                            = "tendermint/event/Tx"
	TendermintEventRoundState                    = "tendermint/event/RoundState"
	TendermintEventNewRound                      = "tendermint/event/NewRound"
	TendermintEventCompleteProposal              = "tendermint/event/CompleteProposal"
	TendermintEventVote                          = "tendermint/event/Vote"
	TendermintEventValidatorSetUpdates           = "tendermint/event/ValidatorSetUpdates"
	TendermintEventProposalString                = "tendermint/event/ProposalString"
	TendermintDuplicateVoteEvidence              = "tendermint/DuplicateVoteEvidence"
	TendermintMockGoodEvidence                   = "tendermint/MockGoodEvidence"
	TendermintMockRandomGoodEvidence             = "tendermint/MockRandomGoodEvidence"
	TendermintMockBadEvidence                    = "tendermint/MockBadEvidence"
)

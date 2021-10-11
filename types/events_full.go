package types

type EventAssetsApprovalCancelled struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventAssetsApprovedTransfer struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Who1   AccountID
	Amount U128
	Topics []Hash
}

type EventAssetsAssetFrozen struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventAssetsAssetStatusChanged struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventAssetsAssetThawed struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventAssetsBurned struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventAssetsCreated struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventAssetsDestroyed struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventAssetsForceCreated struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Topics []Hash
}

type EventAssetsFrozen struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Topics []Hash
}

type EventAssetsIssued struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventAssetsMetadataCleared struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventAssetsMetadataSet struct {
	Phase  Phase
	Index  U32
	Bytes  Bytes
	Bytes1 Bytes
	U8     U8
	Flag   bool
	Topics []Hash
}

type EventAssetsOwnerChanged struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Topics []Hash
}

type EventAssetsTeamChanged struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Who1   AccountID
	Who2   AccountID
	Topics []Hash
}

type EventAssetsThawed struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Topics []Hash
}

type EventAssetsTransferred struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Who1   AccountID
	Amount U128
	Topics []Hash
}

type EventAssetsTransferredApproved struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Who1   AccountID
	Who2   AccountID
	Amount U128
	Topics []Hash
}

type EventBagslistRebagged struct {
	Phase   Phase
	Who     AccountID
	Amount  U128
	Amount1 U128
	Topics  []Hash
}

type EventBountiesBountyAwarded struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Topics []Hash
}

type EventBountiesBountyBecameActive struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventBountiesBountyCanceled struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventBountiesBountyClaimed struct {
	Phase  Phase
	Index  U32
	Amount U128
	Who    AccountID
	Topics []Hash
}

type EventBountiesBountyExtended struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventBountiesBountyProposed struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventBountiesBountyRejected struct {
	Phase  Phase
	Index  U32
	Amount U128
	Topics []Hash
}

type EventContractsCodeRemoved struct {
	Phase  Phase
	Hash   Hash
	Topics []Hash
}

type EventContractsContractEmitted struct {
	Phase  Phase
	Who    AccountID
	Bytes  Bytes
	Topics []Hash
}

type EventContractsTerminated struct {
	Phase  Phase
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventCouncilApproved struct {
	Phase  Phase
	Hash   Hash
	Topics []Hash
}

type EventCouncilClosed struct {
	Phase  Phase
	Hash   Hash
	Index  U32
	Index1 U32
	Topics []Hash
}

type EventCouncilDisapproved struct {
	Phase  Phase
	Hash   Hash
	Topics []Hash
}

type EventCouncilExecuted struct {
	Phase        Phase
	Hash         Hash
	CustomStruct DispatchResult
	Topics       []Hash
}

type EventCouncilMemberExecuted struct {
	Phase        Phase
	Hash         Hash
	CustomStruct DispatchResult
	Topics       []Hash
}

type EventCouncilProposed struct {
	Phase  Phase
	Who    AccountID
	Index  U32
	Hash   Hash
	Index1 U32
	Topics []Hash
}

type EventCouncilVoted struct {
	Phase  Phase
	Who    AccountID
	Hash   Hash
	Flag   bool
	Index  U32
	Index1 U32
	Topics []Hash
}

type EventElectionProviderMultiphaseElectionFinalized struct {
	Phase  Phase
	Option OptionElectionCompute
	Topics []Hash
}

type EventElectionProviderMultiphaseRewarded struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventElectionProviderMultiphaseSignedPhaseStarted struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventElectionProviderMultiphaseSlashed struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventElectionProviderMultiphaseSolutionStored struct {
	Phase        Phase
	CustomStruct ElectionCompute
	Flag         bool
	Topics       []Hash
}

type EventElectionProviderMultiphaseUnsignedPhaseStarted struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventElectionsCandidateSlashed struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventElectionsRenounced struct {
	Phase  Phase
	Who    AccountID
	Topics []Hash
}

type EventElectionsSeatHolderSlashed struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventGiltBidPlaced struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Index  U32
	Topics []Hash
}

type EventGiltBidRetracted struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Index  U32
	Topics []Hash
}

type EventGiltGiltIssued struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventGiltGiltThawed struct {
	Phase   Phase
	Index   U32
	Who     AccountID
	Amount  U128
	Amount1 U128
	Topics  []Hash
}

type EventIdentityIdentityCleared struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventIdentityIdentityKilled struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventIdentityIdentitySet struct {
	Phase  Phase
	Who    AccountID
	Topics []Hash
}

type EventImonlineAllGood struct {
	Phase  Phase
	Topics []Hash
}

type EventLotteryCallsUpdated struct {
	Phase  Phase
	Topics []Hash
}

type EventLotteryLotteryStarted struct {
	Phase  Phase
	Topics []Hash
}

type EventLotteryTicketBought struct {
	Phase     Phase
	Who       AccountID
	NewStruct struct {
		A1 U8
		A2 U8
	}

	Topics []Hash
}

type EventLotteryWinner struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventProxyProxyAdded struct {
	Phase        Phase
	Who          AccountID
	Who1         AccountID
	CustomStruct ProxyType
	Index        U32
	Topics       []Hash
}

type EventRecoveryRecoveryClosed struct {
	Phase  Phase
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventRecoveryRecoveryCreated struct {
	Phase  Phase
	Who    AccountID
	Topics []Hash
}

type EventRecoveryRecoveryInitiated struct {
	Phase  Phase
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventRecoveryRecoveryRemoved struct {
	Phase  Phase
	Who    AccountID
	Topics []Hash
}

type EventRecoveryRecoveryVouched struct {
	Phase  Phase
	Who    AccountID
	Who1   AccountID
	Who2   AccountID
	Topics []Hash
}

type EventStakingChilled struct {
	Phase  Phase
	Who    AccountID
	Topics []Hash
}

type EventStakingEraPaid struct {
	Phase   Phase
	Index   U32
	Amount  U128
	Amount1 U128
	Topics  []Hash
}

type EventStakingKicked struct {
	Phase  Phase
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventStakingPayoutStarted struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Topics []Hash
}

type EventStakingRewarded struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventStakingSlashed struct {
	Phase  Phase
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventStakingStakersElected struct {
	Phase  Phase
	Topics []Hash
}

type EventStakingStakingElectionFailed struct {
	Phase  Phase
	Topics []Hash
}

type EventSudoSudoAsDone struct {
	Phase        Phase
	CustomStruct DispatchResult
	Topics       []Hash
}

type EventSystemRemarked struct {
	Phase  Phase
	Who    AccountID
	Hash   Hash
	Topics []Hash
}

type EventTechnicalcommitteeApproved struct {
	Phase  Phase
	Hash   Hash
	Topics []Hash
}

type EventTechnicalcommitteeClosed struct {
	Phase  Phase
	Hash   Hash
	Index  U32
	Index1 U32
	Topics []Hash
}

type EventTechnicalcommitteeDisapproved struct {
	Phase  Phase
	Hash   Hash
	Topics []Hash
}

type EventTechnicalcommitteeExecuted struct {
	Phase        Phase
	Hash         Hash
	CustomStruct DispatchResult
	Topics       []Hash
}

type EventTechnicalcommitteeMemberExecuted struct {
	Phase        Phase
	Hash         Hash
	CustomStruct DispatchResult
	Topics       []Hash
}

type EventTechnicalcommitteeProposed struct {
	Phase  Phase
	Who    AccountID
	Index  U32
	Hash   Hash
	Index1 U32
	Topics []Hash
}

type EventTechnicalcommitteeVoted struct {
	Phase  Phase
	Who    AccountID
	Hash   Hash
	Flag   bool
	Index  U32
	Index1 U32
	Topics []Hash
}

type EventTechnicalmembershipDummy struct {
	Phase  Phase
	Topics []Hash
}

type EventTechnicalmembershipKeyChanged struct {
	Phase  Phase
	Topics []Hash
}

type EventTechnicalmembershipMemberAdded struct {
	Phase  Phase
	Topics []Hash
}

type EventTechnicalmembershipMemberRemoved struct {
	Phase  Phase
	Topics []Hash
}

type EventTechnicalmembershipMembersReset struct {
	Phase  Phase
	Topics []Hash
}

type EventTechnicalmembershipMembersSwapped struct {
	Phase  Phase
	Topics []Hash
}

type EventTipsNewTip struct {
	Phase  Phase
	Hash   Hash
	Topics []Hash
}

type EventTipsTipClosed struct {
	Phase  Phase
	Hash   Hash
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventTipsTipClosing struct {
	Phase  Phase
	Hash   Hash
	Topics []Hash
}

type EventTipsTipRetracted struct {
	Phase  Phase
	Hash   Hash
	Topics []Hash
}

type EventTipsTipSlashed struct {
	Phase  Phase
	Hash   Hash
	Who    AccountID
	Amount U128
	Topics []Hash
}

type EventTransactionstorageProofChecked struct {
	Phase  Phase
	Topics []Hash
}

type EventTransactionstorageRenewed struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventTransactionstorageStored struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventUniquesApprovalCancelled struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventUniquesApprovedTransfer struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventUniquesAssetStatusChanged struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventUniquesAttributeCleared struct {
	Phase  Phase
	Index  U32
	Option OptionU32
	Bytes  Bytes
	Topics []Hash
}

type EventUniquesAttributeSet struct {
	Phase  Phase
	Index  U32
	Option OptionU32
	Bytes  Bytes
	Bytes1 Bytes
	Topics []Hash
}

type EventUniquesBurned struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Who    AccountID
	Topics []Hash
}

type EventUniquesClassFrozen struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventUniquesClassMetadataCleared struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventUniquesClassMetadataSet struct {
	Phase  Phase
	Index  U32
	Bytes  Bytes
	Flag   bool
	Topics []Hash
}

type EventUniquesClassThawed struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventUniquesCreated struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventUniquesDestroyed struct {
	Phase  Phase
	Index  U32
	Topics []Hash
}

type EventUniquesForceCreated struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Topics []Hash
}

type EventUniquesFrozen struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Topics []Hash
}

type EventUniquesIssued struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Who    AccountID
	Topics []Hash
}

type EventUniquesMetadataCleared struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Topics []Hash
}

type EventUniquesMetadataSet struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Bytes  Bytes
	Flag   bool
	Topics []Hash
}

type EventUniquesOwnerChanged struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Topics []Hash
}

type EventUniquesRedeposited struct {
	Phase  Phase
	Index  U32
	Vec    []U32
	Topics []Hash
}

type EventUniquesTeamChanged struct {
	Phase  Phase
	Index  U32
	Who    AccountID
	Who1   AccountID
	Who2   AccountID
	Topics []Hash
}

type EventUniquesThawed struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Topics []Hash
}

type EventUniquesTransferred struct {
	Phase  Phase
	Index  U32
	Index1 U32
	Who    AccountID
	Who1   AccountID
	Topics []Hash
}

type EventUtilityItemCompleted struct {
	Phase  Phase
	Topics []Hash
}

//type EventRecords struct {
//	Assets_ApprovalCancelled                        []EventAssetsApprovalCancelled
//	Assets_ApprovedTransfer                         []EventAssetsApprovedTransfer
//	Assets_AssetFrozen                              []EventAssetsAssetFrozen
//	Assets_AssetStatusChanged                       []EventAssetsAssetStatusChanged
//	Assets_AssetThawed                              []EventAssetsAssetThawed
//	Assets_Burned                                   []EventAssetsBurned
//	Assets_Created                                  []EventAssetsCreated
//	Assets_Destroyed                                []EventAssetsDestroyed
//	Assets_ForceCreated                             []EventAssetsForceCreated
//	Assets_Frozen                                   []EventAssetsFrozen
//	Assets_Issued                                   []EventAssetsIssued
//	Assets_MetadataCleared                          []EventAssetsMetadataCleared
//	Assets_MetadataSet                              []EventAssetsMetadataSet
//	Assets_OwnerChanged                             []EventAssetsOwnerChanged
//	Assets_TeamChanged                              []EventAssetsTeamChanged
//	Assets_Thawed                                   []EventAssetsThawed
//	Assets_Transferred                              []EventAssetsTransferred
//	Assets_TransferredApproved                      []EventAssetsTransferredApproved
//	Bagslist_Rebagged                               []EventBagslistRebagged
//	Balances_BalanceSet                             []EventBalancesBalanceSet
//	Balances_Deposit                                []EventBalancesDeposit
//	Balances_DustLost                               []EventBalancesDustLost
//	Balances_Endowed                                []EventBalancesEndowed
//	Balances_Reserved                               []EventBalancesReserved
//	Balances_ReserveRepatriated                     []EventBalancesReserveRepatriated
//	Balances_Transfer                               []EventBalancesTransfer
//	Balances_Unreserved                             []EventBalancesUnreserved
//	Bounties_BountyAwarded                          []EventBountiesBountyAwarded
//	Bounties_BountyBecameActive                     []EventBountiesBountyBecameActive
//	Bounties_BountyCanceled                         []EventBountiesBountyCanceled
//	Bounties_BountyClaimed                          []EventBountiesBountyClaimed
//	Bounties_BountyExtended                         []EventBountiesBountyExtended
//	Bounties_BountyProposed                         []EventBountiesBountyProposed
//	Bounties_BountyRejected                         []EventBountiesBountyRejected
//	Contracts_CodeRemoved                           []EventContractsCodeRemoved
//	Contracts_CodeStored                            []EventContractsCodeStored
//	Contracts_ContractEmitted                       []EventContractsContractEmitted
//	Contracts_Instantiated                          []EventContractsInstantiated
//	Contracts_ScheduleUpdated                       []EventContractsScheduleUpdated
//	Contracts_Terminated                            []EventContractsTerminated
//	Council_Approved                                []EventCouncilApproved
//	Council_Closed                                  []EventCouncilClosed
//	Council_Disapproved                             []EventCouncilDisapproved
//	Council_Executed                                []EventCouncilExecuted
//	Council_MemberExecuted                          []EventCouncilMemberExecuted
//	Council_Proposed                                []EventCouncilProposed
//	Council_Voted                                   []EventCouncilVoted
//	Democracy_Blacklisted                           []EventDemocracyBlacklisted
//	Democracy_Cancelled                             []EventDemocracyCancelled
//	Democracy_Delegated                             []EventDemocracyDelegated
//	Democracy_Executed                              []EventDemocracyExecuted
//	Democracy_ExternalTabled                        []EventDemocracyExternalTabled
//	Democracy_NotPassed                             []EventDemocracyNotPassed
//	Democracy_Passed                                []EventDemocracyPassed
//	Democracy_PreimageInvalid                       []EventDemocracyPreimageInvalid
//	Democracy_PreimageMissing                       []EventDemocracyPreimageMissing
//	Democracy_PreimageNoted                         []EventDemocracyPreimageNoted
//	Democracy_PreimageReaped                        []EventDemocracyPreimageReaped
//	Democracy_PreimageUsed                          []EventDemocracyPreimageUsed
//	Democracy_Proposed                              []EventDemocracyProposed
//	Democracy_Started                               []EventDemocracyStarted
//	Democracy_Tabled                                []EventDemocracyTabled
//	Democracy_Undelegated                           []EventDemocracyUndelegated
//	Democracy_Vetoed                                []EventDemocracyVetoed
//	Electionprovidermultiphase_ElectionFinalized    []EventElectionProviderMultiphaseElectionFinalized
//	Electionprovidermultiphase_Rewarded             []EventElectionProviderMultiphaseRewarded
//	Electionprovidermultiphase_SignedPhaseStarted   []EventElectionProviderMultiphaseSignedPhaseStarted
//	Electionprovidermultiphase_Slashed              []EventElectionProviderMultiphaseSlashed
//	Electionprovidermultiphase_SolutionStored       []EventElectionProviderMultiphaseSolutionStored
//	Electionprovidermultiphase_UnsignedPhaseStarted []EventElectionProviderMultiphaseUnsignedPhaseStarted
//	Elections_CandidateSlashed                      []EventElectionsCandidateSlashed
//	Elections_ElectionError                         []EventElectionsElectionError
//	Elections_EmptyTerm                             []EventElectionsEmptyTerm
//	Elections_MemberKicked                          []EventElectionsMemberKicked
//	Elections_NewTerm                               []EventElectionsNewTerm
//	Elections_Renounced                             []EventElectionsRenounced
//	Elections_SeatHolderSlashed                     []EventElectionsSeatHolderSlashed
//	Gilt_BidPlaced                                  []EventGiltBidPlaced
//	Gilt_BidRetracted                               []EventGiltBidRetracted
//	Gilt_GiltIssued                                 []EventGiltGiltIssued
//	Gilt_GiltThawed                                 []EventGiltGiltThawed
//	Grandpa_NewAuthorities                          []EventGrandpaNewAuthorities
//	Grandpa_Paused                                  []EventGrandpaPaused
//	Grandpa_Resumed                                 []EventGrandpaResumed
//	Identity_IdentityCleared                        []EventIdentityIdentityCleared
//	Identity_IdentityKilled                         []EventIdentityIdentityKilled
//	Identity_IdentitySet                            []EventIdentityIdentitySet
//	Identity_JudgementGiven                         []EventIdentityJudgementGiven
//	Identity_JudgementRequested                     []EventIdentityJudgementRequested
//	Identity_JudgementUnrequested                   []EventIdentityJudgementUnrequested
//	Identity_RegistrarAdded                         []EventIdentityRegistrarAdded
//	Identity_SubIdentityAdded                       []EventIdentitySubIdentityAdded
//	Identity_SubIdentityRemoved                     []EventIdentitySubIdentityRemoved
//	Identity_SubIdentityRevoked                     []EventIdentitySubIdentityRevoked
//	Imonline_AllGood                                []EventImonlineAllGood
//	Imonline_HeartbeatReceived                      []EventImonlineHeartbeatReceived
//	Imonline_SomeOffline                            []EventImonlineSomeOffline
//	Indices_IndexAssigned                           []EventIndicesIndexAssigned
//	Indices_IndexFreed                              []EventIndicesIndexFreed
//	Indices_IndexFrozen                             []EventIndicesIndexFrozen
//	Lottery_CallsUpdated                            []EventLotteryCallsUpdated
//	Lottery_LotteryStarted                          []EventLotteryLotteryStarted
//	Lottery_TicketBought                            []EventLotteryTicketBought
//	Lottery_Winner                                  []EventLotteryWinner
//	Multisig_MultisigApproval                       []EventMultisigMultisigApproval
//	Multisig_MultisigCancelled                      []EventMultisigMultisigCancelled
//	Multisig_MultisigExecuted                       []EventMultisigMultisigExecuted
//	Multisig_NewMultisig                            []EventMultisigNewMultisig
//	Offences_Offence                                []EventOffencesOffence
//	Proxy_Announced                                 []EventProxyAnnounced
//	Proxy_AnonymousCreated                          []EventProxyAnonymousCreated
//	Proxy_ProxyAdded                                []EventProxyProxyAdded
//	Proxy_ProxyExecuted                             []EventProxyProxyExecuted
//	Recovery_AccountRecovered                       []EventRecoveryAccountRecovered
//	Recovery_RecoveryClosed                         []EventRecoveryRecoveryClosed
//	Recovery_RecoveryCreated                        []EventRecoveryRecoveryCreated
//	Recovery_RecoveryInitiated                      []EventRecoveryRecoveryInitiated
//	Recovery_RecoveryRemoved                        []EventRecoveryRecoveryRemoved
//	Recovery_RecoveryVouched                        []EventRecoveryRecoveryVouched
//	Scheduler_Canceled                              []EventSchedulerCanceled
//	Scheduler_Dispatched                            []EventSchedulerDispatched
//	Scheduler_Scheduled                             []EventSchedulerScheduled
//	Session_NewSession                              []EventSessionNewSession
//	Society_AutoUnbid                               []EventSocietyAutoUnbid
//	Society_Bid                                     []EventSocietyBid
//	Society_CandidateSuspended                      []EventSocietyCandidateSuspended
//	Society_Challenged                              []EventSocietyChallenged
//	Society_DefenderVote                            []EventSocietyDefenderVote
//	Society_Deposit                                 []EventSocietyDeposit
//	Society_Founded                                 []EventSocietyFounded
//	Society_Inducted                                []EventSocietyInducted
//	Society_MemberSuspended                         []EventSocietyMemberSuspended
//	Society_NewMaxMembers                           []EventSocietyNewMaxMembers
//	Society_SuspendedMemberJudgement                []EventSocietySuspendedMemberJudgement
//	Society_Unbid                                   []EventSocietyUnbid
//	Society_Unfounded                               []EventSocietyUnfounded
//	Society_Unvouch                                 []EventSocietyUnvouch
//	Society_Vote                                    []EventSocietyVote
//	Society_Vouch                                   []EventSocietyVouch
//	Staking_Bonded                                  []EventStakingBonded
//	Staking_Chilled                                 []EventStakingChilled
//	Staking_EraPaid                                 []EventStakingEraPaid
//	Staking_Kicked                                  []EventStakingKicked
//	Staking_OldSlashingReportDiscarded              []EventStakingOldSlashingReportDiscarded
//	Staking_PayoutStarted                           []EventStakingPayoutStarted
//	Staking_Rewarded                                []EventStakingRewarded
//	Staking_Slashed                                 []EventStakingSlashed
//	Staking_StakersElected                          []EventStakingStakersElected
//	Staking_StakingElectionFailed                   []EventStakingStakingElectionFailed
//	Staking_Unbonded                                []EventStakingUnbonded
//	Staking_Withdrawn                               []EventStakingWithdrawn
//	Sudo_KeyChanged                                 []EventSudoKeyChanged
//	Sudo_Sudid                                      []EventSudoSudid
//	Sudo_SudoAsDone                                 []EventSudoSudoAsDone
//	System_CodeUpdated                              []EventSystemCodeUpdated
//	System_ExtrinsicFailed                          []EventSystemExtrinsicFailed
//	System_ExtrinsicSuccess                         []EventSystemExtrinsicSuccess
//	System_KilledAccount                            []EventSystemKilledAccount
//	System_NewAccount                               []EventSystemNewAccount
//	System_Remarked                                 []EventSystemRemarked
//	Technicalcommittee_Approved                     []EventTechnicalcommitteeApproved
//	Technicalcommittee_Closed                       []EventTechnicalcommitteeClosed
//	Technicalcommittee_Disapproved                  []EventTechnicalcommitteeDisapproved
//	Technicalcommittee_Executed                     []EventTechnicalcommitteeExecuted
//	Technicalcommittee_MemberExecuted               []EventTechnicalcommitteeMemberExecuted
//	Technicalcommittee_Proposed                     []EventTechnicalcommitteeProposed
//	Technicalcommittee_Voted                        []EventTechnicalcommitteeVoted
//	Technicalmembership_Dummy                       []EventTechnicalmembershipDummy
//	Technicalmembership_KeyChanged                  []EventTechnicalmembershipKeyChanged
//	Technicalmembership_MemberAdded                 []EventTechnicalmembershipMemberAdded
//	Technicalmembership_MemberRemoved               []EventTechnicalmembershipMemberRemoved
//	Technicalmembership_MembersReset                []EventTechnicalmembershipMembersReset
//	Technicalmembership_MembersSwapped              []EventTechnicalmembershipMembersSwapped
//	Tips_NewTip                                     []EventTipsNewTip
//	Tips_TipClosed                                  []EventTipsTipClosed
//	Tips_TipClosing                                 []EventTipsTipClosing
//	Tips_TipRetracted                               []EventTipsTipRetracted
//	Tips_TipSlashed                                 []EventTipsTipSlashed
//	Transactionstorage_ProofChecked                 []EventTransactionstorageProofChecked
//	Transactionstorage_Renewed                      []EventTransactionstorageRenewed
//	Transactionstorage_Stored                       []EventTransactionstorageStored
//	Treasury_Awarded                                []EventTreasuryAwarded
//	Treasury_Burnt                                  []EventTreasuryBurnt
//	Treasury_Deposit                                []EventTreasuryDeposit
//	Treasury_Proposed                               []EventTreasuryProposed
//	Treasury_Rejected                               []EventTreasuryRejected
//	Treasury_Rollover                               []EventTreasuryRollover
//	Treasury_Spending                               []EventTreasurySpending
//	Uniques_ApprovalCancelled                       []EventUniquesApprovalCancelled
//	Uniques_ApprovedTransfer                        []EventUniquesApprovedTransfer
//	Uniques_AssetStatusChanged                      []EventUniquesAssetStatusChanged
//	Uniques_AttributeCleared                        []EventUniquesAttributeCleared
//	Uniques_AttributeSet                            []EventUniquesAttributeSet
//	Uniques_Burned                                  []EventUniquesBurned
//	Uniques_ClassFrozen                             []EventUniquesClassFrozen
//	Uniques_ClassMetadataCleared                    []EventUniquesClassMetadataCleared
//	Uniques_ClassMetadataSet                        []EventUniquesClassMetadataSet
//	Uniques_ClassThawed                             []EventUniquesClassThawed
//	Uniques_Created                                 []EventUniquesCreated
//	Uniques_Destroyed                               []EventUniquesDestroyed
//	Uniques_ForceCreated                            []EventUniquesForceCreated
//	Uniques_Frozen                                  []EventUniquesFrozen
//	Uniques_Issued                                  []EventUniquesIssued
//	Uniques_MetadataCleared                         []EventUniquesMetadataCleared
//	Uniques_MetadataSet                             []EventUniquesMetadataSet
//	Uniques_OwnerChanged                            []EventUniquesOwnerChanged
//	Uniques_Redeposited                             []EventUniquesRedeposited
//	Uniques_TeamChanged                             []EventUniquesTeamChanged
//	Uniques_Thawed                                  []EventUniquesThawed
//	Uniques_Transferred                             []EventUniquesTransferred
//	Utility_BatchCompleted                          []EventUtilityBatchCompleted
//	Utility_BatchInterrupted                        []EventUtilityBatchInterrupted
//	Utility_ItemCompleted                           []EventUtilityItemCompleted
//	Vesting_VestingCompleted                        []EventVestingVestingCompleted
//	Vesting_VestingUpdated                          []EventVestingVestingUpdated
//}

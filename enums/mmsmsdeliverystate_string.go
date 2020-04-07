// Code generated by "stringer -type=MMSmsDeliveryState"; DO NOT EDIT.

package enums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSmsDeliveryStateCompletedReceived-0]
	_ = x[MmSmsDeliveryStateCompletedForwardedUnconfirmed-1]
	_ = x[MmSmsDeliveryStateCompletedReplacedBySc-2]
	_ = x[MmSmsDeliveryStateTemporaryErrorCongestion-32]
	_ = x[MmSmsDeliveryStateTemporaryErrorSmeBusy-33]
	_ = x[MmSmsDeliveryStateTemporaryErrorNoResponseFromSme-34]
	_ = x[MmSmsDeliveryStateTemporaryErrorServiceRejected-35]
	_ = x[MmSmsDeliveryStateTemporaryErrorQosNotAvailable-36]
	_ = x[MmSmsDeliveryStateTemporaryErrorInSme-37]
	_ = x[MmSmsDeliveryStateErrorRemoteProcedure-64]
	_ = x[MmSmsDeliveryStateErrorIncompatibleDestination-65]
	_ = x[MmSmsDeliveryStateErrorConnectionRejected-66]
	_ = x[MmSmsDeliveryStateErrorNotObtainable-67]
	_ = x[MmSmsDeliveryStateErrorQosNotAvailable-68]
	_ = x[MmSmsDeliveryStateErrorNoInterworkingAvailable-69]
	_ = x[MmSmsDeliveryStateErrorValidityPeriodExpired-70]
	_ = x[MmSmsDeliveryStateErrorDeletedByOriginatingSme-71]
	_ = x[MmSmsDeliveryStateErrorDeletedByScAdministration-72]
	_ = x[MmSmsDeliveryStateErrorMessageDoesNotExist-73]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorCongestion-96]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorSmeBusy-97]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorNoResponseFromSme-98]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorServiceRejected-99]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorQosNotAvailable-100]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorInSme-101]
	_ = x[MmSmsDeliveryStateNetworkProblemAddressVacant-512]
	_ = x[MmSmsDeliveryStateNetworkProblemAddressTranslationFailure-513]
	_ = x[MmSmsDeliveryStateNetworkProblemNetworkResourceOutage-514]
	_ = x[MmSmsDeliveryStateNetworkProblemNetworkFailure-515]
	_ = x[MmSmsDeliveryStateNetworkProblemInvalidTeleserviceId-516]
	_ = x[MmSmsDeliveryStateNetworkProblemOther-517]
	_ = x[MmSmsDeliveryStateTerminalProblemNoPageResponse-544]
	_ = x[MmSmsDeliveryStateTerminalProblemDestinationBusy-545]
	_ = x[MmSmsDeliveryStateTerminalProblemNoAcknowledgment-546]
	_ = x[MmSmsDeliveryStateTerminalProblemDestinationResourceShortage-547]
	_ = x[MmSmsDeliveryStateTerminalProblemSmsDeliveryPostponed-548]
	_ = x[MmSmsDeliveryStateTerminalProblemDestinationOutOfService-549]
	_ = x[MmSmsDeliveryStateTerminalProblemDestinationNoLongerAtThisAddress-550]
	_ = x[MmSmsDeliveryStateTerminalProblemOther-551]
	_ = x[MmSmsDeliveryStateRadioInterfaceProblemResourceShortage-576]
	_ = x[MmSmsDeliveryStateRadioInterfaceProblemIncompatibility-577]
	_ = x[MmSmsDeliveryStateRadioInterfaceProblemOther-578]
	_ = x[MmSmsDeliveryStateGeneralProblemEncoding-608]
	_ = x[MmSmsDeliveryStateGeneralProblemSmsOriginationDenied-609]
	_ = x[MmSmsDeliveryStateGeneralProblemSmsTerminationDenied-610]
	_ = x[MmSmsDeliveryStateGeneralProblemSupplementaryServiceNotSupported-611]
	_ = x[MmSmsDeliveryStateGeneralProblemSmsNotSupported-612]
	_ = x[MmSmsDeliveryStateGeneralProblemMissingExpectedParameter-614]
	_ = x[MmSmsDeliveryStateGeneralProblemMissingMandatoryParameter-615]
	_ = x[MmSmsDeliveryStateGeneralProblemUnrecognizedParameterValue-616]
	_ = x[MmSmsDeliveryStateGeneralProblemUnexpectedParameterValue-617]
	_ = x[MmSmsDeliveryStateGeneralProblemUserDataSizeError-618]
	_ = x[MmSmsDeliveryStateGeneralProblemOther-619]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemAddressVacant-768]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemAddressTranslationFailure-769]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemNetworkResourceOutage-770]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemNetworkFailure-771]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemInvalidTeleserviceId-772]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemOther-773]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemNoPageResponse-800]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemDestinationBusy-801]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemNoAcknowledgment-802]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemDestinationResourceShortage-803]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemSmsDeliveryPostponed-804]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemDestinationOutOfService-805]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemDestinationNoLongerAtThisAddress-806]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemOther-807]
	_ = x[MmSmsDeliveryStateTemporaryRadioInterfaceProblemResourceShortage-832]
	_ = x[MmSmsDeliveryStateTemporaryRadioInterfaceProblemIncompatibility-833]
	_ = x[MmSmsDeliveryStateTemporaryRadioInterfaceProblemOther-834]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemEncoding-864]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemSmsOriginationDenied-865]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemSmsTerminationDenied-866]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemSupplementaryServiceNotSupported-867]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemSmsNotSupported-868]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemMissingExpectedParameter-870]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemMissingMandatoryParameter-871]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemUnrecognizedParameterValue-872]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemUnexpectedParameterValue-873]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemUserDataSizeError-874]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemOther-875]
}

const _MMSmsDeliveryState_name = "MmSmsDeliveryStateCompletedReceivedMmSmsDeliveryStateCompletedForwardedUnconfirmedMmSmsDeliveryStateCompletedReplacedByScMmSmsDeliveryStateTemporaryErrorCongestionMmSmsDeliveryStateTemporaryErrorSmeBusyMmSmsDeliveryStateTemporaryErrorNoResponseFromSmeMmSmsDeliveryStateTemporaryErrorServiceRejectedMmSmsDeliveryStateTemporaryErrorQosNotAvailableMmSmsDeliveryStateTemporaryErrorInSmeMmSmsDeliveryStateErrorRemoteProcedureMmSmsDeliveryStateErrorIncompatibleDestinationMmSmsDeliveryStateErrorConnectionRejectedMmSmsDeliveryStateErrorNotObtainableMmSmsDeliveryStateErrorQosNotAvailableMmSmsDeliveryStateErrorNoInterworkingAvailableMmSmsDeliveryStateErrorValidityPeriodExpiredMmSmsDeliveryStateErrorDeletedByOriginatingSmeMmSmsDeliveryStateErrorDeletedByScAdministrationMmSmsDeliveryStateErrorMessageDoesNotExistMmSmsDeliveryStateTemporaryFatalErrorCongestionMmSmsDeliveryStateTemporaryFatalErrorSmeBusyMmSmsDeliveryStateTemporaryFatalErrorNoResponseFromSmeMmSmsDeliveryStateTemporaryFatalErrorServiceRejectedMmSmsDeliveryStateTemporaryFatalErrorQosNotAvailableMmSmsDeliveryStateTemporaryFatalErrorInSmeMmSmsDeliveryStateNetworkProblemAddressVacantMmSmsDeliveryStateNetworkProblemAddressTranslationFailureMmSmsDeliveryStateNetworkProblemNetworkResourceOutageMmSmsDeliveryStateNetworkProblemNetworkFailureMmSmsDeliveryStateNetworkProblemInvalidTeleserviceIdMmSmsDeliveryStateNetworkProblemOtherMmSmsDeliveryStateTerminalProblemNoPageResponseMmSmsDeliveryStateTerminalProblemDestinationBusyMmSmsDeliveryStateTerminalProblemNoAcknowledgmentMmSmsDeliveryStateTerminalProblemDestinationResourceShortageMmSmsDeliveryStateTerminalProblemSmsDeliveryPostponedMmSmsDeliveryStateTerminalProblemDestinationOutOfServiceMmSmsDeliveryStateTerminalProblemDestinationNoLongerAtThisAddressMmSmsDeliveryStateTerminalProblemOtherMmSmsDeliveryStateRadioInterfaceProblemResourceShortageMmSmsDeliveryStateRadioInterfaceProblemIncompatibilityMmSmsDeliveryStateRadioInterfaceProblemOtherMmSmsDeliveryStateGeneralProblemEncodingMmSmsDeliveryStateGeneralProblemSmsOriginationDeniedMmSmsDeliveryStateGeneralProblemSmsTerminationDeniedMmSmsDeliveryStateGeneralProblemSupplementaryServiceNotSupportedMmSmsDeliveryStateGeneralProblemSmsNotSupportedMmSmsDeliveryStateGeneralProblemMissingExpectedParameterMmSmsDeliveryStateGeneralProblemMissingMandatoryParameterMmSmsDeliveryStateGeneralProblemUnrecognizedParameterValueMmSmsDeliveryStateGeneralProblemUnexpectedParameterValueMmSmsDeliveryStateGeneralProblemUserDataSizeErrorMmSmsDeliveryStateGeneralProblemOtherMmSmsDeliveryStateTemporaryNetworkProblemAddressVacantMmSmsDeliveryStateTemporaryNetworkProblemAddressTranslationFailureMmSmsDeliveryStateTemporaryNetworkProblemNetworkResourceOutageMmSmsDeliveryStateTemporaryNetworkProblemNetworkFailureMmSmsDeliveryStateTemporaryNetworkProblemInvalidTeleserviceIdMmSmsDeliveryStateTemporaryNetworkProblemOtherMmSmsDeliveryStateTemporaryTerminalProblemNoPageResponseMmSmsDeliveryStateTemporaryTerminalProblemDestinationBusyMmSmsDeliveryStateTemporaryTerminalProblemNoAcknowledgmentMmSmsDeliveryStateTemporaryTerminalProblemDestinationResourceShortageMmSmsDeliveryStateTemporaryTerminalProblemSmsDeliveryPostponedMmSmsDeliveryStateTemporaryTerminalProblemDestinationOutOfServiceMmSmsDeliveryStateTemporaryTerminalProblemDestinationNoLongerAtThisAddressMmSmsDeliveryStateTemporaryTerminalProblemOtherMmSmsDeliveryStateTemporaryRadioInterfaceProblemResourceShortageMmSmsDeliveryStateTemporaryRadioInterfaceProblemIncompatibilityMmSmsDeliveryStateTemporaryRadioInterfaceProblemOtherMmSmsDeliveryStateTemporaryGeneralProblemEncodingMmSmsDeliveryStateTemporaryGeneralProblemSmsOriginationDeniedMmSmsDeliveryStateTemporaryGeneralProblemSmsTerminationDeniedMmSmsDeliveryStateTemporaryGeneralProblemSupplementaryServiceNotSupportedMmSmsDeliveryStateTemporaryGeneralProblemSmsNotSupportedMmSmsDeliveryStateTemporaryGeneralProblemMissingExpectedParameterMmSmsDeliveryStateTemporaryGeneralProblemMissingMandatoryParameterMmSmsDeliveryStateTemporaryGeneralProblemUnrecognizedParameterValueMmSmsDeliveryStateTemporaryGeneralProblemUnexpectedParameterValueMmSmsDeliveryStateTemporaryGeneralProblemUserDataSizeErrorMmSmsDeliveryStateTemporaryGeneralProblemOther"

var _MMSmsDeliveryState_map = map[MMSmsDeliveryState]string{
	0:   _MMSmsDeliveryState_name[0:35],
	1:   _MMSmsDeliveryState_name[35:82],
	2:   _MMSmsDeliveryState_name[82:121],
	32:  _MMSmsDeliveryState_name[121:163],
	33:  _MMSmsDeliveryState_name[163:202],
	34:  _MMSmsDeliveryState_name[202:251],
	35:  _MMSmsDeliveryState_name[251:298],
	36:  _MMSmsDeliveryState_name[298:345],
	37:  _MMSmsDeliveryState_name[345:382],
	64:  _MMSmsDeliveryState_name[382:420],
	65:  _MMSmsDeliveryState_name[420:466],
	66:  _MMSmsDeliveryState_name[466:507],
	67:  _MMSmsDeliveryState_name[507:543],
	68:  _MMSmsDeliveryState_name[543:581],
	69:  _MMSmsDeliveryState_name[581:627],
	70:  _MMSmsDeliveryState_name[627:671],
	71:  _MMSmsDeliveryState_name[671:717],
	72:  _MMSmsDeliveryState_name[717:765],
	73:  _MMSmsDeliveryState_name[765:807],
	96:  _MMSmsDeliveryState_name[807:854],
	97:  _MMSmsDeliveryState_name[854:898],
	98:  _MMSmsDeliveryState_name[898:952],
	99:  _MMSmsDeliveryState_name[952:1004],
	100: _MMSmsDeliveryState_name[1004:1056],
	101: _MMSmsDeliveryState_name[1056:1098],
	512: _MMSmsDeliveryState_name[1098:1143],
	513: _MMSmsDeliveryState_name[1143:1200],
	514: _MMSmsDeliveryState_name[1200:1253],
	515: _MMSmsDeliveryState_name[1253:1299],
	516: _MMSmsDeliveryState_name[1299:1351],
	517: _MMSmsDeliveryState_name[1351:1388],
	544: _MMSmsDeliveryState_name[1388:1435],
	545: _MMSmsDeliveryState_name[1435:1483],
	546: _MMSmsDeliveryState_name[1483:1532],
	547: _MMSmsDeliveryState_name[1532:1592],
	548: _MMSmsDeliveryState_name[1592:1645],
	549: _MMSmsDeliveryState_name[1645:1701],
	550: _MMSmsDeliveryState_name[1701:1766],
	551: _MMSmsDeliveryState_name[1766:1804],
	576: _MMSmsDeliveryState_name[1804:1859],
	577: _MMSmsDeliveryState_name[1859:1913],
	578: _MMSmsDeliveryState_name[1913:1957],
	608: _MMSmsDeliveryState_name[1957:1997],
	609: _MMSmsDeliveryState_name[1997:2049],
	610: _MMSmsDeliveryState_name[2049:2101],
	611: _MMSmsDeliveryState_name[2101:2165],
	612: _MMSmsDeliveryState_name[2165:2212],
	614: _MMSmsDeliveryState_name[2212:2268],
	615: _MMSmsDeliveryState_name[2268:2325],
	616: _MMSmsDeliveryState_name[2325:2383],
	617: _MMSmsDeliveryState_name[2383:2439],
	618: _MMSmsDeliveryState_name[2439:2488],
	619: _MMSmsDeliveryState_name[2488:2525],
	768: _MMSmsDeliveryState_name[2525:2579],
	769: _MMSmsDeliveryState_name[2579:2645],
	770: _MMSmsDeliveryState_name[2645:2707],
	771: _MMSmsDeliveryState_name[2707:2762],
	772: _MMSmsDeliveryState_name[2762:2823],
	773: _MMSmsDeliveryState_name[2823:2869],
	800: _MMSmsDeliveryState_name[2869:2925],
	801: _MMSmsDeliveryState_name[2925:2982],
	802: _MMSmsDeliveryState_name[2982:3040],
	803: _MMSmsDeliveryState_name[3040:3109],
	804: _MMSmsDeliveryState_name[3109:3171],
	805: _MMSmsDeliveryState_name[3171:3236],
	806: _MMSmsDeliveryState_name[3236:3310],
	807: _MMSmsDeliveryState_name[3310:3357],
	832: _MMSmsDeliveryState_name[3357:3421],
	833: _MMSmsDeliveryState_name[3421:3484],
	834: _MMSmsDeliveryState_name[3484:3537],
	864: _MMSmsDeliveryState_name[3537:3586],
	865: _MMSmsDeliveryState_name[3586:3647],
	866: _MMSmsDeliveryState_name[3647:3708],
	867: _MMSmsDeliveryState_name[3708:3781],
	868: _MMSmsDeliveryState_name[3781:3837],
	870: _MMSmsDeliveryState_name[3837:3902],
	871: _MMSmsDeliveryState_name[3902:3968],
	872: _MMSmsDeliveryState_name[3968:4035],
	873: _MMSmsDeliveryState_name[4035:4100],
	874: _MMSmsDeliveryState_name[4100:4158],
	875: _MMSmsDeliveryState_name[4158:4204],
}

func (i MMSmsDeliveryState) String() string {
	if str, ok := _MMSmsDeliveryState_map[i]; ok {
		return str
	}
	return "MMSmsDeliveryState(" + strconv.FormatInt(int64(i), 10) + ")"
}
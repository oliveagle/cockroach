// Code generated by "stringer -type=Method"; DO NOT EDIT

package roachpb

import "fmt"

const _Method_name = "GetPutConditionalPutIncrementDeleteDeleteRangeScanReverseScanBeginTransactionEndTransactionAdminSplitAdminMergeHeartbeatTxnGCPushTxnRangeLookupResolveIntentResolveIntentRangeNoopMergeTruncateLogLeaderLeaseBatch"

var _Method_index = [...]uint8{0, 3, 6, 20, 29, 35, 46, 50, 61, 77, 91, 101, 111, 123, 125, 132, 143, 156, 174, 178, 183, 194, 205, 210}

func (i Method) String() string {
	if i < 0 || i >= Method(len(_Method_index)-1) {
		return fmt.Sprintf("Method(%d)", i)
	}
	return _Method_name[_Method_index[i]:_Method_index[i+1]]
}

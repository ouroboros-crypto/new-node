import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateParams } from "./types/../../../../go/pkg/mod/github.com/cosmos/ibc-go/v8@v8.5.1/proto/ibc/applications/interchain_accounts/host/v1/tx";
import { MsgUpdateParamsResponse } from "./types/../../../../go/pkg/mod/github.com/cosmos/ibc-go/v8@v8.5.1/proto/ibc/applications/interchain_accounts/host/v1/tx";
import { MsgModuleQuerySafe } from "./types/../../../../go/pkg/mod/github.com/cosmos/ibc-go/v8@v8.5.1/proto/ibc/applications/interchain_accounts/host/v1/tx";
import { MsgModuleQuerySafeResponse } from "./types/../../../../go/pkg/mod/github.com/cosmos/ibc-go/v8@v8.5.1/proto/ibc/applications/interchain_accounts/host/v1/tx";
import { QueryParamsResponse } from "./types/../../../../go/pkg/mod/github.com/cosmos/ibc-go/v8@v8.5.1/proto/ibc/applications/interchain_accounts/host/v1/query";
import { QueryRequest } from "./types/../../../../go/pkg/mod/github.com/cosmos/ibc-go/v8@v8.5.1/proto/ibc/applications/interchain_accounts/host/v1/host";
import { QueryParamsRequest } from "./types/../../../../go/pkg/mod/github.com/cosmos/ibc-go/v8@v8.5.1/proto/ibc/applications/interchain_accounts/host/v1/query";
import { Params } from "./types/../../../../go/pkg/mod/github.com/cosmos/ibc-go/v8@v8.5.1/proto/ibc/applications/interchain_accounts/host/v1/host";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/ibc.applications.interchain_accounts.host.v1.MsgUpdateParams", MsgUpdateParams],
    ["/ibc.applications.interchain_accounts.host.v1.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/ibc.applications.interchain_accounts.host.v1.MsgModuleQuerySafe", MsgModuleQuerySafe],
    ["/ibc.applications.interchain_accounts.host.v1.MsgModuleQuerySafeResponse", MsgModuleQuerySafeResponse],
    ["/ibc.applications.interchain_accounts.host.v1.QueryParamsResponse", QueryParamsResponse],
    ["/ibc.applications.interchain_accounts.host.v1.QueryRequest", QueryRequest],
    ["/ibc.applications.interchain_accounts.host.v1.QueryParamsRequest", QueryParamsRequest],
    ["/ibc.applications.interchain_accounts.host.v1.Params", Params],
    
];

export { msgTypes }
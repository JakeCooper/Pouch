module Messages exposing (Msg(..))

import Model exposing (CloudObject, Ordering, FileUrlObject)
import RemoteData exposing (WebData)
import Time exposing (Time)
import Http


type Msg
    = OnPollForObjects (WebData (List CloudObject))
    | OrderObjects Ordering
    | UpdateQuery String
    | Tick Time
    | UpdateCurrentPath String
    | OnClickFile CloudObject
    | OnReceiveFileUrl (Result Http.Error FileUrlObject)

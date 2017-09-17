module Messages exposing (Msg(..))

import Model exposing (CloudObject, Ordering)
import RemoteData exposing (WebData)
import Time exposing (Time)


type Msg
    = OnPollForObjects (WebData (List CloudObject))
    | OrderObjects Ordering
    | UpdateQuery String
    | Tick Time

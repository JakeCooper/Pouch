module Messages exposing (Msg(..))

import Model exposing (CloudObject, Ordering)
import RemoteData exposing (WebData)


type Msg
    = OnFetchObjects (WebData (List CloudObject))
    | OrderObjects Ordering
    | UpdateQuery String

module Messages exposing (Msg(..))

import Model exposing (CloudObject)
import RemoteData exposing (WebData)


type Msg
    = OnFetchObjects (WebData (List CloudObject))

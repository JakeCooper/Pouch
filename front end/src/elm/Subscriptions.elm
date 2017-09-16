module Subscriptions exposing (..)

import Messages exposing (Msg(..))
import Model exposing (Model)


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none

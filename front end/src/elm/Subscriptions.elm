module Subscriptions exposing (..)

import Messages exposing (Msg(..))
import Model exposing (Model)
import Time exposing (every, second)


subscriptions : Model -> Sub Msg
subscriptions model =
    every second Tick

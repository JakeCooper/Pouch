module Subscriptions exposing (..)

import Messages exposing (Msg(..))
import Model exposing (Model)
import Time exposing (every, minute)


subscriptions : Model -> Sub Msg
subscriptions model =
    every minute Tick

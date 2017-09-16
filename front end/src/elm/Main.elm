module Main exposing (..)

import Html exposing (Html, program)
import Subscriptions exposing (subscriptions)
import Messages exposing (Msg(..))
import Model exposing (Model, initialModel)
import Update exposing (update)
import View exposing (view)


init : ( Model, Cmd Msg )
init =
    ( initialModel, Cmd.none )


main : Program Never Model Msg
main =
    program
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }

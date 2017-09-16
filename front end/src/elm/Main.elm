module Main exposing (..)

import Html exposing (Html, program)
import Subscriptions exposing (subscriptions)
import Messages exposing (Msg(..))
import Model exposing (Model, initialModel)
import Update exposing (update)
import View exposing (view)
import Commands exposing (fetchObjects)


init : ( Model, Cmd Msg )
init =
    ( initialModel, fetchObjects )


main : Program Never Model Msg
main =
    program
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }

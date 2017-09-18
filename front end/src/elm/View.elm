module View exposing (view)

import Html exposing (Html, div, section)
import Html.Attributes exposing (class)
import Messages exposing (Msg(..))
import Model exposing (Model)
import Components.Header
import Components.FileList
import Components.Footer


view : Model -> Html Msg
view model =
    section [ class "hero is-fullheight" ]
        [ Components.Header.view
        , Components.FileList.view model
        , Components.Footer.view
        ]

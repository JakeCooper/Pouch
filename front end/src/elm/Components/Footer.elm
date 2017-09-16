module Components.Footer exposing (view)

import Html exposing (Html, div, text)
import Html.Attributes exposing (class)
import Messages exposing (Msg(..))


view : Html Msg
view =
    div [ class "hero-foot" ]
        [ div [ class "container has-text-centered" ]
            [ text "Hack the North 2017" ]
        ]

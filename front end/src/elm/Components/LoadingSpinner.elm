module Components.LoadingSpinner exposing (view)

import Html exposing (Html, div)
import Html.Attributes exposing (class)
import Messages exposing (Msg(..))


view : Html Msg
view =
    div [ class "loader" ] []

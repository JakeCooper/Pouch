module Components.FileList exposing (view)

import Html exposing (Html, div)
import Html.Attributes exposing (class, src)
import Messages exposing (Msg(..))
import Model exposing (Model)


view : Model -> Html Msg
view model =
    div [ class "hero-body" ]
        []

module Components.Header exposing (view)

import Html exposing (Html, div, header, a, img, i)
import Html.Attributes exposing (class, src, href, attribute)
import Messages exposing (Msg(..))
import Model exposing (Model)


view : Model -> Html Msg
view model =
    div [ class "hero-head" ]
        [ header [ class "nav" ]
            [ div [ class "container" ]
                [ div [ class "nav-left" ]
                    [ a [ class "nav-item", href "/" ]
                        [ img [ src "/static/logo.png" ]
                            []
                        ]
                    ]
                , div [ class "nav-right" ]
                    [ a [ class "nav-item", href "" ]
                        [ i [ attribute "aria-hidden" "true", class "fa fa-plus" ]
                            []
                        ]
                    ]
                ]
            ]
        ]

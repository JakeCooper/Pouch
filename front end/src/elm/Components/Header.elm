module Components.Header exposing (view)

import Html exposing (Html, div, header, a, img, i, input)
import Html.Attributes exposing (class, src, href, attribute, id, type_, placeholder)
import Messages exposing (Msg(..))
import Html.Events exposing (onInput)


view : Html Msg
view =
    div [ class "hero-head" ]
        [ header [ class "nav" ]
            [ div [ class "container" ]
                [ div [ class "nav-left" ]
                    [ a [ class "nav-item", href "/" ]
                        [ img [ src "/static/logo.png" ]
                            []
                        ]
                    ]
                , div [ class "nav-center" ]
                    [ div [ class "nav-item" ]
                        [ input
                            [ class "input"
                            , placeholder "Filter"
                            , onInput UpdateQuery
                            ]
                            []
                        ]
                    ]
                , div [ class "nav-right" ]
                    [ a [ class "nav-item", href "#", attribute "onclick" "javascript:document.getElementById('fileUpload').click()" ]
                        [ i [ class "fa fa-plus", attribute "aria-hidden" "true" ] [] ]
                    , input [ id "fileUpload", type_ "file" ] []
                    ]
                ]
            ]
        ]

module Components.FileList exposing (view)

import Html exposing (Html, div, ul, li, i, text, button, a, span)
import Html.Attributes exposing (class, src, attribute, href)
import Messages exposing (Msg(..))
import Model exposing (Model, CloudObject, ObjectType(..))
import RemoteData exposing (WebData)
import Components.LoadingSpinner exposing (view)
import Date


view : Model -> Html Msg
view model =
    let
        objectsResult =
            case model.objects of
                RemoteData.NotAsked ->
                    text ""

                RemoteData.Loading ->
                    Components.LoadingSpinner.view

                RemoteData.Success objects ->
                    div [ class "hero-body" ]
                        [ div [ class "box" ]
                            [ ul [ class "objects-list" ]
                                (List.map viewObject objects)
                            ]
                        ]

                RemoteData.Failure error ->
                    div [ class "notification is-danger" ]
                        [ button [ class "delete" ] []
                        , text (toString error)
                        ]
    in
        div [ class "hero-body" ]
            [ div [ class "container" ]
                [ div [ class "columns" ]
                    [ div [ class "column is-8 is-offset-2" ]
                        [ objectsResult ]
                    ]
                ]
            ]


viewObject : CloudObject -> Html Msg
viewObject object =
    li [ class "object" ]
        [ div []
            [ i [ attribute "aria-hidden" "true", class (iconForObjectType object.objectType) ]
                []
            , a [ class "file-link", href "/" ]
                [ text object.name ]
            , i [ attribute "aria-hidden" "true", class "fa fa-ellipsis-h options" ]
                []
            , span [ class "modified" ]
                [ text (dateStringFromModified object.modified) ]
            ]
        ]


dateStringFromModified : Int -> String
dateStringFromModified modified =
    let
        date =
            Date.fromTime (toFloat modified)

        month =
            toString (Date.month date)

        day =
            toString (Date.day date)

        year =
            toString (Date.year date)

        hour =
            toString (Date.hour date)

        minute =
            toString (Date.minute date)
    in
        month ++ " " ++ day ++ " " ++ year ++ " " ++ hour ++ ":" ++ minute


iconForObjectType : String -> String
iconForObjectType objectType =
    let
        faClass =
            case objectType of
                "file" ->
                    "fa fa-file-o"

                "folder" ->
                    "fa fa-folder-o"

                _ ->
                    "fa fa-question"
    in
        faClass ++ " icon"

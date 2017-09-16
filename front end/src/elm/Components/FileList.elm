module Components.FileList exposing (view)

import Html exposing (Html, div, ul, li, i, text, button, a, span, p, br)
import Html.Attributes exposing (class, src, attribute, href)
import Html.Events exposing (onClick)
import Messages exposing (Msg(..))
import Model exposing (Model, CloudObject, ObjectType(..), Order(..), Ordering, Field(..))
import RemoteData exposing (WebData)
import Components.LoadingSpinner exposing (view)
import Date


view : Model -> Html Msg
view model =
    let
        objectsResult =
            case model.filteredObjects of
                RemoteData.NotAsked ->
                    text ""

                RemoteData.Loading ->
                    Components.LoadingSpinner.view

                RemoteData.Success objects ->
                    viewFilesBox model objects

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


viewFilesBox : Model -> List CloudObject -> Html Msg
viewFilesBox model objects =
    if not (List.isEmpty objects) then
        div []
            [ sortingOptions model
            , div [ class "box" ]
                [ ul [ class "objects-list" ]
                    (List.map viewObject objects)
                ]
            ]
    else
        text "No files found"


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
            toString ((Date.hour date) % 12)

        minute =
            toString (Date.minute date)

        meridiem =
            if Date.hour date < 12 then
                "AM"
            else
                "PM"
    in
        month ++ " " ++ day ++ " " ++ year ++ " " ++ hour ++ ":" ++ minute ++ " " ++ meridiem


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


sortingOptions : Model -> Html Msg
sortingOptions model =
    div [ class "sortingOptions" ]
        [ span [ class "sortName", onClick (OrderObjects (Ordering Name (reverseOrder model))) ]
            [ text "Name"
            , nameSortIcon model
            ]
        , span [ class "sortModified", onClick (OrderObjects (Ordering Modified (reverseOrder model))) ]
            [ text "Modified"
            , modifiedSortIcon model
            ]
        ]


nameSortIcon : Model -> Html Msg
nameSortIcon model =
    if model.ordering.field == Name then
        if model.ordering.order == Ascending then
            i [ attribute "aria-hidden" "true", class "fa fa-sort-asc icon" ] []
        else
            i [ attribute "aria-hidden" "true", class "fa fa-sort-desc icon" ] []
    else
        text ""


modifiedSortIcon : Model -> Html Msg
modifiedSortIcon model =
    if model.ordering.field == Modified then
        if model.ordering.order == Ascending then
            i [ attribute "aria-hidden" "true", class "fa fa-sort-asc icon" ] []
        else
            i [ attribute "aria-hidden" "true", class "fa fa-sort-desc icon" ] []
    else
        text ""


reverseOrder : Model -> Model.Order
reverseOrder model =
    if model.ordering.order == Ascending then
        Descending
    else
        Ascending

module Update exposing (update)

import Messages exposing (Msg(..))
import Model exposing (Model, Field(..), Order(..), CloudObject)
import RemoteData exposing (WebData, succeed)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        OnFetchObjects response ->
            ( { model | objects = response, filteredObjects = response }, Cmd.none )

        OrderObjects newOrdering ->
            case model.objects of
                RemoteData.NotAsked ->
                    ( model, Cmd.none )

                RemoteData.Loading ->
                    ( model, Cmd.none )

                RemoteData.Success objects ->
                    let
                        newModel =
                            case newOrdering.field of
                                Name ->
                                    if newOrdering.order == Ascending then
                                        { model | ordering = newOrdering, filteredObjects = succeed (List.sortBy .name objects) }
                                    else
                                        { model | ordering = newOrdering, filteredObjects = succeed (List.reverse (List.sortBy .name objects)) }

                                Modified ->
                                    if newOrdering.order == Ascending then
                                        { model | ordering = newOrdering, filteredObjects = succeed (List.sortBy .modified objects) }
                                    else
                                        { model | ordering = newOrdering, filteredObjects = succeed (List.reverse (List.sortBy .modified objects)) }
                    in
                        ( newModel, Cmd.none )

                RemoteData.Failure error ->
                    ( model, Cmd.none )

        UpdateQuery newQuery ->
            case model.objects of
                RemoteData.NotAsked ->
                    ( model, Cmd.none )

                RemoteData.Loading ->
                    ( model, Cmd.none )

                RemoteData.Success objects ->
                    ( { model | query = newQuery, filteredObjects = succeed (List.filter (filterFunction newQuery) objects) }, Cmd.none )

                RemoteData.Failure error ->
                    ( model, Cmd.none )


filterFunction : String -> CloudObject -> Bool
filterFunction query object =
    String.contains query object.name

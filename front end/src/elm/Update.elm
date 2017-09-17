module Update exposing (update)

import Messages exposing (Msg(..))
import Model exposing (Model, Field(..), Order(..), CloudObject)
import RemoteData exposing (WebData, succeed)
import Commands exposing (pollForObjects, fetchSignedFile)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        OnPollForObjects response ->
            ( { model | objects = response, filteredObjects = folderFilter model response }, Cmd.none )

        OrderObjects newOrdering ->
            case model.filteredObjects of
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

        Tick _ ->
            ( model, pollForObjects )

        UpdateCurrentPath newPath ->
            ( { model | currentPath = newPath }, Cmd.none )

        OnClickFile object ->
            ( model, fetchSignedFile object.filePath )

        OnReceiveFileUrl urlObject ->
            case urlObject of
                Ok value ->
                    ( { model | signedUrl = value.url }, Cmd.none )

                Err msg ->
                    ( model, Cmd.none )


folderFilter : Model -> WebData (List CloudObject) -> WebData (List CloudObject)
folderFilter model response =
    case response of
        RemoteData.NotAsked ->
            response

        RemoteData.Loading ->
            response

        RemoteData.Success objects ->
            let
                filteredItems =
                    List.filter (onlyCurrentFolder model) objects
            in
                succeed filteredItems

        RemoteData.Failure error ->
            response


onlyCurrentFolder : Model -> CloudObject -> Bool
onlyCurrentFolder model object =
    String.startsWith model.currentPath object.filePath


filterFunction : String -> CloudObject -> Bool
filterFunction query object =
    String.contains query object.name

module Commands exposing (..)

import Http
import Json.Decode as Decode
import Json.Decode.Pipeline exposing (decode, required)
import Messages exposing (Msg(..))
import Model exposing (CloudObject)
import RemoteData


fetchObjects : Cmd Msg
fetchObjects =
    Http.get fetchObjectsUrl objectsDecoder
        |> RemoteData.sendRequest
        |> Cmd.map OnFetchObjects


fetchObjectsUrl : String
fetchObjectsUrl =
    "http://localhost:3000/objects"


objectsDecoder : Decode.Decoder (List CloudObject)
objectsDecoder =
    Decode.list objectDecoder


objectDecoder : Decode.Decoder CloudObject
objectDecoder =
    decode CloudObject
        |> required "name" Decode.string
        |> required "objectType" Decode.string
        |> required "filePath" Decode.string
        |> required "modified" Decode.int

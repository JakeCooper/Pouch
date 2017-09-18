port module Commands exposing (pollForObjects, download)

import Http
import Json.Decode as Decode
import Json.Decode.Pipeline exposing (decode, required)
import Messages exposing (Msg(..))
import Model exposing (CloudObject)
import RemoteData


port download : String -> Cmd msg


pollForObjects : Cmd Msg
pollForObjects =
    Http.get pollForObjectsUrl objectsDecoder
        |> RemoteData.sendRequest
        |> Cmd.map OnPollForObjects


pollForObjectsUrl : String
pollForObjectsUrl =
    "https://smpzbbu1uk.execute-api.us-west-2.amazonaws.com/prod/pouch_getmetadata"


objectsDecoder : Decode.Decoder (List CloudObject)
objectsDecoder =
    Decode.list objectDecoder


objectDecoder : Decode.Decoder CloudObject
objectDecoder =
    decode CloudObject
        |> required "name" Decode.string
        |> required "objectType" Decode.string
        |> required "filePath" Decode.string
        |> required "modified" Decode.string

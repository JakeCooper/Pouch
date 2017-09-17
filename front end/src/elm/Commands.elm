module Commands exposing (pollForObjects, fetchSignedFile)

import Http
import Json.Encode as Encode
import Json.Decode as Decode
import Json.Decode.Pipeline exposing (decode, required)
import Messages exposing (Msg(..))
import Model exposing (CloudObject, FileUrlObject)
import RemoteData


pollForObjects : Cmd Msg
pollForObjects =
    Http.get pollForObjectsUrl objectsDecoder
        |> RemoteData.sendRequest
        |> Cmd.map OnPollForObjects


pollForObjectsUrl : String
pollForObjectsUrl =
    "https://smpzbbu1uk.execute-api.us-west-2.amazonaws.com/prod/pouch_getmetadata"


fetchSignedFile : String -> Cmd Msg
fetchSignedFile filePath =
    let
        body =
            Http.jsonBody
                (Encode.object
                    [ ( "filename", Encode.string filePath ) ]
                )
    in
        (Http.request
            { body = body
            , expect = Http.expectJson fileUrlDecoder
            , headers = []
            , method = "POST"
            , timeout = Nothing
            , url = fetchSignedFileUrl
            , withCredentials = False
            }
        )
            |> Http.send OnReceiveFileUrl



-- Http.get (fetchSignedFileUrl ++ "/" ++ filePath)
--     |> RemoteData.sendRequest
--     |> Cmd.map OnReceiveFileUrl


fetchSignedFileUrl : String
fetchSignedFileUrl =
    "https://vacry2wzc4.execute-api.us-west-2.amazonaws.com/prod/pouch_signedURL"


fileUrlDecoder : Decode.Decoder FileUrlObject
fileUrlDecoder =
    decode FileUrlObject
        |> required "url" Decode.string


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

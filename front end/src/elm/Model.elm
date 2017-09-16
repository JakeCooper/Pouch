module Model exposing (Model, CloudObject, ObjectType(..), initialModel)

import RemoteData exposing (WebData)


initialModel : Model
initialModel =
    { objects = RemoteData.Loading
    }


type alias Model =
    { objects : WebData (List CloudObject)
    }


type ObjectType
    = File
    | Folder


type alias CloudObject =
    { name : String
    , objectType : String
    , filePath : String
    , modified : Int
    }

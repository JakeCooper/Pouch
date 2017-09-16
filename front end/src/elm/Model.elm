module Model exposing (Model, CloudObject, ObjectType(..))


type alias Model =
    { objects : List CloudObject
    }


type ObjectType
    = File
    | Folder


type alias CloudObject =
    { name : String
    , objectType : ObjectType
    , filePath : String
    , modified : Int
    }

module Model exposing (Model, CloudObject, ObjectType(..), initialModel, Ordering, Order(..), Field(..))

import RemoteData exposing (WebData)


initialModel : Model
initialModel =
    { objects = RemoteData.Loading
    , ordering = Ordering Name Ascending
    }


type alias Model =
    { objects : WebData (List CloudObject)
    , ordering : Ordering
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


type alias Ordering =
    { field : Field
    , order : Order
    }


type Field
    = Name
    | Modified


type Order
    = Ascending
    | Descending

module Model exposing (Model, CloudObject, ObjectType(..), initialModel, Ordering, Order(..), Field(..))

import RemoteData exposing (WebData)


initialModel : Model
initialModel =
    { objects = RemoteData.Loading
    , filteredObjects = RemoteData.Loading
    , ordering = Ordering Name Ascending
    , query = ""
    }


type alias Model =
    { objects : WebData (List CloudObject)
    , filteredObjects : WebData (List CloudObject)
    , ordering : Ordering
    , query : String
    }


type ObjectType
    = File
    | Folder


type alias CloudObject =
    { name : String
    , objectType : String
    , filePath : String
    , modified : String
    , url : String
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

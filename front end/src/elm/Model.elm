module Model exposing (Model, CloudObject, ObjectType(..), initialModel, Ordering, Order(..), Field(..), FileUrlObject)

import RemoteData exposing (WebData)


initialModel : Model
initialModel =
    { objects = RemoteData.Loading
    , filteredObjects = RemoteData.Loading
    , ordering = Ordering Name Ascending
    , query = ""
    , currentPath = ""
    , signedUrl = ""
    }


type alias Model =
    { objects : WebData (List CloudObject)
    , filteredObjects : WebData (List CloudObject)
    , ordering : Ordering
    , query : String
    , currentPath : String
    , signedUrl : String
    }


type ObjectType
    = File
    | Folder


type alias CloudObject =
    { name : String
    , objectType : String
    , filePath : String
    , modified : String
    }


type alias FileUrlObject =
    { url : String
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

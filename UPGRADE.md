#### Upgrade Instructions

##### Upgrade from 2.x to 3.x
* Change your package import paths from `github.com/Kamva/mgm/v2` 
to `github.com/Kamva/v3`.

- built-in `ID` field changed json tag value from `_id` to `id` [f665e15](https://github.com/Kamva/mgm/commit/f665e1592cdac43fb7fd00b1427a91590f14a9ff)  
 
##### Upgrade from 1.x to 2.x
* Change your package import paths from `github.com/Kamva/mgm` 
to `github.com/Kamva/v3`.

* methods `Save` and `SaveWithContext` was removed from collections.
    * To Create an entity use the `Create` or `CreateWithCtx` methods.
    * To update an entity use the `Update` or `UpdateWithCtx` methods.

* method `IsNew` method was removed from the `model` interface,
 don't use it.


  

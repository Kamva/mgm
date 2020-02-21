#### Upgrade Instructions

##### Upgrade from 1.x to 2.x
* Change your package import paths from `github.com/Kamva/mgm` 
to `github.com/Kamva/v2`.

* methods `Save` and `SaveWithContext` was removed from collections.
    * To Create an entity use the `Create` or `CreateWithCtx` methods.
    * To update an entity use the `Update` or `UpdateWithCtx` methods.

* method `IsNew` method was removed from the model interface,
 don't use it.


  
